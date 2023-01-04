package grpc

import (
	context "context"
	"project-go/models/product"

	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type grpcServer struct {
	db *gorm.DB
	UnimplementedGRPCServer
}

func GetGrpcServer(db *gorm.DB) *grpcServer {
	return &grpcServer{db: db}
}

func (g *grpcServer) CreateProduct(ctx context.Context, in *CreateProductRequest) (*CreateProductReply, error) {
	productFirst := &product.Product{}
	productErr := g.db.Where(&product.Product{Name: in.Name}).First(productFirst).Error
	if productErr == nil {
		return nil, status.Errorf(codes.Unimplemented, "Name Already Used")
	}

	x := &product.Product{Name: in.Name, Price: float32(in.Price)}

	err := g.db.Create(x).Error
	if err != nil {
		return nil, err
	}

	return &CreateProductReply{ID: int32(x.ID), Name: x.Name, Price: int32(x.Price), CreatedAt: x.CreatedAt.String(), UpdatedAt: x.UpdatedAt.String()}, nil
}
