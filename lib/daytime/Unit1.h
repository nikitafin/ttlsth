//---------------------------------------------------------------------------
#ifndef Unit1H
#define Unit1H
//---------------------------------------------------------------------------
#include <vector>
#include <array>
#include <list>
//---------------------------------------------------------------------------
class MyStruct_s
{
  public:
    //    static int dob;
    char name[32];
    int id;
    struct Phone_s
    {
        short area;
        short exchange;
        int station;
    };
    std::list<Phone_s> phones;
    //    std::array<Phone_s, 3> phones;
    //    std::vector<Phone_s> phones;
};
//---------------------------------------------------------------------------
#endif

