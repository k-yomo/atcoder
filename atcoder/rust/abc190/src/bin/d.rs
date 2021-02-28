#![allow(unused_imports)]

use std::cmp::*;
use std::collections::*;
use std::io::Write;
use std::ops::Bound::*;

use itertools::*;
use itertools::__std_iter::once;
use itertools_num::ItertoolsNum;
use proconio::*;
use proconio::marker::*;
use superslice::*;

fn main() {
    input! {
        n: i128,
    }

    let mut div_n = n;
    while div_n % 2 == 0 {
        div_n /= 2;
    }

    println!("{}", (1..).take_while(|i| i * i <= div_n).filter(|i| div_n % i == 0).count() * 4);
}

