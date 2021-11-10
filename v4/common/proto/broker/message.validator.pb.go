// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: message.proto

package broker

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *Empty) Validate() error {
	return nil
}
func (this *PublishRequest) Validate() error {
	if this.Message != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Message); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Message", err)
		}
	}
	return nil
}
func (this *SubscribeRequest) Validate() error {
	return nil
}
func (this *Message) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	return nil
}