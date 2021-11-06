/*
 * Copyright (c) 2019-2021. Abstrium SAS <team (at) pydio.com>
 * This file is part of Pydio Cells.
 *
 * Pydio Cells is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * Pydio Cells is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with Pydio Cells.  If not, see <http://www.gnu.org/licenses/>.
 *
 * The latest code can be found at <https://pydio.com>.
 */

package uuid

import (
	"context"
	"path"
	"strings"

	"github.com/pydio/cells/common/nodes/abstract"

	"github.com/micro/go-micro/errors"
	"go.uber.org/zap"

	"github.com/pydio/cells/common/log"
	"github.com/pydio/cells/common/nodes"
	"github.com/pydio/cells/common/nodes/acl"
	"github.com/pydio/cells/common/proto/idm"
	"github.com/pydio/cells/common/proto/tree"
	servicecontext "github.com/pydio/cells/common/service/context"
	context2 "github.com/pydio/cells/common/utils/context"
	"github.com/pydio/cells/common/utils/permissions"
)

func WithWorkspace() nodes.Option {
	return func(options *nodes.RouterOptions) {
		options.Wrappers = append(options.Wrappers, NewUuidNodeHandler())
	}
}

// UuidNodeHandler is an AbstractBranchFilter extracting workspace info based on node UUID.
type UuidNodeHandler struct {
	abstract.AbstractBranchFilter
}

func (h *UuidNodeHandler) Adapt(c nodes.Client, options nodes.RouterOptions) nodes.Client {
	h.Next = c
	h.ClientsPool = options.Pool
	return h
}

func NewUuidNodeHandler() *UuidNodeHandler {
	u := &UuidNodeHandler{}
	u.InputMethod = u.updateInputBranch
	u.OutputMethod = u.updateOutputBranch
	return u
}

func (h *UuidNodeHandler) updateInputBranch(ctx context.Context, node *tree.Node, identifier string) (context.Context, *tree.Node, error) {

	if info, alreadySet := nodes.GetBranchInfo(ctx, identifier); alreadySet && info.Client != nil {
		return ctx, node, nil
	}

	if acl.HasAdminKey(ctx) {
		ws := &idm.Workspace{UUID: "ROOT", RootUUIDs: []string{"ROOT"}, Slug: "ROOT"}
		branchInfo := nodes.BranchInfo{}
		branchInfo.Workspace = *ws
		ctx = nodes.WithBranchInfo(ctx, identifier, branchInfo)
		return ctx, node, nil
	}

	accessList, ok := acl.FromContext(ctx)
	if !ok {
		return ctx, node, errors.InternalServerError(nodes.VIEWS_LIBRARY_NAME, "Cannot load access list")
	}

	// Update Access List with resolved virtual nodes
	virtualManager := abstract.GetVirtualNodesManager()
	cPool := h.ClientsPool
	for _, vNode := range virtualManager.ListNodes() {
		if aclNodeMask, has := accessList.GetNodesBitmasks()[vNode.Uuid]; has {
			if resolvedRoot, err := virtualManager.ResolveInContext(ctx, vNode, cPool, false); err == nil {
				log.Logger(ctx).Debug("Updating Access List with resolved node Uuid", zap.Any("virtual", vNode), zap.Any("resolved", resolvedRoot))
				accessList.ReplicateBitmask(vNode.Uuid, resolvedRoot.Uuid)
				for _, roots := range accessList.GetWorkspacesNodes() {
					for rootId := range roots {
						if rootId == vNode.Uuid {
							delete(roots, vNode.Uuid)
							roots[resolvedRoot.Uuid] = aclNodeMask
						}
					}
				}
			}
		}
	}

	parents, err := nodes.BuildAncestorsList(ctx, h.ClientsPool.GetTreeClient(), node)
	if err != nil {
		return ctx, node, err
	}
	workspaces, _ := accessList.BelongsToWorkspaces(ctx, parents...)
	if len(workspaces) == 0 {
		log.Logger(ctx).Debug("Node des not belong to any accessible workspace!", accessList.Zap())
		return ctx, node, errors.Forbidden(nodes.VIEWS_LIBRARY_NAME, "Node does not belong to any accessible workspace!")
	}
	// Use first workspace by default
	branchInfo := nodes.BranchInfo{
		AncestorsList: make(map[string][]*tree.Node, 1),
		Workspace:     *workspaces[0],
	}
	branchInfo.AncestorsList[node.Path] = parents
	ctx = context2.WithAdditionalMetadata(ctx, map[string]string{
		servicecontext.CtxWorkspaceUuid: branchInfo.Workspace.UUID,
	})
	return nodes.WithBranchInfo(ctx, identifier, branchInfo), node, nil
}

func (h *UuidNodeHandler) updateOutputBranch(ctx context.Context, node *tree.Node, identifier string) (context.Context, *tree.Node, error) {

	var accessList *permissions.AccessList
	var ok bool
	if accessList, ok = acl.FromContext(ctx); !ok {
		return ctx, node, nil
	}
	if _, ancestors, e := nodes.AncestorsListFromContext(ctx, node, identifier, h.ClientsPool, false); e == nil {
		out := node.Clone()
		workspaces, wsRoots := accessList.BelongsToWorkspaces(ctx, ancestors...)
		log.Logger(ctx).Debug("Belongs to workspaces", zap.Int("ws length", len(workspaces)), zap.Any("wsRoots", wsRoots))
		for _, ws := range workspaces {
			if relativePath, e := h.relativePathToWsRoot(ctx, ws, node.Path, wsRoots[ws.UUID]); e == nil {
				out.AppearsIn = append(out.AppearsIn, &tree.WorkspaceRelativePath{
					WsUuid:  ws.UUID,
					WsLabel: ws.Label,
					WsSlug:  ws.Slug,
					WsScope: ws.Scope.String(),
					Path:    relativePath,
				})
			} else {
				log.Logger(ctx).Error("Error while computing relative path to root", zap.Error(e))
			}
		}
		return ctx, out, nil
	}

	return ctx, node, nil

}

func (h *UuidNodeHandler) relativePathToWsRoot(ctx context.Context, ws *idm.Workspace, nodeFullPath string, rootNodeId string) (string, error) {

	if resp, e := h.Next.ReadNode(ctx, &tree.ReadNodeRequest{Node: &tree.Node{Uuid: rootNodeId}}); e == nil {
		rootPath := resp.Node.Path
		if strings.HasPrefix(nodeFullPath, rootPath) {
			relPath := strings.TrimPrefix(nodeFullPath, rootPath)
			if len(ws.RootUUIDs) > 1 {
				// This workspace has multiple root, prepend the fake root key
				rootKey := h.MakeRootKey(resp.Node)
				relPath = path.Join(rootKey, relPath)
			}
			return relPath, nil
		} else {
			return "", errors.NotFound("RouterUuid", "Cannot subtract paths "+nodeFullPath+" - "+rootPath)
		}
	} else {
		return "", e
	}

}