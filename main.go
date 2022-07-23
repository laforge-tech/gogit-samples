package main

import (
	"os"
	
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/laforge-tech/gogit-samples/cmd"
)

func main()  {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "20060102 15:04:05.000"})

	cmd.Execute()
}
