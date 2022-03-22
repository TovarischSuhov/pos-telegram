package main

import (
	"flag"

	"github.com/TovarischSuhov/pos-telegram/storage"
)

func main() {
	flag.Parse()
	storage := storage.GetStorage()
}
