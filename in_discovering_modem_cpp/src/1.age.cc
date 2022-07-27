#include <fstream>
#include <iostream>

int
main ()
{
  auto success{ false };

  int age;
  while (!success)
    {
      std::cout << "Enter your age: ";
      std::cin >> age;
      if (!std::cin.good ())
        {
          std::cout << "Invalid age. Retry." << std::endl;
          std::cin.clear ();
          // Add ignore line
          continue;
        }

      success = true;
    }

  std ::ofstream os ("output.txt");
  if (!os.good ())
    {
      std::cout << "Can't open output.txt. Permission Denied. Exit.";
      return EXIT_FAILURE;
    }
  os << "age: " << age << '\n';
  if (!os.good ())
    {
      std::cout << "Can't white age in output.txt. Exit";
      return EXIT_FAILURE;
    }
  os.close ();

  return EXIT_SUCCESS;
}