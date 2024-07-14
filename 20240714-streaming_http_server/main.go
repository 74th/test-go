package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	httpServer *http.Server
}

type Config struct {
	Addr         string
	CertFilePath string
	KeyFilePath  string
}

func NewServer(conf Config) *Server {
	s := &Server{}
	mux := http.NewServeMux()
	mux.HandleFunc("/call", s.Call)

	s.httpServer = &http.Server{
		Addr:    conf.Addr,
		Handler: mux,
	}

	return s
}

// ストリーミングで出力する処理
func (s *Server) Call(res http.ResponseWriter, req *http.Request) {
	// 1.11 からは req.Context() が使える
	ctx := req.Context()

	log.Print(res, "/call")

	// Streaming をサポートしているか
	flusher, ok := res.(http.Flusher)
	if !ok {
		http.Error(res, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusOK)

	for i := 0; i < 30; i++ {
		// Contextキャンセルでクライアントから接続が閉じられたのを見られる
		if err := ctx.Err(); err != nil {
			log.Printf("connection closed by client: %v", err)
			return
		}
		fmt.Fprintln(res, "Hello World -", i)
		log.Println("Hello World -", i)

		// 現在までをクライアントに送る
		flusher.Flush()
		time.Sleep(time.Second)
	}
}

func (s *Server) Start() error {
	listener, err := net.Listen("tcp", s.httpServer.Addr)
	if err != nil {
		return err
	}

	go func() {
		err := s.httpServer.Serve(listener)
		if err == http.ErrServerClosed {
			panic(err)
		}
	}()

	return nil
}

func (s *Server) Stop() error {
	if err := s.httpServer.Close(); err != http.ErrServerClosed {
		return err
	}
	return nil
}

func main() {
	conf := Config{
		Addr: "0.0.0.0:8080",
	}

	sv := NewServer(conf)
	if err := sv.Start(); err != nil {
		log.Printf("Failed to start server: %v", err)
		os.Exit(1)
	}

	log.Printf("start server")

	sigC := make(chan os.Signal, 1)
	signal.Notify(sigC, syscall.SIGINT, syscall.SIGTERM)
	<-sigC

	if err := sv.Stop(); err != nil {
		log.Printf("Failed to start server: %v", err)
		os.Exit(1)
	}
}
