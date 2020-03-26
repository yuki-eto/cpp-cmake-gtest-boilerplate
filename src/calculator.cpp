#include "calculator.h"

Calculator::Calculator(int x, int y)
{
    a = x;
    b = y;
}

Calculator::~Calculator()
{
    a = 0;
    b = 0;
}

int Calculator::Sum()
{
    return a + b;
}

int Calculator::Sub()
{
    return a - b;
}

int Calculator::Factorial()
{
    int n = 1;
    for (int i = 1; i <= a; i++) {
        n *= i;
    }
    return n;
}
