HASH=$(shell git rev-parse --short HEAD)
VERSION=$(shell git describe --tags --always)

GOCTL_ENHANCED_VERSION = v1.4.2
GOCTL_ENHANCED_TEMPLATE:=$(shell pwd)/template/goctl

APP 			= slash-admin
RELEASE_ROOT 	= release

RELEASE_TAG     = $(VERSION).$(GIT_HASH)


init:
	go install github.com/zeromicro/go-zero/tools/goctl@v${GOCTL_ENHANCED_VERSION}

goctl-admin-api:
	goctls api go --swagger --api ./app/admin/cmd/api/desc/admin.api --dir ./app/admin/cmd/api -style=goZero

scan-admin-swagger:
	swagger generate spec --output=./core.yml --scan-models

start-admin-api:
	lsof -i:9100 | awk 'NR!=1 {print $2}' | xargs killall -9 || true
	modd

serve-admin-swagger:
	lsof -i:36666 | awk 'NR!=1 {print $2}' | xargs killall -9 || true
	swagger serve -F=swagger --port 36666 core.yml

ent-admin-orm:
	cd app/admin/ent && go run -mod=mod ./entc.go

goverter-admin:
	cd app/admin/cmd/api/internal/converter && go run ./goverter.go