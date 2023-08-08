#include <list>
#include <tchar.h>
#include <vcl.h>

class TEmploy {
public:
  int id;
  std::string name;

  struct TPhone {
    short area;
    short exchange;
    int station;
  };
  std::list<TPhone> phones;

  static int staticField;
};

int TEmploy::staticField = 0;

int _tmain(int argc, _TCHAR *argv[]) {
  TEmploy Employ{};

  return 0;
}

