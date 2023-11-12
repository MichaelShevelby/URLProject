package requests

import (
	"UrlProject/app/internal/requests/entities"
	"UrlProject/app/pkg/api/generateshorturl"
	"UrlProject/app/pkg/api/getshorturl"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func StartListeningGRPC() {
	//Запускаем горутину-слушателя grpc-запросов
	go func() {
		lis, _ := net.Listen("tcp", ":9080")

		s := grpc.NewServer()
		reflection.Register(s)
		generateshorturl.RegisterGenerateShortUrlServer(s, &entities.GenerateShortUrlServer{})
		getshorturl.RegisterGetShortUrlServer(s, &entities.GetShortUrlServer{})

		errServe := s.Serve(lis)

		if errServe != nil {
			log.Fatalf("failed to serve: #{errServe}")
		}
	}()
}
