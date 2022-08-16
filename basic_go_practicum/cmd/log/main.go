package main

import l "basic_go_practicum/pkg/logger"

type Some interface {
}

func main() {
	logger := l.NewLogExtended()
	logger.SetLogLevel(l.LogLevelWarning)
	logger.Infoln("Не должно напечататься")
	logger.Warnln("Hello")
	logger.Errorln("World")
	logger.Println("Debug")
}
