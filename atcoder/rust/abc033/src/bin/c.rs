#![allow(unused_imports)]

use std::cmp::*;
use std::collections::*;
use std::io::Write;
use std::ops::Bound::*;

use itertools::*;
use proconio::*;
use proconio::marker::*;
use superslice::*;

fn main() { 
    input! {
        s: String,
    }

    println!("{}", s.split("+").filter(|sub_str| !sub_str.contains("0")).count());
}
