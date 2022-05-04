SOURCE=main.go
SOURCE_DIR_EXAMPLE=example
TARGET_C_ARCHIVE=libGoEncryption.a
TARGET_C_SHARED=libGoEncryption.so
TARGET_C_HEADER=libGoEncryption.h
GO_CMD=go
GO_FLAGS=-O2 -fstack-protector-all -fPIC -D_FORTITY_SOURCE=2
GO_LDFLAGS='-linkmode=external -extldflags -Wl,-z,now,-z,relro,-z,noexecstack -s'

export GO111MODULE=on
export CGO_ENABLED=1
export CGO_CFLAGS=$(GO_FLAGS)
export CGO_CXXFLAGS=$(GO_FLAGS)

.PHONY: all
all: c-archive c-shared example

.PHONY: c-header
c-header:
	$(GO_CMD) tool cgo -exportheader $(TARGET_C_HEADER) $(SOURCE)
	$(RM) -rf _obj

.PHONY: c-archive
c-archive:
	$(GO_CMD) build -buildmode=c-archive -trimpath -ldflags $(GO_LDFLAGS) -o $(TARGET_C_ARCHIVE) $(SOURCE)

.PHONY: c-shared
c-shared:
	$(GO_CMD) build -buildmode=c-shared -trimpath -ldflags $(GO_LDFLAGS) -o $(TARGET_C_SHARED) $(SOURCE)

.PHONY: example
example: c-archive c-shared
	cp -f $(TARGET_C_HEADER) $(TARGET_C_ARCHIVE) $(TARGET_C_SHARED) $(SOURCE_DIR_EXAMPLE)
	$(MAKE) -C $(SOURCE_DIR_EXAMPLE)

.PHONY: clean
clean:
	$(RM) $(TARGET_C_HEADER) $(TARGET_C_ARCHIVE) $(TARGET_C_SHARED)
	$(MAKE) clean -C example
