GOPATH := ${PWD}
GO11MODULE := yes
export GOPLAGS := -mod=vendor

MODULE_PATH := ${PWD}/src/pricer

test:
	go test -v  ./src/basket_pricer/...

# vendor:
# 	cd ${MODULE_PATH} && go mod vendor

# vendor-init:
# 	cd ${MODULE_PATH} && go mod init pricer
