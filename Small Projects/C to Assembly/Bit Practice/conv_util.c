#include <stdlib.h>
#include <string.h>

#include "conv_util.h"

void parse_params(int argc, char **argv, int *ibase, int *obase, char *input_str) {
    for (int i = 1 /* skip argv[0] */; i < argc; i++) {
        if (!strcmp(argv[i], "-o")) {
            *obase = atoi(argv[i+1]);
            i++; /* already consumed the base */
        } else {
            if (!strncmp(argv[i], "0x", 2)) {
                *ibase = 16;
                strncpy(input_str, argv[i] + 2, CONV_MAX_CHARS);
            } else if (!strncmp(argv[i], "0b", 2)) {
                *ibase = 2;
                strncpy(input_str, argv[i] + 2, CONV_MAX_CHARS);
            } else {
                *ibase = 10;
                strncpy(input_str, argv[i], CONV_MAX_CHARS);
            }
        }
    }		
}
