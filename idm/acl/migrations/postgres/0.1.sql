-- +migrate Up
CREATE TABLE IF NOT EXISTS idm_acl_nodes (
    id           BIGINT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY NOT NULL,
    uuid         VARCHAR(500) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS idm_acl_roles (
    id           BIGINT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY NOT NULL,
    uuid         VARCHAR(500) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS idm_acl_workspaces (
    id           BIGINT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY NOT NULL,
    name         VARCHAR(500) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS idm_acls (
    id           BIGINT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY NOT NULL,
    action_name  VARCHAR(500),
    action_value VARCHAR(500),
    role_id      BIGINT NOT NULL DEFAULT 0,
    node_id      BIGINT NOT NULL DEFAULT 0,
    workspace_id BIGINT NOT NULL DEFAULT 0,

    FOREIGN KEY (node_id) REFERENCES idm_acl_nodes(id),
    FOREIGN KEY (workspace_id) REFERENCES idm_acl_workspaces(id),
    FOREIGN KEY (role_id) REFERENCES idm_acl_roles(id),

    CONSTRAINT acls_u1 UNIQUE(node_id, action_name, role_id, workspace_id)
);

-- INSERT INTO idm_acl_workspaces (id, name) VALUES (-1, "") ON DUPLICATE KEY UPDATE name = "";
INSERT INTO idm_acl_nodes (id, uuid) VALUES (-1, '') ON CONFLICT (id) DO UPDATE SET uuid = '';
INSERT INTO idm_acl_workspaces (id, name) VALUES (-1, '') ON CONFLICT (id) DO UPDATE SET name = '';
INSERT INTO idm_acl_roles (id, uuid) VALUES (-1, '') ON CONFLICT (id) DO UPDATE SET uuid = '';

-- +migrate Down
DROP TABLE idm_acl_nodes;
DROP TABLE idm_acl_roles;
DROP TABLE idm_acl_workspaces;
DROP TABLE idm_acls;