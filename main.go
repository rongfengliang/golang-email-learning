package main

import (
	"demoapp/notify"
	"log"
	"net/http"
	_ "net/http/pprof"
	"sync"
)

func main() {

	emailnotidy := notify.NewEailNotidy2()
	// not working tcp out of order
	http.HandleFunc("/send", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("send email"))
		wg := sync.WaitGroup{}
		wg.Add(2)
		for i := 0; i < 2; i++ {
			go func(wg *sync.WaitGroup) {
				defer wg.Done()
				err := emailnotidy.Send("dalong@qq.com", "demoapp", map[string]interface{}{
					"content": "dalongdemoapp",
				})
				if err != nil {
					log.Println("err:", err.Error())
				}
			}(&wg)
		}
		wg.Wait()
	})
	http.HandleFunc("/send2", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("send email"))
		for _, to := range []string{
			"to1@example1.com",
			"to3@example2.com",
			"to4@example3.com",
		} {
			err := emailnotidy.Send(to, "demoapp", map[string]interface{}{
				"content": "dalongdemoapp",
			})
			if err != nil {
				log.Println("err:", err.Error())
			}
		}
	})
	http.ListenAndServe(":9090", nil)
}
