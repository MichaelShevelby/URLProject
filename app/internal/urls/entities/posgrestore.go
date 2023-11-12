package entities

import (
	"UrlProject/app/internal/db"
	"log"
)

type PostgreStore struct {
}

func (fs PostgreStore) SaveUrlsToStore(urls *map[string]string, url string, shorturl string) {
	log.Println("Сохранение данных в базу данных")

	db.ExecuteQuery("INSERT INTO urls (url, shorturl) VALUES ('"+url+"', '"+shorturl+"')", nil)
}

func (ps PostgreStore) LoadUrlFromStore(urls *map[string]string) {
	log.Println("Загрузка данных из базы данных")

	type UrlData struct {
		Url      string
		Shorturl string
	}

	var urlDatas []UrlData

	db.ExecuteSelectQueryMultipleResults("select url, shorturl from urls", &urlDatas)

	for _, urlData := range urlDatas {
		(*urls)[urlData.Url] = urlData.Shorturl
		log.Printf("Выгружена пара значений url: %s, shorturl: %s \n", urlData.Url, urlData.Shorturl)
	}
}
