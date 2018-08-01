# api
--
    import "github.com/KarelKubat/runtime-metrics/api"

Package api is the protobuf-derived network API that the reporter's server and
client employ. This package is not for public consumption.

## Usage

#### func  RegisterReporterServer

```go
func RegisterReporterServer(s *grpc.Server, srv ReporterServer)
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

#### type EmptyRequest

```go
type EmptyRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```

Discovery of all the names

#### func (*EmptyRequest) Descriptor

```go
func (*EmptyRequest) Descriptor() ([]byte, []int)
```

#### func (*EmptyRequest) ProtoMessage

```go
func (*EmptyRequest) ProtoMessage()
```

#### func (*EmptyRequest) Reset

```go
func (m *EmptyRequest) Reset()
```

#### func (*EmptyRequest) String

```go
func (m *EmptyRequest) String() string
```

#### func (*EmptyRequest) XXX_DiscardUnknown

```go
func (m *EmptyRequest) XXX_DiscardUnknown()
```

#### func (*EmptyRequest) XXX_Marshal

```go
func (m *EmptyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*EmptyRequest) XXX_Merge

```go
func (dst *EmptyRequest) XXX_Merge(src proto.Message)
```

#### func (*EmptyRequest) XXX_Size

```go
func (m *EmptyRequest) XXX_Size() int
```

#### func (*EmptyRequest) XXX_Unmarshal

```go
func (m *EmptyRequest) XXX_Unmarshal(b []byte) error
```

#### type FullDumpResponse

```go
type FullDumpResponse struct {
	NamedAverages            []*NamedAverage            `protobuf:"bytes,1,rep,name=NamedAverages,proto3" json:"NamedAverages,omitempty"`
	NamedAveragesPerDuration []*NamedAveragePerDuration `protobuf:"bytes,2,rep,name=NamedAveragesPerDuration,proto3" json:"NamedAveragesPerDuration,omitempty"`
	NamedCounts              []*NamedCount              `protobuf:"bytes,3,rep,name=NamedCounts,proto3" json:"NamedCounts,omitempty"`
	NamedCountsPerDuration   []*NamedCountPerDuration   `protobuf:"bytes,4,rep,name=NamedCountsPerDuration,proto3" json:"NamedCountsPerDuration,omitempty"`
	NamedSums                []*NamedSum                `protobuf:"bytes,5,rep,name=NamedSums,proto3" json:"NamedSums,omitempty"`
	NamedSumsPerDuration     []*NamedSumPerDuration     `protobuf:"bytes,6,rep,name=NamedSumsPerDuration,proto3" json:"NamedSumsPerDuration,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}                   `json:"-"`
	XXX_unrecognized         []byte                     `json:"-"`
	XXX_sizecache            int32                      `json:"-"`
}
```


#### func (*FullDumpResponse) Descriptor

```go
func (*FullDumpResponse) Descriptor() ([]byte, []int)
```

#### func (*FullDumpResponse) GetNamedAverages

```go
func (m *FullDumpResponse) GetNamedAverages() []*NamedAverage
```

#### func (*FullDumpResponse) GetNamedAveragesPerDuration

```go
func (m *FullDumpResponse) GetNamedAveragesPerDuration() []*NamedAveragePerDuration
```

#### func (*FullDumpResponse) GetNamedCounts

```go
func (m *FullDumpResponse) GetNamedCounts() []*NamedCount
```

#### func (*FullDumpResponse) GetNamedCountsPerDuration

```go
func (m *FullDumpResponse) GetNamedCountsPerDuration() []*NamedCountPerDuration
```

#### func (*FullDumpResponse) GetNamedSums

```go
func (m *FullDumpResponse) GetNamedSums() []*NamedSum
```

#### func (*FullDumpResponse) GetNamedSumsPerDuration

```go
func (m *FullDumpResponse) GetNamedSumsPerDuration() []*NamedSumPerDuration
```

#### func (*FullDumpResponse) ProtoMessage

```go
func (*FullDumpResponse) ProtoMessage()
```

#### func (*FullDumpResponse) Reset

```go
func (m *FullDumpResponse) Reset()
```

#### func (*FullDumpResponse) String

```go
func (m *FullDumpResponse) String() string
```

#### func (*FullDumpResponse) XXX_DiscardUnknown

```go
func (m *FullDumpResponse) XXX_DiscardUnknown()
```

#### func (*FullDumpResponse) XXX_Marshal

```go
func (m *FullDumpResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*FullDumpResponse) XXX_Merge

```go
func (dst *FullDumpResponse) XXX_Merge(src proto.Message)
```

#### func (*FullDumpResponse) XXX_Size

```go
func (m *FullDumpResponse) XXX_Size() int
```

#### func (*FullDumpResponse) XXX_Unmarshal

```go
func (m *FullDumpResponse) XXX_Unmarshal(b []byte) error
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

#### type NamedAverage

```go
type NamedAverage struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Value                float64  `protobuf:"fixed64,2,opt,name=Value,proto3" json:"Value,omitempty"`
	N                    int64    `protobuf:"varint,3,opt,name=N,proto3" json:"N,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func (*NamedAverage) Descriptor

```go
func (*NamedAverage) Descriptor() ([]byte, []int)
```

#### func (*NamedAverage) GetN

```go
func (m *NamedAverage) GetN() int64
```

#### func (*NamedAverage) GetName

```go
func (m *NamedAverage) GetName() string
```

#### func (*NamedAverage) GetValue

```go
func (m *NamedAverage) GetValue() float64
```

#### func (*NamedAverage) ProtoMessage

```go
func (*NamedAverage) ProtoMessage()
```

#### func (*NamedAverage) Reset

```go
func (m *NamedAverage) Reset()
```

#### func (*NamedAverage) String

```go
func (m *NamedAverage) String() string
```

#### func (*NamedAverage) XXX_DiscardUnknown

```go
func (m *NamedAverage) XXX_DiscardUnknown()
```

#### func (*NamedAverage) XXX_Marshal

```go
func (m *NamedAverage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*NamedAverage) XXX_Merge

```go
func (dst *NamedAverage) XXX_Merge(src proto.Message)
```

#### func (*NamedAverage) XXX_Size

```go
func (m *NamedAverage) XXX_Size() int
```

#### func (*NamedAverage) XXX_Unmarshal

```go
func (m *NamedAverage) XXX_Unmarshal(b []byte) error
```

#### type NamedAveragePerDuration

```go
type NamedAveragePerDuration struct {
	Name                 string               `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Value                float64              `protobuf:"fixed64,2,opt,name=Value,proto3" json:"Value,omitempty"`
	N                    int64                `protobuf:"varint,3,opt,name=N,proto3" json:"N,omitempty"`
	Until                *timestamp.Timestamp `protobuf:"bytes,4,opt,name=Until,proto3" json:"Until,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}
```


#### func (*NamedAveragePerDuration) Descriptor

```go
func (*NamedAveragePerDuration) Descriptor() ([]byte, []int)
```

#### func (*NamedAveragePerDuration) GetN

```go
func (m *NamedAveragePerDuration) GetN() int64
```

#### func (*NamedAveragePerDuration) GetName

```go
func (m *NamedAveragePerDuration) GetName() string
```

#### func (*NamedAveragePerDuration) GetUntil

```go
func (m *NamedAveragePerDuration) GetUntil() *timestamp.Timestamp
```

#### func (*NamedAveragePerDuration) GetValue

```go
func (m *NamedAveragePerDuration) GetValue() float64
```

#### func (*NamedAveragePerDuration) ProtoMessage

```go
func (*NamedAveragePerDuration) ProtoMessage()
```

#### func (*NamedAveragePerDuration) Reset

```go
func (m *NamedAveragePerDuration) Reset()
```

#### func (*NamedAveragePerDuration) String

```go
func (m *NamedAveragePerDuration) String() string
```

#### func (*NamedAveragePerDuration) XXX_DiscardUnknown

```go
func (m *NamedAveragePerDuration) XXX_DiscardUnknown()
```

#### func (*NamedAveragePerDuration) XXX_Marshal

```go
func (m *NamedAveragePerDuration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*NamedAveragePerDuration) XXX_Merge

```go
func (dst *NamedAveragePerDuration) XXX_Merge(src proto.Message)
```

#### func (*NamedAveragePerDuration) XXX_Size

```go
func (m *NamedAveragePerDuration) XXX_Size() int
```

#### func (*NamedAveragePerDuration) XXX_Unmarshal

```go
func (m *NamedAveragePerDuration) XXX_Unmarshal(b []byte) error
```

#### type NamedCount

```go
type NamedCount struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Value                int64    `protobuf:"varint,2,opt,name=Value,proto3" json:"Value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func (*NamedCount) Descriptor

```go
func (*NamedCount) Descriptor() ([]byte, []int)
```

#### func (*NamedCount) GetName

```go
func (m *NamedCount) GetName() string
```

#### func (*NamedCount) GetValue

```go
func (m *NamedCount) GetValue() int64
```

#### func (*NamedCount) ProtoMessage

```go
func (*NamedCount) ProtoMessage()
```

#### func (*NamedCount) Reset

```go
func (m *NamedCount) Reset()
```

#### func (*NamedCount) String

```go
func (m *NamedCount) String() string
```

#### func (*NamedCount) XXX_DiscardUnknown

```go
func (m *NamedCount) XXX_DiscardUnknown()
```

#### func (*NamedCount) XXX_Marshal

```go
func (m *NamedCount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*NamedCount) XXX_Merge

```go
func (dst *NamedCount) XXX_Merge(src proto.Message)
```

#### func (*NamedCount) XXX_Size

```go
func (m *NamedCount) XXX_Size() int
```

#### func (*NamedCount) XXX_Unmarshal

```go
func (m *NamedCount) XXX_Unmarshal(b []byte) error
```

#### type NamedCountPerDuration

```go
type NamedCountPerDuration struct {
	Name                 string               `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Value                int64                `protobuf:"varint,2,opt,name=Value,proto3" json:"Value,omitempty"`
	Until                *timestamp.Timestamp `protobuf:"bytes,3,opt,name=Until,proto3" json:"Until,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}
```


#### func (*NamedCountPerDuration) Descriptor

```go
func (*NamedCountPerDuration) Descriptor() ([]byte, []int)
```

#### func (*NamedCountPerDuration) GetName

```go
func (m *NamedCountPerDuration) GetName() string
```

#### func (*NamedCountPerDuration) GetUntil

```go
func (m *NamedCountPerDuration) GetUntil() *timestamp.Timestamp
```

#### func (*NamedCountPerDuration) GetValue

```go
func (m *NamedCountPerDuration) GetValue() int64
```

#### func (*NamedCountPerDuration) ProtoMessage

```go
func (*NamedCountPerDuration) ProtoMessage()
```

#### func (*NamedCountPerDuration) Reset

```go
func (m *NamedCountPerDuration) Reset()
```

#### func (*NamedCountPerDuration) String

```go
func (m *NamedCountPerDuration) String() string
```

#### func (*NamedCountPerDuration) XXX_DiscardUnknown

```go
func (m *NamedCountPerDuration) XXX_DiscardUnknown()
```

#### func (*NamedCountPerDuration) XXX_Marshal

```go
func (m *NamedCountPerDuration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*NamedCountPerDuration) XXX_Merge

```go
func (dst *NamedCountPerDuration) XXX_Merge(src proto.Message)
```

#### func (*NamedCountPerDuration) XXX_Size

```go
func (m *NamedCountPerDuration) XXX_Size() int
```

#### func (*NamedCountPerDuration) XXX_Unmarshal

```go
func (m *NamedCountPerDuration) XXX_Unmarshal(b []byte) error
```

#### type NamedSum

```go
type NamedSum struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Value                float64  `protobuf:"fixed64,2,opt,name=Value,proto3" json:"Value,omitempty"`
	N                    int64    `protobuf:"varint,3,opt,name=N,proto3" json:"N,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func (*NamedSum) Descriptor

```go
func (*NamedSum) Descriptor() ([]byte, []int)
```

#### func (*NamedSum) GetN

```go
func (m *NamedSum) GetN() int64
```

#### func (*NamedSum) GetName

```go
func (m *NamedSum) GetName() string
```

#### func (*NamedSum) GetValue

```go
func (m *NamedSum) GetValue() float64
```

#### func (*NamedSum) ProtoMessage

```go
func (*NamedSum) ProtoMessage()
```

#### func (*NamedSum) Reset

```go
func (m *NamedSum) Reset()
```

#### func (*NamedSum) String

```go
func (m *NamedSum) String() string
```

#### func (*NamedSum) XXX_DiscardUnknown

```go
func (m *NamedSum) XXX_DiscardUnknown()
```

#### func (*NamedSum) XXX_Marshal

```go
func (m *NamedSum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*NamedSum) XXX_Merge

```go
func (dst *NamedSum) XXX_Merge(src proto.Message)
```

#### func (*NamedSum) XXX_Size

```go
func (m *NamedSum) XXX_Size() int
```

#### func (*NamedSum) XXX_Unmarshal

```go
func (m *NamedSum) XXX_Unmarshal(b []byte) error
```

#### type NamedSumPerDuration

```go
type NamedSumPerDuration struct {
	Name                 string               `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Value                float64              `protobuf:"fixed64,2,opt,name=Value,proto3" json:"Value,omitempty"`
	N                    int64                `protobuf:"varint,3,opt,name=N,proto3" json:"N,omitempty"`
	Until                *timestamp.Timestamp `protobuf:"bytes,4,opt,name=Until,proto3" json:"Until,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}
```


#### func (*NamedSumPerDuration) Descriptor

```go
func (*NamedSumPerDuration) Descriptor() ([]byte, []int)
```

#### func (*NamedSumPerDuration) GetN

```go
func (m *NamedSumPerDuration) GetN() int64
```

#### func (*NamedSumPerDuration) GetName

```go
func (m *NamedSumPerDuration) GetName() string
```

#### func (*NamedSumPerDuration) GetUntil

```go
func (m *NamedSumPerDuration) GetUntil() *timestamp.Timestamp
```

#### func (*NamedSumPerDuration) GetValue

```go
func (m *NamedSumPerDuration) GetValue() float64
```

#### func (*NamedSumPerDuration) ProtoMessage

```go
func (*NamedSumPerDuration) ProtoMessage()
```

#### func (*NamedSumPerDuration) Reset

```go
func (m *NamedSumPerDuration) Reset()
```

#### func (*NamedSumPerDuration) String

```go
func (m *NamedSumPerDuration) String() string
```

#### func (*NamedSumPerDuration) XXX_DiscardUnknown

```go
func (m *NamedSumPerDuration) XXX_DiscardUnknown()
```

#### func (*NamedSumPerDuration) XXX_Marshal

```go
func (m *NamedSumPerDuration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*NamedSumPerDuration) XXX_Merge

```go
func (dst *NamedSumPerDuration) XXX_Merge(src proto.Message)
```

#### func (*NamedSumPerDuration) XXX_Size

```go
func (m *NamedSumPerDuration) XXX_Size() int
```

#### func (*NamedSumPerDuration) XXX_Unmarshal

```go
func (m *NamedSumPerDuration) XXX_Unmarshal(b []byte) error
```

#### type ReporterClient

```go
type ReporterClient interface {
	AllNames(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*AllNamesResponse, error)
	FullDump(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*FullDumpResponse, error)
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
	AllNames(context.Context, *EmptyRequest) (*AllNamesResponse, error)
	FullDump(context.Context, *EmptyRequest) (*FullDumpResponse, error)
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
