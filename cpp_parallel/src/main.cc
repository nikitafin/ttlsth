#include <chrono>
#include <iostream>
#include <thread>

using namespace std::chrono_literals;

int
main()
{
  std::cout << "123" << std::endl;
  std::this_thread::sleep_for(10'000ms);
  std::cout << "123" << std::endl;
}
