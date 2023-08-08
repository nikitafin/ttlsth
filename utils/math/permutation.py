import sys
from math import factorial


def permutation(n: int) -> int:
    return factorial(n)


def main():
    arg_count = len(sys.argv)
    if arg_count < 2:
        if arg_count != 0:
            arg_count -= 1
        print(f'invlid arg count({arg_count}), expect 1')
        return

    n = int(sys.argv[1])
    print(f'permutation: {permutation(n)}')


if __name__ == '__main__':
    main()
