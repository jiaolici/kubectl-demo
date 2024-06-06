
export GO111MODULE=on

.PHONY: test
test:
	go test ./pkg/... ./cmd/... -coverprofile cover.out

.PHONY: bin
bin-ctr: fmt vet
	go build -o bin/kubectl-ctr.exe github.com/jiaolici/kubectl-demo/cmd/plugin/ctr

.PHONY: fmt
fmt:
	go fmt ./pkg/... ./cmd/...

.PHONY: vet
vet:
	go vet ./pkg/... ./cmd/...

.PHONY: kubernetes-deps
kubernetes-deps:
	go get k8s.io/client-go@v11.0.0
	go get k8s.io/api@kubernetes-1.14.0
	go get k8s.io/apimachinery@kubernetes-1.14.0
	go get k8s.io/cli-runtime@kubernetes-1.14.0

.PHONY: setup
setup:
	make -C setup