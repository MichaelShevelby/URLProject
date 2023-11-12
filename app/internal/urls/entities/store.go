package entities

type Store interface {
	LoadUrlFromStore(urls *map[string]string)
	SaveUrlsToStore(urls *map[string]string, url string, shorturl string)
}
