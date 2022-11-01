package main

//esta se abre en el dist014

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	pb "github.com/Kendovvul/Ejemplo/Proto"
	"google.golang.org/grpc"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}

func solicitud_cierre() int {
	fmt.Println("Se ha iniciado el cierre de procesos")
	respuesta := "4"
	hostS := "localhost"
	for true {
		port := ":50072"                                         //puerto de la conexion con el laboratorio
		connS, err := grpc.Dial(hostS+port, grpc.WithInsecure()) //crea la conexion sincrona con el laboratorio
		if err != nil {
			panic("No se pudo conectar con el servidor" + err.Error())
		}
		defer connS.Close()
		serviceCliente := pb.NewMessageServiceClient(connS)
		verificador := 1
		for {
			//envia el mensaje al laboratorio
			if verificador == 1 {
				res, err := serviceCliente.Intercambio(context.Background(),
					&pb.Message{
						Body: respuesta,
					})
				if err != nil {
					panic("No se puede crear el mensaje " + err.Error())
				}
				fmt.Println(res.Body) //respuesta del laboratorio
				verificador = 2
			}
			break
		}
		break
	}

	return 1
}

func solicitud_informacion(respuesta string) int {
	fmt.Println("######### DATOS CLASIFICADOS ###############")
	fmt.Println("")
	hostS := "localhost"
	for true {
		port := ":50072"                                         //puerto de la conexion con el laboratorio
		connS, err := grpc.Dial(hostS+port, grpc.WithInsecure()) //crea la conexion sincrona con el laboratorio
		if err != nil {
			panic("No se pudo conectar con el servidor" + err.Error())
		}
		defer connS.Close()
		serviceCliente := pb.NewMessageServiceClient(connS)
		verificador := 1
		for {
			//envia el mensaje al laboratorio
			if verificador == 1 {
				res, err := serviceCliente.Intercambio(context.Background(),
					&pb.Message{
						Body: respuesta,
					})
				if err != nil {
					panic("No se puede crear el mensaje " + err.Error())
				}
				fmt.Println(res.Body) //respuesta del laboratorio
				verificador = 2
			} else if verificador == 2 {
				res, err := serviceCliente.Intercambio(context.Background(),
					&pb.Message{
						Body: "recibido",
					})
				if err != nil {
					panic("No se puede crear el mensaje " + err.Error())
				}
				fmt.Println(res.Body) //respuesta del laboratorio
			}

			break
			//time.Sleep(10 * time.Second) //espera de 5 segundos
		}
		fmt.Println("########################################")
		break
	}

	return 1
}
func main() {
	fmt.Println("Consola de Rebeldes encendido")
	reader := bufio.NewReader(os.Stdin)
	for true {
		fmt.Println("-----------------------------------------------------------------------")
		fmt.Println("Seleccione la opción deseada: ")
		fmt.Println("1) Solicitar información")
		fmt.Println("2) Solicitar cierre")
		//menu()
		//go lab1()
		respuesta, _ := getInput("Respuesta: ", reader)
		if respuesta == "1" {
			for true {
				fmt.Println("Seleccione la información que requiere: ")
				fmt.Println("1) Logística")
				fmt.Println("2) Financiera")
				fmt.Println("3) Militar")
				fmt.Println("4) Atrás")
				respuesta, _ = getInput("Respuesta: ", reader)
				if respuesta == "1" || respuesta == "2" || respuesta == "3" {
					solicitud_informacion(respuesta)
				} else if respuesta == "4" {
					break
				}
			}
		} else if respuesta == "2" {
			fmt.Println("¿Estás seguro que quieres solicitar el cierre de todo?")
			respuesta, _ := getInput("Respuesta (1: si, 2: no): ", reader)
			if respuesta == "1" {
				solicitud_cierre()
			}
		} else {
			fmt.Println("Hay un problema")
			break
		}
	}

}
