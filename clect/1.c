#if 0
#include <stdio.h>

inline void
my_assert(int b, const char *str) {
    if (b)
        return;
    fprintf(stderr, "Assertion failed: %s\n", str);
    abort();
}

int
main(int argc, char **argv) {
    my_assert(argc > 0, "argc <= 0");
    return 0;
}
#elif 1
#include <stdio.h>

int
foo(int x) {
    int a = (x += 0) + 3;
    return a;
}
#endif