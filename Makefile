all: bin/go-git-staged

PLATFORM=local

.PHONY: bin/go-git-staged
bin/go-git-staged:
	@docker build . --target bin \
	--output bin/ \
	--platform ${PLATFORM}

.PHONY: setup
setup:
	@docker run --rm -it \
	--volume $(PWD)/src \
	--workdir /src \
	golang:1.16-alpine ash
