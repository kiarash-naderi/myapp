// Copyright 2023-2024 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: buf/registry/module/v1/graph_service.proto

package modulev1connect

import (
	v1 "buf.build/gen/go/bufbuild/registry/protocolbuffers/go/buf/registry/module/v1"
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// GraphServiceName is the fully-qualified name of the GraphService service.
	GraphServiceName = "buf.registry.module.v1.GraphService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// GraphServiceGetGraphProcedure is the fully-qualified name of the GraphService's GetGraph RPC.
	GraphServiceGetGraphProcedure = "/buf.registry.module.v1.GraphService/GetGraph"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	graphServiceServiceDescriptor        = v1.File_buf_registry_module_v1_graph_service_proto.Services().ByName("GraphService")
	graphServiceGetGraphMethodDescriptor = graphServiceServiceDescriptor.Methods().ByName("GetGraph")
)

// GraphServiceClient is a client for the buf.registry.module.v1.GraphService service.
type GraphServiceClient interface {
	// Get a dependency graph that includes the given Commits.
	//
	// Commits will be resolved via the given ResourceRefs, and all Commits will be included in the
	// graph, along with their dependencies.
	//
	// A dependency graph is a directed acyclic graph.
	GetGraph(context.Context, *connect.Request[v1.GetGraphRequest]) (*connect.Response[v1.GetGraphResponse], error)
}

// NewGraphServiceClient constructs a client for the buf.registry.module.v1.GraphService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewGraphServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) GraphServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &graphServiceClient{
		getGraph: connect.NewClient[v1.GetGraphRequest, v1.GetGraphResponse](
			httpClient,
			baseURL+GraphServiceGetGraphProcedure,
			connect.WithSchema(graphServiceGetGraphMethodDescriptor),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
	}
}

// graphServiceClient implements GraphServiceClient.
type graphServiceClient struct {
	getGraph *connect.Client[v1.GetGraphRequest, v1.GetGraphResponse]
}

// GetGraph calls buf.registry.module.v1.GraphService.GetGraph.
func (c *graphServiceClient) GetGraph(ctx context.Context, req *connect.Request[v1.GetGraphRequest]) (*connect.Response[v1.GetGraphResponse], error) {
	return c.getGraph.CallUnary(ctx, req)
}

// GraphServiceHandler is an implementation of the buf.registry.module.v1.GraphService service.
type GraphServiceHandler interface {
	// Get a dependency graph that includes the given Commits.
	//
	// Commits will be resolved via the given ResourceRefs, and all Commits will be included in the
	// graph, along with their dependencies.
	//
	// A dependency graph is a directed acyclic graph.
	GetGraph(context.Context, *connect.Request[v1.GetGraphRequest]) (*connect.Response[v1.GetGraphResponse], error)
}

// NewGraphServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewGraphServiceHandler(svc GraphServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	graphServiceGetGraphHandler := connect.NewUnaryHandler(
		GraphServiceGetGraphProcedure,
		svc.GetGraph,
		connect.WithSchema(graphServiceGetGraphMethodDescriptor),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	return "/buf.registry.module.v1.GraphService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case GraphServiceGetGraphProcedure:
			graphServiceGetGraphHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedGraphServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedGraphServiceHandler struct{}

func (UnimplementedGraphServiceHandler) GetGraph(context.Context, *connect.Request[v1.GetGraphRequest]) (*connect.Response[v1.GetGraphResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("buf.registry.module.v1.GraphService.GetGraph is not implemented"))
}
