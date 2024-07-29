use std::io::{Read, stdin};

/// Determine the length of the collatz sequence beginning at `n`.
fn collatz_length(mut n: i32) -> u32 {
    if n < 1 {
        panic!("n must be >= 1");
    }

    let mut len = 0;
    loop {
        if n == 1 {
            return len;
        }

        if n % 2 == 0 {
            n /= 2;
        } else {
            n = 3 * n + 1;
        }

        len += 1;
    }
}

fn main() {
    let mut input = String::new();
    stdin().read_to_string(&mut input).expect("Invalid input");
    let n = input.trim().parse::<i32>().expect("Invalid input");

    let len = collatz_length(n);
    println!("The length of the collatz sequence beginning at {} is {}.", n, len)
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    #[should_panic]
    fn invalid1() {
        assert_eq!(collatz_length(0), 0);
    }

    #[test]
    #[should_panic]
    fn invalid2() {
        assert_eq!(collatz_length(-1), 0);
    }

    #[test]
    fn test_collatz_length() {
        assert_eq!(collatz_length(1), 0);
        assert_eq!(collatz_length(2), 1);
        assert_eq!(collatz_length(3), 7);

        assert_eq!(collatz_length(1_000_000), 152);
        assert_eq!(collatz_length(100), 25);
        assert_eq!(collatz_length(20), 7);
        assert_eq!(collatz_length(21), 7);
    }
}