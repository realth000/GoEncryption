SOURCE=test.c
CC=gcc
TARGET_WITH_ARCHIVE=test_archive
TARGET_WITH_SHARED=test_shared
GO_LIB_HEADER=libGoEncryption.h
GO_LIB_ARCHIVE=libGoEncryption.a
GO_LIB_SHARED=GoEncryption

FLAGS=-fstack-protector-all -D_FORTIFY_SOURCE=2 -Wl,-z,now,-z,relro,-z,noexecstack -fPIC -s -pthread

.PHONY: all
all: with-archive with-shared

.PHONY: with-archive
with-archive:
	$(CC) -o $(TARGET_WITH_ARCHIVE) $(SOURCE) $(GO_LIB_ARCHIVE) $(FLAGS)

.PHONY: with-shared
with-shared:
	$(CC) -o $(TARGET_WITH_SHARED) $(SOURCE) -I./ -L./ -l$(GO_LIB_SHARED) $(FLAGS) -Wl,-rpath=./

.PHONY: clean
clean:
	$(RM) $(TARGET_WITH_ARCHIVE) $(TARGET_WITH_SHARED)
