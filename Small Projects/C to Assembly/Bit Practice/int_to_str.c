#include <stdint.h>
#include <stdio.h>

#include "conv_util.h"

uint32_t str_to_int_c(char *, int);

void int_to_str_c(uint32_t, char *, uint32_t);
void int_to_str_s(uint32_t, char *, uint32_t);


int main(int argc, char *argv[]) {
    int ibase;
    int obase;
    uint32_t value;
    char input_str[CONV_MAX_CHARS];
    char output_str[CONV_MAX_CHARS];

    parse_params(argc, argv, &ibase, &obase, input_str);
    value = str_to_int_c(input_str, ibase);
    
    int_to_str_c(value, output_str, obase);
    printf("C: %s\n", output_str);

    int_to_str_s(value, output_str, obase);
    printf("Asm: %s\n", output_str);
}
