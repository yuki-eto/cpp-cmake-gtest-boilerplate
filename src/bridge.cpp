#include "bridge.h"
#include "calculator.h"

void *NewCalculator(int a, int b)
{
    auto c = new Calculator(a, b);
    return c;
}

Calculator *AsCalculator(void *c)
{
    return (Calculator*)c;
}

int CalcSum(void *c)
{
    return AsCalculator(c)->Sum();
}

int CalcSub(void *c)
{
    return AsCalculator(c)->Sub();
}

int CalcFactorial(void *c)
{
    return AsCalculator(c)->Factorial();
}

void FreeCalculator(void *c)
{
    delete AsCalculator(c);
}