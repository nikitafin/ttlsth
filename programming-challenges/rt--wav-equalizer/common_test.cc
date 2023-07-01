#include <gtest/gtest.h>

import Common;

namespace {
    using namespace Pgp::Types;
    constexpr std::string_view kFileName = "data\\blaster-shot.wav";
    constexpr u32 kChunkNo               = 1380533830;
} // namespace

namespace Tests {
    class TValidWavFileFixture : public ::testing::Test {
      protected:
        std::fstream m_SoundFile;

      protected:
        void SetUp() override {
            m_SoundFile = Pgp::OpenFileIfExistInBinary(kFileName);
        }
    };

    TEST_F(TValidWavFileFixture, ParseChunkNo) {
        auto ChunkNoOpt = Pgp::ReadBigEndianFromBinary<u32>(m_SoundFile);
        EXPECT_TRUE(ChunkNoOpt.has_value());
        EXPECT_EQ(ChunkNoOpt, kChunkNo);
    }
} // namespace Tests