package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func isPrime(value int) bool {
	for i := 2; i <= value/2; i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}

func main() {
	m := gin.Default()
	m.LoadHTMLGlob("templates/*")
	srv := &http.Server{Handler: m, ReadHeaderTimeout: 10 * time.Second}

	m.GET("/", func(c *gin.Context) {
		waitQuery := c.Request.URL.Query().Get("wait")
		primeQuery := c.Request.URL.Query().Get("prime")

		if waitQuery != "" {
			sleep, _ := strconv.Atoi(waitQuery)
			log.Printf("Sleep for %d seconds\n", sleep)
			time.Sleep(time.Duration(sleep) * time.Second)
		}
		if primeQuery != "" {
			val, _ := strconv.Atoi(primeQuery)
			log.Printf("Is %d prime: %t", val, isPrime(val))
		}
		c.HTML(http.StatusOK, "index.tmpl", nil)
	})

	if os.Getenv("PANIC") == "true" {
		panic("this is crashing")
	}

	port := "3000"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		panic(err)
	}

	go srv.Serve(listener)
	log.Println("Listening on 0.0.0.0:" + port)

	sigs := make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGTERM)
	<-sigs
	fmt.Println("SIGTERM, time to shutdown")
	listener.Close()
}
