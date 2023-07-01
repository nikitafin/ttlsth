export module Common;

export import std.core;
export import std.filesystem;

namespace fs = std::filesystem;


export namespace Pgp::Types {
    using u8  = unsigned char;
    using u16 = unsigned short;
    using u32 = unsigned int;
    using u64 = unsigned long long;

    using i8  = char;
    using i16 = short;
    using i32 = int;
    using i64 = long long;
} // namespace Pgp::Types

namespace Pgp::Types {
    static_assert(sizeof(u8) == 1, "Invalid u8 sizeof");
    static_assert(sizeof(u16) == 2, "Invalid u16 sizeof");
    static_assert(sizeof(u32) == 4, "Invalid u32 sizeof");
    static_assert(sizeof(u64) == 8, "Invalid u64 sizeof");

    static_assert(sizeof(i8) == 1, "Invalid i8 sizeof");
    static_assert(sizeof(i16) == 2, "Invalid i16 sizeof");
    static_assert(sizeof(i32) == 4, "Invalid i32 sizeof");
    static_assert(sizeof(i64) == 8, "Invalid i64 sizeof");
} // namespace Pgp::Types

export namespace Pgp {
    constexpr auto kDefaultOpenMode = std::ios_base::in | std::ios_base::out;

    std::fstream OpenFileIfExist(fs::path const &FilePath, std::ios_base::openmode OpenMode = kDefaultOpenMode) {
        if (!fs::exists(FilePath)) {
            throw std::runtime_error{"Could not open file: " + FilePath.string()};
        }

        std::fstream OpenedFile{FilePath, OpenMode};
        if (!OpenedFile.is_open()) {
            throw std::runtime_error{"Could not open file: " + FilePath.string()};
        }

        return OpenedFile;
    }

    std::fstream OpenFileIfExistInBinary(fs::path const &FilePath) {
        constexpr auto OpenMode = kDefaultOpenMode | std::ios_base::binary;
        return OpenFileIfExist(FilePath, OpenMode);
    }


    template <typename DataType>
    concept ValidDataType = std::is_standard_layout_v<DataType>;

    template <ValidDataType DataType>
    std::optional<DataType> ReadBigEndianFromBinary(std::istream &Input) {
        constexpr auto DataTypeSize = sizeof(DataType);

        std::array<char, DataTypeSize> Data;
        if (!Input.read(Data.data(), DataTypeSize)) {
            return std::nullopt;
        }
        std::reverse(Data.begin(), Data.end());

        DataType Result{};
        std::memcpy(&Result, Data.data(), DataTypeSize);

        return Result;
    }
} // namespace Pgp