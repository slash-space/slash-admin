HASH=$(shell git rev-parse --short HEAD)
VERSION=$(shell git describe --tags --always)

GOCTL_ENHANCED_VERSION = v1.4.2
GOCTL_ENHANCED_TEMPLATE:=$(shell pwd)/template/goctl

APP 			= slash-admin
RELEASE_ROOT 	= release

RELEASE_TAG     = $(VERSION).$(GIT_HASH)


init:
	go install github.com/zeromicro/go-zero/tools/goctl@v${GOCTL_ENHANCED_VERSION}

gen-admin-api:
	goctls api go --api ./app/admin/cmd/api/desc/admin.api --dir ./app/admin/cmd/api -style=goZero

gen-admin-swagger:
	swagger generate spec --output=./core.yml --scan-models