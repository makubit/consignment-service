package main

import (
	/*"context"
	"log"
	"net"
	"sync"
	pb "local.libraries/shippy-service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"*/
	"fmt"
	//"google.golang.org/grpc"
	//"google.golang.org/grpc/reflection"
	pb "github.com/makubit/shippy-service/proto/consignment"
	"github.com/micro/go-micro"
	"context"
	//"log"
	//"net"
	//"sync"
)

/*const (
	port=":50051"
)*/

type repository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}

type Repository struct {
	//mu sync.RWMutex
	consignments []*pb.Consignment
}

func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	//repo.mu.Lock()
	updated := append(repo.consignments, consignment) //adding new request to list of consignments (it's not in proto)
	repo.consignments = updated
	//repo.mu.Unlock()

	return consignment, nil
}

type service struct {
	repo repository
}

func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) (error) {
	consignment, err := s.repo.Create(req)
	if err != nil {
		return err
	}

	//return &pb.Response{Created: true, Consignment: consignment}, nil
	res.Created = true
	res.Consignment = consignment
	return nil
}

func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) (error) {
	consignments := s.repo.GetAll()
	//return &pb.Response{Created: true, Consignments: consignments}, nil
	res.Consignments = consignments
	return nil
}

func (repo *Repository) GetAll() []*pb.Consignment {
	return repo.consignments
}

func main() {
	repo := &Repository{}

	/*lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterShippingServiceServer(s, &service{repo})

	reflection.Register(s)

	log.Println("Running on port: ", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	 */
	srv := micro.NewService(
		micro.Name("consignment.service"),
		)
	srv.Init()

	pb.RegisterShippingServiceHandler(srv.Server(), &service{repo})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}