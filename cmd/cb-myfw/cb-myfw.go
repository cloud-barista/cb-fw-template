package main

import (
	restapi "github.com/cloud-barista/cb-fw-template/pkg/api/rest"
	"github.com/cloud-barista/cb-fw-template/pkg/utils/config"
)

// @title CB-Myfw REST API
// @version 0.3.0-espresso
// @description CB-Myfw REST API
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://cloud-barista.github.io
// @contact.email contact-to-cloud-barista@googlegroups.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:4321
// @BasePath /myfw
func main() {

	config.Setup()
	restapi.Server()

}
