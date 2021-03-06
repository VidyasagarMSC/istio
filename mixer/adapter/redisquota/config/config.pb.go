// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mixer/adapter/redisquota/config/config.proto

/*
	Package config is a generated protocol buffer package.

	The `redisquota` adapter can be used to support Istio's quota management
	system. It depends on a Redis server to store quota values.

	This adapter supports the [quota template](https://istio.io/docs/reference/config/policy-and-telemetry/templates/quota.html).

	It is generated from these files:
		mixer/adapter/redisquota/config/config.proto

	It has these top-level messages:
		Params
*/
package config

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/types"
import _ "github.com/gogo/protobuf/gogoproto"

import time "time"

import strconv "strconv"

import types "github.com/gogo/protobuf/types"

import strings "strings"
import reflect "reflect"
import sortkeys "github.com/gogo/protobuf/sortkeys"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Algorithms for rate-limiting:
type Params_QuotaAlgorithm int32

const (
	// FIXED_WINDOW The fixed window approach can allow 2x peak specified rate, whereas the rolling-window doesn't.
	FIXED_WINDOW Params_QuotaAlgorithm = 0
	// ROLLING_WINDOW The rolling window algorithm's additional precision comes at the cost of increased redis resource usage.
	ROLLING_WINDOW Params_QuotaAlgorithm = 1
)

var Params_QuotaAlgorithm_name = map[int32]string{
	0: "FIXED_WINDOW",
	1: "ROLLING_WINDOW",
}
var Params_QuotaAlgorithm_value = map[string]int32{
	"FIXED_WINDOW":   0,
	"ROLLING_WINDOW": 1,
}

func (Params_QuotaAlgorithm) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorConfig, []int{0, 0}
}

// redisquota adapter supports the rate limit quota using either fixed or
// rolling window algorithm. And it is using Redis as a shared data storage.
//
// Example configuration:
//
// ```yaml
// redisServerUrl: localhost:6379
// connectionPoolSize: 10
// quotas:
//   - name: requestCount.quota.istio-system
//     maxAmount: 50
//     validDuration: 60s
//     bucketDuration: 1s
//     rateLimitAlgorithm: ROLLING_WINDOW
//     overrides:
//       - dimensions:
//           destination: ratings
//           source: reviews
//         maxAmount: 12
//       - dimensions:
//           destination: reviews
//         maxAmount: 5
// ```
type Params struct {
	// The set of known quotas. At least one quota configuration is required
	Quotas []Params_Quota `protobuf:"bytes,1,rep,name=quotas" json:"quotas"`
	// Redis connection string <hostname>:<port number>
	// ex) localhost:6379
	RedisServerUrl string `protobuf:"bytes,2,opt,name=redis_server_url,json=redisServerUrl,proto3" json:"redis_server_url,omitempty"`
	// Maximum number of idle connections to redis
	// Default is 10 connections per every CPU as reported by runtime.NumCPU.
	ConnectionPoolSize int64 `protobuf:"varint,3,opt,name=connection_pool_size,json=connectionPoolSize,proto3" json:"connection_pool_size,omitempty"`
}

func (m *Params) Reset()                    { *m = Params{} }
func (*Params) ProtoMessage()               {}
func (*Params) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{0} }

type Params_Override struct {
	// The specific dimensions for which this override applies.
	// String representation of instance dimensions is used to check against configured dimensions.
	// dimensions should not be empty
	Dimensions map[string]string `protobuf:"bytes,1,rep,name=dimensions" json:"dimensions" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// The upper limit for this quota override.
	// This value should be bigger than 0
	MaxAmount int64 `protobuf:"varint,2,opt,name=max_amount,json=maxAmount,proto3" json:"max_amount,omitempty"`
}

func (m *Params_Override) Reset()                    { *m = Params_Override{} }
func (*Params_Override) ProtoMessage()               {}
func (*Params_Override) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{0, 0} }

func (m *Params_Override) GetDimensions() map[string]string {
	if m != nil {
		return m.Dimensions
	}
	return nil
}

func (m *Params_Override) GetMaxAmount() int64 {
	if m != nil {
		return m.MaxAmount
	}
	return 0
}

type Params_Quota struct {
	// The name of the quota
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The upper limit for this quota. max_amount should be bigger than 0
	MaxAmount int64 `protobuf:"varint,2,opt,name=max_amount,json=maxAmount,proto3" json:"max_amount,omitempty"`
	// The amount of time allocated quota remains valid before it is
	// automatically released. This is only meaningful for rate limit quotas.
	// value should be 0 < valid_duration
	ValidDuration time.Duration `protobuf:"bytes,3,opt,name=valid_duration,json=validDuration,stdduration" json:"valid_duration"`
	// bucket_duration will be ignored if rate_limit_algorithm is FIXED_WINDOW
	// value should be 0 < bucket_duration < valid_duration
	BucketDuration time.Duration `protobuf:"bytes,4,opt,name=bucket_duration,json=bucketDuration,stdduration" json:"bucket_duration"`
	// Quota management algorithm. The default value is FIXED_WINDOW
	RateLimitAlgorithm Params_QuotaAlgorithm `protobuf:"varint,5,opt,name=rate_limit_algorithm,json=rateLimitAlgorithm,proto3,enum=adapter.redisquota.config.Params_QuotaAlgorithm" json:"rate_limit_algorithm,omitempty"`
	// Overrides associated with this quota.
	// The first matching override is applied.
	Overrides []*Params_Override `protobuf:"bytes,6,rep,name=overrides" json:"overrides,omitempty"`
}

func (m *Params_Quota) Reset()                    { *m = Params_Quota{} }
func (*Params_Quota) ProtoMessage()               {}
func (*Params_Quota) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{0, 1} }

func (m *Params_Quota) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Params_Quota) GetMaxAmount() int64 {
	if m != nil {
		return m.MaxAmount
	}
	return 0
}

func (m *Params_Quota) GetValidDuration() time.Duration {
	if m != nil {
		return m.ValidDuration
	}
	return 0
}

func (m *Params_Quota) GetBucketDuration() time.Duration {
	if m != nil {
		return m.BucketDuration
	}
	return 0
}

func (m *Params_Quota) GetRateLimitAlgorithm() Params_QuotaAlgorithm {
	if m != nil {
		return m.RateLimitAlgorithm
	}
	return FIXED_WINDOW
}

func (m *Params_Quota) GetOverrides() []*Params_Override {
	if m != nil {
		return m.Overrides
	}
	return nil
}

func init() {
	proto.RegisterType((*Params)(nil), "adapter.redisquota.config.Params")
	proto.RegisterType((*Params_Override)(nil), "adapter.redisquota.config.Params.Override")
	proto.RegisterType((*Params_Quota)(nil), "adapter.redisquota.config.Params.Quota")
	proto.RegisterEnum("adapter.redisquota.config.Params_QuotaAlgorithm", Params_QuotaAlgorithm_name, Params_QuotaAlgorithm_value)
}
func (x Params_QuotaAlgorithm) String() string {
	s, ok := Params_QuotaAlgorithm_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Quotas) > 0 {
		for _, msg := range m.Quotas {
			dAtA[i] = 0xa
			i++
			i = encodeVarintConfig(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.RedisServerUrl) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.RedisServerUrl)))
		i += copy(dAtA[i:], m.RedisServerUrl)
	}
	if m.ConnectionPoolSize != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintConfig(dAtA, i, uint64(m.ConnectionPoolSize))
	}
	return i, nil
}

func (m *Params_Override) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params_Override) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Dimensions) > 0 {
		for k, _ := range m.Dimensions {
			dAtA[i] = 0xa
			i++
			v := m.Dimensions[k]
			mapSize := 1 + len(k) + sovConfig(uint64(len(k))) + 1 + len(v) + sovConfig(uint64(len(v)))
			i = encodeVarintConfig(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintConfig(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x12
			i++
			i = encodeVarintConfig(dAtA, i, uint64(len(v)))
			i += copy(dAtA[i:], v)
		}
	}
	if m.MaxAmount != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintConfig(dAtA, i, uint64(m.MaxAmount))
	}
	return i, nil
}

func (m *Params_Quota) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params_Quota) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if m.MaxAmount != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintConfig(dAtA, i, uint64(m.MaxAmount))
	}
	dAtA[i] = 0x1a
	i++
	i = encodeVarintConfig(dAtA, i, uint64(types.SizeOfStdDuration(m.ValidDuration)))
	n1, err := types.StdDurationMarshalTo(m.ValidDuration, dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	dAtA[i] = 0x22
	i++
	i = encodeVarintConfig(dAtA, i, uint64(types.SizeOfStdDuration(m.BucketDuration)))
	n2, err := types.StdDurationMarshalTo(m.BucketDuration, dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	if m.RateLimitAlgorithm != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintConfig(dAtA, i, uint64(m.RateLimitAlgorithm))
	}
	if len(m.Overrides) > 0 {
		for _, msg := range m.Overrides {
			dAtA[i] = 0x32
			i++
			i = encodeVarintConfig(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeVarintConfig(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Params) Size() (n int) {
	var l int
	_ = l
	if len(m.Quotas) > 0 {
		for _, e := range m.Quotas {
			l = e.Size()
			n += 1 + l + sovConfig(uint64(l))
		}
	}
	l = len(m.RedisServerUrl)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	if m.ConnectionPoolSize != 0 {
		n += 1 + sovConfig(uint64(m.ConnectionPoolSize))
	}
	return n
}

func (m *Params_Override) Size() (n int) {
	var l int
	_ = l
	if len(m.Dimensions) > 0 {
		for k, v := range m.Dimensions {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovConfig(uint64(len(k))) + 1 + len(v) + sovConfig(uint64(len(v)))
			n += mapEntrySize + 1 + sovConfig(uint64(mapEntrySize))
		}
	}
	if m.MaxAmount != 0 {
		n += 1 + sovConfig(uint64(m.MaxAmount))
	}
	return n
}

func (m *Params_Quota) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	if m.MaxAmount != 0 {
		n += 1 + sovConfig(uint64(m.MaxAmount))
	}
	l = types.SizeOfStdDuration(m.ValidDuration)
	n += 1 + l + sovConfig(uint64(l))
	l = types.SizeOfStdDuration(m.BucketDuration)
	n += 1 + l + sovConfig(uint64(l))
	if m.RateLimitAlgorithm != 0 {
		n += 1 + sovConfig(uint64(m.RateLimitAlgorithm))
	}
	if len(m.Overrides) > 0 {
		for _, e := range m.Overrides {
			l = e.Size()
			n += 1 + l + sovConfig(uint64(l))
		}
	}
	return n
}

func sovConfig(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozConfig(x uint64) (n int) {
	return sovConfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Params) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Params{`,
		`Quotas:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.Quotas), "Params_Quota", "Params_Quota", 1), `&`, ``, 1) + `,`,
		`RedisServerUrl:` + fmt.Sprintf("%v", this.RedisServerUrl) + `,`,
		`ConnectionPoolSize:` + fmt.Sprintf("%v", this.ConnectionPoolSize) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Params_Override) String() string {
	if this == nil {
		return "nil"
	}
	keysForDimensions := make([]string, 0, len(this.Dimensions))
	for k, _ := range this.Dimensions {
		keysForDimensions = append(keysForDimensions, k)
	}
	sortkeys.Strings(keysForDimensions)
	mapStringForDimensions := "map[string]string{"
	for _, k := range keysForDimensions {
		mapStringForDimensions += fmt.Sprintf("%v: %v,", k, this.Dimensions[k])
	}
	mapStringForDimensions += "}"
	s := strings.Join([]string{`&Params_Override{`,
		`Dimensions:` + mapStringForDimensions + `,`,
		`MaxAmount:` + fmt.Sprintf("%v", this.MaxAmount) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Params_Quota) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Params_Quota{`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`MaxAmount:` + fmt.Sprintf("%v", this.MaxAmount) + `,`,
		`ValidDuration:` + strings.Replace(strings.Replace(this.ValidDuration.String(), "Duration", "google_protobuf.Duration", 1), `&`, ``, 1) + `,`,
		`BucketDuration:` + strings.Replace(strings.Replace(this.BucketDuration.String(), "Duration", "google_protobuf.Duration", 1), `&`, ``, 1) + `,`,
		`RateLimitAlgorithm:` + fmt.Sprintf("%v", this.RateLimitAlgorithm) + `,`,
		`Overrides:` + strings.Replace(fmt.Sprintf("%v", this.Overrides), "Params_Override", "Params_Override", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringConfig(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Quotas", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Quotas = append(m.Quotas, Params_Quota{})
			if err := m.Quotas[len(m.Quotas)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RedisServerUrl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RedisServerUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConnectionPoolSize", wireType)
			}
			m.ConnectionPoolSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ConnectionPoolSize |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Params_Override) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Override: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Override: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Dimensions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Dimensions == nil {
				m.Dimensions = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowConfig
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowConfig
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthConfig
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowConfig
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthConfig
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipConfig(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthConfig
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Dimensions[mapkey] = mapvalue
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxAmount", wireType)
			}
			m.MaxAmount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxAmount |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Params_Quota) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Quota: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Quota: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxAmount", wireType)
			}
			m.MaxAmount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxAmount |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidDuration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := types.StdDurationUnmarshal(&m.ValidDuration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BucketDuration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := types.StdDurationUnmarshal(&m.BucketDuration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RateLimitAlgorithm", wireType)
			}
			m.RateLimitAlgorithm = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RateLimitAlgorithm |= (Params_QuotaAlgorithm(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Overrides", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Overrides = append(m.Overrides, &Params_Override{})
			if err := m.Overrides[len(m.Overrides)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipConfig(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConfig
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthConfig
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowConfig
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipConfig(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthConfig = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConfig   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("mixer/adapter/redisquota/config/config.proto", fileDescriptorConfig) }

var fileDescriptorConfig = []byte{
	// 570 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xb1, 0x6f, 0xd3, 0x4e,
	0x18, 0xf5, 0xd5, 0xa9, 0xd5, 0x5e, 0x7f, 0x3f, 0x37, 0x3a, 0x65, 0x70, 0x23, 0x71, 0x8d, 0xba,
	0x10, 0x21, 0x64, 0x57, 0x45, 0x42, 0x55, 0x25, 0x86, 0x56, 0x09, 0x10, 0x14, 0x35, 0xc5, 0x15,
	0x2a, 0x62, 0x31, 0x97, 0xf8, 0x6a, 0x4e, 0xb5, 0x7d, 0xe1, 0x6c, 0x47, 0x69, 0x27, 0xc6, 0x8e,
	0x8c, 0x8c, 0x8c, 0xfc, 0x1b, 0x6c, 0x19, 0x33, 0x32, 0x01, 0x31, 0x0b, 0x63, 0xff, 0x04, 0xe4,
	0xb3, 0x9d, 0x00, 0x12, 0x6a, 0x26, 0x7f, 0x7e, 0xf7, 0xde, 0xbb, 0xef, 0xfb, 0xde, 0xc1, 0xfb,
	0x01, 0x1b, 0x53, 0x61, 0x11, 0x97, 0x0c, 0x63, 0x2a, 0x2c, 0x41, 0x5d, 0x16, 0xbd, 0x4d, 0x78,
	0x4c, 0xac, 0x01, 0x0f, 0xcf, 0x99, 0x57, 0x7c, 0xcc, 0xa1, 0xe0, 0x31, 0x47, 0x5b, 0x05, 0xcf,
	0x5c, 0xf0, 0xcc, 0x9c, 0x50, 0xc7, 0x1e, 0xe7, 0x9e, 0x4f, 0x2d, 0x49, 0xec, 0x27, 0xe7, 0x96,
	0x9b, 0x08, 0x12, 0x33, 0x1e, 0xe6, 0xd2, 0x7a, 0xcd, 0xe3, 0x1e, 0x97, 0xa5, 0x95, 0x55, 0x39,
	0xba, 0xf3, 0x59, 0x83, 0xda, 0x09, 0x11, 0x24, 0x88, 0x50, 0x1b, 0x6a, 0xd2, 0x30, 0x32, 0x40,
	0x43, 0x6d, 0x6e, 0xec, 0xdd, 0x35, 0xff, 0x79, 0x99, 0x99, 0x4b, 0xcc, 0xe7, 0x19, 0x76, 0x54,
	0x99, 0x7c, 0xdd, 0x56, 0xec, 0x42, 0x8c, 0x9a, 0xb0, 0x2a, 0xf9, 0x4e, 0x44, 0xc5, 0x88, 0x0a,
	0x27, 0x11, 0xbe, 0xb1, 0xd2, 0x00, 0xcd, 0x75, 0x5b, 0x97, 0xf8, 0xa9, 0x84, 0x5f, 0x08, 0x1f,
	0xed, 0xc2, 0xda, 0x80, 0x87, 0x21, 0x1d, 0x64, 0x5d, 0x3a, 0x43, 0xce, 0x7d, 0x27, 0x62, 0x57,
	0xd4, 0x50, 0x1b, 0xa0, 0xa9, 0xda, 0x68, 0x71, 0x76, 0xc2, 0xb9, 0x7f, 0xca, 0xae, 0x68, 0x7d,
	0x0a, 0xe0, 0x5a, 0x6f, 0x44, 0x85, 0x60, 0x2e, 0x45, 0xaf, 0x21, 0x74, 0x59, 0x40, 0xc3, 0x88,
	0xf1, 0xb0, 0xec, 0xf9, 0xe0, 0xf6, 0x9e, 0x4b, 0xbd, 0xd9, 0x9a, 0x8b, 0xdb, 0x61, 0x2c, 0x2e,
	0x8b, 0x31, 0x7e, 0xf3, 0x44, 0x77, 0x20, 0x0c, 0xc8, 0xd8, 0x21, 0x01, 0x4f, 0xc2, 0x58, 0x0e,
	0xa1, 0xda, 0xeb, 0x01, 0x19, 0x1f, 0x4a, 0xa0, 0xfe, 0x08, 0x6e, 0xfe, 0xe5, 0x81, 0xaa, 0x50,
	0xbd, 0xa0, 0x97, 0x06, 0x90, 0xf3, 0x66, 0x25, 0xaa, 0xc1, 0xd5, 0x11, 0xf1, 0x13, 0x5a, 0xec,
	0x20, 0xff, 0x39, 0x58, 0xd9, 0x07, 0x07, 0x95, 0xeb, 0x8f, 0xdb, 0xa0, 0x7e, 0xad, 0xc2, 0x55,
	0xb9, 0x46, 0x84, 0x60, 0x25, 0x24, 0x01, 0x2d, 0xc4, 0xb2, 0xbe, 0xa5, 0x03, 0xf4, 0x0c, 0xea,
	0x23, 0xe2, 0x33, 0xd7, 0x29, 0xb3, 0x96, 0xbb, 0xdb, 0xd8, 0xdb, 0x32, 0xf3, 0xc7, 0x60, 0x96,
	0x8f, 0xc1, 0x6c, 0x15, 0x84, 0xa3, 0xb5, 0x6c, 0xca, 0x0f, 0xdf, 0xb6, 0x81, 0xfd, 0xbf, 0x94,
	0x96, 0x07, 0xa8, 0x0b, 0x37, 0xfb, 0xc9, 0xe0, 0x82, 0xc6, 0x0b, 0xb3, 0xca, 0xf2, 0x66, 0x7a,
	0xae, 0x9d, 0xbb, 0xf5, 0x61, 0x4d, 0x90, 0x98, 0x3a, 0x3e, 0x0b, 0x58, 0xec, 0x10, 0xdf, 0xe3,
	0x82, 0xc5, 0x6f, 0x02, 0x63, 0xb5, 0x01, 0x9a, 0xfa, 0xde, 0xee, 0x92, 0x4f, 0xeb, 0xb0, 0xd4,
	0xd9, 0x28, 0x73, 0xeb, 0x66, 0x66, 0x73, 0x0c, 0x3d, 0x85, 0xeb, 0xbc, 0x08, 0x33, 0x32, 0x34,
	0x99, 0xff, 0xbd, 0xe5, 0xf3, 0xb7, 0x17, 0xe2, 0x3c, 0x8a, 0x9d, 0x87, 0x50, 0xff, 0xf3, 0x56,
	0x54, 0x85, 0xff, 0x3d, 0xee, 0xbc, 0x6c, 0xb7, 0x9c, 0xb3, 0xce, 0x71, 0xab, 0x77, 0x56, 0x55,
	0x10, 0x82, 0xba, 0xdd, 0xeb, 0x76, 0x3b, 0xc7, 0x4f, 0x4a, 0x0c, 0x1c, 0xed, 0x4f, 0x66, 0x58,
	0x99, 0xce, 0xb0, 0xf2, 0x65, 0x86, 0x95, 0x9b, 0x19, 0x56, 0xde, 0xa5, 0x18, 0x7c, 0x4a, 0xb1,
	0x32, 0x49, 0x31, 0x98, 0xa6, 0x18, 0x7c, 0x4f, 0x31, 0xf8, 0x99, 0x62, 0xe5, 0x26, 0xc5, 0xe0,
	0xfd, 0x0f, 0xac, 0xbc, 0xd2, 0xf2, 0x96, 0xfa, 0x9a, 0x5c, 0xe9, 0x83, 0x5f, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x82, 0x98, 0x9b, 0xbf, 0x05, 0x04, 0x00, 0x00,
}
