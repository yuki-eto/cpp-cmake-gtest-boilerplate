#include "library.h"
#include "gtest/gtest.h"

TEST(SumTest, Normal) {
    EXPECT_EQ(10, sum(5, 5));
}

TEST(MulTest, Noraml) {
    EXPECT_EQ(20, mul(10, 2));
}

TEST(FactorialTest, Normal) {
    EXPECT_EQ(3628800, factorial(10));
}
