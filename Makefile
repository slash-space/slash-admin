HASH=$(shell git rev-parse --short HEAD)
VERSION=$(shell git describe --tags --always)

APP 			= slash-admin
RELEASE_ROOT 	= release
RELEASE_TAG     = $(VERSION).$(GIT_HASH)


goctl-admin-api:
	goctls api go --swagger --api ./app/admin/cmd/api/desc/admin.api --dir ./app/admin/cmd/api -style=goZero
	@echo "generate goctl-admin-api success \n"
	swagger generate spec --output=./core.yml --scan-models
	@echo "generate swagger docs success \n"

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