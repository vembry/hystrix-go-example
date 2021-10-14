package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	httpclient_x "playground/common/httpclient"

	"github.com/gin-gonic/gin"
)

func main() {

	// init gin
	r := gin.Default()

	// init handler
	initHandler(r)

	// run
	r.Run("localhost:3001")
}

func initHandler(r *gin.Engine) {

	// handler in-house httpclient lib
	r.GET("/ping-a", func(c *gin.Context) {

		var response map[string]interface{}

		//initialization
		client_x := httpclient_x.NewHttpClient(httpclient_x.Config{
			Host:                  "http://localhost:3002",
			Timeout:               10 * time.Second,
			RetryCount:            5,
			IsUsingCircuitBreaker: true,
			CbConfig: httpclient_x.CircuitBreakerConfig{
				SleepWindow:    10000,
				ErrorThreshold: 10,
				Fallback: func(e error) {
					somerandom("hi!")
				},
			},
		})

		// request initialization
		// header
		header := http.Header{}
		header.Set("x-some-header", "some-header-value")
		header.Set("Content-Type", "application/json")

		// body
		body := map[string]interface{}{
			"someKeyValue1": map[string]string{
				"someInnerKeyValue1": "some-inner-value-1",
			},
			"someKeyValue2": "some-value-2",
		}
		jsonBody, _ := json.Marshal(body)

		//usage
		resp, errResp := client_x.Post("/ping", header, bytes.NewBuffer([]byte(jsonBody)))
		if errResp != nil {
			fmt.Printf("service: %s", errResp)
			fmt.Println()
		} else {
			body, _ := ioutil.ReadAll(resp.Body)
			fmt.Println("service: ", string(body))
			json.Unmarshal(body, &response)
			defer resp.Body.Close()
		}

		if errResp != nil {
			c.JSON(500, gin.H{
				"message": "service unavailable",
				"error":   errResp,
			})
			return
		}
		c.JSON(200, response)
	})

}

func somerandom(str string) {
	fmt.Println("this some random fallback function, saying:", str)
}
