package standard

import (
	"fmt"
	"log"
	"time"
)

func Time() {
	now := time.Now()
	fmt.Printf("%v\n", now.Format("2006.01.02 15:04:05"))
	log.Println("这是一条很普通的日志。")
	v := "很普通的"
	log.Printf("这是一条%s日志。\n", v)
	log.Fatalln("这是一条会触发fatal的日志。")
	log.Panicln("这是一条会触发panic的日志。")
}
