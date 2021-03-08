package main

import (
	"os"

	"github.com/cloud-barista/cb-fw-template/pkg/apiserver"
)

func main() {

	apiserver.SPIDER_URL = os.Getenv("SPIDER_URL")

	// Run API Server
	apiserver.ApiServer()

}
