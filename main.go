package main

import (
	"demoapp/notify"
	"log"
)

func main() {
	emailnotidy := notify.NewEailNotidy()
	for {
		err := emailnotidy.Send("dalong@qq.com", "demoapp", map[string]interface{}{
			"content": "dalongdemoapp",
		})
		if err != nil {
			log.Println("err:", err.Error())
		}
	}
}
