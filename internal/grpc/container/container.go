package container

import (
	"emoLog/internal/grpc/handler"
	pb "emoLog/internal/grpc/proto"
	"emoLog/internal/service"

	"go.uber.org/dig"
	"google.golang.org/grpc"
)

func InitContanier(s *grpc.Server, container *dig.Container) {
	container.Invoke(func(userService service.UserService) {
		// pb.RegisterUserServiceServer(s grpc.ServiceRegistrar, srv pb.UserServiceServer)
		pb.RegisterHelloServiceServer(s, &handler.HelloServer{
			UserService: userService,
		})
	})

}
