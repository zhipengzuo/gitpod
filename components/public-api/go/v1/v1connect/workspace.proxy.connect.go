// Copyright (c) 2023 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

// Code generated by protoc-proxy-gen. DO NOT EDIT.

package v1connect

import (
	context "context"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/gitpod-io/gitpod/components/public-api/go/v1"
)

var _ WorkspaceServiceHandler = (*ProxyWorkspaceServiceHandler)(nil)

type ProxyWorkspaceServiceHandler struct {
	Client v1.WorkspaceServiceClient
	UnimplementedWorkspaceServiceHandler
}

func (s *ProxyWorkspaceServiceHandler) GetWorkspace(ctx context.Context, req *connect_go.Request[v1.GetWorkspaceRequest]) (*connect_go.Response[v1.GetWorkspaceResponse], error) {
	resp, err := s.Client.GetWorkspace(ctx, req.Msg)
	if err != nil {
		// TODO(milan): Convert to correct status code
		return nil, err
	}

	return connect_go.NewResponse(resp), nil
}

func (s *ProxyWorkspaceServiceHandler) ListWorkspaces(ctx context.Context, req *connect_go.Request[v1.ListWorkspacesRequest]) (*connect_go.Response[v1.ListWorkspacesResponse], error) {
	resp, err := s.Client.ListWorkspaces(ctx, req.Msg)
	if err != nil {
		// TODO(milan): Convert to correct status code
		return nil, err
	}

	return connect_go.NewResponse(resp), nil
}

func (s *ProxyWorkspaceServiceHandler) CreateAndStartWorkspace(ctx context.Context, req *connect_go.Request[v1.CreateAndStartWorkspaceRequest]) (*connect_go.Response[v1.CreateAndStartWorkspaceResponse], error) {
	resp, err := s.Client.CreateAndStartWorkspace(ctx, req.Msg)
	if err != nil {
		// TODO(milan): Convert to correct status code
		return nil, err
	}

	return connect_go.NewResponse(resp), nil
}

func (s *ProxyWorkspaceServiceHandler) StartWorkspace(ctx context.Context, req *connect_go.Request[v1.StartWorkspaceRequest]) (*connect_go.Response[v1.StartWorkspaceResponse], error) {
	resp, err := s.Client.StartWorkspace(ctx, req.Msg)
	if err != nil {
		// TODO(milan): Convert to correct status code
		return nil, err
	}

	return connect_go.NewResponse(resp), nil
}

func (s *ProxyWorkspaceServiceHandler) GetWorkspaceDefaultImage(ctx context.Context, req *connect_go.Request[v1.GetWorkspaceDefaultImageRequest]) (*connect_go.Response[v1.GetWorkspaceDefaultImageResponse], error) {
	resp, err := s.Client.GetWorkspaceDefaultImage(ctx, req.Msg)
	if err != nil {
		// TODO(milan): Convert to correct status code
		return nil, err
	}

	return connect_go.NewResponse(resp), nil
}

func (s *ProxyWorkspaceServiceHandler) SendHeartBeat(ctx context.Context, req *connect_go.Request[v1.SendHeartBeatRequest]) (*connect_go.Response[v1.SendHeartBeatResponse], error) {
	resp, err := s.Client.SendHeartBeat(ctx, req.Msg)
	if err != nil {
		// TODO(milan): Convert to correct status code
		return nil, err
	}

	return connect_go.NewResponse(resp), nil
}

func (s *ProxyWorkspaceServiceHandler) GetWorkspaceOwnerToken(ctx context.Context, req *connect_go.Request[v1.GetWorkspaceOwnerTokenRequest]) (*connect_go.Response[v1.GetWorkspaceOwnerTokenResponse], error) {
	resp, err := s.Client.GetWorkspaceOwnerToken(ctx, req.Msg)
	if err != nil {
		// TODO(milan): Convert to correct status code
		return nil, err
	}

	return connect_go.NewResponse(resp), nil
}

func (s *ProxyWorkspaceServiceHandler) GetWorkspaceEditorCredentials(ctx context.Context, req *connect_go.Request[v1.GetWorkspaceEditorCredentialsRequest]) (*connect_go.Response[v1.GetWorkspaceEditorCredentialsResponse], error) {
	resp, err := s.Client.GetWorkspaceEditorCredentials(ctx, req.Msg)
	if err != nil {
		// TODO(milan): Convert to correct status code
		return nil, err
	}

	return connect_go.NewResponse(resp), nil
}