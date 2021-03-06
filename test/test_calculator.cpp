#include "calculator.h"
#include "gtest/gtest.h"


TEST(Calculator, Sum) {
    auto c = new Calculator(10, 5);
    EXPECT_EQ(15, c->Sum());
    delete c;
}

TEST(Calculator, Sub) {
    auto c = new Calculator(10, 5);
    EXPECT_EQ(5, c->Sub());
    delete c;
}

TEST(Calculator, Factorial) {
    auto c = new Calculator(10, 5);
    EXPECT_EQ(3628800, c->Factorial());
    delete c;
}
