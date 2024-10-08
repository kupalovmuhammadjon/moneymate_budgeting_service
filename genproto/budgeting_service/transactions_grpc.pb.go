// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: transactions.proto

package budgeting_service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TransactionServiceClient is the client API for TransactionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransactionServiceClient interface {
	Create(ctx context.Context, in *CreateTransaction, opts ...grpc.CallOption) (*Transaction, error)
	GetById(ctx context.Context, in *PrimaryKey, opts ...grpc.CallOption) (*Transaction, error)
	GetAll(ctx context.Context, in *TransactionFilter, opts ...grpc.CallOption) (*Transactions, error)
	Update(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*Transaction, error)
	Delete(ctx context.Context, in *PrimaryKey, opts ...grpc.CallOption) (*Void, error)
	GenerateSpendingReport(ctx context.Context, in *PrimaryKey, opts ...grpc.CallOption) (*Spendings, error)
	GenerateIncomeReport(ctx context.Context, in *PrimaryKey, opts ...grpc.CallOption) (*Incomes, error)
	GenerateBudgetPerformanceReport(ctx context.Context, in *PrimaryKey, opts ...grpc.CallOption) (*BugetPerformance, error)
	GenerateGoalProgressReport(ctx context.Context, in *PrimaryKey, opts ...grpc.CallOption) (*GoalProgress, error)
}

type transactionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTransactionServiceClient(cc grpc.ClientConnInterface) TransactionServiceClient {
	return &transactionServiceClient{cc}
}

func (c *transactionServiceClient) Create(ctx context.Context, in *CreateTransaction, opts ...grpc.CallOption) (*Transaction, error) {
	out := new(Transaction)
	err := c.cc.Invoke(ctx, "/budgeting_service.TransactionService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) GetById(ctx context.Context, in *PrimaryKey, opts ...grpc.CallOption) (*Transaction, error) {
	out := new(Transaction)
	err := c.cc.Invoke(ctx, "/budgeting_service.TransactionService/GetById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) GetAll(ctx context.Context, in *TransactionFilter, opts ...grpc.CallOption) (*Transactions, error) {
	out := new(Transactions)
	err := c.cc.Invoke(ctx, "/budgeting_service.TransactionService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) Update(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*Transaction, error) {
	out := new(Transaction)
	err := c.cc.Invoke(ctx, "/budgeting_service.TransactionService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) Delete(ctx context.Context, in *PrimaryKey, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/budgeting_service.TransactionService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) GenerateSpendingReport(ctx context.Context, in *PrimaryKey, opts ...grpc.CallOption) (*Spendings, error) {
	out := new(Spendings)
	err := c.cc.Invoke(ctx, "/budgeting_service.TransactionService/GenerateSpendingReport", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) GenerateIncomeReport(ctx context.Context, in *PrimaryKey, opts ...grpc.CallOption) (*Incomes, error) {
	out := new(Incomes)
	err := c.cc.Invoke(ctx, "/budgeting_service.TransactionService/GenerateIncomeReport", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) GenerateBudgetPerformanceReport(ctx context.Context, in *PrimaryKey, opts ...grpc.CallOption) (*BugetPerformance, error) {
	out := new(BugetPerformance)
	err := c.cc.Invoke(ctx, "/budgeting_service.TransactionService/GenerateBudgetPerformanceReport", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) GenerateGoalProgressReport(ctx context.Context, in *PrimaryKey, opts ...grpc.CallOption) (*GoalProgress, error) {
	out := new(GoalProgress)
	err := c.cc.Invoke(ctx, "/budgeting_service.TransactionService/GenerateGoalProgressReport", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransactionServiceServer is the server API for TransactionService service.
// All implementations must embed UnimplementedTransactionServiceServer
// for forward compatibility
type TransactionServiceServer interface {
	Create(context.Context, *CreateTransaction) (*Transaction, error)
	GetById(context.Context, *PrimaryKey) (*Transaction, error)
	GetAll(context.Context, *TransactionFilter) (*Transactions, error)
	Update(context.Context, *Transaction) (*Transaction, error)
	Delete(context.Context, *PrimaryKey) (*Void, error)
	GenerateSpendingReport(context.Context, *PrimaryKey) (*Spendings, error)
	GenerateIncomeReport(context.Context, *PrimaryKey) (*Incomes, error)
	GenerateBudgetPerformanceReport(context.Context, *PrimaryKey) (*BugetPerformance, error)
	GenerateGoalProgressReport(context.Context, *PrimaryKey) (*GoalProgress, error)
	mustEmbedUnimplementedTransactionServiceServer()
}

// UnimplementedTransactionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTransactionServiceServer struct {
}

func (UnimplementedTransactionServiceServer) Create(context.Context, *CreateTransaction) (*Transaction, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedTransactionServiceServer) GetById(context.Context, *PrimaryKey) (*Transaction, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedTransactionServiceServer) GetAll(context.Context, *TransactionFilter) (*Transactions, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedTransactionServiceServer) Update(context.Context, *Transaction) (*Transaction, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedTransactionServiceServer) Delete(context.Context, *PrimaryKey) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedTransactionServiceServer) GenerateSpendingReport(context.Context, *PrimaryKey) (*Spendings, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateSpendingReport not implemented")
}
func (UnimplementedTransactionServiceServer) GenerateIncomeReport(context.Context, *PrimaryKey) (*Incomes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateIncomeReport not implemented")
}
func (UnimplementedTransactionServiceServer) GenerateBudgetPerformanceReport(context.Context, *PrimaryKey) (*BugetPerformance, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateBudgetPerformanceReport not implemented")
}
func (UnimplementedTransactionServiceServer) GenerateGoalProgressReport(context.Context, *PrimaryKey) (*GoalProgress, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateGoalProgressReport not implemented")
}
func (UnimplementedTransactionServiceServer) mustEmbedUnimplementedTransactionServiceServer() {}

// UnsafeTransactionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransactionServiceServer will
// result in compilation errors.
type UnsafeTransactionServiceServer interface {
	mustEmbedUnimplementedTransactionServiceServer()
}

func RegisterTransactionServiceServer(s grpc.ServiceRegistrar, srv TransactionServiceServer) {
	s.RegisterService(&TransactionService_ServiceDesc, srv)
}

func _TransactionService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTransaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/budgeting_service.TransactionService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).Create(ctx, req.(*CreateTransaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/budgeting_service.TransactionService/GetById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).GetById(ctx, req.(*PrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransactionFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/budgeting_service.TransactionService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).GetAll(ctx, req.(*TransactionFilter))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/budgeting_service.TransactionService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).Update(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/budgeting_service.TransactionService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).Delete(ctx, req.(*PrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_GenerateSpendingReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).GenerateSpendingReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/budgeting_service.TransactionService/GenerateSpendingReport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).GenerateSpendingReport(ctx, req.(*PrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_GenerateIncomeReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).GenerateIncomeReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/budgeting_service.TransactionService/GenerateIncomeReport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).GenerateIncomeReport(ctx, req.(*PrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_GenerateBudgetPerformanceReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).GenerateBudgetPerformanceReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/budgeting_service.TransactionService/GenerateBudgetPerformanceReport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).GenerateBudgetPerformanceReport(ctx, req.(*PrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_GenerateGoalProgressReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).GenerateGoalProgressReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/budgeting_service.TransactionService/GenerateGoalProgressReport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).GenerateGoalProgressReport(ctx, req.(*PrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

// TransactionService_ServiceDesc is the grpc.ServiceDesc for TransactionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TransactionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "budgeting_service.TransactionService",
	HandlerType: (*TransactionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _TransactionService_Create_Handler,
		},
		{
			MethodName: "GetById",
			Handler:    _TransactionService_GetById_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _TransactionService_GetAll_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _TransactionService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _TransactionService_Delete_Handler,
		},
		{
			MethodName: "GenerateSpendingReport",
			Handler:    _TransactionService_GenerateSpendingReport_Handler,
		},
		{
			MethodName: "GenerateIncomeReport",
			Handler:    _TransactionService_GenerateIncomeReport_Handler,
		},
		{
			MethodName: "GenerateBudgetPerformanceReport",
			Handler:    _TransactionService_GenerateBudgetPerformanceReport_Handler,
		},
		{
			MethodName: "GenerateGoalProgressReport",
			Handler:    _TransactionService_GenerateGoalProgressReport_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "transactions.proto",
}
