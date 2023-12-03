#include "stdio.h"
#include "string.h"
#include "stdlib.h"

#define __USE_POSIX 1

int is_prefix(const char *pre, const char *str) {
    return strncmp(pre, str, strlen(pre)) == 0;
}

int str_includes(const char *str, const char *find) {
    return strstr(str, find) != NULL;
}

int main() {
    char str[] = " 2 blue, 2 green, 7 red; 3 red, 5 blue; 7 green, 14 blue, 3 red\n";
    const char *delim = ";";
    char *next_token;
    const char * token = strtok_r(str, delim, &next_token);
    while(token) {
        printf("%s||", token);
        token = strtok_r(NULL, delim, &next_token);
    }
}