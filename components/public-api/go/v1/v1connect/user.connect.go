// Copyright (c) 2023 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: gitpod/v1/user.proto

package v1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/gitpod-io/gitpod/components/public-api/go/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// UserServiceName is the fully-qualified name of the UserService service.
	UserServiceName = "gitpod.v1.UserService"
)

// UserServiceClient is a client for the gitpod.v1.UserService service.
type UserServiceClient interface {
	// GetAuthenticatedUser allows to retrieve the current user.
	GetAuthenticatedUser(context.Context, *connect_go.Request[v1.GetAuthenticatedUserRequest]) (*connect_go.Response[v1.GetAuthenticatedUserResponse], error)
	// UpdateUser updates the properties of a user.
	UpdateUser(context.Context, *connect_go.Request[v1.UpdateUserRequest]) (*connect_go.Response[v1.UpdateUserResponse], error)
	// SetWorkspaceAutoStartOptions updates the auto start options for the Gitpod Dashboard.
	// +internal - only used by the Gitpod Dashboard.
	SetWorkspaceAutoStartOptions(context.Context, *connect_go.Request[v1.SetWorkspaceAutoStartOptionsRequest]) (*connect_go.Response[v1.SetWorkspaceAutoStartOptionsResponse], error)
}

// NewUserServiceClient constructs a client for the gitpod.v1.UserService service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewUserServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) UserServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &userServiceClient{
		getAuthenticatedUser: connect_go.NewClient[v1.GetAuthenticatedUserRequest, v1.GetAuthenticatedUserResponse](
			httpClient,
			baseURL+"/gitpod.v1.UserService/GetAuthenticatedUser",
			opts...,
		),
		updateUser: connect_go.NewClient[v1.UpdateUserRequest, v1.UpdateUserResponse](
			httpClient,
			baseURL+"/gitpod.v1.UserService/UpdateUser",
			opts...,
		),
		setWorkspaceAutoStartOptions: connect_go.NewClient[v1.SetWorkspaceAutoStartOptionsRequest, v1.SetWorkspaceAutoStartOptionsResponse](
			httpClient,
			baseURL+"/gitpod.v1.UserService/SetWorkspaceAutoStartOptions",
			opts...,
		),
	}
}

// userServiceClient implements UserServiceClient.
type userServiceClient struct {
	getAuthenticatedUser         *connect_go.Client[v1.GetAuthenticatedUserRequest, v1.GetAuthenticatedUserResponse]
	updateUser                   *connect_go.Client[v1.UpdateUserRequest, v1.UpdateUserResponse]
	setWorkspaceAutoStartOptions *connect_go.Client[v1.SetWorkspaceAutoStartOptionsRequest, v1.SetWorkspaceAutoStartOptionsResponse]
}

// GetAuthenticatedUser calls gitpod.v1.UserService.GetAuthenticatedUser.
func (c *userServiceClient) GetAuthenticatedUser(ctx context.Context, req *connect_go.Request[v1.GetAuthenticatedUserRequest]) (*connect_go.Response[v1.GetAuthenticatedUserResponse], error) {
	return c.getAuthenticatedUser.CallUnary(ctx, req)
}

// UpdateUser calls gitpod.v1.UserService.UpdateUser.
func (c *userServiceClient) UpdateUser(ctx context.Context, req *connect_go.Request[v1.UpdateUserRequest]) (*connect_go.Response[v1.UpdateUserResponse], error) {
	return c.updateUser.CallUnary(ctx, req)
}

// SetWorkspaceAutoStartOptions calls gitpod.v1.UserService.SetWorkspaceAutoStartOptions.
func (c *userServiceClient) SetWorkspaceAutoStartOptions(ctx context.Context, req *connect_go.Request[v1.SetWorkspaceAutoStartOptionsRequest]) (*connect_go.Response[v1.SetWorkspaceAutoStartOptionsResponse], error) {
	return c.setWorkspaceAutoStartOptions.CallUnary(ctx, req)
}

// UserServiceHandler is an implementation of the gitpod.v1.UserService service.
type UserServiceHandler interface {
	// GetAuthenticatedUser allows to retrieve the current user.
	GetAuthenticatedUser(context.Context, *connect_go.Request[v1.GetAuthenticatedUserRequest]) (*connect_go.Response[v1.GetAuthenticatedUserResponse], error)
	// UpdateUser updates the properties of a user.
	UpdateUser(context.Context, *connect_go.Request[v1.UpdateUserRequest]) (*connect_go.Response[v1.UpdateUserResponse], error)
	// SetWorkspaceAutoStartOptions updates the auto start options for the Gitpod Dashboard.
	// +internal - only used by the Gitpod Dashboard.
	SetWorkspaceAutoStartOptions(context.Context, *connect_go.Request[v1.SetWorkspaceAutoStartOptionsRequest]) (*connect_go.Response[v1.SetWorkspaceAutoStartOptionsResponse], error)
}

// NewUserServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewUserServiceHandler(svc UserServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/gitpod.v1.UserService/GetAuthenticatedUser", connect_go.NewUnaryHandler(
		"/gitpod.v1.UserService/GetAuthenticatedUser",
		svc.GetAuthenticatedUser,
		opts...,
	))
	mux.Handle("/gitpod.v1.UserService/UpdateUser", connect_go.NewUnaryHandler(
		"/gitpod.v1.UserService/UpdateUser",
		svc.UpdateUser,
		opts...,
	))
	mux.Handle("/gitpod.v1.UserService/SetWorkspaceAutoStartOptions", connect_go.NewUnaryHandler(
		"/gitpod.v1.UserService/SetWorkspaceAutoStartOptions",
		svc.SetWorkspaceAutoStartOptions,
		opts...,
	))
	return "/gitpod.v1.UserService/", mux
}

// UnimplementedUserServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedUserServiceHandler struct{}

func (UnimplementedUserServiceHandler) GetAuthenticatedUser(context.Context, *connect_go.Request[v1.GetAuthenticatedUserRequest]) (*connect_go.Response[v1.GetAuthenticatedUserResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("gitpod.v1.UserService.GetAuthenticatedUser is not implemented"))
}

func (UnimplementedUserServiceHandler) UpdateUser(context.Context, *connect_go.Request[v1.UpdateUserRequest]) (*connect_go.Response[v1.UpdateUserResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("gitpod.v1.UserService.UpdateUser is not implemented"))
}

func (UnimplementedUserServiceHandler) SetWorkspaceAutoStartOptions(context.Context, *connect_go.Request[v1.SetWorkspaceAutoStartOptionsRequest]) (*connect_go.Response[v1.SetWorkspaceAutoStartOptionsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("gitpod.v1.UserService.SetWorkspaceAutoStartOptions is not implemented"))
}