#.PHONY: all docs build
#
#all: docs build
#
#docs:
#	go run main.go doc-gen
#
#build: 
#	go build -o app main.go

.PHONY: release $(TARGETS)
TARGETS := linux/amd64 linux/arm64 windows/amd64

temp = $(subst /, ,$@)
os = $(word 1, $(temp))
arch = $(word 2, $(temp))

release: $(TARGETS)

$(TARGETS):
	GOOS=$(os) GOARCH=$(arch) go build -o 'build/pod-$(os)-$(arch)' main.go
