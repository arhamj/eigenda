// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.4
// source: churner/churner.proto

package churner

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ChurnRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Ethereum address (in hex like "0x123abcdef...") of the operator.
	OperatorAddress string `protobuf:"bytes,1,opt,name=operator_address,json=operatorAddress,proto3" json:"operator_address,omitempty"`
	// The operator making the churn request.
	OperatorToRegisterPubkeyG1 []byte `protobuf:"bytes,2,opt,name=operator_to_register_pubkey_g1,json=operatorToRegisterPubkeyG1,proto3" json:"operator_to_register_pubkey_g1,omitempty"`
	OperatorToRegisterPubkeyG2 []byte `protobuf:"bytes,3,opt,name=operator_to_register_pubkey_g2,json=operatorToRegisterPubkeyG2,proto3" json:"operator_to_register_pubkey_g2,omitempty"`
	// The operator's BLS signature signed on the keccak256 hash of
	// concat("ChurnRequest", operator address, g1, g2, salt).
	OperatorRequestSignature []byte `protobuf:"bytes,4,opt,name=operator_request_signature,json=operatorRequestSignature,proto3" json:"operator_request_signature,omitempty"`
	// The salt used as part of the message to sign on for operator_request_signature.
	Salt []byte `protobuf:"bytes,5,opt,name=salt,proto3" json:"salt,omitempty"`
	// The quorums to register for.
	// Note:
	//   - If any of the quorum here has already been registered, this entire request
	//     will fail to proceed.
	//   - If any of the quorum fails to register, this entire request will fail.
	//   - Regardless of whether the specified quorums are full or not, the Churner
	//     will return parameters for all quorums specified here. The smart contract will
	//     determine whether it needs to churn out existing operators based on whether
	//     the quorums have available space.
	//
	// The IDs must be in range [0, 254].
	QuorumIds []uint32 `protobuf:"varint,6,rep,packed,name=quorum_ids,json=quorumIds,proto3" json:"quorum_ids,omitempty"`
}

func (x *ChurnRequest) Reset() {
	*x = ChurnRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_churner_churner_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChurnRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChurnRequest) ProtoMessage() {}

func (x *ChurnRequest) ProtoReflect() protoreflect.Message {
	mi := &file_churner_churner_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChurnRequest.ProtoReflect.Descriptor instead.
func (*ChurnRequest) Descriptor() ([]byte, []int) {
	return file_churner_churner_proto_rawDescGZIP(), []int{0}
}

func (x *ChurnRequest) GetOperatorAddress() string {
	if x != nil {
		return x.OperatorAddress
	}
	return ""
}

func (x *ChurnRequest) GetOperatorToRegisterPubkeyG1() []byte {
	if x != nil {
		return x.OperatorToRegisterPubkeyG1
	}
	return nil
}

func (x *ChurnRequest) GetOperatorToRegisterPubkeyG2() []byte {
	if x != nil {
		return x.OperatorToRegisterPubkeyG2
	}
	return nil
}

func (x *ChurnRequest) GetOperatorRequestSignature() []byte {
	if x != nil {
		return x.OperatorRequestSignature
	}
	return nil
}

func (x *ChurnRequest) GetSalt() []byte {
	if x != nil {
		return x.Salt
	}
	return nil
}

func (x *ChurnRequest) GetQuorumIds() []uint32 {
	if x != nil {
		return x.QuorumIds
	}
	return nil
}

type ChurnReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The signature signed by the Churner.
	SignatureWithSaltAndExpiry *SignatureWithSaltAndExpiry `protobuf:"bytes,1,opt,name=signature_with_salt_and_expiry,json=signatureWithSaltAndExpiry,proto3" json:"signature_with_salt_and_expiry,omitempty"`
	// A list of existing operators that get churned out.
	// This list will contain the target operators to be churned out for all quorums specified
	// in the ChurnRequest even if some quorums may not have any churned out operators.
	// It is smart contract's responsibility to determine whether it needs to churn out
	// these target operators based on whether the quorums have available space.
	//
	// For example, if the ChurnRequest specifies quorums 0 and 1 where quorum 0 is full
	// and quorum 1 has available space, the ChurnReply will contain the operators to be
	// churned out for both quorums 0 and 1 (operators with lowest stake). However,
	// smart contract should only churn out the operators for quorum 0 because quorum 1
	// has available space without having any operators churned.
	// Note: it's possible an operator gets churned out just for one or more quorums
	// (rather than entirely churned out for all quorums).
	OperatorsToChurn []*OperatorToChurn `protobuf:"bytes,2,rep,name=operators_to_churn,json=operatorsToChurn,proto3" json:"operators_to_churn,omitempty"`
}

func (x *ChurnReply) Reset() {
	*x = ChurnReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_churner_churner_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChurnReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChurnReply) ProtoMessage() {}

func (x *ChurnReply) ProtoReflect() protoreflect.Message {
	mi := &file_churner_churner_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChurnReply.ProtoReflect.Descriptor instead.
func (*ChurnReply) Descriptor() ([]byte, []int) {
	return file_churner_churner_proto_rawDescGZIP(), []int{1}
}

func (x *ChurnReply) GetSignatureWithSaltAndExpiry() *SignatureWithSaltAndExpiry {
	if x != nil {
		return x.SignatureWithSaltAndExpiry
	}
	return nil
}

func (x *ChurnReply) GetOperatorsToChurn() []*OperatorToChurn {
	if x != nil {
		return x.OperatorsToChurn
	}
	return nil
}

type SignatureWithSaltAndExpiry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Churner's signature on the Operator's attributes.
	Signature []byte `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	// Salt is the keccak256 hash of
	// concat("churn", time.Now(), operatorToChurn's OperatorID, Churner's ECDSA private key)
	Salt []byte `protobuf:"bytes,2,opt,name=salt,proto3" json:"salt,omitempty"`
	// When this churn decision will expire.
	Expiry int64 `protobuf:"varint,3,opt,name=expiry,proto3" json:"expiry,omitempty"`
}

func (x *SignatureWithSaltAndExpiry) Reset() {
	*x = SignatureWithSaltAndExpiry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_churner_churner_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignatureWithSaltAndExpiry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignatureWithSaltAndExpiry) ProtoMessage() {}

func (x *SignatureWithSaltAndExpiry) ProtoReflect() protoreflect.Message {
	mi := &file_churner_churner_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignatureWithSaltAndExpiry.ProtoReflect.Descriptor instead.
func (*SignatureWithSaltAndExpiry) Descriptor() ([]byte, []int) {
	return file_churner_churner_proto_rawDescGZIP(), []int{2}
}

func (x *SignatureWithSaltAndExpiry) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *SignatureWithSaltAndExpiry) GetSalt() []byte {
	if x != nil {
		return x.Salt
	}
	return nil
}

func (x *SignatureWithSaltAndExpiry) GetExpiry() int64 {
	if x != nil {
		return x.Expiry
	}
	return 0
}

// This describes an operator to churn out for a quorum.
type OperatorToChurn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the quorum of the operator to churn out.
	QuorumId uint32 `protobuf:"varint,1,opt,name=quorum_id,json=quorumId,proto3" json:"quorum_id,omitempty"`
	// The address of the operator.
	Operator []byte `protobuf:"bytes,2,opt,name=operator,proto3" json:"operator,omitempty"`
	// BLS pubkey (G1 point) of the operator.
	Pubkey []byte `protobuf:"bytes,3,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
}

func (x *OperatorToChurn) Reset() {
	*x = OperatorToChurn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_churner_churner_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperatorToChurn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperatorToChurn) ProtoMessage() {}

func (x *OperatorToChurn) ProtoReflect() protoreflect.Message {
	mi := &file_churner_churner_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperatorToChurn.ProtoReflect.Descriptor instead.
func (*OperatorToChurn) Descriptor() ([]byte, []int) {
	return file_churner_churner_proto_rawDescGZIP(), []int{3}
}

func (x *OperatorToChurn) GetQuorumId() uint32 {
	if x != nil {
		return x.QuorumId
	}
	return 0
}

func (x *OperatorToChurn) GetOperator() []byte {
	if x != nil {
		return x.Operator
	}
	return nil
}

func (x *OperatorToChurn) GetPubkey() []byte {
	if x != nil {
		return x.Pubkey
	}
	return nil
}

var File_churner_churner_proto protoreflect.FileDescriptor

var file_churner_churner_proto_rawDesc = []byte{
	0x0a, 0x15, 0x63, 0x68, 0x75, 0x72, 0x6e, 0x65, 0x72, 0x2f, 0x63, 0x68, 0x75, 0x72, 0x6e, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x68, 0x75, 0x72, 0x6e, 0x65, 0x72,
	0x22, 0xb2, 0x02, 0x0a, 0x0c, 0x43, 0x68, 0x75, 0x72, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x29, 0x0a, 0x10, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x42, 0x0a, 0x1e,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x74, 0x6f, 0x5f, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x5f, 0x70, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x5f, 0x67, 0x31, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x1a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x54, 0x6f,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x50, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x47, 0x31,
	0x12, 0x42, 0x0a, 0x1e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x74, 0x6f, 0x5f,
	0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x70, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x5f,
	0x67, 0x32, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x1a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x54, 0x6f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x50, 0x75, 0x62, 0x6b,
	0x65, 0x79, 0x47, 0x32, 0x12, 0x3c, 0x0a, 0x1a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x18, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x61, 0x6c, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x73, 0x61, 0x6c, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x71, 0x75, 0x6f, 0x72, 0x75, 0x6d,
	0x5f, 0x69, 0x64, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x09, 0x71, 0x75, 0x6f, 0x72,
	0x75, 0x6d, 0x49, 0x64, 0x73, 0x22, 0xbd, 0x01, 0x0a, 0x0a, 0x43, 0x68, 0x75, 0x72, 0x6e, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x67, 0x0a, 0x1e, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x73, 0x61, 0x6c, 0x74, 0x5f, 0x61, 0x6e, 0x64, 0x5f,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63,
	0x68, 0x75, 0x72, 0x6e, 0x65, 0x72, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x57, 0x69, 0x74, 0x68, 0x53, 0x61, 0x6c, 0x74, 0x41, 0x6e, 0x64, 0x45, 0x78, 0x70, 0x69, 0x72,
	0x79, 0x52, 0x1a, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x57, 0x69, 0x74, 0x68,
	0x53, 0x61, 0x6c, 0x74, 0x41, 0x6e, 0x64, 0x45, 0x78, 0x70, 0x69, 0x72, 0x79, 0x12, 0x46, 0x0a,
	0x12, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x5f, 0x74, 0x6f, 0x5f, 0x63, 0x68,
	0x75, 0x72, 0x6e, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x68, 0x75, 0x72,
	0x6e, 0x65, 0x72, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x54, 0x6f, 0x43, 0x68,
	0x75, 0x72, 0x6e, 0x52, 0x10, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x54, 0x6f,
	0x43, 0x68, 0x75, 0x72, 0x6e, 0x22, 0x66, 0x0a, 0x1a, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x57, 0x69, 0x74, 0x68, 0x53, 0x61, 0x6c, 0x74, 0x41, 0x6e, 0x64, 0x45, 0x78, 0x70,
	0x69, 0x72, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x61, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x04, 0x73, 0x61, 0x6c, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x22, 0x62, 0x0a,
	0x0f, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x54, 0x6f, 0x43, 0x68, 0x75, 0x72, 0x6e,
	0x12, 0x1b, 0x0a, 0x09, 0x71, 0x75, 0x6f, 0x72, 0x75, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x08, 0x71, 0x75, 0x6f, 0x72, 0x75, 0x6d, 0x49, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x75, 0x62,
	0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x70, 0x75, 0x62, 0x6b, 0x65,
	0x79, 0x32, 0x40, 0x0a, 0x07, 0x43, 0x68, 0x75, 0x72, 0x6e, 0x65, 0x72, 0x12, 0x35, 0x0a, 0x05,
	0x43, 0x68, 0x75, 0x72, 0x6e, 0x12, 0x15, 0x2e, 0x63, 0x68, 0x75, 0x72, 0x6e, 0x65, 0x72, 0x2e,
	0x43, 0x68, 0x75, 0x72, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x63,
	0x68, 0x75, 0x72, 0x6e, 0x65, 0x72, 0x2e, 0x43, 0x68, 0x75, 0x72, 0x6e, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x00, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x4c, 0x61, 0x79, 0x72, 0x2d, 0x4c, 0x61, 0x62, 0x73, 0x2f, 0x65, 0x69, 0x67, 0x65,
	0x6e, 0x64, 0x61, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x68, 0x75,
	0x72, 0x6e, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_churner_churner_proto_rawDescOnce sync.Once
	file_churner_churner_proto_rawDescData = file_churner_churner_proto_rawDesc
)

func file_churner_churner_proto_rawDescGZIP() []byte {
	file_churner_churner_proto_rawDescOnce.Do(func() {
		file_churner_churner_proto_rawDescData = protoimpl.X.CompressGZIP(file_churner_churner_proto_rawDescData)
	})
	return file_churner_churner_proto_rawDescData
}

var file_churner_churner_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_churner_churner_proto_goTypes = []interface{}{
	(*ChurnRequest)(nil),               // 0: churner.ChurnRequest
	(*ChurnReply)(nil),                 // 1: churner.ChurnReply
	(*SignatureWithSaltAndExpiry)(nil), // 2: churner.SignatureWithSaltAndExpiry
	(*OperatorToChurn)(nil),            // 3: churner.OperatorToChurn
}
var file_churner_churner_proto_depIdxs = []int32{
	2, // 0: churner.ChurnReply.signature_with_salt_and_expiry:type_name -> churner.SignatureWithSaltAndExpiry
	3, // 1: churner.ChurnReply.operators_to_churn:type_name -> churner.OperatorToChurn
	0, // 2: churner.Churner.Churn:input_type -> churner.ChurnRequest
	1, // 3: churner.Churner.Churn:output_type -> churner.ChurnReply
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_churner_churner_proto_init() }
func file_churner_churner_proto_init() {
	if File_churner_churner_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_churner_churner_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChurnRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_churner_churner_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChurnReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_churner_churner_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignatureWithSaltAndExpiry); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_churner_churner_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperatorToChurn); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_churner_churner_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_churner_churner_proto_goTypes,
		DependencyIndexes: file_churner_churner_proto_depIdxs,
		MessageInfos:      file_churner_churner_proto_msgTypes,
	}.Build()
	File_churner_churner_proto = out.File
	file_churner_churner_proto_rawDesc = nil
	file_churner_churner_proto_goTypes = nil
	file_churner_churner_proto_depIdxs = nil
}
