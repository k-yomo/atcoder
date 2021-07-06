#![allow(unused_imports)]

use std::cmp::*;
use std::collections::*;
use std::io::Write;
use std::ops::Bound::*;

use itertools::__std_iter::once;
use itertools::*;
use itertools_num::ItertoolsNum;
use proconio::marker::*;
use proconio::*;
use superslice::*;

fn main() {
    input! {
        n: usize,
        ab: [(usize, usize); n],
    }

    let mut count = 0;
    for (a, b) in ab.iter().rev() {
        let m = (a + count) % b;
        if m != 0 {
            count += b - m
        }
    }
    println!("{}", count)
}
