package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"sync"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"

	"github.com/Razzle131/grpc-service/internal/dns"
	desc "github.com/Razzle131/grpc-service/internal/generated"
	"github.com/Razzle131/grpc-service/internal/host"
)

const (
	grpcAddress string = "localhost:50051"
	httpAddress string = "localhost:8080"
)

func main() {
	ctx := context.Background()

	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()

		if err := startGrpcServer(); err != nil {
			log.Println("Cannot start grpc server")
		}
	}()

	go func() {
		defer wg.Done()

		if err := startHttpServer(ctx); err != nil {
			log.Println("Cannot start http server")
		}
	}()

	wg.Wait()
}

func startGrpcServer() error {
	grpcServer := grpc.NewServer(
		grpc.Creds(insecure.NewCredentials()),
	)

	reflection.Register(grpcServer)

	desc.RegisterCustomizerServer(grpcServer, &server{})

	list, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		return err
	}

	log.Printf("gRPC server listening at %v\n", grpcAddress)

	return grpcServer.Serve(list)
}

func startHttpServer(ctx context.Context) error {
	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	err := desc.RegisterCustomizerHandlerFromEndpoint(ctx, mux, grpcAddress, opts)
	if err != nil {
		return err
	}

	log.Printf("http server listening at %v\n", httpAddress)

	return http.ListenAndServe(httpAddress, mux)
}

type server struct {
	desc.UnimplementedCustomizerServer
}

func (s *server) GetHostName(ctx context.Context, req *desc.GetHostRequest) (*desc.HostResponse, error) {
	hostname, err := host.GetHost()
	if err != nil {
		return nil, err
	}

	return &desc.HostResponse{
		CurHostname: hostname,
	}, nil
}

func (s *server) SetHostName(ctx context.Context, req *desc.SetHostRequest) (*desc.HostResponse, error) {
	hostname, err := host.SetHost(req.GetNewHostname())
	if err != nil {
		return nil, err
	}

	return &desc.HostResponse{
		CurHostname: hostname,
	}, nil
}

func (s *server) GetDNS(context.Context, *desc.GetDnsRequest) (*desc.GetDnsResponse, error) {
	res, err := dns.GetDNS()
	if err != nil {
		return nil, err
	}

	return &desc.GetDnsResponse{
		DnsIps: res,
	}, nil
}

func (s *server) AddDNS(ctx context.Context, req *desc.SetDnsRequest) (*desc.SetDnsResponse, error) {
	if err := dns.AddDns(req.DnsIp); err != nil {
		return nil, err
	}

	return &desc.SetDnsResponse{}, nil
}

func (s *server) RemoveDNS(ctx context.Context, req *desc.SetDnsRequest) (*desc.SetDnsResponse, error) {
	if err := dns.RemoveDns(req.DnsIp); err != nil {
		return nil, err
	}

	return &desc.SetDnsResponse{}, nil
}
