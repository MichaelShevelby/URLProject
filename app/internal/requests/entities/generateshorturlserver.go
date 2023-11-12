package entities

import (
	"UrlProject/app/internal/urls"
	"UrlProject/app/internal/utils"
	"UrlProject/app/pkg/api/generateshorturl"
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GenerateShortUrlServer struct {
	generateshorturl.UnimplementedGenerateShortUrlServer
}

func (s *GenerateShortUrlServer) GenerateShortUrl(ctx context.Context, req *generateshorturl.GenerateShortUrlRequest) (*generateshorturl.GenerateShortUrlResponse, error) {

	//Достать ссылку из запроса
	url := (*req).Url

	fmt.Printf("Ссылка из запроса: %s \n", url)

	if !utils.IsUrlCorrect(url) {
		return nil, status.Errorf(codes.InvalidArgument, "Адрес '%s' имеет неверный формат: отсутствует префикс 'http://' или 'https://', обнаружены лишние символы", url)
	}

	//Проверить, что ссылка уже сохранена
	existshorturl, containsurl := (*urls.UrlManager.Urls)[url]

	if containsurl {
		return &generateshorturl.GenerateShortUrlResponse{Shorturl: existshorturl, StatusCode: 1}, nil
	}

	//Генерируем новую ссылку и проверяем на уникальность
	var shorturl string = ""

	for shorturl == "" {
		generatedshorturl := "cl.ru/" + utils.GenerateSequence(4)

		_, error := urls.UrlManager.GetUrlByShortUrl(shorturl)
		for error != nil {
			shorturl = generatedshorturl
			break
		}
	}

	urls.UrlManager.SaveUrl(url, shorturl)

	//return nil, status.Errorf(codes.Internal, "Internal server error")

	return &generateshorturl.GenerateShortUrlResponse{Shorturl: shorturl, StatusCode: 0}, nil
}
