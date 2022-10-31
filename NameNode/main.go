package main

//esta se abre en el dist014
import (
	"bufio"
	"context"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strings"
	"time"

	pb "github.com/Kendovvul/Ejemplo/Proto"
	"google.golang.org/grpc"
)

func numeroAleatorio(valorMin int, valorMax int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return valorMin + rand.Intn(valorMax-valorMin)
}

type server struct {
	pb.UnimplementedMessageServiceServer
}

func revisar_id(id string) int {
	readFile, err := os.Open("DATA.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	contador := 0
	for fileScanner.Scan() {
		contador++
		if fileScanner.Text() != "" && (contador%2 == 1) {
			array := strings.Split(fileScanner.Text(), " : ")
			id_texto := array[1]
			if id_texto == id {
				return (-1)
			}
		}
	}
	readFile.Close()
	return (1)
}

func (s *server) Intercambio(ctx context.Context, msg *pb.Message) (*pb.Message, error) {
	remitente := "Combine"

	file, err := os.OpenFile("DATA.txt", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	if err != nil {
		fmt.Println("Error al leer el fichero ", err)
	}

	if msg.Body == "-1" {
		fmt.Println("Se ha ingresado una entrada invalida, no se ha guardado.")
		return &pb.Message{Body: "Se ingresó una opción invalida"}, nil
	}

	array := strings.Split(msg.Body, " : ")
	tipo, id, texto_msje := array[0], array[1], array[2]
	if revisar_id(id) == -1 {
		fmt.Println("Solicitud de " + remitente + " recibida, entrada invalida")
		return &pb.Message{Body: "Entrada invalida, el id ya existe"}, nil
	}
	fmt.Println("Solicitud de " + remitente + " recibida, mensaje enviado: " + msg.Body)
	//buscar si el id está en el txt

	dataNode_random := numeroAleatorio(1, 2)
	if dataNode_random == 1 {
		fmt.Println("Solicitud de Combine recibido, enviado a dataNode1 (Grunth)")
		texto := (tipo + " : " + id + " : " + "dataNode1 (Grunth)" + "\n")
		file.WriteString(texto)
		if err != nil {
			fmt.Println(err)
		}
		file.Close()
	} else if dataNode_random == 2 {
		fmt.Println("Solicitud de Combine recibido, enviado a dataNode2 (Synth)")
		texto := (tipo + " : " + id + " : " + "dataNode2 (Synth)" + "\n")
		file.WriteString(texto)
		if err != nil {
			fmt.Println(err)
		}
		file.Close()
	} else {
		fmt.Println("Solicitud de Combine recibido, enviado a dataNode3 (Cremator)")
		texto := (tipo + " : " + id + " : " + "dataNode3 (Cremator)" + "\n")
		file.WriteString(texto)
		if err != nil {
			fmt.Println(err)
		}
		file.Close()
	}
	texto := tipo + " " + id + " " + "'" + texto_msje + "'" + "\n"
	envio_datanode(dataNode_random, texto)
	//fmt.Println(texto)
	return &pb.Message{Body: "Se ha guardado el mensaje con exito."}, nil
}

func envio_datanode(datanode int, texto string) {
	hostS := "localhost"
	port := ":50056" //puerto de la conexion con el laboratorio
	if datanode == 1 {
		hostS = "dist013"
		hostS = "localhost"
		port = ":50057" //puerto de la conexion con el laboratorio
	} else if datanode == 2 {
		hostS = "dist015"
		hostS = "localhost"
		port = ":50058" //puerto de la conexion con el laboratorio
	} else {
		hostS = "dist016"
		hostS = "localhost"
		port = ":50059" //puerto de la conexion con el laboratorio
	}
	for true {
		connS, err := grpc.Dial(hostS+port, grpc.WithInsecure()) //crea la conexion sincrona con el laboratorio
		if err != nil {
			panic("No se pudo conectar con el servidor" + err.Error())
		}
		defer connS.Close()
		serviceCliente := pb.NewMessageServiceClient(connS)
		for {
			//envia el mensaje al laboratorio
			res, err := serviceCliente.Intercambio(context.Background(),
				&pb.Message{
					Body: texto,
				})
			if err != nil {
				panic("No se puede crear el mensaje " + err.Error())
			}
			fmt.Println(res.Body) //respuesta del laboratorio
			break
			//time.Sleep(10 * time.Second) //espera de 5 segundos
		}
		break
	}
}

func main() {
	fmt.Println("NameNode Encendido")
	file, _ := os.Create("DATA.txt")
	defer file.Close()
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
