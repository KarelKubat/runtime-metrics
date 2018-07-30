# api
--
    import "github.com/KarelKubat/runtime-metrics/api"

api is the protobuf-derived network API that the reporter's server and client
employ. This package is not for public consumption.

## Usage

#### func  RegisterReporterServer

```go
func RegisterReporterServer(s *grpc.Server, srv ReporterServer)
```

#### type AllNamesRequest

```go
type AllNamesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```

Discovery of all the names

#### func (*AllNamesRequest) Descriptor

```go
func (*AllNamesRequest) Descriptor() ([]byte, []int)
```

#### func (*AllNamesRequest) ProtoMessage

```go
func (*AllNamesRequest) ProtoMessage()
```

#### func (*AllNamesRequest) Reset

```go
func (m *AllNamesRequest) Reset()
```

#### func (*AllNamesRequest) String

```go
func (m *AllNamesRequest) String() string
```

#### func (*AllNamesRequest) XXX_DiscardUnknown

```go
func (m *AllNamesRequest) XXX_DiscardUnknown()
```

#### func (*AllNamesRequest) XXX_Marshal

```go
func (m *AllNamesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*AllNamesRequest) XXX_Merge

```go
func (dst *AllNamesRequest) XXX_Merge(src proto.Message)
```

#### func (*AllNamesRequest) XXX_Size

```go
func (m *AllNamesRequest) XXX_Size() int
```

#### func (*AllNamesRequest) XXX_Unmarshal

```go
func (m *AllNamesRequest) XXX_Unmarshal(b []byte) error
```

#### type AllNamesResponse

```go
type AllNamesResponse struct {
	AverageNames            []string `protobuf:"bytes,1,rep,name=AverageNames,proto3" json:"AverageNames,omitempty"`
	AveragePerDurationNames []string `protobuf:"bytes,2,rep,name=AveragePerDurationNames,proto3" json:"AveragePerDurationNames,omitempty"`
	CountNames              []string `protobuf:"bytes,3,rep,name=CountNames,proto3" json:"CountNames,omitempty"`
	CountPerDurationNames   []string `protobuf:"bytes,4,rep,name=CountPerDurationNames,proto3" json:"CountPerDurationNames,omitempty"`
	SumNames                []string `protobuf:"bytes,5,rep,name=SumNames,proto3" json:"SumNames,omitempty"`
	SumPerDurationNames     []string `protobuf:"bytes,6,rep,name=SumPerDurationNames,proto3" json:"SumPerDurationNames,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}
```


#### func (*AllNamesResponse) Descriptor

```go
func (*AllNamesResponse) Descriptor() ([]byte, []int)
```

#### func (*AllNamesResponse) GetAverageNames

```go
func (m *AllNamesResponse) GetAverageNames() []string
```

#### func (*AllNamesResponse) GetAveragePerDurationNames

```go
func (m *AllNamesResponse) GetAveragePerDurationNames() []string
```

#### func (*AllNamesResponse) GetCountNames

```go
func (m *AllNamesResponse) GetCountNames() []string
```

#### func (*AllNamesResponse) GetCountPerDurationNames

```go
func (m *AllNamesResponse) GetCountPerDurationNames() []string
```

#### func (*AllNamesResponse) GetSumNames

```go
func (m *AllNamesResponse) GetSumNames() []string
```

#### func (*AllNamesResponse) GetSumPerDurationNames

```go
func (m *AllNamesResponse) GetSumPerDurationNames() []string
```

#### func (*AllNamesResponse) ProtoMessage

```go
func (*AllNamesResponse) ProtoMessage()
```

#### func (*AllNamesResponse) Reset

```go
func (m *AllNamesResponse) Reset()
```

#### func (*AllNamesResponse) String

```go
func (m *AllNamesResponse) String() string
```

#### func (*AllNamesResponse) XXX_DiscardUnknown

```go
func (m *AllNamesResponse) XXX_DiscardUnknown()
```

#### func (*AllNamesResponse) XXX_Marshal

```go
func (m *AllNamesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*AllNamesResponse) XXX_Merge

```go
func (dst *AllNamesResponse) XXX_Merge(src proto.Message)
```

#### func (*AllNamesResponse) XXX_Size

```go
func (m *AllNamesResponse) XXX_Size() int
```

#### func (*AllNamesResponse) XXX_Unmarshal

```go
func (m *AllNamesResponse) XXX_Unmarshal(b []byte) error
```

#### type AveragePerDurationResponse

```go
type AveragePerDurationResponse struct {
	Average              float64              `protobuf:"fixed64,1,opt,name=Average,proto3" json:"Average,omitempty"`
	N                    int64                `protobuf:"varint,2,opt,name=N,proto3" json:"N,omitempty"`
	Until                *timestamp.Timestamp `protobuf:"bytes,3,opt,name=Until,proto3" json:"Until,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}
```

Responding to an average per duration

#### func (*AveragePerDurationResponse) Descriptor

```go
func (*AveragePerDurationResponse) Descriptor() ([]byte, []int)
```

#### func (*AveragePerDurationResponse) GetAverage

```go
func (m *AveragePerDurationResponse) GetAverage() float64
```

#### func (*AveragePerDurationResponse) GetN

```go
func (m *AveragePerDurationResponse) GetN() int64
```

#### func (*AveragePerDurationResponse) GetUntil

```go
func (m *AveragePerDurationResponse) GetUntil() *timestamp.Timestamp
```

#### func (*AveragePerDurationResponse) ProtoMessage

```go
func (*AveragePerDurationResponse) ProtoMessage()
```

#### func (*AveragePerDurationResponse) Reset

```go
func (m *AveragePerDurationResponse) Reset()
```

#### func (*AveragePerDurationResponse) String

```go
func (m *AveragePerDurationResponse) String() string
```

#### func (*AveragePerDurationResponse) XXX_DiscardUnknown

```go
func (m *AveragePerDurationResponse) XXX_DiscardUnknown()
```

#### func (*AveragePerDurationResponse) XXX_Marshal

```go
func (m *AveragePerDurationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*AveragePerDurationResponse) XXX_Merge

```go
func (dst *AveragePerDurationResponse) XXX_Merge(src proto.Message)
```

#### func (*AveragePerDurationResponse) XXX_Size

```go
func (m *AveragePerDurationResponse) XXX_Size() int
```

#### func (*AveragePerDurationResponse) XXX_Unmarshal

```go
func (m *AveragePerDurationResponse) XXX_Unmarshal(b []byte) error
```

#### type AverageResponse

```go
type AverageResponse struct {
	Average              float64  `protobuf:"fixed64,1,opt,name=Average,proto3" json:"Average,omitempty"`
	N                    int64    `protobuf:"varint,2,opt,name=N,proto3" json:"N,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```

Responding to an average

#### func (*AverageResponse) Descriptor

```go
func (*AverageResponse) Descriptor() ([]byte, []int)
```

#### func (*AverageResponse) GetAverage

```go
func (m *AverageResponse) GetAverage() float64
```

#### func (*AverageResponse) GetN

```go
func (m *AverageResponse) GetN() int64
```

#### func (*AverageResponse) ProtoMessage

```go
func (*AverageResponse) ProtoMessage()
```

#### func (*AverageResponse) Reset

```go
func (m *AverageResponse) Reset()
```

#### func (*AverageResponse) String

```go
func (m *AverageResponse) String() string
```

#### func (*AverageResponse) XXX_DiscardUnknown

```go
func (m *AverageResponse) XXX_DiscardUnknown()
```

#### func (*AverageResponse) XXX_Marshal

```go
func (m *AverageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*AverageResponse) XXX_Merge

```go
func (dst *AverageResponse) XXX_Merge(src proto.Message)
```

#### func (*AverageResponse) XXX_Size

```go
func (m *AverageResponse) XXX_Size() int
```

#### func (*AverageResponse) XXX_Unmarshal

```go
func (m *AverageResponse) XXX_Unmarshal(b []byte) error
```

#### type CountPerDurationResponse

```go
type CountPerDurationResponse struct {
	Count                int64                `protobuf:"varint,1,opt,name=Count,proto3" json:"Count,omitempty"`
	Until                *timestamp.Timestamp `protobuf:"bytes,2,opt,name=Until,proto3" json:"Until,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}
```


#### func (*CountPerDurationResponse) Descriptor

```go
func (*CountPerDurationResponse) Descriptor() ([]byte, []int)
```

#### func (*CountPerDurationResponse) GetCount

```go
func (m *CountPerDurationResponse) GetCount() int64
```

#### func (*CountPerDurationResponse) GetUntil

```go
func (m *CountPerDurationResponse) GetUntil() *timestamp.Timestamp
```

#### func (*CountPerDurationResponse) ProtoMessage

```go
func (*CountPerDurationResponse) ProtoMessage()
```

#### func (*CountPerDurationResponse) Reset

```go
func (m *CountPerDurationResponse) Reset()
```

#### func (*CountPerDurationResponse) String

```go
func (m *CountPerDurationResponse) String() string
```

#### func (*CountPerDurationResponse) XXX_DiscardUnknown

```go
func (m *CountPerDurationResponse) XXX_DiscardUnknown()
```

#### func (*CountPerDurationResponse) XXX_Marshal

```go
func (m *CountPerDurationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*CountPerDurationResponse) XXX_Merge

```go
func (dst *CountPerDurationResponse) XXX_Merge(src proto.Message)
```

#### func (*CountPerDurationResponse) XXX_Size

```go
func (m *CountPerDurationResponse) XXX_Size() int
```

#### func (*CountPerDurationResponse) XXX_Unmarshal

```go
func (m *CountPerDurationResponse) XXX_Unmarshal(b []byte) error
```

#### type CountResponse

```go
type CountResponse struct {
	Count                int64    `protobuf:"varint,1,opt,name=Count,proto3" json:"Count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func (*CountResponse) Descriptor

```go
func (*CountResponse) Descriptor() ([]byte, []int)
```

#### func (*CountResponse) GetCount

```go
func (m *CountResponse) GetCount() int64
```

#### func (*CountResponse) ProtoMessage

```go
func (*CountResponse) ProtoMessage()
```

#### func (*CountResponse) Reset

```go
func (m *CountResponse) Reset()
```

#### func (*CountResponse) String

```go
func (m *CountResponse) String() string
```

#### func (*CountResponse) XXX_DiscardUnknown

```go
func (m *CountResponse) XXX_DiscardUnknown()
```

#### func (*CountResponse) XXX_Marshal

```go
func (m *CountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*CountResponse) XXX_Merge

```go
func (dst *CountResponse) XXX_Merge(src proto.Message)
```

#### func (*CountResponse) XXX_Size

```go
func (m *CountResponse) XXX_Size() int
```

#### func (*CountResponse) XXX_Unmarshal

```go
func (m *CountResponse) XXX_Unmarshal(b []byte) error
```

#### type NameRequest

```go
type NameRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```

Requesting a metric by name

#### func (*NameRequest) Descriptor

```go
func (*NameRequest) Descriptor() ([]byte, []int)
```

#### func (*NameRequest) GetName

```go
func (m *NameRequest) GetName() string
```

#### func (*NameRequest) ProtoMessage

```go
func (*NameRequest) ProtoMessage()
```

#### func (*NameRequest) Reset

```go
func (m *NameRequest) Reset()
```

#### func (*NameRequest) String

```go
func (m *NameRequest) String() string
```

#### func (*NameRequest) XXX_DiscardUnknown

```go
func (m *NameRequest) XXX_DiscardUnknown()
```

#### func (*NameRequest) XXX_Marshal

```go
func (m *NameRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*NameRequest) XXX_Merge

```go
func (dst *NameRequest) XXX_Merge(src proto.Message)
```

#### func (*NameRequest) XXX_Size

```go
func (m *NameRequest) XXX_Size() int
```

#### func (*NameRequest) XXX_Unmarshal

```go
func (m *NameRequest) XXX_Unmarshal(b []byte) error
```

#### type ReporterClient

```go
type ReporterClient interface {
	AllNames(ctx context.Context, in *AllNamesRequest, opts ...grpc.CallOption) (*AllNamesResponse, error)
	Average(ctx context.Context, in *NameRequest, opts ...grpc.CallOption) (*AverageResponse, error)
	AveragePerDuration(ctx context.Context, in *NameRequest, opts ...grpc.CallOption) (*AveragePerDurationResponse, error)
	Count(ctx context.Context, in *NameRequest, opts ...grpc.CallOption) (*CountResponse, error)
	CountPerDuration(ctx context.Context, in *NameRequest, opts ...grpc.CallOption) (*CountPerDurationResponse, error)
	Sum(ctx context.Context, in *NameRequest, opts ...grpc.CallOption) (*SumResponse, error)
	SumPerDuration(ctx context.Context, in *NameRequest, opts ...grpc.CallOption) (*SumPerDurationResponse, error)
}
```

ReporterClient is the client API for Reporter service.

For semantics around ctx use and closing/ending streaming RPCs, please refer to
https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.

#### func  NewReporterClient

```go
func NewReporterClient(cc *grpc.ClientConn) ReporterClient
```

#### type ReporterServer

```go
type ReporterServer interface {
	AllNames(context.Context, *AllNamesRequest) (*AllNamesResponse, error)
	Average(context.Context, *NameRequest) (*AverageResponse, error)
	AveragePerDuration(context.Context, *NameRequest) (*AveragePerDurationResponse, error)
	Count(context.Context, *NameRequest) (*CountResponse, error)
	CountPerDuration(context.Context, *NameRequest) (*CountPerDurationResponse, error)
	Sum(context.Context, *NameRequest) (*SumResponse, error)
	SumPerDuration(context.Context, *NameRequest) (*SumPerDurationResponse, error)
}
```

ReporterServer is the server API for Reporter service.

#### type SumPerDurationResponse

```go
type SumPerDurationResponse struct {
	Sum                  float64              `protobuf:"fixed64,1,opt,name=Sum,proto3" json:"Sum,omitempty"`
	N                    int64                `protobuf:"varint,2,opt,name=N,proto3" json:"N,omitempty"`
	Until                *timestamp.Timestamp `protobuf:"bytes,3,opt,name=Until,proto3" json:"Until,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}
```


#### func (*SumPerDurationResponse) Descriptor

```go
func (*SumPerDurationResponse) Descriptor() ([]byte, []int)
```

#### func (*SumPerDurationResponse) GetN

```go
func (m *SumPerDurationResponse) GetN() int64
```

#### func (*SumPerDurationResponse) GetSum

```go
func (m *SumPerDurationResponse) GetSum() float64
```

#### func (*SumPerDurationResponse) GetUntil

```go
func (m *SumPerDurationResponse) GetUntil() *timestamp.Timestamp
```

#### func (*SumPerDurationResponse) ProtoMessage

```go
func (*SumPerDurationResponse) ProtoMessage()
```

#### func (*SumPerDurationResponse) Reset

```go
func (m *SumPerDurationResponse) Reset()
```

#### func (*SumPerDurationResponse) String

```go
func (m *SumPerDurationResponse) String() string
```

#### func (*SumPerDurationResponse) XXX_DiscardUnknown

```go
func (m *SumPerDurationResponse) XXX_DiscardUnknown()
```

#### func (*SumPerDurationResponse) XXX_Marshal

```go
func (m *SumPerDurationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*SumPerDurationResponse) XXX_Merge

```go
func (dst *SumPerDurationResponse) XXX_Merge(src proto.Message)
```

#### func (*SumPerDurationResponse) XXX_Size

```go
func (m *SumPerDurationResponse) XXX_Size() int
```

#### func (*SumPerDurationResponse) XXX_Unmarshal

```go
func (m *SumPerDurationResponse) XXX_Unmarshal(b []byte) error
```

#### type SumResponse

```go
type SumResponse struct {
	Sum                  float64  `protobuf:"fixed64,1,opt,name=Sum,proto3" json:"Sum,omitempty"`
	N                    int64    `protobuf:"varint,2,opt,name=N,proto3" json:"N,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func (*SumResponse) Descriptor

```go
func (*SumResponse) Descriptor() ([]byte, []int)
```

#### func (*SumResponse) GetN

```go
func (m *SumResponse) GetN() int64
```

#### func (*SumResponse) GetSum

```go
func (m *SumResponse) GetSum() float64
```

#### func (*SumResponse) ProtoMessage

```go
func (*SumResponse) ProtoMessage()
```

#### func (*SumResponse) Reset

```go
func (m *SumResponse) Reset()
```

#### func (*SumResponse) String

```go
func (m *SumResponse) String() string
```

#### func (*SumResponse) XXX_DiscardUnknown

```go
func (m *SumResponse) XXX_DiscardUnknown()
```

#### func (*SumResponse) XXX_Marshal

```go
func (m *SumResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*SumResponse) XXX_Merge

```go
func (dst *SumResponse) XXX_Merge(src proto.Message)
```

#### func (*SumResponse) XXX_Size

```go
func (m *SumResponse) XXX_Size() int
```

#### func (*SumResponse) XXX_Unmarshal

```go
func (m *SumResponse) XXX_Unmarshal(b []byte) error
```
