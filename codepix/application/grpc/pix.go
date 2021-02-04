package grpc

import (
	"context"

	"github.com/erickrodrigs/codepix/codepix-go/application/grpc/pb"
	"github.com/erickrodrigs/codepix/codepix-go/application/usecase"
)

// PixGrpcService ...
type PixGrpcService struct {
	PixUseCase usecase.PixUseCase
	pb.UnimplementedPixServiceServer
}

// RegisterPixKey ...
func (service *PixGrpcService) RegisterPixKey(ctx context.Context, in *pb.PixKeyRegistration) (*pb.PixKeyCreatedResult, error) {
	key, err := service.PixUseCase.RegisterKey(in.Key, in.Kind, in.AccountId)

	if err != nil {
		return &pb.PixKeyCreatedResult{
			Status: "not created",
			Error:  err.Error(),
		}, err
	}

	return &pb.PixKeyCreatedResult{
		Id:     key.ID,
		Status: "created",
	}, nil
}

// Find ...
func (service *PixGrpcService) Find(ctx context.Context, in *pb.PixKey) (*pb.PixKeyInfo, error) {
	pixKey, err := service.PixUseCase.FindKey(in.Key, in.Kind)

	if err != nil {
		return &pb.PixKeyInfo{}, err
	}

	return &pb.PixKeyInfo{
		Id:   pixKey.ID,
		Kind: pixKey.Kind,
		Key:  pixKey.Key,
		Account: &pb.Account{
			AccountId:     pixKey.AccountID,
			AccountNumber: pixKey.Account.Number,
			BankId:        pixKey.Account.Bank.ID,
			BankName:      pixKey.Account.Bank.Name,
			OwnerName:     pixKey.Account.OwnerName,
			CreatedAt:     pixKey.Account.CreatedAt.String(),
		},
		CreatedAt: pixKey.CreatedAt.String(),
	}, nil
}

// NewPixGrpcService ...
func NewPixGrpcService(usecase usecase.PixUseCase) *PixGrpcService {
	return &PixGrpcService{
		PixUseCase: usecase,
	}
}
