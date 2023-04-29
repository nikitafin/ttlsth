#ifndef PACARDOLIB_H
#define PACARDOLIB_H

namespace pacardolib {
    namespace types {
        using u8  = unsigned char;
        using u16 = unsigned short;
        using u32 = unsigned int;
        using u64 = unsigned long long;

        using i8  = char;
        using i16 = short;
        using i32 = int;
        using i64 = long long;
    } // namespace types

    [[nodiscard("TODO: MSG")]] bool is_prime(types::u64 n) noexcept;
    [[nodiscard("TODO: MSG")]] types::u64 prime(types::u64 n) noexcept;
} // namespace pacardolib

#endif // PACARDOLIB_H
