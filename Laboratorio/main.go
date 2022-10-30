package main

import (
	"fmt"
	"context"
	"net"
	"google.golang.org/grpc"
	pb "github.com/Kendovvul/Ejemplo/Proto"
)

type server struct {
	pb.UnimplementedMessageServiceServer
}

func (s *server) Intercambio (ctx context.Context, msg *pb.Message) (*pb.Message, error){
	fmt.Println(msg.Body)
	return &pb.Message{Body: "NO",}, nil
}

func lab1(){

}

func main() {
	LabName := "Laboratiorio Pripyat" //nombre del laboratorio

	fmt.Println(LabName)

	listener, err := net.Listen("tcp", ":50055") //conexion sincrona
	if err != nil {
		panic("La conexion no se pudo crear" + err.Error())
	}

	serv := grpc.NewServer()
	for {
		pb.RegisterMessageServiceServer(serv, &server{})
		if err = serv.Serve(listener); err != nil {
			panic("El server no se pudo iniciar" + err.Error())
		}
	}
}