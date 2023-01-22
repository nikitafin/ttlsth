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
    fn array_stack_new_test1() {
        let v1 = ArrayStack::<u8>::new(10);
        let test_v = Vec::<u8>::with_capacity(10);
        assert_eq!(v1.data, test_v);
        assert_eq!(v1.data.len(), test_v.len());
    }

    #[test]
    fn array_stack_new_test2() {
        let v1 = ArrayStack::<u8>::new(11);
        let test_v = Vec::<u8>::with_capacity(10);
        assert_eq!(v1.data, test_v);
        assert_eq!(v1.data.len(), test_v.len());
        assert_ne!(v1.data.capacity(), test_v.capacity());
    }

    #[test]
    fn array_stack_clone_test1() {
        let mut v1 = ArrayStack::<u8>::new(10);
        v1.data.push(1);
        v1.data.push(2);

        let v2 = v1.clone();
        assert_eq!(v1.data, v2.data);
        assert_eq!(v1.data.len(), 2);
        assert!(v1.data.capacity() >= 10);
    }
}
