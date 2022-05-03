#include <stdio.h>
#include <string.h>
#include "libGoEncryption.h"

int main(int argc, char *argv[]) {
       if (argc < 2 ) {
               printf("Usage: %s [data]\n", argv[0]);
               return 0;
       }

       char *data = argv[1];
       char *key = MakeAES256KeyToBase64();
       char *c = GoEncryptToBase64(data, key);
       char *p = GoDecryptToBase64(c, key);
       printf("AES-256 result =%s\n", p);
}
