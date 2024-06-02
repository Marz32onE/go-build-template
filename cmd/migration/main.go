package main

import (
	config "github.com/marz32one/go-build-template/internal"
	"github.com/marz32one/go-build-template/pkg/storage"
	data "github.com/marz32one/go-build-template/pkg/storage/data"
)

func main() {
	config.Load()
	storage.DB.AutoMigrate(&data.Item{})
}
