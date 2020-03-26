#ifndef CALCULATOR_BRIDGE_H
#define CALCULATOR_BRIDGE_H

#ifdef __cplusplus
extern "C" {
#endif
    void *NewCalculator(int a, int b);
    int CalcSum(void *c);
    int CalcSub(void *c);
    int CalcFactorial(void *c);
    void FreeCalculator(void *c);
#ifdef __cplusplus
}
#endif

#endif
