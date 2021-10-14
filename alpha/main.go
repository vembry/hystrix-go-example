package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/gin-gonic/gin"
)

func main() {

	// init gin
	r := gin.Default()

	// hystrix configs
	hystrix.ConfigureCommand("something", hystrix.CommandConfig{
		SleepWindow:            10000, // in ms
		RequestVolumeThreshold: 5,
	})

	// init handler
	initHandler(r)

	// run
	r.Run("localhost:3000")
}

func initHandler(r *gin.Engine) {

	// handler without circuit breaker
	r.GET("/ping-a", func(c *gin.Context) {

		//do task
		obj, err := doTask()
		if err != nil {
			c.JSON(500, gin.H{
				"message": "service not available",
				"error":   err,
			})
			return
		}

		c.JSON(200, obj)
	})

	// handler with circuit breaker
	r.GET("/ping-b", func(c *gin.Context) {
		var response map[string]string

		//circuit breaker starts
		err := hystrix.Do("something", func() error {
			//circuit breaker scopes

			//do tasks
			obj, err1 := doTask()
			if err1 != nil {
				return err1
			}
			response = obj
			return nil
		}, func(e error) error {

			//circuit breaker fallback function
			fmt.Println("hystrix fallback")
			fmt.Println(fmt.Sprintf("message: %v", e))
			return errors.New("hystrix fallback")
		})
		//circuit breaker ends

		if err != nil {
			c.JSON(500, gin.H{
				"message": "service not available",
				"error":   err,
			})
			return
		}

		c.JSON(200, response)
	})
}

//dummy task
func doTask() (map[string]string, error) {
	fmt.Println("doing task-a")
	fmt.Println("doing task-b")

	var response map[string]string
	resp, err1 := http.Get("http://localhost:3001/ping")
	if err1 != nil {
		fmt.Println("failing request")
		return response, errors.New("failing request")
	}

	body, err2 := io.ReadAll(resp.Body)
	if err2 != nil {
		fmt.Printf("error reading responses, message: %v", err1)
		fmt.Println("failing parse")
		return response, errors.New("failing parse")
	}
	json.Unmarshal(body, &response)

	fmt.Println("doing task-c")
	fmt.Println("doing task-d")

	return response, nil
}
