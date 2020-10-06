#![allow(unused_imports)]
use proconio::*;
use proconio::marker::*;
use std::cmp::*;
use std::collections::*;
use std::io::Write;
use std::ops::Bound::*;
 
fn main() { 
    input! {
        n: usize,
        a: [usize; n],
    }

    let mut count = 1000000;
    for i in a {
        if i % 2 != 0 {
            println!("0");
            return
        }
        let factors = prime_factors(i);
        if factors[&2] < count {
            count = factors[&2];
        }
    }
    println!("{}", count)
}

fn prime_factors(mut n: usize) -> HashMap<usize, usize> {
    let mut primes = sieve_of_eratosthenes(f64::sqrt(n as f64) as usize);
    let mut hm_primes = HashMap::new();

    for prime in primes {
        while n % prime == 0 {
            n /= prime;
            if hm_primes.contains_key(&prime) {
                let mut x = hm_primes.get_mut(&prime).unwrap();
                *x += 1;
            } else {
                hm_primes.insert(prime, 1);
            }
        }
    }

    if n > 1 {
        if hm_primes.contains_key(&n) {
            let mut x = hm_primes.get_mut(&n).unwrap();
            *x += 1;
        } else {
            hm_primes.insert(n, 1);
        }
    }
    hm_primes
}

fn sieve_of_eratosthenes(n: usize) -> Vec<usize> {
    let mut spf = vec![None; n+1];
    let mut is_prime = vec![true; n+1];
    let mut primes = Vec::new();

    is_prime[0] = false;
    is_prime[1] = false;

    for i in 2..n+1 {
        if is_prime[i] {
            primes.push(i);
            spf[i] = Some(i);
        }

        for prime in &primes {
            if i * prime >= n + 1 || prime > &spf[i].unwrap() {
                break;
            }

            is_prime[i * prime] = false;

            spf[i * prime] = Some(*prime);
        }
    }
    primes
}