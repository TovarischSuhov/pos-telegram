package main

import (
	"flag"

	"github.com/TovarischSuhov/pos-telegram/storage"
	_ "github.com/mailru/easyjson/gen"
)

func main() {
	flag.Parse()
	storage := storage.GetStorage()
}
