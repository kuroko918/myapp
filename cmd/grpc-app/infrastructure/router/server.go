package router

import (
	"net"
	"log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"go.uber.org/zap"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"

	"github.com/kuroko918/myapp/cmd/grpc-app/interfaces/servers"
	"github.com/kuroko918/myapp/cmd/grpc-app/interfaces/interceptor"
	"github.com/kuroko918/myapp/cmd/grpc-app/infrastructure"

	messagepb "github.com/kuroko918/myapp/cmd/grpc-app/proto/message"
)

func Server() {
	listenPort, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal(err)
	}

	authInterceptor := interceptor.NewAuthInterceptor(infrastructure.NewDbHandler())
	server := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				grpc_ctxtags.UnaryServerInterceptor(),
				grpc_opentracing.UnaryServerInterceptor(),
				grpc_prometheus.UnaryServerInterceptor,
				grpc_zap.UnaryServerInterceptor(logger),
				grpc_auth.UnaryServerInterceptor(authInterceptor.Auth),
				grpc_validator.UnaryServerInterceptor(),
				grpc_recovery.UnaryServerInterceptor(),
			),
		),
	)

	MessagesServer := servers.NewMessagesServer(infrastructure.NewDbHandler())
	messagepb.RegisterMessageServiceServer(server, MessagesServer)

	reflection.Register(server)
	go func() {
		if err := server.Serve(listenPort); err != nil {
		log.Fatal(err)
	}}()
}
