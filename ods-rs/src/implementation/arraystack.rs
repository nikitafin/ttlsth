use std::fmt::Display;

use crate::interface::list::List;

pub struct ArrayStack<T: Copy + Display> {
    data: Vec<T>,
}

impl<T: Copy + Display> ArrayStack<T> {
    // TODO(Nikita): if no memory ?
    pub fn new(sz: usize) -> Self {
        Self {
            data: Vec::with_capacity(sz),
        }
    }
}

impl<T: Copy + Display> Clone for ArrayStack<T> {
    fn clone(&self) -> Self {
        Self {
            data: self.data.clone(),
        }
    }
}

impl<T: Copy + Display> List<T> for ArrayStack<T> {
    type Idx = isize;

    fn add(&self, i: Self::Idx, value: T) {
        println!("{i}, {value}");
    }
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn arraystack_clone_test() {
        let v1 = ArrayStack::<u8>::new(10);
        let test_v = Vec::<u8>::with_capacity(10);
        assert_eq!(v1.data, test_v);
    }
}
