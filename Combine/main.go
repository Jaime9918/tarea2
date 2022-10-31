package main

//esta se abre en el dist013

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	pb "github.com/Kendovvul/Ejemplo/Proto"
	"google.golang.org/grpc"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}
func menu() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Bienvenido")
	for true {
		fmt.Println("-----------------------------------------------------------------------")
		fmt.Println("Seleccione el numero del tipo de información que desea ingresar: ")
		fmt.Println("1) Logística")
		fmt.Println("2) Financiera")
		fmt.Println("3) Militar")
		respuesta, _ := getInput("Respuesta: ", reader)

		respuesta1, _ := getInput("Ingrese el ID: ", reader)

		respuesta2, _ := getInput("Ingrese el texto: ", reader)
		if respuesta == "1" {
			respuesta = "LOGÍSTICA"
		} else if respuesta == "2" {
			respuesta = "FINANCIERA"
		} else if respuesta == "3" {
			respuesta = "MILITAR"
		} else {
			respuesta = "-1"
			fmt.Println("Entrada invalida")
		}
		fmt.Println("Se ha enviado el mensaje al NameNode")

		if respuesta == "-1" {
			envio_namenode(respuesta)

		} else {
			respuesta_final := respuesta + " : " + respuesta1 + " : " + respuesta2
			envio_namenode(respuesta_final)
		}

		time.Sleep(5 * time.Second) //espera de 5 segundos
	}

}

func envio_namenode(mensaje string) {
	//central - Combine
	//hostS := "dist014" //Host de un Laboratorio
	hostS := "localhost"
	for true {
		port := ":50055"                                         //puerto de la conexion con el laboratorio
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
					Body: mensaje,
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
	for true {
		menu()
		//go lab1()
		time.Sleep(5 * time.Second) //espera de 5 segundos
	}

}
