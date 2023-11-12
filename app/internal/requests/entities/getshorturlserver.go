package entities

import (
	"UrlProject/app/internal/urls"
	"UrlProject/app/pkg/api/getshorturl"
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
)

type GetShortUrlServer struct {
	getshorturl.UnimplementedGetShortUrlServer
}

func (c *GetShortUrlServer) GetShortUrl(ctx context.Context, req *getshorturl.GetShortUrlRequest) (*getshorturl.GetShortUrlResponse, error) {
	//Достать ссылку из запроса
	shorturl := (*req).Shorturl

	fmt.Printf("Ссылка из запроса: %s \n", shorturl)

	if !strings.HasPrefix(shorturl, "cl.ru/") {
		return nil, status.Errorf(codes.InvalidArgument, "Некорректный сокращенный адрес в запросе")
	}

	//Проверяем, существует ли ссылка с данной shorturl
	fullurl, errorNotExist := urls.UrlManager.GetUrlByShortUrl(shorturl)

	if errorNotExist != nil {
		return nil, status.Errorf(codes.NotFound, "Сокращенный адрес не существует")
	}

	return &getshorturl.GetShortUrlResponse{Url: fullurl, StatusCode: 0}, nil
}
