package storage

import (
	pb "github.com/Mubinabd/RestaurantService/genproto"
)
type StorageI interface {
	Orders() OrdersI
}

type OrdersI interface {
	CreateOrder(req *pb.Order) (*pb.Void, error)
	UpdateOrder(req *pb.Order) (*pb.Void, error)
	DeleteOrder(id *pb.ById) (*pb.Void, error)
	GetOrder(req *pb.ById) (*pb.Order, error)
	GetAllOrders(req *pb.Void) (*pb.Orders, error)
}
