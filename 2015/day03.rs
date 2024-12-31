use num::complex::Complex;
use std::collections::HashSet;
use std::fs;

pub fn main() {
    println!("Part1: {}", part1());
    println!("Part2: {}", part2());
}

pub fn part1() -> String {
    let input = fs::read_to_string("./2015/day03/input.txt").expect("Error reading input.txt");

    let mut position = Complex::new(0, 0);
    let mut visited: HashSet<Complex<i32>> = HashSet::new();

    visited.insert(position);

    for c in input.chars() {
        match c {
            '^' => position += Complex::new(0, 1),
            'v' => position += Complex::new(0, -1),
            '>' => position += Complex::new(1, 0),
            '<' => position += Complex::new(-1, 0),
            _ => panic!("Invalid input"),
        }
        visited.insert(position);
    }

    return format!("{}", visited.len());
}

pub fn part2() -> String {
    let input = fs::read_to_string("./2015/day03/input.txt").expect("Error reading input.txt");
    let chars: Vec<char> = input.chars().collect();

    let mut santa_position = (0, 0);
    let mut robot_position = (0, 0);
    let mut visited: HashSet<(i32, i32)> = HashSet::new();

    visited.insert(santa_position);

    for (i, c) in chars.iter().enumerate() {
        let (x, y) = if i % 2 == 0 {
            &mut santa_position
        } else {
            &mut robot_position
        };

        match c {
            '^' => *y += 1,
            'v' => *y -= 1,
            '>' => *x += 1,
            '<' => *x -= 1,
            _ => panic!("Invalid input"),
        }

        visited.insert((*x, *y));
    }

    return format!("{}", visited.len());
}
