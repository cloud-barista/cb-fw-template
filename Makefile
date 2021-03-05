default:
	cd cmd/cb-myfw && $(MAKE)
swag:
	swag i -g ./cmd/cb-myfw/cb-myfw.go -o ./pkg/api/rest/swaggerDocs/
run:
	cd cmd/cb-myfw && $(MAKE) run
clean:
	cd cmd/cb-myfw && $(MAKE) clean
