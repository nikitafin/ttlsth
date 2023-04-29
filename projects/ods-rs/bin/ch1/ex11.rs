use std::io;
// use std::collections::VecDeque;

fn main() -> io::Result<()> {
    println!("ex1.1");

    for res in io::stdin().lines() {
        match res {
            Ok(line) => {
                println!(
                    "Read \"{}\", with size {}",
                    line.trim_end(),
                    line.chars().count()
                );
            }
            Err(_) => {
                println!("err")
            }
        }
    }

    Ok(())
}

// other way
// while let Ok(buffer_size) = stdin.read_line(&mut buffer) {
//     if buffer_size == 0 {
//         break;
//     }
//     let pretty_str = buffer.trim_end();
//     println!(
//         "Read \"{}\", with size {} bytes",
//         pretty_str,
//         pretty_str.chars().count()
//     );
//     buffer.clear();
// }
