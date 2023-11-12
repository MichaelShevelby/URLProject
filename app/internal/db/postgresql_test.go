package db

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSelect(t *testing.T) {
	ExecuteQuery("insert into urls (url, shorturl) VALUES ('example.com', 'cl.ru/TEST')", nil)

	type TestStruct struct {
		Id       int
		Url      string
		Shorturl string
	}

	var testStructs []TestStruct

	ExecuteSelectQueryMultipleResults("select * from urls order by id desc", &testStructs)

	assert.Equal(t, "example.com", testStructs[0].Url)

	t.Cleanup(func() {
		ExecuteQuery("DELETE from urls where url = 'example.com'", nil)
	})
}
