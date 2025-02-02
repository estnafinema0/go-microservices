check_install:
	which swagger || go install github.com/go-swagger/go-swagger/cmd/swagger@latest

swagger: check_install
	GOR111MODULE=off swagger generate spec -o ./swagger.yaml --scan-models

generate_client:
	cd sdk && swagger generate client -f ../swagger.yaml -A product-api