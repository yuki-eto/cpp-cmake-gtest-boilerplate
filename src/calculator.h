#ifndef CALCULATOR_CALCULATOR_H
#define CALCULATOR_CALCULATOR_H

class Calculator
{
public:
    Calculator(int x, int y);
    ~Calculator();

    int Sum();
    int Sub();
    int Factorial();
private:
    int a;
    int b;
};

#endif
