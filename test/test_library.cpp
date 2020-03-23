#include "library.h"
#include "gtest/gtest.h"

TEST(SumTest, Normal) {
    EXPECT_EQ(10, sum(5, 5));
}
