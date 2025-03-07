// Code generated by protoc-gen-psrpc v0.2.1, DO NOT EDIT.
// source: pkg/service/rpc/egress.proto

package rpc

import context "context"

import psrpc1 "github.com/livekit/psrpc"
import livekit "github.com/livekit/protocol/livekit"
import livekit3 "github.com/livekit/protocol/livekit"

// ===============================
// EgressInternal Client Interface
// ===============================

type EgressInternalClient interface {
	StartEgress(context.Context, *livekit3.StartEgressRequest, ...psrpc1.RequestOption) (*livekit.EgressInfo, error)

	ListActiveEgress(context.Context, *ListActiveEgressRequest, ...psrpc1.RequestOption) (<-chan *psrpc1.Response[*ListActiveEgressResponse], error)
}

// ===================================
// EgressInternal ServerImpl Interface
// ===================================

type EgressInternalServerImpl interface {
	StartEgress(context.Context, *livekit3.StartEgressRequest) (*livekit.EgressInfo, error)
	StartEgressAffinity(*livekit3.StartEgressRequest) float32

	ListActiveEgress(context.Context, *ListActiveEgressRequest) (*ListActiveEgressResponse, error)
}

// ===============================
// EgressInternal Server Interface
// ===============================

type EgressInternalServer interface {
	// Close and wait for pending RPCs to complete
	Shutdown()

	// Close immediately, without waiting for pending RPCs
	Kill()
}

// =====================
// EgressInternal Client
// =====================

type egressInternalClient struct {
	client *psrpc1.RPCClient
}

// NewEgressInternalClient creates a psrpc client that implements the EgressInternalClient interface.
func NewEgressInternalClient(clientID string, bus psrpc1.MessageBus, opts ...psrpc1.ClientOption) (EgressInternalClient, error) {
	rpcClient, err := psrpc1.NewRPCClient("EgressInternal", clientID, bus, opts...)
	if err != nil {
		return nil, err
	}

	return &egressInternalClient{
		client: rpcClient,
	}, nil
}

func (c *egressInternalClient) StartEgress(ctx context.Context, req *livekit3.StartEgressRequest, opts ...psrpc1.RequestOption) (*livekit.EgressInfo, error) {
	return psrpc1.RequestSingle[*livekit.EgressInfo](ctx, c.client, "StartEgress", "", req, opts...)
}

func (c *egressInternalClient) ListActiveEgress(ctx context.Context, req *ListActiveEgressRequest, opts ...psrpc1.RequestOption) (<-chan *psrpc1.Response[*ListActiveEgressResponse], error) {
	return psrpc1.RequestMulti[*ListActiveEgressResponse](ctx, c.client, "ListActiveEgress", "", req, opts...)
}

// =====================
// EgressInternal Server
// =====================

type egressInternalServer struct {
	svc EgressInternalServerImpl
	rpc *psrpc1.RPCServer
}

// NewEgressInternalServer builds a RPCServer that will route requests
// to the corresponding method in the provided svc implementation.
func NewEgressInternalServer(serverID string, svc EgressInternalServerImpl, bus psrpc1.MessageBus, opts ...psrpc1.ServerOption) (EgressInternalServer, error) {
	s := psrpc1.NewRPCServer("EgressInternal", serverID, bus, opts...)

	var err error
	err = psrpc1.RegisterHandler(s, "StartEgress", "", svc.StartEgress, svc.StartEgressAffinity)
	if err != nil {
		s.Close(false)
		return nil, err
	}

	err = psrpc1.RegisterHandler(s, "ListActiveEgress", "", svc.ListActiveEgress, nil)
	if err != nil {
		s.Close(false)
		return nil, err
	}

	return &egressInternalServer{
		svc: svc,
		rpc: s,
	}, nil
}

func (s *egressInternalServer) Shutdown() {
	s.rpc.Close(false)
}

func (s *egressInternalServer) Kill() {
	s.rpc.Close(true)
}

// ==============================
// EgressHandler Client Interface
// ==============================

type EgressHandlerClient interface {
	UpdateStream(context.Context, string, *livekit.UpdateStreamRequest, ...psrpc1.RequestOption) (*livekit.EgressInfo, error)

	StopEgress(context.Context, string, *livekit.StopEgressRequest, ...psrpc1.RequestOption) (*livekit.EgressInfo, error)

	SubscribeInfoUpdate(context.Context) (psrpc1.Subscription[*livekit.EgressInfo], error)
}

// ==================================
// EgressHandler ServerImpl Interface
// ==================================

type EgressHandlerServerImpl interface {
	UpdateStream(context.Context, *livekit.UpdateStreamRequest) (*livekit.EgressInfo, error)

	StopEgress(context.Context, *livekit.StopEgressRequest) (*livekit.EgressInfo, error)
}

// ==============================
// EgressHandler Server Interface
// ==============================

type EgressHandlerServer interface {
	RegisterUpdateStreamTopic(string) error
	DeregisterUpdateStreamTopic(string)

	RegisterStopEgressTopic(string) error
	DeregisterStopEgressTopic(string)

	PublishInfoUpdate(context.Context, *livekit.EgressInfo) error

	// Close and wait for pending RPCs to complete
	Shutdown()

	// Close immediately, without waiting for pending RPCs
	Kill()
}

// ====================
// EgressHandler Client
// ====================

type egressHandlerClient struct {
	client *psrpc1.RPCClient
}

// NewEgressHandlerClient creates a psrpc client that implements the EgressHandlerClient interface.
func NewEgressHandlerClient(clientID string, bus psrpc1.MessageBus, opts ...psrpc1.ClientOption) (EgressHandlerClient, error) {
	rpcClient, err := psrpc1.NewRPCClient("EgressHandler", clientID, bus, opts...)
	if err != nil {
		return nil, err
	}

	return &egressHandlerClient{
		client: rpcClient,
	}, nil
}

func (c *egressHandlerClient) UpdateStream(ctx context.Context, topic string, req *livekit.UpdateStreamRequest, opts ...psrpc1.RequestOption) (*livekit.EgressInfo, error) {
	return psrpc1.RequestSingle[*livekit.EgressInfo](ctx, c.client, "UpdateStream", topic, req, opts...)
}

func (c *egressHandlerClient) StopEgress(ctx context.Context, topic string, req *livekit.StopEgressRequest, opts ...psrpc1.RequestOption) (*livekit.EgressInfo, error) {
	return psrpc1.RequestSingle[*livekit.EgressInfo](ctx, c.client, "StopEgress", topic, req, opts...)
}

func (c *egressHandlerClient) SubscribeInfoUpdate(ctx context.Context) (psrpc1.Subscription[*livekit.EgressInfo], error) {
	return psrpc1.JoinQueue[*livekit.EgressInfo](ctx, c.client, "InfoUpdate", "")
}

// ====================
// EgressHandler Server
// ====================

type egressHandlerServer struct {
	svc EgressHandlerServerImpl
	rpc *psrpc1.RPCServer
}

// NewEgressHandlerServer builds a RPCServer that will route requests
// to the corresponding method in the provided svc implementation.
func NewEgressHandlerServer(serverID string, svc EgressHandlerServerImpl, bus psrpc1.MessageBus, opts ...psrpc1.ServerOption) (EgressHandlerServer, error) {
	s := psrpc1.NewRPCServer("EgressHandler", serverID, bus, opts...)

	return &egressHandlerServer{
		svc: svc,
		rpc: s,
	}, nil
}

func (s *egressHandlerServer) RegisterUpdateStreamTopic(topic string) error {
	return psrpc1.RegisterHandler(s.rpc, "UpdateStream", topic, s.svc.UpdateStream, nil)
}

func (s *egressHandlerServer) DeregisterUpdateStreamTopic(topic string) {
	s.rpc.DeregisterHandler("UpdateStream", topic)
}

func (s *egressHandlerServer) RegisterStopEgressTopic(topic string) error {
	return psrpc1.RegisterHandler(s.rpc, "StopEgress", topic, s.svc.StopEgress, nil)
}

func (s *egressHandlerServer) DeregisterStopEgressTopic(topic string) {
	s.rpc.DeregisterHandler("StopEgress", topic)
}

func (s *egressHandlerServer) PublishInfoUpdate(ctx context.Context, msg *livekit.EgressInfo) error {
	return s.rpc.Publish(ctx, "InfoUpdate", "", msg)
}

func (s *egressHandlerServer) Shutdown() {
	s.rpc.Close(false)
}

func (s *egressHandlerServer) Kill() {
	s.rpc.Close(true)
}

var psrpcFileDescriptor0 = []byte{
	// 329 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xd1, 0x4a, 0x02, 0x41,
	0x14, 0x65, 0x92, 0x2c, 0x6f, 0x19, 0x32, 0x05, 0xd9, 0xa6, 0x20, 0xfb, 0x14, 0x11, 0xbb, 0x60,
	0xf4, 0xd0, 0x63, 0x81, 0x94, 0xd0, 0x93, 0x12, 0x41, 0x2f, 0xb2, 0xce, 0xde, 0x6c, 0x50, 0x77,
	0xa6, 0x99, 0xab, 0xd0, 0x27, 0xf4, 0x3b, 0x7e, 0x4d, 0x9f, 0x13, 0xed, 0x8c, 0xb2, 0x18, 0x56,
	0x4f, 0xc3, 0x9c, 0x73, 0xef, 0xb9, 0xe7, 0x70, 0x2f, 0x34, 0xf4, 0x78, 0x14, 0x5b, 0x34, 0x73,
	0x29, 0x30, 0x36, 0x5a, 0xc4, 0x38, 0x32, 0x68, 0x6d, 0xa4, 0x8d, 0x22, 0xc5, 0x4b, 0x46, 0x8b,
	0xa0, 0xaa, 0x34, 0x49, 0x95, 0x79, 0x2c, 0x08, 0x26, 0x72, 0x8e, 0x63, 0x49, 0x03, 0xa3, 0xc5,
	0x40, 0x66, 0x84, 0x26, 0x4b, 0x26, 0x9e, 0x3b, 0x5a, 0x72, 0x45, 0x95, 0x70, 0x07, 0xb6, 0x3b,
	0x53, 0x4d, 0xef, 0xe1, 0x09, 0x1c, 0x3f, 0x48, 0x4b, 0x37, 0x82, 0xe4, 0x1c, 0x3b, 0x79, 0x49,
	0x0f, 0xdf, 0x66, 0x68, 0x29, 0xbc, 0x86, 0xfa, 0x4f, 0xca, 0x6a, 0x95, 0x59, 0xe4, 0x4d, 0x00,
	0xa7, 0x37, 0x90, 0xa9, 0xad, 0xb3, 0x56, 0xe9, 0xac, 0xd2, 0xab, 0x38, 0xa4, 0x9b, 0xda, 0xf6,
	0x82, 0xc1, 0x81, 0xeb, 0xe8, 0x7a, 0x37, 0xfc, 0x0e, 0xf6, 0xfa, 0x94, 0x18, 0x72, 0x30, 0x3f,
	0x8d, 0xbc, 0xaf, 0xa8, 0x80, 0xfa, 0xc9, 0xc1, 0xe1, 0x8a, 0x5c, 0x8a, 0xbc, 0xa8, 0xb0, 0xbc,
	0xf8, 0x60, 0x5b, 0x2d, 0xc6, 0x9f, 0xa0, 0xb6, 0x6e, 0x8b, 0x37, 0x22, 0xa3, 0x45, 0xb4, 0x21,
	0x48, 0xd0, 0xdc, 0xc0, 0xba, 0x2c, 0x4e, 0x78, 0x97, 0xb5, 0x3f, 0x19, 0x54, 0x1d, 0x75, 0x9f,
	0x64, 0xe9, 0x04, 0x0d, 0xef, 0xc2, 0xfe, 0xa3, 0x4e, 0x13, 0xc2, 0x3e, 0x19, 0x4c, 0xa6, 0xbc,
	0xb1, 0xf2, 0x55, 0x84, 0xff, 0x76, 0x5d, 0x67, 0xbc, 0x03, 0xd0, 0x27, 0xa5, 0xbd, 0xdf, 0xa0,
	0x90, 0x7e, 0x09, 0xfe, 0x4b, 0xe6, 0x0a, 0xe0, 0xfb, 0xef, 0xc6, 0x73, 0xc8, 0x83, 0xe5, 0x8b,
	0xfc, 0xa5, 0xad, 0xc6, 0x6e, 0x2f, 0x9e, 0xcf, 0x47, 0x92, 0x5e, 0x67, 0xc3, 0x48, 0xa8, 0x69,
	0xec, 0x0b, 0x57, 0xef, 0xda, 0xbd, 0x0d, 0xcb, 0xf9, 0x8d, 0x5c, 0x7e, 0x05, 0x00, 0x00, 0xff,
	0xff, 0x26, 0x11, 0xba, 0xff, 0x89, 0x02, 0x00, 0x00,
}
