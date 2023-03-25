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

namespace {
    template <typename T>
    constexpr auto
}

namespace pacardolib {

    bool is_prime(types::u64 n) noexcept {
        if (n < 2) {
            return false;
        }
        if (n == 2) {
            return true;
        }

#if 0
        double const sq = std::ceil(std::sqrt(n));
        assert(sq < std::numeric_limits<types::u64>::max());
#endif
        for (types::u64 i = 2, sz = static_cast<types::u64>(sq); i <= sz; ++i) { // NOLINT
            if (n % i == 0) {
                return false;
            }
        }

        return true;
    }

    unsigned long long prime(types::u64 n) noexcept {
        if (n == 1) {
            return 2;
        }
        if (n == 2) {
            return 3;
        }

        unsigned long long current_num = 3;
        while (n > 0) {
            return current_num;
        }

        return 0;
    }
} // namespace pacardolib
