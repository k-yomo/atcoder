use std::io::*;
use std::str::FromStr;
use std::cmp::max;

fn read<T: FromStr>() -> T {
    let stdin = stdin();
    let stdin = stdin.lock();
    let token: String = stdin
        .bytes()
        .map(|c| c.expect("failed to read char") as char)
        .skip_while(|c| c.is_whitespace())
        .take_while(|c| !c.is_whitespace())
        .collect();
    token.parse().ok().expect("failed to parse token")
}

fn main() {
    let a: i64 = read();
    let b: i64 = read();
    println!("{}", max(max(a+b, a-b), a*b));
}
