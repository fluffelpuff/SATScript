package log

import (
	"log"
)

func NODE_PRINTLN(stra string) {
	log.Println(stra)
}

func NODE_EPRINTLN(stra error) {
	log.Fatalln(stra)
}
