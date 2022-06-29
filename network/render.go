package network

import (
	"log"
	"os"

	"github.com/valyala/fasttemplate"
)

func RenderString(dir string, data map[string]interface{}) string {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	file, err := os.ReadFile(dir)
	if err != nil {
		log.Println(err)
	}
	tmp := fasttemplate.New(string(file), "{{", "}}")
	return tmp.ExecuteString(data)
}
