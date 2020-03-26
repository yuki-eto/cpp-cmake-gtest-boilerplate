#ifndef CALCULATOR_BRIDGE_H
#define CALCULATOR_BRIDGE_H

extern "C" {
    void *NewCalculator(int a, int b);
    int CalcSum(void *c);
    int CalcSub(void *c);
    int CalcFactorial(void *c);
    void FreeCalculator(void *c);
}

#endif
