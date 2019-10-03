CONSENTPROVIDER_VERSION=1.0.0
GITBRANCH=$(shell git branch | grep \* | cut -d ' ' -f2)

.PHONY: clean

image-build:
	@docker build -t consentprovider:$(CONSENTPROVIDER_VERSION)-$(GITBRANCH) -f Dockerfile-build .
	@docker images | grep none | awk '{print $3}' | xargs -I {} docker rmi {} || exit 0

image-copy:
	@CGO_ENABLED=0 GOOS=linux go build -o bin/loginprovider github.com/ArnoSen/hydraloginconsentprovider/cmd/loginprovider
	@docker build -t consentprovider:$(CONSENTPROVIDER_VERSION)-$(GITBRANCH) -f Dockerfile-copy .
	@docker images | grep none | awk '{print $3}' | xargs -I {} docker rmi {} || exit 0
