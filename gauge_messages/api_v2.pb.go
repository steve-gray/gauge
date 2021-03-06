// Code generated by protoc-gen-go.
// source: api_v2.proto
// DO NOT EDIT!

package gauge_messages

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Types of log level that gauge supports
type ExecutionRequest_LogLevel int32

const (
	ExecutionRequest_INFO    ExecutionRequest_LogLevel = 0
	ExecutionRequest_DEBUG   ExecutionRequest_LogLevel = 1
	ExecutionRequest_WARNING ExecutionRequest_LogLevel = 2
	ExecutionRequest_ERROR   ExecutionRequest_LogLevel = 3
)

var ExecutionRequest_LogLevel_name = map[int32]string{
	0: "INFO",
	1: "DEBUG",
	2: "WARNING",
	3: "ERROR",
}
var ExecutionRequest_LogLevel_value = map[string]int32{
	"INFO":    0,
	"DEBUG":   1,
	"WARNING": 2,
	"ERROR":   3,
}

func (x ExecutionRequest_LogLevel) Enum() *ExecutionRequest_LogLevel {
	p := new(ExecutionRequest_LogLevel)
	*p = x
	return p
}
func (x ExecutionRequest_LogLevel) String() string {
	return proto.EnumName(ExecutionRequest_LogLevel_name, int32(x))
}
func (x *ExecutionRequest_LogLevel) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ExecutionRequest_LogLevel_value, data, "ExecutionRequest_LogLevel")
	if err != nil {
		return err
	}
	*x = ExecutionRequest_LogLevel(value)
	return nil
}
func (ExecutionRequest_LogLevel) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0, 0} }

// Strategies for parallelization
type ExecutionRequest_Strategy int32

const (
	ExecutionRequest_LAZY  ExecutionRequest_Strategy = 0
	ExecutionRequest_EAGER ExecutionRequest_Strategy = 1
)

var ExecutionRequest_Strategy_name = map[int32]string{
	0: "LAZY",
	1: "EAGER",
}
var ExecutionRequest_Strategy_value = map[string]int32{
	"LAZY":  0,
	"EAGER": 1,
}

func (x ExecutionRequest_Strategy) Enum() *ExecutionRequest_Strategy {
	p := new(ExecutionRequest_Strategy)
	*p = x
	return p
}
func (x ExecutionRequest_Strategy) String() string {
	return proto.EnumName(ExecutionRequest_Strategy_name, int32(x))
}
func (x *ExecutionRequest_Strategy) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ExecutionRequest_Strategy_value, data, "ExecutionRequest_Strategy")
	if err != nil {
		return err
	}
	*x = ExecutionRequest_Strategy(value)
	return nil
}
func (ExecutionRequest_Strategy) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0, 1} }

// Specifies the execution status
type Result_Status int32

const (
	Result_PASSED  Result_Status = 0
	Result_FAILED  Result_Status = 1
	Result_SKIPPED Result_Status = 2
)

var Result_Status_name = map[int32]string{
	0: "PASSED",
	1: "FAILED",
	2: "SKIPPED",
}
var Result_Status_value = map[string]int32{
	"PASSED":  0,
	"FAILED":  1,
	"SKIPPED": 2,
}

func (x Result_Status) Enum() *Result_Status {
	p := new(Result_Status)
	*p = x
	return p
}
func (x Result_Status) String() string {
	return proto.EnumName(Result_Status_name, int32(x))
}
func (x *Result_Status) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Result_Status_value, data, "Result_Status")
	if err != nil {
		return err
	}
	*x = Result_Status(value)
	return nil
}
func (Result_Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{1, 0} }

// Types of Execution respone
type ExecutionResponse_Type int32

const (
	ExecutionResponse_SuiteStart    ExecutionResponse_Type = 0
	ExecutionResponse_SpecStart     ExecutionResponse_Type = 1
	ExecutionResponse_ScenarioStart ExecutionResponse_Type = 2
	ExecutionResponse_ScenarioEnd   ExecutionResponse_Type = 3
	ExecutionResponse_SpecEnd       ExecutionResponse_Type = 4
	ExecutionResponse_SuiteEnd      ExecutionResponse_Type = 5
	ExecutionResponse_ErrorResult   ExecutionResponse_Type = 6
)

var ExecutionResponse_Type_name = map[int32]string{
	0: "SuiteStart",
	1: "SpecStart",
	2: "ScenarioStart",
	3: "ScenarioEnd",
	4: "SpecEnd",
	5: "SuiteEnd",
	6: "ErrorResult",
}
var ExecutionResponse_Type_value = map[string]int32{
	"SuiteStart":    0,
	"SpecStart":     1,
	"ScenarioStart": 2,
	"ScenarioEnd":   3,
	"SpecEnd":       4,
	"SuiteEnd":      5,
	"ErrorResult":   6,
}

func (x ExecutionResponse_Type) Enum() *ExecutionResponse_Type {
	p := new(ExecutionResponse_Type)
	*p = x
	return p
}
func (x ExecutionResponse_Type) String() string {
	return proto.EnumName(ExecutionResponse_Type_name, int32(x))
}
func (x *ExecutionResponse_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ExecutionResponse_Type_value, data, "ExecutionResponse_Type")
	if err != nil {
		return err
	}
	*x = ExecutionResponse_Type(value)
	return nil
}
func (ExecutionResponse_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{2, 0} }

// ExecutionRequest defines the structure of requests that the API's consumers can send to request execution of specs/scenarios
type ExecutionRequest struct {
	// Each ExecutionRequest can ask to execute multiple spec files or multiple scenarios in a spec or a directory or all
	Specs []string `protobuf:"bytes,1,rep,name=specs" json:"specs,omitempty"`
	// Tag expression to filter specs and scenarios during execution. Default: ""
	Tags *string `protobuf:"bytes,2,opt,name=tags" json:"tags,omitempty"`
	// The working directory for gauge.
	WorkingDir *string `protobuf:"bytes,3,opt,name=workingDir" json:"workingDir,omitempty"`
	// Environment to choose for gauge.
	Env      *string                    `protobuf:"bytes,4,opt,name=env" json:"env,omitempty"`
	LogLevel *ExecutionRequest_LogLevel `protobuf:"varint,5,opt,name=logLevel,enum=gauge.messages.ExecutionRequest_LogLevel" json:"logLevel,omitempty"`
	// Whether to run gauge in parallel mode.
	IsParallel *bool `protobuf:"varint,6,opt,name=isParallel" json:"isParallel,omitempty"`
	// If isParallel is true, this specifies how many parallel streams to run.
	ParallelStreams *int32 `protobuf:"varint,7,opt,name=parallelStreams" json:"parallelStreams,omitempty"`
	// If true, sorts execution of specifications alphabetically
	Sort     *bool                      `protobuf:"varint,8,opt,name=sort" json:"sort,omitempty"`
	Strategy *ExecutionRequest_Strategy `protobuf:"varint,9,opt,name=strategy,enum=gauge.messages.ExecutionRequest_Strategy" json:"strategy,omitempty"`
	// Specify against which rows of datatable the scenarios should be executed
	TableRows *string `protobuf:"bytes,10,opt,name=tableRows" json:"tableRows,omitempty"`
	// Signals Gauge that the execution is being debugged, usually via an identifier
	Debug            *bool  `protobuf:"varint,11,opt,name=debug" json:"debug,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ExecutionRequest) Reset()                    { *m = ExecutionRequest{} }
func (m *ExecutionRequest) String() string            { return proto.CompactTextString(m) }
func (*ExecutionRequest) ProtoMessage()               {}
func (*ExecutionRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *ExecutionRequest) GetSpecs() []string {
	if m != nil {
		return m.Specs
	}
	return nil
}

func (m *ExecutionRequest) GetTags() string {
	if m != nil && m.Tags != nil {
		return *m.Tags
	}
	return ""
}

func (m *ExecutionRequest) GetWorkingDir() string {
	if m != nil && m.WorkingDir != nil {
		return *m.WorkingDir
	}
	return ""
}

func (m *ExecutionRequest) GetEnv() string {
	if m != nil && m.Env != nil {
		return *m.Env
	}
	return ""
}

func (m *ExecutionRequest) GetLogLevel() ExecutionRequest_LogLevel {
	if m != nil && m.LogLevel != nil {
		return *m.LogLevel
	}
	return ExecutionRequest_INFO
}

func (m *ExecutionRequest) GetIsParallel() bool {
	if m != nil && m.IsParallel != nil {
		return *m.IsParallel
	}
	return false
}

func (m *ExecutionRequest) GetParallelStreams() int32 {
	if m != nil && m.ParallelStreams != nil {
		return *m.ParallelStreams
	}
	return 0
}

func (m *ExecutionRequest) GetSort() bool {
	if m != nil && m.Sort != nil {
		return *m.Sort
	}
	return false
}

func (m *ExecutionRequest) GetStrategy() ExecutionRequest_Strategy {
	if m != nil && m.Strategy != nil {
		return *m.Strategy
	}
	return ExecutionRequest_LAZY
}

func (m *ExecutionRequest) GetTableRows() string {
	if m != nil && m.TableRows != nil {
		return *m.TableRows
	}
	return ""
}

func (m *ExecutionRequest) GetDebug() bool {
	if m != nil && m.Debug != nil {
		return *m.Debug
	}
	return false
}

type Result struct {
	Status *Result_Status `protobuf:"varint,1,opt,name=status,enum=gauge.messages.Result_Status" json:"status,omitempty"`
	// Contains the Execution errors and its details
	Errors []*Result_ExecutionError `protobuf:"bytes,2,rep,name=errors" json:"errors,omitempty"`
	// Specifies the execution time
	ExecutionTime *int64 `protobuf:"varint,3,opt,name=executionTime" json:"executionTime,omitempty"`
	// Contains the console output messages
	Stdout *string `protobuf:"bytes,4,opt,name=stdout" json:"stdout,omitempty"`
	// Holds the before hook failure
	BeforeHookFailure *Result_ExecutionError `protobuf:"bytes,5,opt,name=beforeHookFailure" json:"beforeHookFailure,omitempty"`
	// Holds the after hook failure
	AfterHookFailure *Result_ExecutionError `protobuf:"bytes,6,opt,name=afterHookFailure" json:"afterHookFailure,omitempty"`
	// Holds the table row number (starting with 1) against which the scenario was executed.
	// This field is populated only for data table driven scenarios.
	TableRowNumber   *int64 `protobuf:"varint,7,opt,name=tableRowNumber" json:"tableRowNumber,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Result) Reset()                    { *m = Result{} }
func (m *Result) String() string            { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()               {}
func (*Result) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *Result) GetStatus() Result_Status {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return Result_PASSED
}

func (m *Result) GetErrors() []*Result_ExecutionError {
	if m != nil {
		return m.Errors
	}
	return nil
}

func (m *Result) GetExecutionTime() int64 {
	if m != nil && m.ExecutionTime != nil {
		return *m.ExecutionTime
	}
	return 0
}

func (m *Result) GetStdout() string {
	if m != nil && m.Stdout != nil {
		return *m.Stdout
	}
	return ""
}

func (m *Result) GetBeforeHookFailure() *Result_ExecutionError {
	if m != nil {
		return m.BeforeHookFailure
	}
	return nil
}

func (m *Result) GetAfterHookFailure() *Result_ExecutionError {
	if m != nil {
		return m.AfterHookFailure
	}
	return nil
}

func (m *Result) GetTableRowNumber() int64 {
	if m != nil && m.TableRowNumber != nil {
		return *m.TableRowNumber
	}
	return 0
}

// ExecutionError represents the failure during execution along with details of failure
type Result_ExecutionError struct {
	// Error message from the failure
	ErrorMessage *string `protobuf:"bytes,1,opt,name=errorMessage" json:"errorMessage,omitempty"`
	// Stacktrace from the failure
	StackTrace *string `protobuf:"bytes,2,opt,name=stackTrace" json:"stackTrace,omitempty"`
	// Byte array holding the screenshot taken at the time of failure.
	Screenshot       []byte `protobuf:"bytes,3,opt,name=screenshot" json:"screenshot,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Result_ExecutionError) Reset()                    { *m = Result_ExecutionError{} }
func (m *Result_ExecutionError) String() string            { return proto.CompactTextString(m) }
func (*Result_ExecutionError) ProtoMessage()               {}
func (*Result_ExecutionError) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1, 0} }

func (m *Result_ExecutionError) GetErrorMessage() string {
	if m != nil && m.ErrorMessage != nil {
		return *m.ErrorMessage
	}
	return ""
}

func (m *Result_ExecutionError) GetStackTrace() string {
	if m != nil && m.StackTrace != nil {
		return *m.StackTrace
	}
	return ""
}

func (m *Result_ExecutionError) GetScreenshot() []byte {
	if m != nil {
		return m.Screenshot
	}
	return nil
}

// ExecutionResponse defines the structure of response for ExecutionRequest message
type ExecutionResponse struct {
	// Response type
	Type *ExecutionResponse_Type `protobuf:"varint,1,req,name=type,enum=gauge.messages.ExecutionResponse_Type" json:"type,omitempty"`
	// An identifier for the current execution result. This field is populated only for spec/scenario result.
	// For spec, the value will be the filename.
	// For scenario, the value will be filename:scenario_heading_line_num.
	ID *string `protobuf:"bytes,2,opt,name=ID" json:"ID,omitempty"`
	// Contains all the result details. This field is populated only for ScenarioStart, ScenaioEnd, SpecEnd, SuiteEnd, ErrorResult
	Result           *Result `protobuf:"bytes,3,opt,name=result" json:"result,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ExecutionResponse) Reset()                    { *m = ExecutionResponse{} }
func (m *ExecutionResponse) String() string            { return proto.CompactTextString(m) }
func (*ExecutionResponse) ProtoMessage()               {}
func (*ExecutionResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *ExecutionResponse) GetType() ExecutionResponse_Type {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ExecutionResponse_SuiteStart
}

func (m *ExecutionResponse) GetID() string {
	if m != nil && m.ID != nil {
		return *m.ID
	}
	return ""
}

func (m *ExecutionResponse) GetResult() *Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func init() {
	proto.RegisterType((*ExecutionRequest)(nil), "gauge.messages.ExecutionRequest")
	proto.RegisterType((*Result)(nil), "gauge.messages.Result")
	proto.RegisterType((*Result_ExecutionError)(nil), "gauge.messages.Result.ExecutionError")
	proto.RegisterType((*ExecutionResponse)(nil), "gauge.messages.ExecutionResponse")
	proto.RegisterEnum("gauge.messages.ExecutionRequest_LogLevel", ExecutionRequest_LogLevel_name, ExecutionRequest_LogLevel_value)
	proto.RegisterEnum("gauge.messages.ExecutionRequest_Strategy", ExecutionRequest_Strategy_name, ExecutionRequest_Strategy_value)
	proto.RegisterEnum("gauge.messages.Result_Status", Result_Status_name, Result_Status_value)
	proto.RegisterEnum("gauge.messages.ExecutionResponse_Type", ExecutionResponse_Type_name, ExecutionResponse_Type_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Execution service

type ExecutionClient interface {
	// Bind RPC method
	Execute(ctx context.Context, in *ExecutionRequest, opts ...grpc.CallOption) (Execution_ExecuteClient, error)
}

type executionClient struct {
	cc *grpc.ClientConn
}

func NewExecutionClient(cc *grpc.ClientConn) ExecutionClient {
	return &executionClient{cc}
}

func (c *executionClient) Execute(ctx context.Context, in *ExecutionRequest, opts ...grpc.CallOption) (Execution_ExecuteClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Execution_serviceDesc.Streams[0], c.cc, "/gauge.messages.Execution/execute", opts...)
	if err != nil {
		return nil, err
	}
	x := &executionExecuteClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Execution_ExecuteClient interface {
	Recv() (*ExecutionResponse, error)
	grpc.ClientStream
}

type executionExecuteClient struct {
	grpc.ClientStream
}

func (x *executionExecuteClient) Recv() (*ExecutionResponse, error) {
	m := new(ExecutionResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Execution service

type ExecutionServer interface {
	// Bind RPC method
	Execute(*ExecutionRequest, Execution_ExecuteServer) error
}

func RegisterExecutionServer(s *grpc.Server, srv ExecutionServer) {
	s.RegisterService(&_Execution_serviceDesc, srv)
}

func _Execution_Execute_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ExecutionRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ExecutionServer).Execute(m, &executionExecuteServer{stream})
}

type Execution_ExecuteServer interface {
	Send(*ExecutionResponse) error
	grpc.ServerStream
}

type executionExecuteServer struct {
	grpc.ServerStream
}

func (x *executionExecuteServer) Send(m *ExecutionResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _Execution_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gauge.messages.Execution",
	HandlerType: (*ExecutionServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "execute",
			Handler:       _Execution_Execute_Handler,
			ServerStreams: true,
		},
	},
	Metadata: fileDescriptor1,
}

func init() { proto.RegisterFile("api_v2.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 714 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x54, 0xcd, 0x8e, 0xe3, 0x44,
	0x10, 0x8e, 0xed, 0xc4, 0x13, 0x57, 0x32, 0x59, 0xa7, 0x85, 0x56, 0xd6, 0x8a, 0x9f, 0x60, 0xc1,
	0xca, 0x1c, 0xb0, 0x50, 0x24, 0x84, 0x84, 0xc4, 0x21, 0xab, 0x78, 0x42, 0x44, 0xc8, 0x86, 0x76,
	0x10, 0x82, 0xcb, 0xaa, 0xe3, 0xa9, 0x35, 0xd6, 0x38, 0xb1, 0xe9, 0x6e, 0xcf, 0x30, 0x2f, 0xc0,
	0xc3, 0x70, 0xe3, 0xad, 0x78, 0x0c, 0xd4, 0x6d, 0x27, 0x24, 0x19, 0x60, 0x86, 0x5b, 0xd7, 0xd7,
	0xf5, 0x55, 0x7f, 0x55, 0xf5, 0xd9, 0xd0, 0x67, 0x65, 0xf6, 0xe6, 0x76, 0x1c, 0x96, 0xbc, 0x90,
	0x05, 0x19, 0xa4, 0xac, 0x4a, 0x31, 0xdc, 0xa2, 0x10, 0x2c, 0x45, 0xe1, 0xff, 0x69, 0x81, 0x1b,
	0xfd, 0x8a, 0x49, 0x25, 0xb3, 0x62, 0x47, 0xf1, 0x97, 0x0a, 0x85, 0x24, 0xef, 0x40, 0x47, 0x94,
	0x98, 0x08, 0xcf, 0x18, 0x59, 0x81, 0x43, 0xeb, 0x80, 0x10, 0x68, 0x4b, 0x96, 0x0a, 0xcf, 0x1c,
	0x19, 0x81, 0x43, 0xf5, 0x99, 0xbc, 0x0f, 0x70, 0x57, 0xf0, 0x9b, 0x6c, 0x97, 0x4e, 0x33, 0xee,
	0x59, 0xfa, 0xe6, 0x08, 0x21, 0x2e, 0x58, 0xb8, 0xbb, 0xf5, 0xda, 0xfa, 0x42, 0x1d, 0x49, 0x04,
	0xdd, 0xbc, 0x48, 0x17, 0x78, 0x8b, 0xb9, 0xd7, 0x19, 0x19, 0xc1, 0x60, 0xfc, 0x49, 0x78, 0xaa,
	0x29, 0x3c, 0xd7, 0x13, 0x2e, 0x1a, 0x02, 0x3d, 0x50, 0xd5, 0xc3, 0x99, 0x58, 0x31, 0xce, 0xf2,
	0x1c, 0x73, 0xcf, 0x1e, 0x19, 0x41, 0x97, 0x1e, 0x21, 0x24, 0x80, 0x67, 0x65, 0x73, 0x8e, 0x25,
	0x47, 0xb6, 0x15, 0xde, 0xc5, 0xc8, 0x08, 0x3a, 0xf4, 0x1c, 0x56, 0x6d, 0x89, 0x82, 0x4b, 0xaf,
	0xab, 0x6b, 0xe8, 0xb3, 0x12, 0x29, 0x24, 0x67, 0x12, 0xd3, 0x7b, 0xcf, 0x79, 0xa2, 0xc8, 0xb8,
	0x21, 0xd0, 0x03, 0x95, 0xbc, 0x0b, 0x8e, 0x64, 0x9b, 0x1c, 0x69, 0x71, 0x27, 0x3c, 0xd0, 0x33,
	0xf8, 0x1b, 0x50, 0x53, 0xbe, 0xc6, 0x4d, 0x95, 0x7a, 0x3d, 0xfd, 0x72, 0x1d, 0xf8, 0x5f, 0x40,
	0x77, 0xdf, 0x2e, 0xe9, 0x42, 0x7b, 0xbe, 0xbc, 0x7a, 0xed, 0xb6, 0x88, 0x03, 0x9d, 0x69, 0xf4,
	0xea, 0xfb, 0x99, 0x6b, 0x90, 0x1e, 0x5c, 0xfc, 0x30, 0xa1, 0xcb, 0xf9, 0x72, 0xe6, 0x9a, 0x0a,
	0x8f, 0x28, 0x7d, 0x4d, 0x5d, 0xcb, 0xff, 0x00, 0xba, 0x7b, 0x09, 0x8a, 0xb8, 0x98, 0xfc, 0xf4,
	0x63, 0x4d, 0x8c, 0x26, 0xb3, 0x88, 0xba, 0x86, 0xff, 0x47, 0x1b, 0x6c, 0x8a, 0xa2, 0xca, 0x25,
	0xf9, 0x1c, 0x6c, 0x21, 0x99, 0xac, 0xd4, 0x86, 0x55, 0x77, 0xef, 0x9d, 0x77, 0x57, 0xe7, 0x85,
	0xb1, 0x4e, 0xa2, 0x4d, 0x32, 0xf9, 0x0a, 0x6c, 0xe4, 0xbc, 0xe0, 0xca, 0x03, 0x56, 0xd0, 0x1b,
	0x7f, 0xfc, 0x2f, 0xb4, 0xc3, 0x6c, 0x22, 0x95, 0x4d, 0x1b, 0x12, 0xf9, 0x08, 0x2e, 0x71, 0x7f,
	0xb3, 0xce, 0xb6, 0xa8, 0xfd, 0x62, 0xd1, 0x53, 0x90, 0x3c, 0x57, 0xda, 0xae, 0x8b, 0x4a, 0x36,
	0xae, 0x69, 0x22, 0x12, 0xc3, 0x70, 0x83, 0x6f, 0x0b, 0x8e, 0x5f, 0x17, 0xc5, 0xcd, 0x15, 0xcb,
	0xf2, 0x8a, 0xa3, 0x76, 0xd0, 0x93, 0x75, 0x3c, 0xe4, 0x93, 0xef, 0xc0, 0x65, 0x6f, 0x25, 0xf2,
	0xe3, 0x9a, 0xf6, 0xff, 0xa9, 0xf9, 0x80, 0x4e, 0x5e, 0xc2, 0x60, 0xbf, 0xe3, 0x65, 0xb5, 0xdd,
	0x20, 0xd7, 0xc6, 0xb3, 0xe8, 0x19, 0xfa, 0x42, 0xc2, 0xe0, 0xb4, 0x16, 0xf1, 0xa1, 0xaf, 0x27,
	0xf5, 0x6d, 0xfd, 0xa4, 0xde, 0x8d, 0x43, 0x4f, 0x30, 0xe5, 0x7b, 0x21, 0x59, 0x72, 0xb3, 0xe6,
	0x2c, 0xc1, 0xe6, 0x53, 0x3c, 0x42, 0xf4, 0x7d, 0xc2, 0x11, 0x77, 0xe2, 0xe7, 0x42, 0xea, 0x01,
	0xf7, 0xe9, 0x11, 0xe2, 0x7f, 0x0a, 0x76, 0xbd, 0x54, 0x02, 0x60, 0xaf, 0x26, 0x71, 0x1c, 0x4d,
	0xdd, 0x96, 0x3a, 0x5f, 0x4d, 0xe6, 0x8b, 0x68, 0x5a, 0xfb, 0x2b, 0xfe, 0x66, 0xbe, 0x5a, 0x45,
	0x53, 0xd7, 0xf4, 0x7f, 0x33, 0x61, 0x78, 0xe4, 0x74, 0x51, 0x16, 0x3b, 0x81, 0xe4, 0x4b, 0x68,
	0xcb, 0xfb, 0x52, 0x09, 0x34, 0x83, 0xc1, 0xf8, 0xe5, 0x7f, 0x7c, 0x1a, 0x35, 0x21, 0x5c, 0xdf,
	0x97, 0x48, 0x35, 0x87, 0x0c, 0xc0, 0x9c, 0x4f, 0x1b, 0xe1, 0xe6, 0x7c, 0x4a, 0x42, 0xb0, 0xb9,
	0x9e, 0xac, 0x16, 0xdb, 0x1b, 0x3f, 0xff, 0xe7, 0xb9, 0xd3, 0x26, 0xcb, 0xaf, 0xa0, 0xbd, 0xae,
	0xeb, 0x40, 0x5c, 0x65, 0x12, 0x63, 0xc9, 0xb8, 0x74, 0x5b, 0xe4, 0x12, 0x9c, 0xb8, 0xc4, 0xa4,
	0x0e, 0x0d, 0x32, 0x84, 0xcb, 0x38, 0xc1, 0x1d, 0xe3, 0x59, 0x51, 0x43, 0x26, 0x79, 0x06, 0xbd,
	0x3d, 0x14, 0xed, 0xae, 0x5d, 0x4b, 0x77, 0x5a, 0x62, 0xa2, 0x82, 0x36, 0xe9, 0x43, 0x57, 0xd7,
	0x53, 0x51, 0x47, 0xe5, 0xd6, 0xfb, 0xd5, 0x8f, 0xba, 0xf6, 0xf8, 0x0d, 0x38, 0x87, 0xb6, 0x08,
	0x85, 0x8b, 0xda, 0xb3, 0x48, 0x46, 0x8f, 0xfd, 0x17, 0x5e, 0x7c, 0xf8, 0xe8, 0x78, 0xfc, 0xd6,
	0x67, 0xc6, 0xab, 0xe1, 0xef, 0xe6, 0x60, 0xa6, 0x13, 0x9b, 0x55, 0x8b, 0xbf, 0x02, 0x00, 0x00,
	0xff, 0xff, 0x61, 0x7a, 0x7c, 0x1f, 0xba, 0x05, 0x00, 0x00,
}
