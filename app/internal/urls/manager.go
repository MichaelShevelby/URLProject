package urls

import (
	"UrlProject/app/internal/urls/entities"
	"log"
)

var UrlManager entities.URLManager

func InitManager(nameStore string) {

	var store entities.Store
	if nameStore == "file" {
		store = entities.FileStore{}
	} else if nameStore == "postgre" {
		store = entities.PostgreStore{}
	} else {
		log.Fatalf("Заданный при запуске тип хранилища %s не соответствует допустимым значениям (file или postgre) \n", nameStore)
	}

	urls := map[string]string{}

	UrlManager = entities.URLManager{Store: store, Urls: &urls}

	store.LoadUrlFromStore(&urls)
}
