use std::{collections::HashMap, fs};

pub fn main() {
    println!("Part1: {}", part1());
    println!("Part2: {}", part2());
}

pub fn part1() -> String {
    let input = fs::read_to_string("./2015/day05/input.txt").expect("Error reading input.txt");

    let mut total = 0;

    'outer: for line in input.lines() {
        let chars: Vec<char> = line.chars().collect();
        let mut double = false;
        let mut vowels = 0;
        for (i, c) in chars.iter().enumerate() {
            if "aeiou".contains(*c) {
                vowels += 1;
            }
            if i == 0 {
                continue;
            }
            match (chars[i - 1], *c) {
                ('a', 'b') => continue 'outer,
                ('c', 'd') => continue 'outer,
                ('p', 'q') => continue 'outer,
                ('x', 'y') => continue 'outer,
                (a, _) if a == *c => double = true,
                _ => (),
            }
        }

        if !double {
            // println!("Line contains no double: {}", line);
            continue;
        }
        if vowels >= 3 {
            total += 1
        } else {
            // println!("Line contains not enough vowels: {}", line);
        }
    }

    return format!("{total}");
}

pub fn part2() -> String {
    let input = fs::read_to_string("./2015/day05/input.txt").expect("Error reading input.txt");

    let mut total = 0;

    for line in input.lines() {
        let mut seen: HashMap<String, usize> = HashMap::new();
        let chars: Vec<char> = line.chars().collect();
        let mut found_pair = false;
        let mut found_triplet = false;

        for (i, c) in chars.iter().enumerate() {
            if i == 0 {
                continue;
            }
            let pair = format!("{}{}", chars[i - 1], c);
            match seen.get(&pair) {
                Some(index) => {
                    if *index != i - 1 {
                        found_pair = true;
                    }
                }
                None => (),
            }
            seen.insert(pair, i);
            if i == 1 {
                continue;
            }
            match (chars[i - 2], chars[i - 1], *c) {
                (a, b, c) if (a == c) && (a != b) => found_triplet = true,
                _ => (),
            }
        }

        if found_pair && found_triplet {
            total += 1;
        }
    }

    return format!("{total}");
}
