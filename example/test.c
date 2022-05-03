#include <stdio.h>
#include <string.h>
#include "libGoEncryption.h"

//GoSlice makeSlice(char *p) {
//	GoSlice ret = {(void*)*p, strlen(p), strlen(p)};
//	return ret;
//}

//int main_back(int argc, char *argv[]) {
//	printf("1 %d\n", argc);
//	if (argc < 2) {
//		printf("Usage: %s [data]\n", argv[0]);
//		return 0;
//	}
//	printf("2\n");
//	GoSlice data = makeSlice(argv[1]);
//	GoSlice key = MakeAES256Key();
//
//	printf("3\n");
//
//	GoSlice c = GoEncrypt(data, key);
//	GoSlice p = GoDecrypt(c, key);
//	printf("AES-256 result = %s\n", p);
//
//
//}

int main(int argc, char *argv[]) {
	if (argc < 2 ) {
		printf("Usage: %s [data]\n", argv[0]);
		return 0;
	}

	char *data = argv[1];
	char *key = MakeAES256Key();
	printf("data len=%ld key len=%ld\n", strlen(data), strlen(key));
	char *c = GoEncrypt(data, key);
	char *p = GoDecrypt(c, key);
	printf("AES-256 result =%s\n", p);
}
