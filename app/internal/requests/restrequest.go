package requests

import (
	"UrlProject/app/internal/urls"
	"UrlProject/app/internal/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type GenerateShortUrlRq struct {
	URL string `json:"url"`
}

type GetUrlByShortUrlRq struct {
	SHORTURL string `json:"shorturl"`
}

func StartListeningREST() {
	http.HandleFunc("/generateshorturl", generateShortUrlHandler)
	http.HandleFunc("/geturlbyshorturl", getUrlByShortUrlHandler)
	log.Fatal(http.ListenAndServe(":7080", nil))
}

func generateShortUrlHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		return
	}

	//Проверить запрос на корректность
	body, errRead := ioutil.ReadAll(r.Body)

	if errRead != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	var generateShortUrlRq GenerateShortUrlRq

	errUmarshal := json.Unmarshal(body, &generateShortUrlRq)

	if errUmarshal != nil {
		http.Error(w, "Failed to unmarshal json request body", http.StatusInternalServerError)
		return
	}

	//Достать ссылку из запроса
	urlFromRq := generateShortUrlRq.URL

	fmt.Printf("Ссылка из запроса: %s \n", urlFromRq)

	if !utils.IsUrlCorrect(urlFromRq) {
		fmt.Fprintf(w, `{"status": { "code": 1, "desc": "Адрес '%s' имеет неверный формат: отсутствует префикс 'http://' или 'https://', обнаружены лишние символы"}}`, urlFromRq)
		return
	}

	//Проверить, что ссылка уже сохранена
	existshorturl, containsurl := (*urls.UrlManager.Urls)[urlFromRq]

	if containsurl {
		fmt.Fprintf(w, `{"status": { "code": 1, "desc": "Адрес '%s' уже имеет сохраненную сокращенную ссылку"}, "shorturl":"%s"}`, urlFromRq, existshorturl)
		return
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

	urls.UrlManager.SaveUrl(urlFromRq, shorturl)

	//Генерируем новую ссылку
	fmt.Fprintf(w, `{"status": { "code": 0, "desc": "Успех"}, "shorturl":"%s"}`, shorturl)
}

func getUrlByShortUrlHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		return
	}

	//Проверить запрос на корректность
	body, errRead := ioutil.ReadAll(r.Body)

	if errRead != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	var getUrlByShortUrlRq GetUrlByShortUrlRq

	errUmarshal := json.Unmarshal(body, &getUrlByShortUrlRq)

	if errUmarshal != nil {
		http.Error(w, "Failed to unmarshal json request body", http.StatusInternalServerError)
		return
	}

	//Достать ссылку из запроса
	shorturl := getUrlByShortUrlRq.SHORTURL

	fmt.Printf("Ссылка из запроса: %s \n", shorturl)

	//Проверяем, существует ли ссылка с данной shorturl
	fullurl, errorNotExist := urls.UrlManager.GetUrlByShortUrl(shorturl)

	if errorNotExist != nil {
		fmt.Fprintf(w, `{"status": { "code": 1, "desc": "Короткая ссылка '%s' не привязана ни к одному полному адресу"}}`, shorturl)
		return
	}

	fmt.Fprintf(w, `{"status": { "code": 0, "desc": "Успех"}, "url":"%s"}`, fullurl)
}
