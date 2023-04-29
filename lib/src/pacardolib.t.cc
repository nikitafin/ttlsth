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

TEST(NthPrimeTest, Test1) {
    EXPECT_EQ(prime(1), 2);
}

TEST(NthPrimeTest, Test2) {
    EXPECT_EQ(prime(2), 3);
}

TEST(NthPrimeTest, Test5) {
    EXPECT_EQ(prime(5), 11);
}

TEST(NthPrimeTest, Test10) {
    EXPECT_EQ(prime(10), 29);
}

TEST(NthPrimeTest, Test50) {
    EXPECT_EQ(prime(50), 229);
}

TEST(NthPrimeTest, Test100) {
    EXPECT_EQ(prime(100), 541);
}

TEST(NthPrimeTest, Test200) {
    EXPECT_EQ(prime(200), 1223);
}

TEST(NthPrimeTest, Test500) {
    EXPECT_EQ(prime(500), 3571);
}

TEST(NthPrimeTest, Test1000) {
    EXPECT_EQ(prime(1000), 7919);
}

TEST(NthPrimeTest, Test5000) {
    EXPECT_EQ(prime(5000), 48611);
}
