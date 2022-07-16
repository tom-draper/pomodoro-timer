TARGET_SYSTEM ?= $(OS)

EXT =
ifndef GOOS
  ifneq (,$(filter Windows%,$(TARGET_SYSTEM)))
    EXT =.exe
  endif
else
  ifeq ($(GOOS),windows)
    EXT = .exe
  endif
endif

GO_FLAGS   ?=
NAME       := pomodoro
OUTPUT_BIN ?= bin/$(NAME)$(ARCH)$(EXT)
PACKAGE    := github.com/tom-draper/pomodoro-timer
GIT_REV     = $(shell git rev-parse --short HEAD)
VERSION     = $(shell git describe --abbrev=0 --tags)

default: help

# Builds the win-amd64 CLI
build-win-amd: 
	@env GOOS=windows GOARCH=amd64 ARCH=-win-amd64 make build

# Builds the win-arm64 CLI
build-win-arm:
	@env GOOS=windows GOARCH=arm64 ARCH=-win-arm64 make build

# Builds the mac-amd64 CLI
build-mac-amd:
	@env GOOS=darwin GOARCH=amd64 ARCH=-mac-amd64 make build

# Builds the mac-arm64 CLI
build-mac-arm:
	@env GOOS=darwin GOARCH=arm64 ARCH=-mac-arm64 make build

# Builds the linux-amd64 CLI
build-linux-amd:
	@env GOOS=linux GOARCH=amd64 ARCH=-linux-amd64 make build
	
# Builds the CLI
build:
	@go build -trimpath ${GO_FLAGS} \
	-ldflags "-w -s -X 'github.com/tom-draper/pomodoro-timer/cmd.Version=${VERSION}'" \
	-a -tags netgo -o ${OUTPUT_BIN}

# Builds for all architectures
build-all: build-win-amd build-win-arm build-mac-amd build-mac-arm build-linux-amd 