package customers

import (
	"context"
	"grpc-go-srv/common"
	"grpc-go-srv/routes"
)

type Server struct {
	routes.UnimplementedCustomersServer
}

func (s *Server) GetCustomer(context.Context, *routes.CustomerRequest) (*routes.CustomerReply, error) {
	var err error
	cust := routes.CustomerReply{}
	err = nil

	result := common.DB.Where("customer_number = ?", custid).First(&cust)
	if result.Error != nil {
		return cust, result.Error
	}

	return &cust, err
}
