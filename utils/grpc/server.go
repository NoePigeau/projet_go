package grpc

import (
	context "context"
	"project-go/models/payment"
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

	product := &product.Product{Name: in.Name, Price: float32(in.Price)}

	err := g.db.Create(product).Error
	if err != nil {
		return nil, err
	}

	return &CreateProductReply{ID: int32(product.ID), Name: product.Name, Price: int32(product.Price), CreatedAt: product.CreatedAt.String(), UpdatedAt: product.UpdatedAt.String()}, nil
}

func (g *grpcServer) CreatePayment(ctx context.Context, in *CreatePaymentRequest) (*CreatePaymentReply, error) {
	paymentFirst := &payment.Payment{}
	paymentErr := g.db.Where(&product.Product{ID: int(in.ProductID)}).First(&paymentFirst).Error
	if paymentErr != nil {
		return nil, status.Errorf(codes.Unimplemented, "product id not found")
	}

	payment := &payment.Payment{ProductID: int(in.ProductID), PricePaid: float32(in.PricePaid)}

	err := g.db.Create(&payment).Error
	if err != nil {
		return nil, err
	}

	return &CreatePaymentReply{ID: int32(payment.ID), PricePaid: int32(payment.PricePaid), ProductID: int32(payment.ProductID), CreatedAt: payment.CreatedAt.String(), UpdatedAt: payment.UpdatedAt.String()}, nil
}
