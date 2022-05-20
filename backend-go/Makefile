.PHONY: build

default: build

BUILD_TIME=`date +%FT%T%z`
GIT_REVISION=`git rev-parse --short HEAD`
GIT_BRANCH=`git rev-parse --symbolic-full-name --abbrev-ref HEAD`

LDFLAGS=-ldflags "-s -w -X main.buildTime=${BUILD_TIME} -X main.gitRevision=${GIT_REVISION} -X main.gitBranch=${GIT_BRANCH}"

build:
	go build -o cmd/kemenag-scraper/kemenag-scraper ${LDFLAGS}  ./cmd/kemenag-scraper
	go build -o cmd/http-hikarie-islamy-be/http-hikarie-islamy-be ${LDFLAGS}  ./cmd/http-hikarie-islamy-be
