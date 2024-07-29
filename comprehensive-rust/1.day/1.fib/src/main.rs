fn fib(n: u32) -> u32 {
    if n < 2 {
        return n;
    }

    let mut a = 0;
    let mut b = 1;
    for _i in 2..n {
        let c = a + b;
        a = b;
        b = c;
    }

    b
}

fn main() {
    let n = 20;
    println!("fib({n}) = {}", fib(n));
}


#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn zero() {
        assert_eq!(fib(0), 0);
    }

    #[test]
    fn one() {
        assert_eq!(fib(1), 1);
    }

    #[test]
    fn two() {
        assert_eq!(fib(2), 1);
    }

    #[test]
    #[should_panic]
    fn n_1337() {
        assert_eq!(fib(1337), 23564956);
    }

    #[test]
    #[should_panic]
    fn max() {
        assert_eq!(fib(u32::MAX), 0);
    }
}