FROM alpine

ENTRYPOINT ["/usr/bin/go-git-staged"]

COPY go-git-staged /usr/bin/go-git-staged

RUN apk add --no-cache git
