use std::fs;

pub fn main() {
    println!("Part1: {}", part1());
    println!("Part2: {}", part2());
}

pub fn part1() -> String {
    let input = fs::read_to_string("./2015/day02/input.txt").expect("Error reading input.txt");

    let mut total = 0;

    for line in input.lines() {
        let mut dimensions: Vec<i32> = line.split('x').map(|x| x.parse().unwrap()).collect();
        dimensions.sort();
        let l = dimensions[0];
        let w = dimensions[1];
        let h = dimensions[2];
        total += 2 * l * w + 2 * w * h + 2 * h * l;

        let sides = vec![l * w, w * h, h * l];
        let min_value = sides.iter().min();
        match min_value {
            Some(min) => total += min,
            None => (),
        }
    }

    return format!("{total}");
}

pub fn part2() -> String {
    let input = fs::read_to_string("./2015/day02/input.txt").expect("Error reading input.txt");

    let mut total = 0;

    for line in input.lines() {
        let mut dimensions: Vec<i32> = line.split('x').map(|x| x.parse().unwrap()).collect();
        dimensions.sort();

        let l = dimensions[0];
        let w = dimensions[1];
        let h = dimensions[2];
        // At this point the two smallest are l and w

        total += l * 2 + w * 2 + l * w * h
    }
    return format!("{total}");
}
