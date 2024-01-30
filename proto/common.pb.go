// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: common.proto

package investapi

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Тип инструмента.
type InstrumentType int32

const (
	InstrumentType_INSTRUMENT_TYPE_UNSPECIFIED          InstrumentType = 0
	InstrumentType_INSTRUMENT_TYPE_BOND                 InstrumentType = 1  //Облигация.
	InstrumentType_INSTRUMENT_TYPE_SHARE                InstrumentType = 2  //Акция.
	InstrumentType_INSTRUMENT_TYPE_CURRENCY             InstrumentType = 3  //Валюта.
	InstrumentType_INSTRUMENT_TYPE_ETF                  InstrumentType = 4  //Exchange-traded fund. Фонд.
	InstrumentType_INSTRUMENT_TYPE_FUTURES              InstrumentType = 5  //Фьючерс.
	InstrumentType_INSTRUMENT_TYPE_SP                   InstrumentType = 6  //Структурная нота.
	InstrumentType_INSTRUMENT_TYPE_OPTION               InstrumentType = 7  //Опцион.
	InstrumentType_INSTRUMENT_TYPE_CLEARING_CERTIFICATE InstrumentType = 8  //Clearing certificate.
	InstrumentType_INSTRUMENT_TYPE_INDEX                InstrumentType = 9  //Индекс.
	InstrumentType_INSTRUMENT_TYPE_COMMODITY            InstrumentType = 10 //Товар.
)

// Enum value maps for InstrumentType.
var (
	InstrumentType_name = map[int32]string{
		0:  "INSTRUMENT_TYPE_UNSPECIFIED",
		1:  "INSTRUMENT_TYPE_BOND",
		2:  "INSTRUMENT_TYPE_SHARE",
		3:  "INSTRUMENT_TYPE_CURRENCY",
		4:  "INSTRUMENT_TYPE_ETF",
		5:  "INSTRUMENT_TYPE_FUTURES",
		6:  "INSTRUMENT_TYPE_SP",
		7:  "INSTRUMENT_TYPE_OPTION",
		8:  "INSTRUMENT_TYPE_CLEARING_CERTIFICATE",
		9:  "INSTRUMENT_TYPE_INDEX",
		10: "INSTRUMENT_TYPE_COMMODITY",
	}
	InstrumentType_value = map[string]int32{
		"INSTRUMENT_TYPE_UNSPECIFIED":          0,
		"INSTRUMENT_TYPE_BOND":                 1,
		"INSTRUMENT_TYPE_SHARE":                2,
		"INSTRUMENT_TYPE_CURRENCY":             3,
		"INSTRUMENT_TYPE_ETF":                  4,
		"INSTRUMENT_TYPE_FUTURES":              5,
		"INSTRUMENT_TYPE_SP":                   6,
		"INSTRUMENT_TYPE_OPTION":               7,
		"INSTRUMENT_TYPE_CLEARING_CERTIFICATE": 8,
		"INSTRUMENT_TYPE_INDEX":                9,
		"INSTRUMENT_TYPE_COMMODITY":            10,
	}
)

func (x InstrumentType) Enum() *InstrumentType {
	p := new(InstrumentType)
	*p = x
	return p
}

func (x InstrumentType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InstrumentType) Descriptor() protoreflect.EnumDescriptor {
	return file_common_proto_enumTypes[0].Descriptor()
}

func (InstrumentType) Type() protoreflect.EnumType {
	return &file_common_proto_enumTypes[0]
}

func (x InstrumentType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InstrumentType.Descriptor instead.
func (InstrumentType) EnumDescriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{0}
}

// Режим торгов инструмента
type SecurityTradingStatus int32

const (
	SecurityTradingStatus_SECURITY_TRADING_STATUS_UNSPECIFIED                      SecurityTradingStatus = 0  //Торговый статус не определён
	SecurityTradingStatus_SECURITY_TRADING_STATUS_NOT_AVAILABLE_FOR_TRADING        SecurityTradingStatus = 1  //Недоступен для торгов
	SecurityTradingStatus_SECURITY_TRADING_STATUS_OPENING_PERIOD                   SecurityTradingStatus = 2  //Период открытия торгов
	SecurityTradingStatus_SECURITY_TRADING_STATUS_CLOSING_PERIOD                   SecurityTradingStatus = 3  //Период закрытия торгов
	SecurityTradingStatus_SECURITY_TRADING_STATUS_BREAK_IN_TRADING                 SecurityTradingStatus = 4  //Перерыв в торговле
	SecurityTradingStatus_SECURITY_TRADING_STATUS_NORMAL_TRADING                   SecurityTradingStatus = 5  //Нормальная торговля
	SecurityTradingStatus_SECURITY_TRADING_STATUS_CLOSING_AUCTION                  SecurityTradingStatus = 6  //Аукцион закрытия
	SecurityTradingStatus_SECURITY_TRADING_STATUS_DARK_POOL_AUCTION                SecurityTradingStatus = 7  //Аукцион крупных пакетов
	SecurityTradingStatus_SECURITY_TRADING_STATUS_DISCRETE_AUCTION                 SecurityTradingStatus = 8  //Дискретный аукцион
	SecurityTradingStatus_SECURITY_TRADING_STATUS_OPENING_AUCTION_PERIOD           SecurityTradingStatus = 9  //Аукцион открытия
	SecurityTradingStatus_SECURITY_TRADING_STATUS_TRADING_AT_CLOSING_AUCTION_PRICE SecurityTradingStatus = 10 //Период торгов по цене аукциона закрытия
	SecurityTradingStatus_SECURITY_TRADING_STATUS_SESSION_ASSIGNED                 SecurityTradingStatus = 11 //Сессия назначена
	SecurityTradingStatus_SECURITY_TRADING_STATUS_SESSION_CLOSE                    SecurityTradingStatus = 12 //Сессия закрыта
	SecurityTradingStatus_SECURITY_TRADING_STATUS_SESSION_OPEN                     SecurityTradingStatus = 13 //Сессия открыта
	SecurityTradingStatus_SECURITY_TRADING_STATUS_DEALER_NORMAL_TRADING            SecurityTradingStatus = 14 //Доступна торговля в режиме внутренней ликвидности брокера
	SecurityTradingStatus_SECURITY_TRADING_STATUS_DEALER_BREAK_IN_TRADING          SecurityTradingStatus = 15 //Перерыв торговли в режиме внутренней ликвидности брокера
	SecurityTradingStatus_SECURITY_TRADING_STATUS_DEALER_NOT_AVAILABLE_FOR_TRADING SecurityTradingStatus = 16 //Недоступна торговля в режиме внутренней ликвидности брокера
)

// Enum value maps for SecurityTradingStatus.
var (
	SecurityTradingStatus_name = map[int32]string{
		0:  "SECURITY_TRADING_STATUS_UNSPECIFIED",
		1:  "SECURITY_TRADING_STATUS_NOT_AVAILABLE_FOR_TRADING",
		2:  "SECURITY_TRADING_STATUS_OPENING_PERIOD",
		3:  "SECURITY_TRADING_STATUS_CLOSING_PERIOD",
		4:  "SECURITY_TRADING_STATUS_BREAK_IN_TRADING",
		5:  "SECURITY_TRADING_STATUS_NORMAL_TRADING",
		6:  "SECURITY_TRADING_STATUS_CLOSING_AUCTION",
		7:  "SECURITY_TRADING_STATUS_DARK_POOL_AUCTION",
		8:  "SECURITY_TRADING_STATUS_DISCRETE_AUCTION",
		9:  "SECURITY_TRADING_STATUS_OPENING_AUCTION_PERIOD",
		10: "SECURITY_TRADING_STATUS_TRADING_AT_CLOSING_AUCTION_PRICE",
		11: "SECURITY_TRADING_STATUS_SESSION_ASSIGNED",
		12: "SECURITY_TRADING_STATUS_SESSION_CLOSE",
		13: "SECURITY_TRADING_STATUS_SESSION_OPEN",
		14: "SECURITY_TRADING_STATUS_DEALER_NORMAL_TRADING",
		15: "SECURITY_TRADING_STATUS_DEALER_BREAK_IN_TRADING",
		16: "SECURITY_TRADING_STATUS_DEALER_NOT_AVAILABLE_FOR_TRADING",
	}
	SecurityTradingStatus_value = map[string]int32{
		"SECURITY_TRADING_STATUS_UNSPECIFIED":                      0,
		"SECURITY_TRADING_STATUS_NOT_AVAILABLE_FOR_TRADING":        1,
		"SECURITY_TRADING_STATUS_OPENING_PERIOD":                   2,
		"SECURITY_TRADING_STATUS_CLOSING_PERIOD":                   3,
		"SECURITY_TRADING_STATUS_BREAK_IN_TRADING":                 4,
		"SECURITY_TRADING_STATUS_NORMAL_TRADING":                   5,
		"SECURITY_TRADING_STATUS_CLOSING_AUCTION":                  6,
		"SECURITY_TRADING_STATUS_DARK_POOL_AUCTION":                7,
		"SECURITY_TRADING_STATUS_DISCRETE_AUCTION":                 8,
		"SECURITY_TRADING_STATUS_OPENING_AUCTION_PERIOD":           9,
		"SECURITY_TRADING_STATUS_TRADING_AT_CLOSING_AUCTION_PRICE": 10,
		"SECURITY_TRADING_STATUS_SESSION_ASSIGNED":                 11,
		"SECURITY_TRADING_STATUS_SESSION_CLOSE":                    12,
		"SECURITY_TRADING_STATUS_SESSION_OPEN":                     13,
		"SECURITY_TRADING_STATUS_DEALER_NORMAL_TRADING":            14,
		"SECURITY_TRADING_STATUS_DEALER_BREAK_IN_TRADING":          15,
		"SECURITY_TRADING_STATUS_DEALER_NOT_AVAILABLE_FOR_TRADING": 16,
	}
)

func (x SecurityTradingStatus) Enum() *SecurityTradingStatus {
	p := new(SecurityTradingStatus)
	*p = x
	return p
}

func (x SecurityTradingStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SecurityTradingStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_common_proto_enumTypes[1].Descriptor()
}

func (SecurityTradingStatus) Type() protoreflect.EnumType {
	return &file_common_proto_enumTypes[1]
}

func (x SecurityTradingStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SecurityTradingStatus.Descriptor instead.
func (SecurityTradingStatus) EnumDescriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{1}
}

// Тип цены.
type PriceType int32

const (
	PriceType_PRICE_TYPE_UNSPECIFIED PriceType = 0 //Значение не определено.
	PriceType_PRICE_TYPE_POINT       PriceType = 1 //Цена в пунктах (только для фьючерсов и облигаций).
	PriceType_PRICE_TYPE_CURRENCY    PriceType = 2 //Цена в валюте расчётов по инструменту.
)

// Enum value maps for PriceType.
var (
	PriceType_name = map[int32]string{
		0: "PRICE_TYPE_UNSPECIFIED",
		1: "PRICE_TYPE_POINT",
		2: "PRICE_TYPE_CURRENCY",
	}
	PriceType_value = map[string]int32{
		"PRICE_TYPE_UNSPECIFIED": 0,
		"PRICE_TYPE_POINT":       1,
		"PRICE_TYPE_CURRENCY":    2,
	}
)

func (x PriceType) Enum() *PriceType {
	p := new(PriceType)
	*p = x
	return p
}

func (x PriceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PriceType) Descriptor() protoreflect.EnumDescriptor {
	return file_common_proto_enumTypes[2].Descriptor()
}

func (PriceType) Type() protoreflect.EnumType {
	return &file_common_proto_enumTypes[2]
}

func (x PriceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PriceType.Descriptor instead.
func (PriceType) EnumDescriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{2}
}

// Денежная сумма в определенной валюте
type MoneyValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// строковый ISO-код валюты
	Currency string `protobuf:"bytes,1,opt,name=currency,proto3" json:"currency,omitempty"`
	// целая часть суммы, может быть отрицательным числом
	Units int64 `protobuf:"varint,2,opt,name=units,proto3" json:"units,omitempty"`
	// дробная часть суммы, может быть отрицательным числом
	Nano int32 `protobuf:"varint,3,opt,name=nano,proto3" json:"nano,omitempty"`
}

func (x *MoneyValue) Reset() {
	*x = MoneyValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MoneyValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MoneyValue) ProtoMessage() {}

func (x *MoneyValue) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MoneyValue.ProtoReflect.Descriptor instead.
func (*MoneyValue) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{0}
}

func (x *MoneyValue) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *MoneyValue) GetUnits() int64 {
	if x != nil {
		return x.Units
	}
	return 0
}

func (x *MoneyValue) GetNano() int32 {
	if x != nil {
		return x.Nano
	}
	return 0
}

// Котировка — денежная сумма без указания валюты
type Quotation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// целая часть суммы, может быть отрицательным числом
	Units int64 `protobuf:"varint,1,opt,name=units,proto3" json:"units,omitempty"`
	// дробная часть суммы, может быть отрицательным числом
	Nano int32 `protobuf:"varint,2,opt,name=nano,proto3" json:"nano,omitempty"`
}

func (x *Quotation) Reset() {
	*x = Quotation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Quotation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Quotation) ProtoMessage() {}

func (x *Quotation) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Quotation.ProtoReflect.Descriptor instead.
func (*Quotation) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{1}
}

func (x *Quotation) GetUnits() int64 {
	if x != nil {
		return x.Units
	}
	return 0
}

func (x *Quotation) GetNano() int32 {
	if x != nil {
		return x.Nano
	}
	return 0
}

// Проверка активности стрима.
type Ping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time     *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`                         //Время проверки.
	StreamId string                 `protobuf:"bytes,2,opt,name=stream_id,json=streamId,proto3" json:"stream_id,omitempty"` //Идентификатор соединения
}

func (x *Ping) Reset() {
	*x = Ping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ping) ProtoMessage() {}

func (x *Ping) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ping.ProtoReflect.Descriptor instead.
func (*Ping) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{2}
}

func (x *Ping) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *Ping) GetStreamId() string {
	if x != nil {
		return x.StreamId
	}
	return ""
}

type Page struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit      int32 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`                             //Максимальное число возвращаемых записей.
	PageNumber int32 `protobuf:"varint,2,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"` //Порядковый номер страницы, начиная с 0.
}

func (x *Page) Reset() {
	*x = Page{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Page) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Page) ProtoMessage() {}

func (x *Page) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Page.ProtoReflect.Descriptor instead.
func (*Page) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{3}
}

func (x *Page) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *Page) GetPageNumber() int32 {
	if x != nil {
		return x.PageNumber
	}
	return 0
}

type PageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit      int32 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`                             //Максимальное число возвращаемых записей.
	PageNumber int32 `protobuf:"varint,2,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"` //Порядковый номер страницы, начиная с 0.
	TotalCount int32 `protobuf:"varint,3,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"` //Общее количество записей.
}

func (x *PageResponse) Reset() {
	*x = PageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageResponse) ProtoMessage() {}

func (x *PageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageResponse.ProtoReflect.Descriptor instead.
func (*PageResponse) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{4}
}

func (x *PageResponse) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *PageResponse) GetPageNumber() int32 {
	if x != nil {
		return x.PageNumber
	}
	return 0
}

func (x *PageResponse) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

type ResponseMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TrackingId string                 `protobuf:"bytes,42,opt,name=tracking_id,json=trackingId,proto3" json:"tracking_id,omitempty"` //Идентификатор трекинга
	ServerTime *timestamppb.Timestamp `protobuf:"bytes,43,opt,name=server_time,json=serverTime,proto3" json:"server_time,omitempty"` //Серверное время
}

func (x *ResponseMetadata) Reset() {
	*x = ResponseMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseMetadata) ProtoMessage() {}

func (x *ResponseMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseMetadata.ProtoReflect.Descriptor instead.
func (*ResponseMetadata) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{5}
}

func (x *ResponseMetadata) GetTrackingId() string {
	if x != nil {
		return x.TrackingId
	}
	return ""
}

func (x *ResponseMetadata) GetServerTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ServerTime
	}
	return nil
}

var File_common_proto protoreflect.FileDescriptor

var file_common_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x25,
	0x74, 0x69, 0x6e, 0x6b, 0x6f, 0x66, 0x66, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x69,
	0x6e, 0x76, 0x65, 0x73, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x52, 0x0a, 0x0a, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6e, 0x6f, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6e, 0x61, 0x6e, 0x6f, 0x22, 0x35, 0x0a, 0x09, 0x51, 0x75,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x75, 0x6e, 0x69, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6e, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6e, 0x61, 0x6e,
	0x6f, 0x22, 0x53, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x22, 0x3d, 0x0a, 0x04, 0x50, 0x61, 0x67, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x66, 0x0a, 0x0c, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x70, 0x0a,
	0x10, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64,
	0x18, 0x2a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67,
	0x49, 0x64, 0x12, 0x3b, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x2b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x2a,
	0xd2, 0x02, 0x0a, 0x0e, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x1f, 0x0a, 0x1b, 0x49, 0x4e, 0x53, 0x54, 0x52, 0x55, 0x4d, 0x45, 0x4e, 0x54,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x49, 0x4e, 0x53, 0x54, 0x52, 0x55, 0x4d, 0x45, 0x4e,
	0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x42, 0x4f, 0x4e, 0x44, 0x10, 0x01, 0x12, 0x19, 0x0a,
	0x15, 0x49, 0x4e, 0x53, 0x54, 0x52, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x53, 0x48, 0x41, 0x52, 0x45, 0x10, 0x02, 0x12, 0x1c, 0x0a, 0x18, 0x49, 0x4e, 0x53, 0x54,
	0x52, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x55, 0x52, 0x52,
	0x45, 0x4e, 0x43, 0x59, 0x10, 0x03, 0x12, 0x17, 0x0a, 0x13, 0x49, 0x4e, 0x53, 0x54, 0x52, 0x55,
	0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x45, 0x54, 0x46, 0x10, 0x04, 0x12,
	0x1b, 0x0a, 0x17, 0x49, 0x4e, 0x53, 0x54, 0x52, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x46, 0x55, 0x54, 0x55, 0x52, 0x45, 0x53, 0x10, 0x05, 0x12, 0x16, 0x0a, 0x12,
	0x49, 0x4e, 0x53, 0x54, 0x52, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x53, 0x50, 0x10, 0x06, 0x12, 0x1a, 0x0a, 0x16, 0x49, 0x4e, 0x53, 0x54, 0x52, 0x55, 0x4d, 0x45,
	0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4f, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x07,
	0x12, 0x28, 0x0a, 0x24, 0x49, 0x4e, 0x53, 0x54, 0x52, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x43, 0x4c, 0x45, 0x41, 0x52, 0x49, 0x4e, 0x47, 0x5f, 0x43, 0x45, 0x52,
	0x54, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x45, 0x10, 0x08, 0x12, 0x19, 0x0a, 0x15, 0x49, 0x4e,
	0x53, 0x54, 0x52, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e,
	0x44, 0x45, 0x58, 0x10, 0x09, 0x12, 0x1d, 0x0a, 0x19, 0x49, 0x4e, 0x53, 0x54, 0x52, 0x55, 0x4d,
	0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x4f, 0x44, 0x49,
	0x54, 0x59, 0x10, 0x0a, 0x2a, 0xce, 0x06, 0x0a, 0x15, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74,
	0x79, 0x54, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x27,
	0x0a, 0x23, 0x53, 0x45, 0x43, 0x55, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x54, 0x52, 0x41, 0x44, 0x49,
	0x4e, 0x47, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x35, 0x0a, 0x31, 0x53, 0x45, 0x43, 0x55, 0x52,
	0x49, 0x54, 0x59, 0x5f, 0x54, 0x52, 0x41, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45,
	0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x54, 0x52, 0x41, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x2a,
	0x0a, 0x26, 0x53, 0x45, 0x43, 0x55, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x54, 0x52, 0x41, 0x44, 0x49,
	0x4e, 0x47, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4f, 0x50, 0x45, 0x4e, 0x49, 0x4e,
	0x47, 0x5f, 0x50, 0x45, 0x52, 0x49, 0x4f, 0x44, 0x10, 0x02, 0x12, 0x2a, 0x0a, 0x26, 0x53, 0x45,
	0x43, 0x55, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x54, 0x52, 0x41, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x4c, 0x4f, 0x53, 0x49, 0x4e, 0x47, 0x5f, 0x50, 0x45,
	0x52, 0x49, 0x4f, 0x44, 0x10, 0x03, 0x12, 0x2c, 0x0a, 0x28, 0x53, 0x45, 0x43, 0x55, 0x52, 0x49,
	0x54, 0x59, 0x5f, 0x54, 0x52, 0x41, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x5f, 0x42, 0x52, 0x45, 0x41, 0x4b, 0x5f, 0x49, 0x4e, 0x5f, 0x54, 0x52, 0x41, 0x44, 0x49,
	0x4e, 0x47, 0x10, 0x04, 0x12, 0x2a, 0x0a, 0x26, 0x53, 0x45, 0x43, 0x55, 0x52, 0x49, 0x54, 0x59,
	0x5f, 0x54, 0x52, 0x41, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f,
	0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x5f, 0x54, 0x52, 0x41, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x05,
	0x12, 0x2b, 0x0a, 0x27, 0x53, 0x45, 0x43, 0x55, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x54, 0x52, 0x41,
	0x44, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x4c, 0x4f, 0x53,
	0x49, 0x4e, 0x47, 0x5f, 0x41, 0x55, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x06, 0x12, 0x2d, 0x0a,
	0x29, 0x53, 0x45, 0x43, 0x55, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x54, 0x52, 0x41, 0x44, 0x49, 0x4e,
	0x47, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x44, 0x41, 0x52, 0x4b, 0x5f, 0x50, 0x4f,
	0x4f, 0x4c, 0x5f, 0x41, 0x55, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x07, 0x12, 0x2c, 0x0a, 0x28,
	0x53, 0x45, 0x43, 0x55, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x54, 0x52, 0x41, 0x44, 0x49, 0x4e, 0x47,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x44, 0x49, 0x53, 0x43, 0x52, 0x45, 0x54, 0x45,
	0x5f, 0x41, 0x55, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x08, 0x12, 0x32, 0x0a, 0x2e, 0x53, 0x45,
	0x43, 0x55, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x54, 0x52, 0x41, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4f, 0x50, 0x45, 0x4e, 0x49, 0x4e, 0x47, 0x5f, 0x41, 0x55,
	0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x50, 0x45, 0x52, 0x49, 0x4f, 0x44, 0x10, 0x09, 0x12, 0x3c,
	0x0a, 0x38, 0x53, 0x45, 0x43, 0x55, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x54, 0x52, 0x41, 0x44, 0x49,
	0x4e, 0x47, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x54, 0x52, 0x41, 0x44, 0x49, 0x4e,
	0x47, 0x5f, 0x41, 0x54, 0x5f, 0x43, 0x4c, 0x4f, 0x53, 0x49, 0x4e, 0x47, 0x5f, 0x41, 0x55, 0x43,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x50, 0x52, 0x49, 0x43, 0x45, 0x10, 0x0a, 0x12, 0x2c, 0x0a, 0x28,
	0x53, 0x45, 0x43, 0x55, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x54, 0x52, 0x41, 0x44, 0x49, 0x4e, 0x47,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f,
	0x41, 0x53, 0x53, 0x49, 0x47, 0x4e, 0x45, 0x44, 0x10, 0x0b, 0x12, 0x29, 0x0a, 0x25, 0x53, 0x45,
	0x43, 0x55, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x54, 0x52, 0x41, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x4c,
	0x4f, 0x53, 0x45, 0x10, 0x0c, 0x12, 0x28, 0x0a, 0x24, 0x53, 0x45, 0x43, 0x55, 0x52, 0x49, 0x54,
	0x59, 0x5f, 0x54, 0x52, 0x41, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x5f, 0x53, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x4f, 0x50, 0x45, 0x4e, 0x10, 0x0d, 0x12,
	0x31, 0x0a, 0x2d, 0x53, 0x45, 0x43, 0x55, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x54, 0x52, 0x41, 0x44,
	0x49, 0x4e, 0x47, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x44, 0x45, 0x41, 0x4c, 0x45,
	0x52, 0x5f, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x5f, 0x54, 0x52, 0x41, 0x44, 0x49, 0x4e, 0x47,
	0x10, 0x0e, 0x12, 0x33, 0x0a, 0x2f, 0x53, 0x45, 0x43, 0x55, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x54,
	0x52, 0x41, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x44, 0x45,
	0x41, 0x4c, 0x45, 0x52, 0x5f, 0x42, 0x52, 0x45, 0x41, 0x4b, 0x5f, 0x49, 0x4e, 0x5f, 0x54, 0x52,
	0x41, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x0f, 0x12, 0x3c, 0x0a, 0x38, 0x53, 0x45, 0x43, 0x55, 0x52,
	0x49, 0x54, 0x59, 0x5f, 0x54, 0x52, 0x41, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x44, 0x45, 0x41, 0x4c, 0x45, 0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x56,
	0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x54, 0x52, 0x41, 0x44,
	0x49, 0x4e, 0x47, 0x10, 0x10, 0x2a, 0x56, 0x0a, 0x09, 0x50, 0x72, 0x69, 0x63, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x1a, 0x0a, 0x16, 0x50, 0x52, 0x49, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x14,
	0x0a, 0x10, 0x50, 0x52, 0x49, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x4f, 0x49,
	0x4e, 0x54, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x50, 0x52, 0x49, 0x43, 0x45, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x43, 0x55, 0x52, 0x52, 0x45, 0x4e, 0x43, 0x59, 0x10, 0x02, 0x42, 0x61, 0x0a,
	0x1c, 0x72, 0x75, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x6f, 0x66, 0x66, 0x2e, 0x70, 0x69, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a,
	0x0c, 0x2e, 0x2f, 0x3b, 0x69, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x61, 0x70, 0x69, 0xa2, 0x02, 0x05,
	0x54, 0x49, 0x41, 0x50, 0x49, 0xaa, 0x02, 0x14, 0x54, 0x69, 0x6e, 0x6b, 0x6f, 0x66, 0x66, 0x2e,
	0x49, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x41, 0x70, 0x69, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x11, 0x54,
	0x69, 0x6e, 0x6b, 0x6f, 0x66, 0x66, 0x5c, 0x49, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x5c, 0x56, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_proto_rawDescOnce sync.Once
	file_common_proto_rawDescData = file_common_proto_rawDesc
)

func file_common_proto_rawDescGZIP() []byte {
	file_common_proto_rawDescOnce.Do(func() {
		file_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_proto_rawDescData)
	})
	return file_common_proto_rawDescData
}

var file_common_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_common_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_common_proto_goTypes = []interface{}{
	(InstrumentType)(0),           // 0: tinkoff.public.invest.api.contract.v1.InstrumentType
	(SecurityTradingStatus)(0),    // 1: tinkoff.public.invest.api.contract.v1.SecurityTradingStatus
	(PriceType)(0),                // 2: tinkoff.public.invest.api.contract.v1.PriceType
	(*MoneyValue)(nil),            // 3: tinkoff.public.invest.api.contract.v1.MoneyValue
	(*Quotation)(nil),             // 4: tinkoff.public.invest.api.contract.v1.Quotation
	(*Ping)(nil),                  // 5: tinkoff.public.invest.api.contract.v1.Ping
	(*Page)(nil),                  // 6: tinkoff.public.invest.api.contract.v1.Page
	(*PageResponse)(nil),          // 7: tinkoff.public.invest.api.contract.v1.PageResponse
	(*ResponseMetadata)(nil),      // 8: tinkoff.public.invest.api.contract.v1.ResponseMetadata
	(*timestamppb.Timestamp)(nil), // 9: google.protobuf.Timestamp
}
var file_common_proto_depIdxs = []int32{
	9, // 0: tinkoff.public.invest.api.contract.v1.Ping.time:type_name -> google.protobuf.Timestamp
	9, // 1: tinkoff.public.invest.api.contract.v1.ResponseMetadata.server_time:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_common_proto_init() }
func file_common_proto_init() {
	if File_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MoneyValue); i {
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
		file_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Quotation); i {
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
		file_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ping); i {
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
		file_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Page); i {
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
		file_common_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageResponse); i {
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
		file_common_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseMetadata); i {
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
			RawDescriptor: file_common_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_proto_goTypes,
		DependencyIndexes: file_common_proto_depIdxs,
		EnumInfos:         file_common_proto_enumTypes,
		MessageInfos:      file_common_proto_msgTypes,
	}.Build()
	File_common_proto = out.File
	file_common_proto_rawDesc = nil
	file_common_proto_goTypes = nil
	file_common_proto_depIdxs = nil
}
