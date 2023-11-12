package main

import (
	"backend/api/controller"
	"backend/api/model"
	"database/sql"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var Db *sql.DB

func init() {
	model.DbConnect()
}

func main() {
	controller.StartServer()
	// ③ Ctrl+CでHTTPサーバー停止時にDBをクローズする
	closeDBWithSysCall()
}

// ③ Ctrl+CでHTTPサーバー停止時にDBをクローズする
func closeDBWithSysCall() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		s := <-sig
		log.Printf("received syscall, %v", s)

		if err := Db.Close(); err != nil {
			log.Fatal(err)
		}
		log.Printf("success: db.Close()")
		os.Exit(0)
	}()
}
