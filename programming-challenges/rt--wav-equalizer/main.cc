#include <cassert>
import Common;

static constexpr std::string_view kFileName = "data\\blaster-shot.wav";
static constexpr unsigned int kChunkNo = 1380533830;

int main() {
    try {
        std::fstream SoundFile = Pgp::OpenFileIfExistInBinary(kFileName);
        auto ChunkNo = Pgp::ReadBigEndianFromBinary<unsigned int>(SoundFile);
        assert(ChunkNo.has_value());
        assert(ChunkNo == kChunkNo);
    }
    catch (const std::exception &err) {
        std::cerr << err.what() << std::endl;
        return -1;
    }
    return 0;
}