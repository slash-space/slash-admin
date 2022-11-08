HASH=$(shell git rev-parse --short HEAD)
VERSION=$(shell git describe --tags --always)

APP 			= slash-admin
RELEASE_ROOT 	= release
RELEASE_TAG     = $(VERSION).$(GIT_HASH)

RED  =  "\e[31;1m"
GREEN = "\e[32;1m"
YELLOW = "\e[33;1m"
BLUE  = "\e[34;1m"
PURPLE = "\e[35;1m"
CYAN  = "\e[36;1m"

goctl-admin-api:
	goctls api go --swagger --api ./app/admin/cmd/api/desc/admin.api --dir ./app/admin/cmd/api -style=goZero
	@printf $(GREEN)"[SUCCESS] generate goctl-admin-api success"

scan-admin-swagger:
	swagger generate spec --output=./core.yml --scan-models
	@printf $(GREEN)"[SUCCESS] scan-admin-swagger success"

start-admin-api:
	lsof -i:9100 | awk 'NR!=1 {print $2}' | xargs killall -9 || true
	@printf $(GREEN)"[SUCCESS] kill admin-api success"
	modd

ent-admin-orm:
	cd app/admin/ent && go run -mod=mod ./entc.go
	@printf $(GREEN)"[SUCCESS] generate ent-admin-orm success"

goverter-admin:
	cd app/admin/cmd/api/internal/converter && go run ./goverter.go
	@printf $(GREEN)"[SUCCESS] generate goverter-admin success"
