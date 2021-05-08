package main

import (
	"github.com/sirupsen/logrus"
	"log"
)

func main(){
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	l := logrus.WithField("ku", "vu").WithField("some", "field")
	s, err := l.String()
	if err != nil {
		log.Println("err: ", err)
		return
	}

	log.Printf("s: %+v", s)
	b, err := l.Bytes()
	if err != nil {
		log.Println("err: ", err)
		return
	}

	log.Printf("bytes: %+v", string(b))
	log.Printf("l.Message: %+v", l.Message)
}