#--------------------- app --------------------
.PHONY: run
run:
	go run main.go

.PHONY: test
test:
	go test ./... -v

.PHONY: build
build:	
	if [ -z "$(FILE_NAME)" ]; then \
		GOOS=windows GOARCH=amd64 go build -o bin/warrantyapi/warrantyapi.exe main.go; \
	else \
		GOOS=windows GOARCH=amd64 go build -o bin/$(FILE_NAME)/api.exe main.go; \
	fi
#--------------------- k8s --------------------
K8SDIR=k8s

.PHONY: create-file-ns
k8s-create-ns-file:
	[ -d $(K8SDIR) ] || mkdir -p $(K8SDIR)
	if [ -z "$(NS)" ]; then \
		kubectl create namespace my-namespace --dry-run=client -o yaml > ./k8s/namespace.yaml; \
	else \
		kubectl create namespace $(NS) --dry-run=client -o yaml > ./k8s/namespace.yaml; \
	fi
	