#include "pacardolib.h"

#include <cassert>
#include <cmath>

namespace pacardolib::types {
    static_assert(sizeof(u8) == 1, "Invalid u8 sizeof");
    static_assert(sizeof(u16) == 2, "Invalid u16 sizeof");
    static_assert(sizeof(u32) == 4, "Invalid u32 sizeof");
    static_assert(sizeof(u64) == 8, "Invalid u64 sizeof");

    static_assert(sizeof(i8) == 1, "Invalid i8 sizeof");
    static_assert(sizeof(i16) == 2, "Invalid i16 sizeof");
    static_assert(sizeof(i32) == 4, "Invalid i32 sizeof");
    static_assert(sizeof(i64) == 8, "Invalid i64 sizeof");
} // namespace pacardolib::types

namespace pacardolib {

    bool is_prime(types::u64 n) noexcept {
        if (n < 2) {
            return false;
        }
        if (n == 2) {
            return true;
        }

        for (types::u64 i = 2; i * i <= n; ++i) {
            if (n % i == 0) {
                return false;
            }
        }

        return true;
    }

    types::u64 prime(types::u64 n) noexcept {
        if (n == 1) {
            return 2;
        }
        types::u64 current_num = 3;
        types::u64 count       = 1;

        while (count < n) {
            if (current_num % 2 != 0 && is_prime(current_num)) {
                ++count;
            }
            current_num += 2;
        }

        return current_num - 2;
    }
} // namespace pacardolib
