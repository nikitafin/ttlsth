pub trait List<T: Copy> {
    type Idx;
    fn add(i: Self::Idx, value: T);
}