package entities

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type FileStore struct {
}

func (fs FileStore) LoadUrlFromStore(urls *map[string]string) {
	log.Println("Загрузка данных из файла")

	content, _ := os.ReadFile("app/links.json")

	var fileUrls map[string]interface{}

	errUnmarshal := json.Unmarshal(content, &fileUrls)
	if errUnmarshal != nil {
		panic(errUnmarshal)
	}

	for key, value := range fileUrls {
		(*urls)[key] = value.(string)
	}
}

func (fs FileStore) SaveUrlsToStore(urls *map[string]string, url string, shorturl string) {
	log.Println("Сохранение данных в файл")

	file, errCreate := os.Create("app/links.json")

	if errCreate != nil {
		fmt.Println("Ошибка при создании файла: ", errCreate)
		return
	}

	defer file.Close()

	encoder := json.NewEncoder(file)

	errSave := encoder.Encode(urls)

	if errSave != nil {
		fmt.Println("Ошибка при сохранении в файл: ", errSave)
		return
	}
}
