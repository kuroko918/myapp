package router

import (
  "context"
  "flag"
  "net/http"
  "google.golang.org/grpc"
  "github.com/gorilla/handlers"
  "github.com/golang/glog"
  "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

  messagegw "github.com/kuroko918/myapp/cmd/grpc-app/proto/message"
  usergw "github.com/kuroko918/myapp/cmd/grpc-app/proto/user"
)

var grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:50051", "gRPC server endpoint")

func run() error {
  ctx := context.Background()
  ctx, cancel := context.WithCancel(ctx)
  defer cancel()

  mux := runtime.NewServeMux()
  opts := []grpc.DialOption{grpc.WithInsecure()}
  if err := messagegw.RegisterMessageServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts); err != nil {
    return err
  }
  if err := usergw.RegisterUserServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts); err != nil {
    return err
  }

  handler := handlers.CORS(
    handlers.AllowedMethods([]string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"}),
    handlers.AllowedHeaders([]string{"Access-Control-Allow-Headers", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"}),
    handlers.AllowedOrigins([]string{"http://localhost:3000", "https://nuxt-dtlmvofcca-an.a.run.app"}),
  )(mux)

  return http.ListenAndServe(":8080", handler)
}

func ReverseProxyServer() {
  flag.Parse()
  defer glog.Flush()

  if err := run(); err != nil {
    glog.Fatal(err)
  }
}
