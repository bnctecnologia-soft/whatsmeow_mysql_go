// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.21.12
// source: waMsgApplication/WAMsgApplication.proto

package waMsgApplication

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"

	waCommon "github.com/bnctecnologia-soft/whatsmeow_mysql_go/proto/waCommon"

	_ "embed"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MessageApplication_Metadata_ThreadType int32

const (
	MessageApplication_Metadata_DEFAULT               MessageApplication_Metadata_ThreadType = 0
	MessageApplication_Metadata_VANISH_MODE           MessageApplication_Metadata_ThreadType = 1
	MessageApplication_Metadata_DISAPPEARING_MESSAGES MessageApplication_Metadata_ThreadType = 2
)

// Enum value maps for MessageApplication_Metadata_ThreadType.
var (
	MessageApplication_Metadata_ThreadType_name = map[int32]string{
		0: "DEFAULT",
		1: "VANISH_MODE",
		2: "DISAPPEARING_MESSAGES",
	}
	MessageApplication_Metadata_ThreadType_value = map[string]int32{
		"DEFAULT":               0,
		"VANISH_MODE":           1,
		"DISAPPEARING_MESSAGES": 2,
	}
)

func (x MessageApplication_Metadata_ThreadType) Enum() *MessageApplication_Metadata_ThreadType {
	p := new(MessageApplication_Metadata_ThreadType)
	*p = x
	return p
}

func (x MessageApplication_Metadata_ThreadType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MessageApplication_Metadata_ThreadType) Descriptor() protoreflect.EnumDescriptor {
	return file_waMsgApplication_WAMsgApplication_proto_enumTypes[0].Descriptor()
}

func (MessageApplication_Metadata_ThreadType) Type() protoreflect.EnumType {
	return &file_waMsgApplication_WAMsgApplication_proto_enumTypes[0]
}

func (x MessageApplication_Metadata_ThreadType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *MessageApplication_Metadata_ThreadType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = MessageApplication_Metadata_ThreadType(num)
	return nil
}

// Deprecated: Use MessageApplication_Metadata_ThreadType.Descriptor instead.
func (MessageApplication_Metadata_ThreadType) EnumDescriptor() ([]byte, []int) {
	return file_waMsgApplication_WAMsgApplication_proto_rawDescGZIP(), []int{0, 0, 0}
}

type MessageApplication struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload  *MessageApplication_Payload  `protobuf:"bytes,1,opt,name=payload" json:"payload,omitempty"`
	Metadata *MessageApplication_Metadata `protobuf:"bytes,2,opt,name=metadata" json:"metadata,omitempty"`
}

func (x *MessageApplication) Reset() {
	*x = MessageApplication{}
	mi := &file_waMsgApplication_WAMsgApplication_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MessageApplication) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageApplication) ProtoMessage() {}

func (x *MessageApplication) ProtoReflect() protoreflect.Message {
	mi := &file_waMsgApplication_WAMsgApplication_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageApplication.ProtoReflect.Descriptor instead.
func (*MessageApplication) Descriptor() ([]byte, []int) {
	return file_waMsgApplication_WAMsgApplication_proto_rawDescGZIP(), []int{0}
}

func (x *MessageApplication) GetPayload() *MessageApplication_Payload {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *MessageApplication) GetMetadata() *MessageApplication_Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type MessageApplication_Metadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Ephemeral:
	//
	//	*MessageApplication_Metadata_ChatEphemeralSetting
	//	*MessageApplication_Metadata_EphemeralSettingList
	//	*MessageApplication_Metadata_EphemeralSharedSecret
	Ephemeral                 isMessageApplication_Metadata_Ephemeral    `protobuf_oneof:"ephemeral"`
	ForwardingScore           *uint32                                    `protobuf:"varint,5,opt,name=forwardingScore" json:"forwardingScore,omitempty"`
	IsForwarded               *bool                                      `protobuf:"varint,6,opt,name=isForwarded" json:"isForwarded,omitempty"`
	BusinessMetadata          *waCommon.SubProtocol                      `protobuf:"bytes,7,opt,name=businessMetadata" json:"businessMetadata,omitempty"`
	FrankingKey               []byte                                     `protobuf:"bytes,8,opt,name=frankingKey" json:"frankingKey,omitempty"`
	FrankingVersion           *int32                                     `protobuf:"varint,9,opt,name=frankingVersion" json:"frankingVersion,omitempty"`
	QuotedMessage             *MessageApplication_Metadata_QuotedMessage `protobuf:"bytes,10,opt,name=quotedMessage" json:"quotedMessage,omitempty"`
	ThreadType                *MessageApplication_Metadata_ThreadType    `protobuf:"varint,11,opt,name=threadType,enum=WAMsgApplication.MessageApplication_Metadata_ThreadType" json:"threadType,omitempty"`
	ReadonlyMetadataDataclass *string                                    `protobuf:"bytes,12,opt,name=readonlyMetadataDataclass" json:"readonlyMetadataDataclass,omitempty"`
	GroupID                   *string                                    `protobuf:"bytes,13,opt,name=groupID" json:"groupID,omitempty"`
	GroupSize                 *uint32                                    `protobuf:"varint,14,opt,name=groupSize" json:"groupSize,omitempty"`
	GroupIndex                *uint32                                    `protobuf:"varint,15,opt,name=groupIndex" json:"groupIndex,omitempty"`
	BotResponseID             *string                                    `protobuf:"bytes,16,opt,name=botResponseID" json:"botResponseID,omitempty"`
	CollapsibleID             *string                                    `protobuf:"bytes,17,opt,name=collapsibleID" json:"collapsibleID,omitempty"`
	SecondaryOtid             *string                                    `protobuf:"bytes,18,opt,name=secondaryOtid" json:"secondaryOtid,omitempty"`
}

func (x *MessageApplication_Metadata) Reset() {
	*x = MessageApplication_Metadata{}
	mi := &file_waMsgApplication_WAMsgApplication_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MessageApplication_Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageApplication_Metadata) ProtoMessage() {}

func (x *MessageApplication_Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_waMsgApplication_WAMsgApplication_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageApplication_Metadata.ProtoReflect.Descriptor instead.
func (*MessageApplication_Metadata) Descriptor() ([]byte, []int) {
	return file_waMsgApplication_WAMsgApplication_proto_rawDescGZIP(), []int{0, 0}
}

func (m *MessageApplication_Metadata) GetEphemeral() isMessageApplication_Metadata_Ephemeral {
	if m != nil {
		return m.Ephemeral
	}
	return nil
}

func (x *MessageApplication_Metadata) GetChatEphemeralSetting() *MessageApplication_EphemeralSetting {
	if x, ok := x.GetEphemeral().(*MessageApplication_Metadata_ChatEphemeralSetting); ok {
		return x.ChatEphemeralSetting
	}
	return nil
}

func (x *MessageApplication_Metadata) GetEphemeralSettingList() *MessageApplication_Metadata_EphemeralSettingMap {
	if x, ok := x.GetEphemeral().(*MessageApplication_Metadata_EphemeralSettingList); ok {
		return x.EphemeralSettingList
	}
	return nil
}

func (x *MessageApplication_Metadata) GetEphemeralSharedSecret() []byte {
	if x, ok := x.GetEphemeral().(*MessageApplication_Metadata_EphemeralSharedSecret); ok {
		return x.EphemeralSharedSecret
	}
	return nil
}

func (x *MessageApplication_Metadata) GetForwardingScore() uint32 {
	if x != nil && x.ForwardingScore != nil {
		return *x.ForwardingScore
	}
	return 0
}

func (x *MessageApplication_Metadata) GetIsForwarded() bool {
	if x != nil && x.IsForwarded != nil {
		return *x.IsForwarded
	}
	return false
}

func (x *MessageApplication_Metadata) GetBusinessMetadata() *waCommon.SubProtocol {
	if x != nil {
		return x.BusinessMetadata
	}
	return nil
}

func (x *MessageApplication_Metadata) GetFrankingKey() []byte {
	if x != nil {
		return x.FrankingKey
	}
	return nil
}

func (x *MessageApplication_Metadata) GetFrankingVersion() int32 {
	if x != nil && x.FrankingVersion != nil {
		return *x.FrankingVersion
	}
	return 0
}

func (x *MessageApplication_Metadata) GetQuotedMessage() *MessageApplication_Metadata_QuotedMessage {
	if x != nil {
		return x.QuotedMessage
	}
	return nil
}

func (x *MessageApplication_Metadata) GetThreadType() MessageApplication_Metadata_ThreadType {
	if x != nil && x.ThreadType != nil {
		return *x.ThreadType
	}
	return MessageApplication_Metadata_DEFAULT
}

func (x *MessageApplication_Metadata) GetReadonlyMetadataDataclass() string {
	if x != nil && x.ReadonlyMetadataDataclass != nil {
		return *x.ReadonlyMetadataDataclass
	}
	return ""
}

func (x *MessageApplication_Metadata) GetGroupID() string {
	if x != nil && x.GroupID != nil {
		return *x.GroupID
	}
	return ""
}

func (x *MessageApplication_Metadata) GetGroupSize() uint32 {
	if x != nil && x.GroupSize != nil {
		return *x.GroupSize
	}
	return 0
}

func (x *MessageApplication_Metadata) GetGroupIndex() uint32 {
	if x != nil && x.GroupIndex != nil {
		return *x.GroupIndex
	}
	return 0
}

func (x *MessageApplication_Metadata) GetBotResponseID() string {
	if x != nil && x.BotResponseID != nil {
		return *x.BotResponseID
	}
	return ""
}

func (x *MessageApplication_Metadata) GetCollapsibleID() string {
	if x != nil && x.CollapsibleID != nil {
		return *x.CollapsibleID
	}
	return ""
}

func (x *MessageApplication_Metadata) GetSecondaryOtid() string {
	if x != nil && x.SecondaryOtid != nil {
		return *x.SecondaryOtid
	}
	return ""
}

type isMessageApplication_Metadata_Ephemeral interface {
	isMessageApplication_Metadata_Ephemeral()
}

type MessageApplication_Metadata_ChatEphemeralSetting struct {
	ChatEphemeralSetting *MessageApplication_EphemeralSetting `protobuf:"bytes,1,opt,name=chatEphemeralSetting,oneof"`
}

type MessageApplication_Metadata_EphemeralSettingList struct {
	EphemeralSettingList *MessageApplication_Metadata_EphemeralSettingMap `protobuf:"bytes,2,opt,name=ephemeralSettingList,oneof"`
}

type MessageApplication_Metadata_EphemeralSharedSecret struct {
	EphemeralSharedSecret []byte `protobuf:"bytes,3,opt,name=ephemeralSharedSecret,oneof"`
}

func (*MessageApplication_Metadata_ChatEphemeralSetting) isMessageApplication_Metadata_Ephemeral() {}

func (*MessageApplication_Metadata_EphemeralSettingList) isMessageApplication_Metadata_Ephemeral() {}

func (*MessageApplication_Metadata_EphemeralSharedSecret) isMessageApplication_Metadata_Ephemeral() {}

type MessageApplication_Payload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Content:
	//
	//	*MessageApplication_Payload_CoreContent
	//	*MessageApplication_Payload_Signal
	//	*MessageApplication_Payload_ApplicationData
	//	*MessageApplication_Payload_SubProtocol
	Content isMessageApplication_Payload_Content `protobuf_oneof:"content"`
}

func (x *MessageApplication_Payload) Reset() {
	*x = MessageApplication_Payload{}
	mi := &file_waMsgApplication_WAMsgApplication_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MessageApplication_Payload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageApplication_Payload) ProtoMessage() {}

func (x *MessageApplication_Payload) ProtoReflect() protoreflect.Message {
	mi := &file_waMsgApplication_WAMsgApplication_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageApplication_Payload.ProtoReflect.Descriptor instead.
func (*MessageApplication_Payload) Descriptor() ([]byte, []int) {
	return file_waMsgApplication_WAMsgApplication_proto_rawDescGZIP(), []int{0, 1}
}

func (m *MessageApplication_Payload) GetContent() isMessageApplication_Payload_Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (x *MessageApplication_Payload) GetCoreContent() *MessageApplication_Content {
	if x, ok := x.GetContent().(*MessageApplication_Payload_CoreContent); ok {
		return x.CoreContent
	}
	return nil
}

func (x *MessageApplication_Payload) GetSignal() *MessageApplication_Signal {
	if x, ok := x.GetContent().(*MessageApplication_Payload_Signal); ok {
		return x.Signal
	}
	return nil
}

func (x *MessageApplication_Payload) GetApplicationData() *MessageApplication_ApplicationData {
	if x, ok := x.GetContent().(*MessageApplication_Payload_ApplicationData); ok {
		return x.ApplicationData
	}
	return nil
}

func (x *MessageApplication_Payload) GetSubProtocol() *MessageApplication_SubProtocolPayload {
	if x, ok := x.GetContent().(*MessageApplication_Payload_SubProtocol); ok {
		return x.SubProtocol
	}
	return nil
}

type isMessageApplication_Payload_Content interface {
	isMessageApplication_Payload_Content()
}

type MessageApplication_Payload_CoreContent struct {
	CoreContent *MessageApplication_Content `protobuf:"bytes,1,opt,name=coreContent,oneof"`
}

type MessageApplication_Payload_Signal struct {
	Signal *MessageApplication_Signal `protobuf:"bytes,2,opt,name=signal,oneof"`
}

type MessageApplication_Payload_ApplicationData struct {
	ApplicationData *MessageApplication_ApplicationData `protobuf:"bytes,3,opt,name=applicationData,oneof"`
}

type MessageApplication_Payload_SubProtocol struct {
	SubProtocol *MessageApplication_SubProtocolPayload `protobuf:"bytes,4,opt,name=subProtocol,oneof"`
}

func (*MessageApplication_Payload_CoreContent) isMessageApplication_Payload_Content() {}

func (*MessageApplication_Payload_Signal) isMessageApplication_Payload_Content() {}

func (*MessageApplication_Payload_ApplicationData) isMessageApplication_Payload_Content() {}

func (*MessageApplication_Payload_SubProtocol) isMessageApplication_Payload_Content() {}

type MessageApplication_SubProtocolPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to SubProtocol:
	//
	//	*MessageApplication_SubProtocolPayload_ConsumerMessage
	//	*MessageApplication_SubProtocolPayload_BusinessMessage
	//	*MessageApplication_SubProtocolPayload_PaymentMessage
	//	*MessageApplication_SubProtocolPayload_MultiDevice
	//	*MessageApplication_SubProtocolPayload_Voip
	//	*MessageApplication_SubProtocolPayload_Armadillo
	SubProtocol isMessageApplication_SubProtocolPayload_SubProtocol `protobuf_oneof:"subProtocol"`
	FutureProof *waCommon.FutureProofBehavior                       `protobuf:"varint,1,opt,name=futureProof,enum=WACommon.FutureProofBehavior" json:"futureProof,omitempty"`
}

func (x *MessageApplication_SubProtocolPayload) Reset() {
	*x = MessageApplication_SubProtocolPayload{}
	mi := &file_waMsgApplication_WAMsgApplication_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MessageApplication_SubProtocolPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageApplication_SubProtocolPayload) ProtoMessage() {}

func (x *MessageApplication_SubProtocolPayload) ProtoReflect() protoreflect.Message {
	mi := &file_waMsgApplication_WAMsgApplication_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageApplication_SubProtocolPayload.ProtoReflect.Descriptor instead.
func (*MessageApplication_SubProtocolPayload) Descriptor() ([]byte, []int) {
	return file_waMsgApplication_WAMsgApplication_proto_rawDescGZIP(), []int{0, 2}
}

func (m *MessageApplication_SubProtocolPayload) GetSubProtocol() isMessageApplication_SubProtocolPayload_SubProtocol {
	if m != nil {
		return m.SubProtocol
	}
	return nil
}

func (x *MessageApplication_SubProtocolPayload) GetConsumerMessage() *waCommon.SubProtocol {
	if x, ok := x.GetSubProtocol().(*MessageApplication_SubProtocolPayload_ConsumerMessage); ok {
		return x.ConsumerMessage
	}
	return nil
}

func (x *MessageApplication_SubProtocolPayload) GetBusinessMessage() *waCommon.SubProtocol {
	if x, ok := x.GetSubProtocol().(*MessageApplication_SubProtocolPayload_BusinessMessage); ok {
		return x.BusinessMessage
	}
	return nil
}

func (x *MessageApplication_SubProtocolPayload) GetPaymentMessage() *waCommon.SubProtocol {
	if x, ok := x.GetSubProtocol().(*MessageApplication_SubProtocolPayload_PaymentMessage); ok {
		return x.PaymentMessage
	}
	return nil
}

func (x *MessageApplication_SubProtocolPayload) GetMultiDevice() *waCommon.SubProtocol {
	if x, ok := x.GetSubProtocol().(*MessageApplication_SubProtocolPayload_MultiDevice); ok {
		return x.MultiDevice
	}
	return nil
}

func (x *MessageApplication_SubProtocolPayload) GetVoip() *waCommon.SubProtocol {
	if x, ok := x.GetSubProtocol().(*MessageApplication_SubProtocolPayload_Voip); ok {
		return x.Voip
	}
	return nil
}

func (x *MessageApplication_SubProtocolPayload) GetArmadillo() *waCommon.SubProtocol {
	if x, ok := x.GetSubProtocol().(*MessageApplication_SubProtocolPayload_Armadillo); ok {
		return x.Armadillo
	}
	return nil
}

func (x *MessageApplication_SubProtocolPayload) GetFutureProof() waCommon.FutureProofBehavior {
	if x != nil && x.FutureProof != nil {
		return *x.FutureProof
	}
	return waCommon.FutureProofBehavior(0)
}

type isMessageApplication_SubProtocolPayload_SubProtocol interface {
	isMessageApplication_SubProtocolPayload_SubProtocol()
}

type MessageApplication_SubProtocolPayload_ConsumerMessage struct {
	ConsumerMessage *waCommon.SubProtocol `protobuf:"bytes,2,opt,name=consumerMessage,oneof"`
}

type MessageApplication_SubProtocolPayload_BusinessMessage struct {
	BusinessMessage *waCommon.SubProtocol `protobuf:"bytes,3,opt,name=businessMessage,oneof"`
}

type MessageApplication_SubProtocolPayload_PaymentMessage struct {
	PaymentMessage *waCommon.SubProtocol `protobuf:"bytes,4,opt,name=paymentMessage,oneof"`
}

type MessageApplication_SubProtocolPayload_MultiDevice struct {
	MultiDevice *waCommon.SubProtocol `protobuf:"bytes,5,opt,name=multiDevice,oneof"`
}

type MessageApplication_SubProtocolPayload_Voip struct {
	Voip *waCommon.SubProtocol `protobuf:"bytes,6,opt,name=voip,oneof"`
}

type MessageApplication_SubProtocolPayload_Armadillo struct {
	Armadillo *waCommon.SubProtocol `protobuf:"bytes,7,opt,name=armadillo,oneof"`
}

func (*MessageApplication_SubProtocolPayload_ConsumerMessage) isMessageApplication_SubProtocolPayload_SubProtocol() {
}

func (*MessageApplication_SubProtocolPayload_BusinessMessage) isMessageApplication_SubProtocolPayload_SubProtocol() {
}

func (*MessageApplication_SubProtocolPayload_PaymentMessage) isMessageApplication_SubProtocolPayload_SubProtocol() {
}

func (*MessageApplication_SubProtocolPayload_MultiDevice) isMessageApplication_SubProtocolPayload_SubProtocol() {
}

func (*MessageApplication_SubProtocolPayload_Voip) isMessageApplication_SubProtocolPayload_SubProtocol() {
}

func (*MessageApplication_SubProtocolPayload_Armadillo) isMessageApplication_SubProtocolPayload_SubProtocol() {
}

type MessageApplication_ApplicationData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MessageApplication_ApplicationData) Reset() {
	*x = MessageApplication_ApplicationData{}
	mi := &file_waMsgApplication_WAMsgApplication_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MessageApplication_ApplicationData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageApplication_ApplicationData) ProtoMessage() {}

func (x *MessageApplication_ApplicationData) ProtoReflect() protoreflect.Message {
	mi := &file_waMsgApplication_WAMsgApplication_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageApplication_ApplicationData.ProtoReflect.Descriptor instead.
func (*MessageApplication_ApplicationData) Descriptor() ([]byte, []int) {
	return file_waMsgApplication_WAMsgApplication_proto_rawDescGZIP(), []int{0, 3}
}

type MessageApplication_Signal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MessageApplication_Signal) Reset() {
	*x = MessageApplication_Signal{}
	mi := &file_waMsgApplication_WAMsgApplication_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MessageApplication_Signal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageApplication_Signal) ProtoMessage() {}

func (x *MessageApplication_Signal) ProtoReflect() protoreflect.Message {
	mi := &file_waMsgApplication_WAMsgApplication_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageApplication_Signal.ProtoReflect.Descriptor instead.
func (*MessageApplication_Signal) Descriptor() ([]byte, []int) {
	return file_waMsgApplication_WAMsgApplication_proto_rawDescGZIP(), []int{0, 4}
}

type MessageApplication_Content struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MessageApplication_Content) Reset() {
	*x = MessageApplication_Content{}
	mi := &file_waMsgApplication_WAMsgApplication_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MessageApplication_Content) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageApplication_Content) ProtoMessage() {}

func (x *MessageApplication_Content) ProtoReflect() protoreflect.Message {
	mi := &file_waMsgApplication_WAMsgApplication_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageApplication_Content.ProtoReflect.Descriptor instead.
func (*MessageApplication_Content) Descriptor() ([]byte, []int) {
	return file_waMsgApplication_WAMsgApplication_proto_rawDescGZIP(), []int{0, 5}
}

type MessageApplication_EphemeralSetting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EphemeralExpiration       *uint32 `protobuf:"varint,2,opt,name=ephemeralExpiration" json:"ephemeralExpiration,omitempty"`
	EphemeralSettingTimestamp *int64  `protobuf:"varint,3,opt,name=ephemeralSettingTimestamp" json:"ephemeralSettingTimestamp,omitempty"`
	IsEphemeralSettingReset   *bool   `protobuf:"varint,4,opt,name=isEphemeralSettingReset" json:"isEphemeralSettingReset,omitempty"`
}

func (x *MessageApplication_EphemeralSetting) Reset() {
	*x = MessageApplication_EphemeralSetting{}
	mi := &file_waMsgApplication_WAMsgApplication_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MessageApplication_EphemeralSetting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageApplication_EphemeralSetting) ProtoMessage() {}

func (x *MessageApplication_EphemeralSetting) ProtoReflect() protoreflect.Message {
	mi := &file_waMsgApplication_WAMsgApplication_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageApplication_EphemeralSetting.ProtoReflect.Descriptor instead.
func (*MessageApplication_EphemeralSetting) Descriptor() ([]byte, []int) {
	return file_waMsgApplication_WAMsgApplication_proto_rawDescGZIP(), []int{0, 6}
}

func (x *MessageApplication_EphemeralSetting) GetEphemeralExpiration() uint32 {
	if x != nil && x.EphemeralExpiration != nil {
		return *x.EphemeralExpiration
	}
	return 0
}

func (x *MessageApplication_EphemeralSetting) GetEphemeralSettingTimestamp() int64 {
	if x != nil && x.EphemeralSettingTimestamp != nil {
		return *x.EphemeralSettingTimestamp
	}
	return 0
}

func (x *MessageApplication_EphemeralSetting) GetIsEphemeralSettingReset() bool {
	if x != nil && x.IsEphemeralSettingReset != nil {
		return *x.IsEphemeralSettingReset
	}
	return false
}

type MessageApplication_Metadata_QuotedMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StanzaID    *string                     `protobuf:"bytes,1,opt,name=stanzaID" json:"stanzaID,omitempty"`
	RemoteJID   *string                     `protobuf:"bytes,2,opt,name=remoteJID" json:"remoteJID,omitempty"`
	Participant *string                     `protobuf:"bytes,3,opt,name=participant" json:"participant,omitempty"`
	Payload     *MessageApplication_Payload `protobuf:"bytes,4,opt,name=payload" json:"payload,omitempty"`
}

func (x *MessageApplication_Metadata_QuotedMessage) Reset() {
	*x = MessageApplication_Metadata_QuotedMessage{}
	mi := &file_waMsgApplication_WAMsgApplication_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MessageApplication_Metadata_QuotedMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageApplication_Metadata_QuotedMessage) ProtoMessage() {}

func (x *MessageApplication_Metadata_QuotedMessage) ProtoReflect() protoreflect.Message {
	mi := &file_waMsgApplication_WAMsgApplication_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageApplication_Metadata_QuotedMessage.ProtoReflect.Descriptor instead.
func (*MessageApplication_Metadata_QuotedMessage) Descriptor() ([]byte, []int) {
	return file_waMsgApplication_WAMsgApplication_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *MessageApplication_Metadata_QuotedMessage) GetStanzaID() string {
	if x != nil && x.StanzaID != nil {
		return *x.StanzaID
	}
	return ""
}

func (x *MessageApplication_Metadata_QuotedMessage) GetRemoteJID() string {
	if x != nil && x.RemoteJID != nil {
		return *x.RemoteJID
	}
	return ""
}

func (x *MessageApplication_Metadata_QuotedMessage) GetParticipant() string {
	if x != nil && x.Participant != nil {
		return *x.Participant
	}
	return ""
}

func (x *MessageApplication_Metadata_QuotedMessage) GetPayload() *MessageApplication_Payload {
	if x != nil {
		return x.Payload
	}
	return nil
}

type MessageApplication_Metadata_EphemeralSettingMap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatJID          *string                              `protobuf:"bytes,1,opt,name=chatJID" json:"chatJID,omitempty"`
	EphemeralSetting *MessageApplication_EphemeralSetting `protobuf:"bytes,2,opt,name=ephemeralSetting" json:"ephemeralSetting,omitempty"`
}

func (x *MessageApplication_Metadata_EphemeralSettingMap) Reset() {
	*x = MessageApplication_Metadata_EphemeralSettingMap{}
	mi := &file_waMsgApplication_WAMsgApplication_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MessageApplication_Metadata_EphemeralSettingMap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageApplication_Metadata_EphemeralSettingMap) ProtoMessage() {}

func (x *MessageApplication_Metadata_EphemeralSettingMap) ProtoReflect() protoreflect.Message {
	mi := &file_waMsgApplication_WAMsgApplication_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageApplication_Metadata_EphemeralSettingMap.ProtoReflect.Descriptor instead.
func (*MessageApplication_Metadata_EphemeralSettingMap) Descriptor() ([]byte, []int) {
	return file_waMsgApplication_WAMsgApplication_proto_rawDescGZIP(), []int{0, 0, 1}
}

func (x *MessageApplication_Metadata_EphemeralSettingMap) GetChatJID() string {
	if x != nil && x.ChatJID != nil {
		return *x.ChatJID
	}
	return ""
}

func (x *MessageApplication_Metadata_EphemeralSettingMap) GetEphemeralSetting() *MessageApplication_EphemeralSetting {
	if x != nil {
		return x.EphemeralSetting
	}
	return nil
}

var File_waMsgApplication_WAMsgApplication_proto protoreflect.FileDescriptor

//go:embed WAMsgApplication.pb.raw
var file_waMsgApplication_WAMsgApplication_proto_rawDesc []byte

var (
	file_waMsgApplication_WAMsgApplication_proto_rawDescOnce sync.Once
	file_waMsgApplication_WAMsgApplication_proto_rawDescData = file_waMsgApplication_WAMsgApplication_proto_rawDesc
)

func file_waMsgApplication_WAMsgApplication_proto_rawDescGZIP() []byte {
	file_waMsgApplication_WAMsgApplication_proto_rawDescOnce.Do(func() {
		file_waMsgApplication_WAMsgApplication_proto_rawDescData = protoimpl.X.CompressGZIP(file_waMsgApplication_WAMsgApplication_proto_rawDescData)
	})
	return file_waMsgApplication_WAMsgApplication_proto_rawDescData
}

var file_waMsgApplication_WAMsgApplication_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_waMsgApplication_WAMsgApplication_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_waMsgApplication_WAMsgApplication_proto_goTypes = []any{
	(MessageApplication_Metadata_ThreadType)(0),             // 0: WAMsgApplication.MessageApplication.Metadata.ThreadType
	(*MessageApplication)(nil),                              // 1: WAMsgApplication.MessageApplication
	(*MessageApplication_Metadata)(nil),                     // 2: WAMsgApplication.MessageApplication.Metadata
	(*MessageApplication_Payload)(nil),                      // 3: WAMsgApplication.MessageApplication.Payload
	(*MessageApplication_SubProtocolPayload)(nil),           // 4: WAMsgApplication.MessageApplication.SubProtocolPayload
	(*MessageApplication_ApplicationData)(nil),              // 5: WAMsgApplication.MessageApplication.ApplicationData
	(*MessageApplication_Signal)(nil),                       // 6: WAMsgApplication.MessageApplication.Signal
	(*MessageApplication_Content)(nil),                      // 7: WAMsgApplication.MessageApplication.Content
	(*MessageApplication_EphemeralSetting)(nil),             // 8: WAMsgApplication.MessageApplication.EphemeralSetting
	(*MessageApplication_Metadata_QuotedMessage)(nil),       // 9: WAMsgApplication.MessageApplication.Metadata.QuotedMessage
	(*MessageApplication_Metadata_EphemeralSettingMap)(nil), // 10: WAMsgApplication.MessageApplication.Metadata.EphemeralSettingMap
	(*waCommon.SubProtocol)(nil),                            // 11: WACommon.SubProtocol
	(waCommon.FutureProofBehavior)(0),                       // 12: WACommon.FutureProofBehavior
}
var file_waMsgApplication_WAMsgApplication_proto_depIdxs = []int32{
	3,  // 0: WAMsgApplication.MessageApplication.payload:type_name -> WAMsgApplication.MessageApplication.Payload
	2,  // 1: WAMsgApplication.MessageApplication.metadata:type_name -> WAMsgApplication.MessageApplication.Metadata
	8,  // 2: WAMsgApplication.MessageApplication.Metadata.chatEphemeralSetting:type_name -> WAMsgApplication.MessageApplication.EphemeralSetting
	10, // 3: WAMsgApplication.MessageApplication.Metadata.ephemeralSettingList:type_name -> WAMsgApplication.MessageApplication.Metadata.EphemeralSettingMap
	11, // 4: WAMsgApplication.MessageApplication.Metadata.businessMetadata:type_name -> WACommon.SubProtocol
	9,  // 5: WAMsgApplication.MessageApplication.Metadata.quotedMessage:type_name -> WAMsgApplication.MessageApplication.Metadata.QuotedMessage
	0,  // 6: WAMsgApplication.MessageApplication.Metadata.threadType:type_name -> WAMsgApplication.MessageApplication.Metadata.ThreadType
	7,  // 7: WAMsgApplication.MessageApplication.Payload.coreContent:type_name -> WAMsgApplication.MessageApplication.Content
	6,  // 8: WAMsgApplication.MessageApplication.Payload.signal:type_name -> WAMsgApplication.MessageApplication.Signal
	5,  // 9: WAMsgApplication.MessageApplication.Payload.applicationData:type_name -> WAMsgApplication.MessageApplication.ApplicationData
	4,  // 10: WAMsgApplication.MessageApplication.Payload.subProtocol:type_name -> WAMsgApplication.MessageApplication.SubProtocolPayload
	11, // 11: WAMsgApplication.MessageApplication.SubProtocolPayload.consumerMessage:type_name -> WACommon.SubProtocol
	11, // 12: WAMsgApplication.MessageApplication.SubProtocolPayload.businessMessage:type_name -> WACommon.SubProtocol
	11, // 13: WAMsgApplication.MessageApplication.SubProtocolPayload.paymentMessage:type_name -> WACommon.SubProtocol
	11, // 14: WAMsgApplication.MessageApplication.SubProtocolPayload.multiDevice:type_name -> WACommon.SubProtocol
	11, // 15: WAMsgApplication.MessageApplication.SubProtocolPayload.voip:type_name -> WACommon.SubProtocol
	11, // 16: WAMsgApplication.MessageApplication.SubProtocolPayload.armadillo:type_name -> WACommon.SubProtocol
	12, // 17: WAMsgApplication.MessageApplication.SubProtocolPayload.futureProof:type_name -> WACommon.FutureProofBehavior
	3,  // 18: WAMsgApplication.MessageApplication.Metadata.QuotedMessage.payload:type_name -> WAMsgApplication.MessageApplication.Payload
	8,  // 19: WAMsgApplication.MessageApplication.Metadata.EphemeralSettingMap.ephemeralSetting:type_name -> WAMsgApplication.MessageApplication.EphemeralSetting
	20, // [20:20] is the sub-list for method output_type
	20, // [20:20] is the sub-list for method input_type
	20, // [20:20] is the sub-list for extension type_name
	20, // [20:20] is the sub-list for extension extendee
	0,  // [0:20] is the sub-list for field type_name
}

func init() { file_waMsgApplication_WAMsgApplication_proto_init() }
func file_waMsgApplication_WAMsgApplication_proto_init() {
	if File_waMsgApplication_WAMsgApplication_proto != nil {
		return
	}
	file_waMsgApplication_WAMsgApplication_proto_msgTypes[1].OneofWrappers = []any{
		(*MessageApplication_Metadata_ChatEphemeralSetting)(nil),
		(*MessageApplication_Metadata_EphemeralSettingList)(nil),
		(*MessageApplication_Metadata_EphemeralSharedSecret)(nil),
	}
	file_waMsgApplication_WAMsgApplication_proto_msgTypes[2].OneofWrappers = []any{
		(*MessageApplication_Payload_CoreContent)(nil),
		(*MessageApplication_Payload_Signal)(nil),
		(*MessageApplication_Payload_ApplicationData)(nil),
		(*MessageApplication_Payload_SubProtocol)(nil),
	}
	file_waMsgApplication_WAMsgApplication_proto_msgTypes[3].OneofWrappers = []any{
		(*MessageApplication_SubProtocolPayload_ConsumerMessage)(nil),
		(*MessageApplication_SubProtocolPayload_BusinessMessage)(nil),
		(*MessageApplication_SubProtocolPayload_PaymentMessage)(nil),
		(*MessageApplication_SubProtocolPayload_MultiDevice)(nil),
		(*MessageApplication_SubProtocolPayload_Voip)(nil),
		(*MessageApplication_SubProtocolPayload_Armadillo)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_waMsgApplication_WAMsgApplication_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_waMsgApplication_WAMsgApplication_proto_goTypes,
		DependencyIndexes: file_waMsgApplication_WAMsgApplication_proto_depIdxs,
		EnumInfos:         file_waMsgApplication_WAMsgApplication_proto_enumTypes,
		MessageInfos:      file_waMsgApplication_WAMsgApplication_proto_msgTypes,
	}.Build()
	File_waMsgApplication_WAMsgApplication_proto = out.File
	file_waMsgApplication_WAMsgApplication_proto_rawDesc = nil
	file_waMsgApplication_WAMsgApplication_proto_goTypes = nil
	file_waMsgApplication_WAMsgApplication_proto_depIdxs = nil
}
