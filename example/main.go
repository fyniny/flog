package main

import (
	"github.com/fyniny/flog"
)

func main() {
	flog.Info("hello world")

	err := flog.Init(flog.Config{
		Level:      "debug",
		Filename:   "./test.log",
		MaxSize:    1024,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   true,
	})
	if err != nil {
		panic(err)
	}

	flog.Debug("hello debug")
	flog.Info("start")
}
