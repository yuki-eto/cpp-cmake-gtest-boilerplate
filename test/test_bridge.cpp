#include <gtest/gtest.h>
#include "bridge.h"

TEST(Bridge, Calc) {
    auto c = NewCalculator(10, 5);
    EXPECT_EQ(15, CalcSum(c));
    EXPECT_EQ(5, CalcSub(c));
    EXPECT_EQ(3628800, CalcFactorial(c));
    FreeCalculator(c);
}