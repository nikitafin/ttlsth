use std::fmt::Display;

use crate::interface::list::List;

pub struct ArrayStuck<T: Copy + Display> {
    pub data: Vec<T>,
}

impl<T: Copy + Display> ArrayStuck<T> {
    // TODO(Nikita): if no memory ?
    pub fn new(sz: usize) -> Self {
        Self {
            data: Vec::with_capacity(sz),
        }
    }
}



impl<T: Copy + Display> List<T> for ArrayStuck<T> {
    type Idx = isize;

    fn add(i: Self::Idx, value: T) {
        println!("{i}, {value}");
    }
}