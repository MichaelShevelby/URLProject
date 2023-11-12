package entities

import "errors"

type URLManager struct {
	Store Store
	Urls  *map[string]string
}

func (urlManager URLManager) SaveUrl(url string, shorturl string) {
	(*urlManager.Urls)[url] = shorturl
	urlManager.Store.SaveUrlsToStore(urlManager.Urls, url, shorturl)
}

func (urlManager URLManager) GetUrlByShortUrl(shorturl string) (string, error) {
	for k, v := range *urlManager.Urls {
		if v == shorturl {
			return k, nil
		}
	}

	return "", errors.New("not found")
}
