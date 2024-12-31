use std::fs;

pub fn main() {
    println!("Part1: {}", part1());
    println!("Part2: {}", part2());
}

fn find_hash_with_prefix(input: &str, prefix: &str) -> i32 {
    let mut i = 0;
    let mut to_hash: String = input.to_string();

    loop {
        to_hash.truncate(input.len());
        to_hash.push_str(&i.to_string());

        let digest = md5::compute(&to_hash);

        if format!("{:x}", digest).starts_with(prefix) {
            return i;
        }
        i += 1
    }
}

pub fn part1() -> String {
    let input = fs::read_to_string("./2015/day04/input.txt").expect("Error reading input.txt");
    let input = input.trim();

    let result = find_hash_with_prefix(input, "00000");
    format!("{}", result)
}

pub fn part2() -> String {
    let input = fs::read_to_string("./2015/day04/input.txt").expect("Error reading input.txt");
    let input = input.trim();

    let result = find_hash_with_prefix(input, "000000");
    format!("{}", result)
}
