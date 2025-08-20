#include <stdint.h>
#include <stdio.h>

#include "conv_util.h"

uint32_t str_to_int_c(char *, int);
uint32_t str_to_int_s(char *, int);

int main(int argc, char *argv[]) {
    int ibase;
    int obase;
    uint32_t value;
    char input_str[CONV_MAX_CHARS];

    parse_params(argc, argv, &ibase, &obase, input_str);
    value = str_to_int_c(input_str, ibase);
    printf("C: %d\n", value);

    value = str_to_int_s(input_str, ibase);
    printf("Asm: %d\n", value);
}
