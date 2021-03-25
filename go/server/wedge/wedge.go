package main

import (
	"context"
	"log"
	"net"

	"github.com/Wappsto/wedge-api/go/wedge"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type wedgeServer struct{}

// func (*wedgeServer) SetModel(ctx context.Context, req *wedge.SetModelRequest) *wedge.Replay {
// 	return &wedge.Replay{
// 		Ok: true,
// 		Error: nil,
// 	}
// }
// func (*wedgeServer) SetDevice(ctx context.Context, req *wedge.SetDeviceRequest) *wedge.Replay {
// 	return &wedge.Replay{
// 		Ok: true,
// 		Error: nil,
// 	}
// }
// func (*wedgeServer) SetValue(ctx context.Context, req *wedge.SetValueRequest) *wedge.Replay {
// 	return &wedge.Replay{
// 		Ok: true,
// 		Error: nil,
// 	}
// }
// func (*wedgeServer) SetState(ctx context.Context, req *wedge.SetStateRequest) *wedge.Replay {
// 	return &wedge.Replay{
// 		Ok: true,
// 		Error: nil,
// 	}
// }
func (*wedgeServer) SetModel(context.Context, *wedge.SetModelRequest) (*wedge.Replay, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetModel not implemented")
}
func (*wedgeServer) SetDevice(context.Context, *wedge.SetDeviceRequest) (*wedge.Replay, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetDevice not implemented")
}
func (*wedgeServer) SetValue(context.Context, *wedge.SetValueRequest) (*wedge.Replay, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetValue not implemented")
}
func (*wedgeServer) SetState(context.Context, *wedge.SetStateRequest) (*wedge.Replay, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetState not implemented")
}
func (*wedgeServer) GetModel(context.Context, *wedge.GetModelRequest) (*wedge.Model, error) {
	return nil, nil
}
func (*wedgeServer) GetControl(context.Context, *wedge.GetControlRequest) (*wedge.Control, error) {
	return nil, nil
}

func main() {
	l, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatal("Could not create wedge listener on port 50051")
	}
	grpcServ := grpc.NewServer()
	wedge.RegisterWedgeServer(grpcServ, &wedgeServer{})
	if err := grpcServ.Serve(l); err != nil {
		log.Fatal("Failed to serve from wedge server")
	}
	log.Println("Wedge server running on port: 50051...")
}
