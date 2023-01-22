use ods_rs::implementation::arraystack::ArrayStuck;

trait A: Sized {
    fn a1(&self);
}

trait Shape {
    fn perimeter(&self) -> f32;
    fn area(&self) -> f32;
}

struct Rectangle {
    // pub width: f32,
    // pub height: f32,
}

struct Triangle {
    // pub side: f32,
}

struct Circle {
    // pub radius: f32,
}

impl Shape for Rectangle {
    fn perimeter(&self) -> f32 {
        // self.width * 2.0 + self.height * 2.0
        1.0
    }
    fn area(&self) -> f32 {
        // self.width * self.height
        1.0
    }
}

impl Shape for Triangle {
    fn perimeter(&self) -> f32 {
        // self.side * 3.0
        1.0
    }
    fn area(&self) -> f32 {
        // self.side * 0.5 * 3.0_f32.sqrt() / 2.0 * self.side
        1.0
    }
}

impl Shape for Circle {
    fn perimeter(&self) -> f32 {
        // self.radius * 2.0 * std::f32::consts::PI
        1.0
    }
    fn area(&self) -> f32 {
        // self.radius * self.radius * std::f32::consts::PI
        1.0
    }
}

fn main() {
    let st = ArrayStuck::<u8>::new(10);
    assert_eq!(st.data.capacity(), 10);
    let st1: u8 = 10;
    println!("usage {st1}");

    let st2 = Vec::<&dyn Shape>::new();
}