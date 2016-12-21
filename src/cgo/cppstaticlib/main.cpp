// g++ -o cpp main.cpp lib.a
#include <stdio.h>
#include "lib.h"

int main(int argc, char *argv[]) {
    printf("get version: %#x\n", get_version());
    return 0;
}
