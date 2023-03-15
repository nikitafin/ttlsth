export module Common;

export int Fn() { return 10; }

// namespace fs = std::filesystem;

// std::ifstream openFileIfExists(const fs::path &filePath)
// {
//     if (fs::exists(filePath))
//     {
//         std::ifstream file(filePath);
//         return file;
//     }
//     else
//     {
//         throw std::runtime_error("Could not open file: " + filePath.string());
//     }
// }
