CONSENTPROVIDER_VERSION=0.0.2

.PHONY: clean

image-build:
	@docker build -t loginprovider:$(CONSENTPROVIDER_VERSION) -f Dockerfile-build .
	@docker images | grep none | awk '{print $3}' | xargs -I {} docker rmi {} || exit 0

image-copy:
	@CGO_ENABLED=0 GOOS=linux go build -o bin/loginprovider github.com/ArnoSen/hydraloginconsentprovider/cmd/loginprovider
	@docker build -t loginprovider:$(CONSENTPROVIDER_VERSION) -f Dockerfile-copy .
	@docker images | grep none | awk '{print $3}' | xargs -I {} docker rmi {} || exit 0
