use std::fs;

pub fn main() {
    println!("Part1: {}", part1());
    println!("Part2: {}", part2());
}

pub fn part1() -> String {
    let input = fs::read_to_string("./2015/day01/input.txt").expect("Error reading input.txt");

    let mut floor = 0;

    for c in input.chars() {
        match c {
            '(' => floor += 1,
            ')' => floor -= 1,
            _ => (),
        }
    }
    return format!("{}", floor);
}

pub fn part2() -> String {
    let input = fs::read_to_string("./2015/day01/input.txt").expect("Error reading input.txt");

    let mut floor = 0;
    let mut pos = 0;

    for c in input.chars() {
        match c {
            '(' => floor += 1,
            ')' => floor -= 1,
            _ => (),
        }
        pos += 1;
        if floor < 0 {
            return format!("{}", pos);
        }
    }

    return format!("");
}
