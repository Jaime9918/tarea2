package main

import (
	"fmt"
	"context"
	"time"
	"google.golang.org/grpc"
	pb "github.com/Kendovvul/Ejemplo/Proto"
)

func lab1(){

	hostS := "dist014" //Host de un Laboratorio
	fmt.Println("Esperando Emergencias")

	for true{
		port := ":50055"  //puerto de la conexion con el laboratorio
		connS, err := grpc.Dial(hostS + port, grpc.WithInsecure()) //crea la conexion sincrona con el laboratorio

		if err != nil {
			panic("No se pudo conectar con el servidor" + err.Error())
		}
	
		defer connS.Close()
	
		serviceCliente := pb.NewMessageServiceClient(connS)
	
		for {
			//envia el mensaje al laboratorio
			res, err := serviceCliente.Intercambio(context.Background(), 
				&pb.Message{
					Body: "Equipo listo?",
				})
	
			if err != nil {
				panic("No se puede crear el mensaje " + err.Error())
			}


			fmt.Println(res.Body) //respuesta del laboratorio
			time.Sleep(10 * time.Second) //espera de 5 segundos
		}
	}


}

func main () {
	for true{
		go lab1()

		fmt.Println("olaallalallala")
		time.Sleep(100 * time.Second) //espera de 5 segundos
	}

}