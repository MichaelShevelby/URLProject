package main

import (
	"UrlProject/app/internal/requests"
	"UrlProject/app/internal/urls"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func main() {
	var storeType string
	if len(os.Args) > 1 {
		storeType = os.Args[1]
	} else {
		log.Fatal("Не задан тип храненения информации (file или postgre)")
	}

	//Определяем типа хранилища
	urls.InitManager(storeType)

	//Включаем обработку REST и GRPC запросов
	requests.StartListeningGRPC()
	requests.StartListeningREST()
}
