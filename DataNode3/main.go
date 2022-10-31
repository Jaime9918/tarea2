//esta se abre en el dist 016

package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"strings"

	pb "github.com/Kendovvul/Ejemplo/Proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedMessageServiceServer
}

func (s *server) Intercambio(ctx context.Context, msg *pb.Message) (*pb.Message, error) {
	fmt.Println(msg.Body)
	array := strings.Split(msg.Body, "'")
	mensaje := array[1]
	fmt.Println("Solicitud de DataNode3 recibida, mensaje enviado: " + mensaje)
	file, err := os.OpenFile("DATA.txt", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.WriteString(msg.Body)
	if err != nil {
		fmt.Println(err)
	}
	file.Close()
	fmt.Println("Se ha guardado el mensaje con exito.")

	return &pb.Message{Body: "Se ha recibido y guardado el mensaje en DataNode1 con exito."}, nil
}

func main() {
	fmt.Println("DataNode3 encendido")
	listener, err := net.Listen("tcp", ":50059") //conexion sincrona
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
