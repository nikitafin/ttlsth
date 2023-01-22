fn main() {
    let result = Some("string");
    match result {
        Some(s) => println!("String value is: {s}"),
        None => todo!(),
    }
}
