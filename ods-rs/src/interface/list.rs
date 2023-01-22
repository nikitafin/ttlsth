pub trait List<T: Copy> {
    type Idx;
    fn add(&self, i: Self::Idx, value: T);
}
