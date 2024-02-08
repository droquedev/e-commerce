#define standard colors
BLACK  := "\e[1;30m%s\e[0m\n"
RED    := "\e[1;31m%s\e[0m\n"
GREEN  := "\e[1;32m%s\e[0m\n"
YELLOW := "\e[1;33m%s\e[0m\n"
BLUE   := "\e[1;34m%s\e[0m\n"
PURPLE := "\e[1;35m%s\e[0m\n"
CYAN   := "\e[1;36m%s\e[0m\n"
GRAY   := "\e[1;37m%s\e[0m\n"

ifndef GOOS
GOOS := $(shell go env GOOS)
endif

ifndef GOARCH
GOARCH := $(shell go env GOARCH)
endif

bins: clean-bins products-service users-service notification-service

.PHONY: clean-bins
clean-bins:
	@printf $(YELLOW) "Delete old binaries..."
	@rm -fr bins

.PHONY: products-service
products-service:
	@printf $(BLUE) "Building products-service with OS: $(GOOS), ARCH: $(GOARCH)..."
	@mkdir -p bins
	@go build -o bins/products-service products-service/cmd/main.go

.PHONY: users-service
users-service:
	@printf $(BLUE) "Building users-service with OS: $(GOOS), ARCH: $(GOARCH)..."
	@mkdir -p bins
	@go build -o bins/users-service users-service/cmd/main.go

.PHONY: notification-service
notification-service:
	@printf $(BLUE) "Building notification-service with OS: $(GOOS), ARCH: $(GOARCH)..."
	@mkdir -p bins
	@go build -o bins/notification-service notification-service/cmd/main.go