package main

import (
	"fmt"

	customLogger "github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
	formatter "github.com/x-cray/logrus-prefixed-formatter"

	"jobqueue/router"
)

const TIME = "2006-01-02 15:04:05"

func main() {
	logger := customLogger.New()
	logger.Formatter = &formatter.TextFormatter{
		ForceFormatting: true,
		FullTimestamp:   true,
		TimestampFormat: TIME,
	}

	n := negroni.New()
	n.UseHandler(router.NewRouter(logger))
	n.Run(fmt.Sprintf(":%d", 8000))
}
