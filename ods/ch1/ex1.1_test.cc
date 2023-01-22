#include <gtest/gtest.h>

namespace domain {
    int SomeBody(int a) {
        if (a > 0) {
            return a;
        }

        return -a;
    }
}

namespace tests {
    using domain::SomeBody;
    TEST(TestSuiteName, TestName) {
        EXPECT_EQ(SomeBody(10), 10);
    }

    TEST(TestSuiteName, TestName1) {
        EXPECT_EQ(SomeBody(-10), 10);
    }
}

