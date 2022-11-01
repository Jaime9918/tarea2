//esta se abre en el dist 015

package main

import (
	"bufio"
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

func revisar_id(id string) string {
	readFile, err := os.Open("DATA.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	contador := 0
	for fileScanner.Scan() {
		contador++
		if fileScanner.Text() != "" && (contador%2 == 0) {
			array := strings.Split(fileScanner.Text(), "'")
			new_array := strings.Split(array[0], " ")
			if new_array[1] == id {
				resultado := id + " '" + array[1] + "'"
				return (resultado)
			}
		}
	}
	readFile.Close()
	return ("Hubo un error")
}
func (s *server) Intercambio(ctx context.Context, msg *pb.Message) (*pb.Message, error) {
	if len(msg.Body) < 15 {
		fmt.Println("Solicitud de NameNode recibida, mensaje enviado: " + msg.Body)
		resultado_busqueda := revisar_id(msg.Body)
		fmt.Println("Se ha enviado la información solicitado al NameNode")
		return &pb.Message{Body: resultado_busqueda}, nil

	} else {
		array := strings.Split(msg.Body, "'")
		mensaje := array[1]
		fmt.Println("Solicitud de NameNode recibida, mensaje enviado: " + mensaje)
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

		return &pb.Message{Body: "Se ha recibido y guardado el mensaje en DataNode2 con exito."}, nil
	}
}

func main() {
	//Este se ejecuta en el dist015
	fmt.Println("DataNode2 encendido")
	listener, err := net.Listen("tcp", ":50058") //conexion sincrona
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
