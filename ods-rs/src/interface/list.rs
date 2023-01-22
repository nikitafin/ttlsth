pub trait List<T: Copy>: Clone {
    type Idx;
    fn add(&self, i: Self::Idx, value: T);
}
