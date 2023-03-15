//---------------------------------------------------------------------------
#ifndef uDayTimeH
#define uDayTimeH
//---------------------------------------------------------------------------
#include <vector>
//---------------------------------------------------------------------------
namespace lib
{
    class DayTime
    {
      public:
        static const long FACTOR;
      private:
        struct SomeStruct
        {
            int a;
            int b;
        };
        std::vector<SomeStruct> FVect;
    };
} // namespace lib
//---------------------------------------------------------------------------
#endif

