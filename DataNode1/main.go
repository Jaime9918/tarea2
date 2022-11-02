//esta se abre en el dist 013

package main

import (
	"bufio"
	"context"
	"fmt"
	"net"
	"os"
	"strings"
	"time"

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
	for fileScanner.Scan() {
		if fileScanner.Text() != "" {
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

func exit(texto string) {
	time.Sleep(3 * time.Second)
	fmt.Println(texto)
	defer os.Exit(0)
}

func (s *server) Intercambio(ctx context.Context, msg *pb.Message) (*pb.Message, error) {
	if msg.Body == "cierre" {
		fmt.Println("Solicitud de NameNode recibida, mensaje enviado: cierre de procesos")
		go exit("Se ha cerrado el dataNode1 (Grunth) satisfactoriamente")
		return &pb.Message{Body: "Se confirma cierre del dataNode"}, nil
	} else if len(msg.Body) < 15 {
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

		return &pb.Message{Body: "Se ha recibido y guardado el mensaje en DataNode1 con exito."}, nil
	}
}

func main() {
	fmt.Println("DataNode1 encendido")
	listener, err := net.Listen("tcp", ":50057") //conexion sincrona
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
