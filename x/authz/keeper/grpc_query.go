package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	proto "github.com/gogo/protobuf/proto"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/authz/types"
)

var _ types.QueryServer = Keeper{}

// Authorizations implements the Query/Authorizations gRPC method.
func (k Keeper) Authorizations(c context.Context, req *types.QueryAuthorizationsRequest) (*types.QueryAuthorizationsResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	if req.GranterAddr == "" {
		return nil, status.Errorf(codes.InvalidArgument, "invalid granter addr")
	}

	if req.GranteeAddr == "" {
		return nil, status.Errorf(codes.InvalidArgument, "invalid grantee addr")
	}

	granterAddr, err := sdk.AccAddressFromBech32(req.GranterAddr)

	if err != nil {
		return nil, err
	}
	granteeAddr, err := sdk.AccAddressFromBech32(req.GranteeAddr)

	if err != nil {
		return nil, err
	}
	ctx := sdk.UnwrapSDKContext(c)

	authorizations := k.GetAuthorizations(ctx, granteeAddr, granterAddr)

	if len(authorizations) == 0 {
		return nil, status.Errorf(codes.NotFound, "No authorizations found")
	}
	result := make([]*codectypes.Any, len(authorizations))
	for i, authorization := range authorizations {
		msg, ok := authorization.(proto.Message)
		if !ok {
			return nil, status.Errorf(codes.Internal, "can't protomarshal %T", msg)
		}

		authorizationAny, err := codectypes.NewAnyWithValue(msg)
		if err != nil {
			return nil, status.Errorf(codes.Internal, err.Error())
		}
		result[i] = authorizationAny
	}
	return &types.QueryAuthorizationsResponse{
		Authorization: result,
	}, nil
}

// Authorization implements the Query/Authorization gRPC method.
func (k Keeper) Authorization(c context.Context, req *types.QueryAuthorizationRequest) (*types.QueryAuthorizationResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	if req.GranterAddr == "" {
		return nil, status.Errorf(codes.InvalidArgument, "invalid granter addr")
	}

	if req.GranteeAddr == "" {
		return nil, status.Errorf(codes.InvalidArgument, "invalid grantee addr")
	}

	if req.MsgType == "" {
		return nil, status.Errorf(codes.InvalidArgument, "invalid msg-type")
	}

	granterAddr, err := sdk.AccAddressFromBech32(req.GranterAddr)

	if err != nil {
		return nil, err
	}
	granteeAddr, err := sdk.AccAddressFromBech32(req.GranteeAddr)

	if err != nil {
		return nil, err
	}
	ctx := sdk.UnwrapSDKContext(c)

	authorization, _ := k.GetAuthorization(ctx, granteeAddr, granterAddr, req.MsgType)
	if authorization == nil {
		return nil, status.Errorf(codes.NotFound, "no authorization found for %s type", req.MsgType)
	}

	msg, ok := authorization.(proto.Message)
	if !ok {
		return nil, status.Errorf(codes.Internal, "can't protomarshal %T", msg)
	}

	authorizationAny, err := codectypes.NewAnyWithValue(msg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &types.QueryAuthorizationResponse{Authorization: authorizationAny}, nil

}
