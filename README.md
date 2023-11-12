# UrlProject
Генерация короткого URL:
http://localhost:7080/generateshorturl

{
	"url": "https://google.com"
}
Получение короткого URL по полному:
http://localhost:7080/geturlbyshorturl
{
	"shorturl": "cl.ru/lM8L"
}
grpc также принимают url и shorturl по обоим запросам соответственно

Хост для отправки grpc запросов:
localhost:9080

Proto-файлы хранятся в проекте в папке app/api

Для запуска с сохранением в файл: go build file
Для запроска с сохранением в бд: go build postgre

При разработке бд была запущена через docker-compose(файл в корне проекта docker-compose.yml)
Необходима бд, ее данные нужно прописать в файле \app\internal\db\postgresql.go
В ней необходима таблица urls с полями id - автоинкремент, url - varchar, shorturl - varchar
