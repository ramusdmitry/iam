package main

import (
	"auth/internal/app/authservice"
	api "auth/pkg/auth-api"
	"context"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	ctx := context.Background()
	run(ctx)
}

func run(ctx context.Context) {
	wg := &sync.WaitGroup{}
	ctx, cancel := context.WithCancelCause(ctx)
	defer cancel(nil)

	db, err := sqlx.Open("pgx", "postgres")
	if err != nil {
		log.Fatalf("failed to open db: %v", err)
	}
	if err := db.PingContext(ctx); err != nil {
		log.Fatalf("failed to ping db: %v", err)
	}

	portGRPC, srvGRPC := startGRPCServer()
	authService := authservice.NewAuthService()
	api.RegisterAuthServiceServer(srvGRPC, authService)

	if err := srvGRPC.Serve(portGRPC); err != nil {
		log.Fatalf("failed to serve grpc server: %v", err)
	}
	log.Printf("grpc server listening on port %d", portGRPC)

	gracefulStop(ctx, cancel, wg, srvGRPC)
}

func startGRPCServer() (net.Listener, *grpc.Server) {
	portGRPC, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	return portGRPC, grpcServer
}

func gracefulStop(ctx context.Context, cancel context.CancelCauseFunc, wg *sync.WaitGroup, grpcServer *grpc.Server) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(ch)

	select {
	case sig := <-ch:
		log.Printf("received signal: %s", sig)
	case <-ctx.Done():
		log.Printf("context done: %s", ctx.Err())
	}

	cancel(nil)
	grpcServer.GracefulStop()
	wg.Wait()
}
