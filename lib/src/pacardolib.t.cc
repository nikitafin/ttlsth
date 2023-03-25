#include <gtest/gtest.h>

#include <pacardo.h>

using namespace pacardolib;

TEST(IsPrimeTest, Test1) {
    EXPECT_FALSE(is_prime(0));
}

TEST(IsPrimeTest, Test2) {
    EXPECT_FALSE(is_prime(1));
}

TEST(IsPrimeTest, Test3) {
    EXPECT_TRUE(is_prime(2));
}

TEST(IsPrimeTest, Test4) {
    EXPECT_TRUE(is_prime(3));
}

TEST(IsPrimeTest, Test5) {
    EXPECT_FALSE(is_prime(4));
}

TEST(IsPrimeTest, Test6) {
    EXPECT_TRUE(is_prime(5));
}

TEST(IsPrimeTest, Test7) {
    EXPECT_FALSE(is_prime(6));
}

TEST(IsPrimeTest, Test8) {
    EXPECT_TRUE(is_prime(7));
}

TEST(IsPrimeTest, Test9) {
    EXPECT_FALSE(is_prime(8));
}

TEST(IsPrimeTest, Test10) {
    EXPECT_FALSE(is_prime(9));
}

TEST(IsPrimeTest, Test11) {
    EXPECT_TRUE(is_prime(1000000007));
}

TEST(IsPrimeTest, Test12) {
    EXPECT_TRUE(is_prime(100000000003));
}

TEST(IsPrimeTest, Test13) {
    EXPECT_FALSE(is_prime(std::numeric_limits<types::u64>::max()));
}