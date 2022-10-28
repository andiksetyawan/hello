package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	addr := flag.String("addr", ":80", ":3000")
	flag.Parse()

	c := make(chan os.Signal)
	signal.Notify(c)

	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	s := http.Server{Addr: *addr, Handler: mux}
	go func() {
		log.Println("server is running on ", *addr)
		err := s.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	<-c

	err := s.Shutdown(context.TODO())
	if err != nil {
		panic(err)
	}

	log.Println("server has been shutdown properly")
}
