#include "library.h"

int sum(int a, int b) {
    return a + b;
}

int mul(int a, int b) {
    return a * b;
}

int factorial(int a) {
    int n = 1;
    for (int i = 1; i <= a; i++) {
        n *= i;
    }
    return n;
}
