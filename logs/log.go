package logs

import "log"

type Logger interface {
	Println(v ...interface{})
	Printf(fmt string, v ...interface{})
}

var logger Logger = &stdLog{}

func Init(l Logger) {
	logger = l
}

func Println(v ...interface{}) {
	logger.Println(v...)
}

func Printf(fmt string, v ...interface{}) {
	logger.Printf(fmt, v...)
}

type stdLog struct{}

func (s *stdLog) Println(v ...interface{}) {
	log.Println(v...)
}

func (s *stdLog) Printf(fmt string, v ...interface{}) {
	log.Printf(fmt, v...)
}
