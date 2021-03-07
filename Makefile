all: bin/go-git-staged

PLATFORM=local

.PHONY: bin/go-git-staged
bin/go-git-staged:
	@docker build . --target bin \
	--output bin/ \
	--platform ${PLATFORM}
