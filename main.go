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
	"context"
	"github.com/go-acme/lego/log"
	"sync"

	//"google.golang.org/grpc"
	//"google.golang.org/grpc/reflection"
	pb "github.com/makubit/shippy-service/proto/consignment"
	vesselProto "github.com/makubit/shippy-service/vessel-service/proto/vessel"
	"github.com/micro/go-micro"
)

const (
	port=":50051"
)

type repository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}

type Repository struct {
	mu sync.RWMutex
	consignments []*pb.Consignment
}

func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	repo.mu.Lock()
	updated := append(repo.consignments, consignment) //adding new request to list of consignments (it's not in proto)
	repo.consignments = updated
	repo.mu.Unlock()

	return consignment, nil
}

type service struct {
	repo repository
	vesselClient vesselProto.VesselServiceClient
}

func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) (error) {
	/*consignment, err := s.repo.Create(req)
	if err != nil {
		return err
	}

	//return &pb.Response{Created: true, Consignment: consignment}, nil
	res.Created = true
	res.Consignment = consignment
	return nil*/
	vesselResponse, err := s.vesselClient.FindAvalilable(context.Background(), &vesselProto.Specification{
		MaxWeight: req.Weight,
		Capacity: int32(len(req.Containers)),
	})
	log.Printf("Found vessel: %s\n", vesselResponse.Vessel.Name)
	if err != nil {
		return err
	}

	req.VesselId = vesselResponse.Vessel.Id

	consignment, err := s.repo.Create(req)
	if err != nil {
		return err
	}

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

	srv := micro.NewService(
		micro.Name("shippy-service.consignment.service"),
		)
	srv.Init()

	vesselClient := vesselProto.NewVesselServiceClient("shippy-service.vessel.service", srv.Client())

	pb.RegisterShippingServiceHandler(srv.Server(), &service{repo, vesselClient})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}