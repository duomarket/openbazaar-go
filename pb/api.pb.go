// Code generated by protoc-gen-go.
// source: api.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	api.proto
	contracts.proto
	countrycodes.proto
	message.proto
	moderator.proto
	orders.proto
	profile.proto

It has these top-level messages:
	Coupon
	OrderRespApi
	CaseRespApi
	TransactionRecord
	PeerAndProfile
	PeerAndProfileWithID
	RatingWithID
	RicardianContract
	Listing
	Order
	OrderConfirmation
	OrderReject
	RatingSignature
	BitcoinSignature
	OrderFulfillment
	OrderCompletion
	Rating
	Dispute
	DisputeResolution
	DisputeAcceptance
	Outpoint
	Refund
	ID
	Signature
	SignedListing
	Message
	Envelope
	Chat
	SignedData
	Moderator
	DisputeUpdate
	Profile
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Coupon struct {
	Hash string `protobuf:"bytes,1,opt,name=hash" json:"hash,omitempty"`
	Code string `protobuf:"bytes,2,opt,name=code" json:"code,omitempty"`
}

func (m *Coupon) Reset()                    { *m = Coupon{} }
func (m *Coupon) String() string            { return proto.CompactTextString(m) }
func (*Coupon) ProtoMessage()               {}
func (*Coupon) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Coupon) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *Coupon) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type OrderRespApi struct {
	Contract                   *RicardianContract   `protobuf:"bytes,1,opt,name=contract" json:"contract,omitempty"`
	State                      OrderState           `protobuf:"varint,2,opt,name=state,enum=OrderState" json:"state,omitempty"`
	Read                       bool                 `protobuf:"varint,3,opt,name=read" json:"read,omitempty"`
	Funded                     bool                 `protobuf:"varint,4,opt,name=funded" json:"funded,omitempty"`
	UnreadChatMessages         uint64               `protobuf:"varint,5,opt,name=unreadChatMessages" json:"unreadChatMessages,omitempty"`
	PaymentAddressTransactions []*TransactionRecord `protobuf:"bytes,6,rep,name=paymentAddressTransactions" json:"paymentAddressTransactions,omitempty"`
	RefundAddressTransaction   *TransactionRecord   `protobuf:"bytes,7,opt,name=refundAddressTransaction" json:"refundAddressTransaction,omitempty"`
}

func (m *OrderRespApi) Reset()                    { *m = OrderRespApi{} }
func (m *OrderRespApi) String() string            { return proto.CompactTextString(m) }
func (*OrderRespApi) ProtoMessage()               {}
func (*OrderRespApi) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *OrderRespApi) GetContract() *RicardianContract {
	if m != nil {
		return m.Contract
	}
	return nil
}

func (m *OrderRespApi) GetState() OrderState {
	if m != nil {
		return m.State
	}
	return OrderState_PENDING
}

func (m *OrderRespApi) GetRead() bool {
	if m != nil {
		return m.Read
	}
	return false
}

func (m *OrderRespApi) GetFunded() bool {
	if m != nil {
		return m.Funded
	}
	return false
}

func (m *OrderRespApi) GetUnreadChatMessages() uint64 {
	if m != nil {
		return m.UnreadChatMessages
	}
	return 0
}

func (m *OrderRespApi) GetPaymentAddressTransactions() []*TransactionRecord {
	if m != nil {
		return m.PaymentAddressTransactions
	}
	return nil
}

func (m *OrderRespApi) GetRefundAddressTransaction() *TransactionRecord {
	if m != nil {
		return m.RefundAddressTransaction
	}
	return nil
}

type CaseRespApi struct {
	Timestamp                      *google_protobuf.Timestamp `protobuf:"bytes,1,opt,name=timestamp" json:"timestamp,omitempty"`
	BuyerContract                  *RicardianContract         `protobuf:"bytes,2,opt,name=buyerContract" json:"buyerContract,omitempty"`
	VendorContract                 *RicardianContract         `protobuf:"bytes,3,opt,name=vendorContract" json:"vendorContract,omitempty"`
	BuyerContractValidationErrors  []string                   `protobuf:"bytes,4,rep,name=buyerContractValidationErrors" json:"buyerContractValidationErrors,omitempty"`
	VendorContractValidationErrors []string                   `protobuf:"bytes,5,rep,name=vendorContractValidationErrors" json:"vendorContractValidationErrors,omitempty"`
	State                          OrderState                 `protobuf:"varint,6,opt,name=state,enum=OrderState" json:"state,omitempty"`
	Read                           bool                       `protobuf:"varint,7,opt,name=read" json:"read,omitempty"`
	BuyerOpened                    bool                       `protobuf:"varint,8,opt,name=buyerOpened" json:"buyerOpened,omitempty"`
	Claim                          string                     `protobuf:"bytes,9,opt,name=claim" json:"claim,omitempty"`
	UnreadChatMessages             uint64                     `protobuf:"varint,10,opt,name=unreadChatMessages" json:"unreadChatMessages,omitempty"`
	Resolution                     *DisputeResolution         `protobuf:"bytes,11,opt,name=resolution" json:"resolution,omitempty"`
}

func (m *CaseRespApi) Reset()                    { *m = CaseRespApi{} }
func (m *CaseRespApi) String() string            { return proto.CompactTextString(m) }
func (*CaseRespApi) ProtoMessage()               {}
func (*CaseRespApi) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CaseRespApi) GetTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *CaseRespApi) GetBuyerContract() *RicardianContract {
	if m != nil {
		return m.BuyerContract
	}
	return nil
}

func (m *CaseRespApi) GetVendorContract() *RicardianContract {
	if m != nil {
		return m.VendorContract
	}
	return nil
}

func (m *CaseRespApi) GetBuyerContractValidationErrors() []string {
	if m != nil {
		return m.BuyerContractValidationErrors
	}
	return nil
}

func (m *CaseRespApi) GetVendorContractValidationErrors() []string {
	if m != nil {
		return m.VendorContractValidationErrors
	}
	return nil
}

func (m *CaseRespApi) GetState() OrderState {
	if m != nil {
		return m.State
	}
	return OrderState_PENDING
}

func (m *CaseRespApi) GetRead() bool {
	if m != nil {
		return m.Read
	}
	return false
}

func (m *CaseRespApi) GetBuyerOpened() bool {
	if m != nil {
		return m.BuyerOpened
	}
	return false
}

func (m *CaseRespApi) GetClaim() string {
	if m != nil {
		return m.Claim
	}
	return ""
}

func (m *CaseRespApi) GetUnreadChatMessages() uint64 {
	if m != nil {
		return m.UnreadChatMessages
	}
	return 0
}

func (m *CaseRespApi) GetResolution() *DisputeResolution {
	if m != nil {
		return m.Resolution
	}
	return nil
}

type TransactionRecord struct {
	Txid          string                     `protobuf:"bytes,1,opt,name=txid" json:"txid,omitempty"`
	Value         int64                      `protobuf:"varint,2,opt,name=value" json:"value,omitempty"`
	Confirmations uint32                     `protobuf:"varint,3,opt,name=confirmations" json:"confirmations,omitempty"`
	Height        uint32                     `protobuf:"varint,4,opt,name=height" json:"height,omitempty"`
	Timestamp     *google_protobuf.Timestamp `protobuf:"bytes,5,opt,name=timestamp" json:"timestamp,omitempty"`
}

func (m *TransactionRecord) Reset()                    { *m = TransactionRecord{} }
func (m *TransactionRecord) String() string            { return proto.CompactTextString(m) }
func (*TransactionRecord) ProtoMessage()               {}
func (*TransactionRecord) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *TransactionRecord) GetTxid() string {
	if m != nil {
		return m.Txid
	}
	return ""
}

func (m *TransactionRecord) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *TransactionRecord) GetConfirmations() uint32 {
	if m != nil {
		return m.Confirmations
	}
	return 0
}

func (m *TransactionRecord) GetHeight() uint32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *TransactionRecord) GetTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

type PeerAndProfile struct {
	PeerId  string   `protobuf:"bytes,1,opt,name=peerId" json:"peerId,omitempty"`
	Profile *Profile `protobuf:"bytes,2,opt,name=profile" json:"profile,omitempty"`
}

func (m *PeerAndProfile) Reset()                    { *m = PeerAndProfile{} }
func (m *PeerAndProfile) String() string            { return proto.CompactTextString(m) }
func (*PeerAndProfile) ProtoMessage()               {}
func (*PeerAndProfile) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *PeerAndProfile) GetPeerId() string {
	if m != nil {
		return m.PeerId
	}
	return ""
}

func (m *PeerAndProfile) GetProfile() *Profile {
	if m != nil {
		return m.Profile
	}
	return nil
}

type PeerAndProfileWithID struct {
	Id      string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	PeerId  string   `protobuf:"bytes,2,opt,name=peerId" json:"peerId,omitempty"`
	Profile *Profile `protobuf:"bytes,3,opt,name=profile" json:"profile,omitempty"`
}

func (m *PeerAndProfileWithID) Reset()                    { *m = PeerAndProfileWithID{} }
func (m *PeerAndProfileWithID) String() string            { return proto.CompactTextString(m) }
func (*PeerAndProfileWithID) ProtoMessage()               {}
func (*PeerAndProfileWithID) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *PeerAndProfileWithID) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PeerAndProfileWithID) GetPeerId() string {
	if m != nil {
		return m.PeerId
	}
	return ""
}

func (m *PeerAndProfileWithID) GetProfile() *Profile {
	if m != nil {
		return m.Profile
	}
	return nil
}

type RatingWithID struct {
	Id       string  `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	RatingId string  `protobuf:"bytes,2,opt,name=ratingId" json:"ratingId,omitempty"`
	Rating   *Rating `protobuf:"bytes,3,opt,name=rating" json:"rating,omitempty"`
}

func (m *RatingWithID) Reset()                    { *m = RatingWithID{} }
func (m *RatingWithID) String() string            { return proto.CompactTextString(m) }
func (*RatingWithID) ProtoMessage()               {}
func (*RatingWithID) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *RatingWithID) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *RatingWithID) GetRatingId() string {
	if m != nil {
		return m.RatingId
	}
	return ""
}

func (m *RatingWithID) GetRating() *Rating {
	if m != nil {
		return m.Rating
	}
	return nil
}

func init() {
	proto.RegisterType((*Coupon)(nil), "Coupon")
	proto.RegisterType((*OrderRespApi)(nil), "OrderRespApi")
	proto.RegisterType((*CaseRespApi)(nil), "CaseRespApi")
	proto.RegisterType((*TransactionRecord)(nil), "TransactionRecord")
	proto.RegisterType((*PeerAndProfile)(nil), "PeerAndProfile")
	proto.RegisterType((*PeerAndProfileWithID)(nil), "PeerAndProfileWithID")
	proto.RegisterType((*RatingWithID)(nil), "RatingWithID")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 625 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x54, 0xcd, 0x6e, 0x13, 0x31,
	0x10, 0x56, 0xfe, 0x93, 0x49, 0x13, 0x84, 0x55, 0xa1, 0x55, 0x24, 0x68, 0x58, 0x71, 0xc8, 0x69,
	0x8b, 0xca, 0xa5, 0xe2, 0x56, 0x52, 0x90, 0x2a, 0x01, 0xad, 0x4c, 0x05, 0x12, 0x9c, 0x9c, 0xf5,
	0x24, 0xb1, 0x94, 0xd8, 0x2b, 0xdb, 0x5b, 0xd1, 0x97, 0xe1, 0x2d, 0x78, 0x20, 0xde, 0x04, 0xd9,
	0xeb, 0x4d, 0x13, 0xd2, 0x6d, 0xc5, 0xcd, 0x33, 0xf3, 0xcd, 0x37, 0xb3, 0x33, 0xdf, 0x2c, 0xf4,
	0x58, 0x26, 0x92, 0x4c, 0x2b, 0xab, 0x46, 0x4f, 0x52, 0x25, 0xad, 0x66, 0xa9, 0x35, 0xc1, 0x71,
	0xa0, 0x34, 0x47, 0x5d, 0x5a, 0x83, 0x4c, 0xab, 0xb9, 0x58, 0x61, 0x30, 0x8f, 0x16, 0x4a, 0x2d,
	0x56, 0x78, 0xec, 0xad, 0x59, 0x3e, 0x3f, 0xb6, 0x62, 0x8d, 0xc6, 0xb2, 0x75, 0x56, 0x00, 0xe2,
	0xd7, 0xd0, 0x9e, 0xaa, 0x3c, 0x53, 0x92, 0x10, 0x68, 0x2e, 0x99, 0x59, 0x46, 0xb5, 0x71, 0x6d,
	0xd2, 0xa3, 0xfe, 0xed, 0x7c, 0xa9, 0xe2, 0x18, 0xd5, 0x0b, 0x9f, 0x7b, 0xc7, 0x7f, 0xea, 0x70,
	0x70, 0xe9, 0x4a, 0x52, 0x34, 0xd9, 0x59, 0x26, 0x48, 0x02, 0xdd, 0xb2, 0x27, 0x9f, 0xdc, 0x3f,
	0x21, 0x09, 0x15, 0x29, 0xd3, 0x5c, 0x30, 0x39, 0x0d, 0x11, 0xba, 0xc1, 0x90, 0x97, 0xd0, 0x32,
	0x96, 0xd9, 0x82, 0x75, 0x78, 0xd2, 0x4f, 0x3c, 0xdb, 0x17, 0xe7, 0xa2, 0x45, 0xc4, 0xd5, 0xd5,
	0xc8, 0x78, 0xd4, 0x18, 0xd7, 0x26, 0x5d, 0xea, 0xdf, 0xe4, 0x19, 0xb4, 0xe7, 0xb9, 0xe4, 0xc8,
	0xa3, 0xa6, 0xf7, 0x06, 0x8b, 0x24, 0x40, 0x72, 0xe9, 0x10, 0xd3, 0x25, 0xb3, 0x9f, 0xd0, 0x18,
	0xb6, 0x40, 0x13, 0xb5, 0xc6, 0xb5, 0x49, 0x93, 0xde, 0x13, 0x21, 0x14, 0x46, 0x19, 0xbb, 0x5d,
	0xa3, 0xb4, 0x67, 0x9c, 0x6b, 0x34, 0xe6, 0x5a, 0x33, 0x69, 0x58, 0x6a, 0x85, 0x92, 0x26, 0x6a,
	0x8f, 0x1b, 0xfe, 0x03, 0xb6, 0x9c, 0x14, 0x53, 0xa5, 0x39, 0x7d, 0x20, 0x8b, 0x7c, 0x86, 0x48,
	0xa3, 0xeb, 0x67, 0x3f, 0x18, 0x75, 0xc2, 0x48, 0xf6, 0x19, 0x2b, 0x73, 0xe2, 0x5f, 0x4d, 0xe8,
	0x4f, 0x99, 0xc1, 0x72, 0xc4, 0xa7, 0xd0, 0xdb, 0x2c, 0x2e, 0xcc, 0x78, 0x94, 0x14, 0xab, 0x4d,
	0xca, 0xd5, 0x26, 0xd7, 0x25, 0x82, 0xde, 0x81, 0xc9, 0x29, 0x0c, 0x66, 0xf9, 0x2d, 0xea, 0x72,
	0x0f, 0x7e, 0xe8, 0xf7, 0x6f, 0x68, 0x17, 0x48, 0xde, 0xc2, 0xf0, 0x06, 0x25, 0x57, 0x77, 0xa9,
	0x8d, 0xca, 0xd4, 0x7f, 0x90, 0xe4, 0x1c, 0x9e, 0xef, 0x90, 0x7d, 0x65, 0x2b, 0xc1, 0x99, 0xfb,
	0xb4, 0xf7, 0x5a, 0x2b, 0x6d, 0xa2, 0xe6, 0xb8, 0x31, 0xe9, 0xd1, 0x87, 0x41, 0xe4, 0x03, 0xbc,
	0xd8, 0xe5, 0xdd, 0xa3, 0x69, 0x79, 0x9a, 0x47, 0x50, 0x77, 0x82, 0x6b, 0x3f, 0x2a, 0xb8, 0xce,
	0x96, 0xe0, 0xc6, 0xd0, 0xf7, 0xfd, 0x5d, 0x66, 0x28, 0x91, 0x47, 0x5d, 0x1f, 0xda, 0x76, 0x91,
	0x43, 0x68, 0xa5, 0x2b, 0x26, 0xd6, 0x51, 0xcf, 0xdf, 0x47, 0x61, 0x54, 0x08, 0x12, 0x2a, 0x05,
	0x79, 0x02, 0xa0, 0xd1, 0xa8, 0x55, 0xee, 0xe5, 0xd2, 0x0f, 0x43, 0x3e, 0x17, 0x26, 0xcb, 0xad,
	0x53, 0x40, 0x88, 0xd0, 0x2d, 0x54, 0xfc, 0xbb, 0x06, 0x4f, 0xf7, 0x04, 0xe5, 0xbe, 0xc2, 0xfe,
	0x14, 0xbc, 0x3c, 0x61, 0xf7, 0x76, 0x3d, 0xde, 0xb0, 0x55, 0x5e, 0x5c, 0x5b, 0x83, 0x16, 0x06,
	0x79, 0x05, 0x83, 0x54, 0xc9, 0xb9, 0xd0, 0x6b, 0x56, 0xe8, 0xde, 0xed, 0x76, 0x40, 0x77, 0x9d,
	0xee, 0xe4, 0x96, 0x28, 0x16, 0x4b, 0xeb, 0x4f, 0x6e, 0x40, 0x83, 0xb5, 0x2b, 0xc7, 0xd6, 0x7f,
	0xc8, 0x31, 0xfe, 0x08, 0xc3, 0x2b, 0x44, 0x7d, 0x26, 0xf9, 0x55, 0xf1, 0x9f, 0x72, 0x35, 0x32,
	0x44, 0x7d, 0x51, 0x76, 0x1d, 0x2c, 0x12, 0x43, 0x27, 0xfc, 0xca, 0x82, 0x64, 0xbb, 0x49, 0x48,
	0xa1, 0x65, 0x20, 0x9e, 0xc1, 0xe1, 0x2e, 0xdb, 0x37, 0x61, 0x97, 0x17, 0xe7, 0x64, 0x08, 0xf5,
	0xcd, 0x14, 0xea, 0x82, 0x6f, 0xd5, 0xa8, 0x57, 0xd5, 0x68, 0x54, 0xd5, 0xf8, 0x01, 0x07, 0x94,
	0x59, 0x21, 0x17, 0x15, 0xdc, 0x23, 0xe8, 0x6a, 0x1f, 0xdf, 0xb0, 0x6f, 0x6c, 0x72, 0x04, 0xed,
	0xe2, 0x1d, 0xe8, 0x3b, 0x49, 0x41, 0x45, 0x83, 0xfb, 0x5d, 0xf3, 0x7b, 0x3d, 0x9b, 0xcd, 0xda,
	0x7e, 0x66, 0x6f, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0x3d, 0x96, 0xc3, 0x9b, 0xe6, 0x05, 0x00,
	0x00,
}
