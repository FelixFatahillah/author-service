package author

import (
	"author-service/internal/domain/author/services"
	"author-service/internal/infrastructure/pb"
	"context"
)

type AuthorHandlerGrpc struct {
	pb.UnimplementedAuthorServiceServer
	author services.AuthorService
}

func NewAuthorHandlerGrpcHandler(author services.AuthorService) *AuthorHandlerGrpc {
	authorServer := AuthorHandlerGrpc{
		author: author,
	}
	return &authorServer
}

func (grpc AuthorHandlerGrpc) GetAuthorById(ctx context.Context, req *pb.GetAuthorByIdRequest) (*pb.GetAuthorByIdResponse, error) {
	record, err := grpc.author.FindById(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.GetAuthorByIdResponse{
		FirstName:   record.FirstName,
		LastName:    *record.LastName,
		PhoneNumber: *record.PhoneNumber,
		Email:       record.Email,
	}, nil
}
