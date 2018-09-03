// Code generated by protoc-gen-go. DO NOT EDIT.
// source: coprocess_mini_request_object.proto

package coprocess

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type MiniRequestObject struct {
	Headers         map[string]string `protobuf:"bytes,1,rep,name=headers" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	SetHeaders      map[string]string `protobuf:"bytes,2,rep,name=set_headers,json=setHeaders" json:"set_headers,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	DeleteHeaders   []string          `protobuf:"bytes,3,rep,name=delete_headers,json=deleteHeaders" json:"delete_headers,omitempty"`
	Body            string            `protobuf:"bytes,4,opt,name=body" json:"body,omitempty"`
	Url             string            `protobuf:"bytes,5,opt,name=url" json:"url,omitempty"`
	Params          map[string]string `protobuf:"bytes,6,rep,name=params" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	AddParams       map[string]string `protobuf:"bytes,7,rep,name=add_params,json=addParams" json:"add_params,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	ExtendedParams  map[string]string `protobuf:"bytes,8,rep,name=extended_params,json=extendedParams" json:"extended_params,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	DeleteParams    []string          `protobuf:"bytes,9,rep,name=delete_params,json=deleteParams" json:"delete_params,omitempty"`
	ReturnOverrides *ReturnOverrides  `protobuf:"bytes,10,opt,name=return_overrides,json=returnOverrides" json:"return_overrides,omitempty"`
	Method          string            `protobuf:"bytes,11,opt,name=method" json:"method,omitempty"`
	RequestUri      string            `protobuf:"bytes,12,opt,name=request_uri,json=requestUri" json:"request_uri,omitempty"`
	Scheme          string            `protobuf:"bytes,13,opt,name=scheme" json:"scheme,omitempty"`
	RawBody         []byte            `protobuf:"bytes,14,opt,name=raw_body,json=rawBody,proto3" json:"raw_body,omitempty"`
}

func (m *MiniRequestObject) Reset()                    { *m = MiniRequestObject{} }
func (m *MiniRequestObject) String() string            { return proto.CompactTextString(m) }
func (*MiniRequestObject) ProtoMessage()               {}
func (*MiniRequestObject) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *MiniRequestObject) GetHeaders() map[string]string {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *MiniRequestObject) GetSetHeaders() map[string]string {
	if m != nil {
		return m.SetHeaders
	}
	return nil
}

func (m *MiniRequestObject) GetDeleteHeaders() []string {
	if m != nil {
		return m.DeleteHeaders
	}
	return nil
}

func (m *MiniRequestObject) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *MiniRequestObject) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *MiniRequestObject) GetParams() map[string]string {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *MiniRequestObject) GetAddParams() map[string]string {
	if m != nil {
		return m.AddParams
	}
	return nil
}

func (m *MiniRequestObject) GetExtendedParams() map[string]string {
	if m != nil {
		return m.ExtendedParams
	}
	return nil
}

func (m *MiniRequestObject) GetDeleteParams() []string {
	if m != nil {
		return m.DeleteParams
	}
	return nil
}

func (m *MiniRequestObject) GetReturnOverrides() *ReturnOverrides {
	if m != nil {
		return m.ReturnOverrides
	}
	return nil
}

func (m *MiniRequestObject) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *MiniRequestObject) GetRequestUri() string {
	if m != nil {
		return m.RequestUri
	}
	return ""
}

func (m *MiniRequestObject) GetScheme() string {
	if m != nil {
		return m.Scheme
	}
	return ""
}

func (m *MiniRequestObject) GetRawBody() []byte {
	if m != nil {
		return m.RawBody
	}
	return nil
}

func init() {
	proto.RegisterType((*MiniRequestObject)(nil), "coprocess.MiniRequestObject")
}

func init() { proto.RegisterFile("coprocess_mini_request_object.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 452 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xd1, 0x8b, 0xd3, 0x40,
	0x10, 0xc6, 0x49, 0x7b, 0xd7, 0x5e, 0x26, 0xbd, 0xf6, 0x5c, 0x45, 0xd6, 0xbe, 0x18, 0x3c, 0x84,
	0x88, 0x52, 0xe4, 0x7c, 0xd1, 0x43, 0xc1, 0x53, 0x0a, 0x22, 0x1c, 0x27, 0x11, 0x1f, 0x7c, 0x0a,
	0xdb, 0xee, 0x40, 0xe3, 0xb5, 0xd9, 0x3a, 0xbb, 0xb9, 0x33, 0xff, 0xb3, 0x7f, 0x84, 0x74, 0x37,
	0x89, 0xe9, 0x29, 0x81, 0xbc, 0xed, 0x7c, 0xf9, 0xbe, 0x1f, 0xc3, 0xcc, 0x04, 0x4e, 0x97, 0x6a,
	0x4b, 0x6a, 0x89, 0x5a, 0x27, 0x9b, 0x34, 0x4b, 0x13, 0xc2, 0x9f, 0x39, 0x6a, 0x93, 0xa8, 0xc5,
	0x0f, 0x5c, 0x9a, 0xd9, 0x96, 0x94, 0x51, 0xcc, 0xaf, 0x4d, 0xd3, 0xf0, 0xaf, 0x9f, 0xd0, 0xe4,
	0x94, 0x25, 0xea, 0x06, 0x89, 0x52, 0x89, 0xda, 0x99, 0x9f, 0xfc, 0x1e, 0xc2, 0xbd, 0xcb, 0x34,
	0x4b, 0x63, 0x47, 0xba, 0xb2, 0x20, 0xf6, 0x11, 0x86, 0x2b, 0x14, 0x12, 0x49, 0x73, 0x2f, 0xec,
	0x47, 0xc1, 0xd9, 0xb3, 0x59, 0x4d, 0x9a, 0xfd, 0x63, 0x9f, 0x7d, 0x72, 0xde, 0x79, 0x66, 0xa8,
	0x88, 0xab, 0x24, 0xbb, 0x84, 0x40, 0xa3, 0x49, 0x2a, 0x50, 0xcf, 0x82, 0x5e, 0xb4, 0x82, 0xbe,
	0xa2, 0xd9, 0x63, 0x81, 0xae, 0x05, 0xf6, 0x14, 0xc6, 0x12, 0xd7, 0x68, 0xb0, 0x26, 0xf6, 0xc3,
	0x7e, 0xe4, 0xc7, 0xc7, 0x4e, 0xad, 0x6c, 0x0c, 0x0e, 0x16, 0x4a, 0x16, 0xfc, 0x20, 0xf4, 0x22,
	0x3f, 0xb6, 0x6f, 0x76, 0x02, 0xfd, 0x9c, 0xd6, 0xfc, 0xd0, 0x4a, 0xbb, 0x27, 0x7b, 0x0f, 0x83,
	0xad, 0x20, 0xb1, 0xd1, 0x7c, 0x60, 0xdb, 0x8a, 0x5a, 0xdb, 0xfa, 0x62, 0xad, 0xae, 0xa5, 0x32,
	0xc7, 0x3e, 0x03, 0x08, 0x29, 0x93, 0x92, 0x32, 0xb4, 0x94, 0xe7, 0xad, 0x94, 0x0b, 0x29, 0x9b,
	0x20, 0x5f, 0x54, 0x35, 0xfb, 0x0e, 0x13, 0xfc, 0x65, 0x30, 0x93, 0x58, 0x03, 0x8f, 0x2c, 0xf0,
	0x65, 0x2b, 0x70, 0x5e, 0x66, 0x9a, 0xd4, 0x31, 0xee, 0x89, 0xec, 0x14, 0xca, 0xf9, 0x54, 0x60,
	0xdf, 0x0e, 0x6d, 0xe4, 0xc4, 0xd2, 0x34, 0x87, 0x93, 0xbb, 0xe7, 0xc1, 0x21, 0xf4, 0xa2, 0xe0,
	0x6c, 0xda, 0x68, 0x20, 0xb6, 0x96, 0xab, 0xca, 0x11, 0x4f, 0x68, 0x5f, 0x60, 0x0f, 0x61, 0xb0,
	0x41, 0xb3, 0x52, 0x92, 0x07, 0x76, 0xd2, 0x65, 0xc5, 0x1e, 0x43, 0x50, 0x1d, 0x6a, 0x4e, 0x29,
	0x1f, 0xd9, 0x8f, 0x50, 0x4a, 0xdf, 0x28, 0xdd, 0x05, 0xf5, 0x72, 0x85, 0x1b, 0xe4, 0xc7, 0x2e,
	0xe8, 0x2a, 0xf6, 0x08, 0x8e, 0x48, 0xdc, 0x26, 0x76, 0x9f, 0xe3, 0xd0, 0x8b, 0x46, 0xf1, 0x90,
	0xc4, 0xed, 0x07, 0x25, 0x8b, 0xe9, 0x39, 0x8c, 0x9a, 0x97, 0xb2, 0x5b, 0xf1, 0x35, 0x16, 0xdc,
	0x73, 0x2b, 0xbe, 0xc6, 0x82, 0x3d, 0x80, 0xc3, 0x1b, 0xb1, 0xce, 0x91, 0xf7, 0xac, 0xe6, 0x8a,
	0xf3, 0xde, 0x6b, 0x6f, 0xfa, 0x0e, 0x26, 0x77, 0x0e, 0xad, 0x53, 0xfc, 0x0d, 0x04, 0x8d, 0x89,
	0x77, 0x8a, 0xbe, 0x85, 0xf1, 0xfe, 0x15, 0x74, 0x4a, 0x5f, 0xc0, 0xfd, 0xff, 0xac, 0xbc, 0x0b,
	0x62, 0x31, 0xb0, 0x7f, 0xfd, 0xab, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa0, 0xe0, 0x90, 0xb7,
	0x49, 0x04, 0x00, 0x00,
}