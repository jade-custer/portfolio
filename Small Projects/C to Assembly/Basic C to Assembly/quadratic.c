#include <stdio.h>
#include <stdlib.h>

int quadratic_c(int, int,int,int);
int quadratic_s(int, int,int,int);

int main(int argc, char **argv) {
    int a = atoi(argv[1]);
    int b = atoi(argv[2]);
    int c = atoi(argv[3]);
    int d = atoi(argv[4]);

    int c_result = quadratic_c(a, b,c,d);
    printf("C: %d\n", c_result);

    int s_result = quadratic_s(a, b,c,d);
    printf("Asm: %d\n", s_result);

    return 0;
}
