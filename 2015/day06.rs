use std::{
    collections::{HashMap, HashSet},
    fs,
};

pub fn main() {
    println!("Part1: {}", part1());
    println!("Part2: {}", part2());
}

#[derive(Eq, Hash, PartialEq)]
struct Light(i32, i32);

pub fn part1() -> String {
    let input = fs::read_to_string("./2015/day06/input.txt").expect("Error reading input.txt");

    let mut grid: HashSet<Light> = HashSet::new();

    for line in input.lines() {
        let instructions: Vec<_> = line.split_ascii_whitespace().collect();

        if instructions.len() == 4 {
            let corner1_vect: Vec<_> = instructions[1].split(",").collect();
            let corner2_vect: Vec<_> = instructions[3].split(",").collect();

            let corner1 = Light(
                corner1_vect[0].parse().expect("invalid input"),
                corner1_vect[1].parse().expect("invalid input"),
            );
            let corner2 = Light(
                corner2_vect[0].parse().expect("invalid input"),
                corner2_vect[1].parse().expect("invalid input"),
            );

            grid = toggle(grid, corner1, corner2)
        } else {
            if instructions[1] == "on" {
                let corner1_vect: Vec<_> = instructions[2].split(",").collect();
                let corner2_vect: Vec<_> = instructions[4].split(",").collect();

                let corner1 = Light(
                    corner1_vect[0].parse().expect("invalid input"),
                    corner1_vect[1].parse().expect("invalid input"),
                );
                let corner2 = Light(
                    corner2_vect[0].parse().expect("invalid input"),
                    corner2_vect[1].parse().expect("invalid input"),
                );

                grid = on(grid, corner1, corner2)
            } else {
                let corner1_vect: Vec<_> = instructions[2].split(",").collect();
                let corner2_vect: Vec<_> = instructions[4].split(",").collect();

                let corner1 = Light(
                    corner1_vect[0].parse().expect("invalid input"),
                    corner1_vect[1].parse().expect("invalid input"),
                );
                let corner2 = Light(
                    corner2_vect[0].parse().expect("invalid input"),
                    corner2_vect[1].parse().expect("invalid input"),
                );

                grid = off(grid, corner1, corner2)
            }
        }
    }

    return format!("{}", grid.len());
}

fn toggle(mut grid: HashSet<Light>, corner1: Light, corner2: Light) -> HashSet<Light> {
    for x in corner1.0..corner2.0 + 1 {
        for y in corner1.1..corner2.1 + 1 {
            let light = Light(x, y);
            if grid.contains(&light) {
                grid.remove(&light);
            } else {
                grid.insert(light);
            }
        }
    }
    return grid;
}

fn on(mut grid: HashSet<Light>, corner1: Light, corner2: Light) -> HashSet<Light> {
    for x in corner1.0..corner2.0 + 1 {
        for y in corner1.1..corner2.1 + 1 {
            let light = Light(x, y);
            grid.insert(light);
        }
    }
    return grid;
}

fn off(mut grid: HashSet<Light>, corner1: Light, corner2: Light) -> HashSet<Light> {
    for x in corner1.0..corner2.0 + 1 {
        for y in corner1.1..corner2.1 + 1 {
            let light = Light(x, y);
            grid.remove(&light);
        }
    }
    return grid;
}

pub fn part2() -> String {
    let input = fs::read_to_string("./2015/day06/input.txt").expect("Error reading input.txt");

    let mut grid: HashMap<Light, i32> = HashMap::new();

    for line in input.lines() {
        let instructions: Vec<_> = line.split_ascii_whitespace().collect();

        if instructions.len() == 4 {
            let corner1_vect: Vec<_> = instructions[1].split(",").collect();
            let corner2_vect: Vec<_> = instructions[3].split(",").collect();

            let corner1 = Light(
                corner1_vect[0].parse().expect("invalid input"),
                corner1_vect[1].parse().expect("invalid input"),
            );
            let corner2 = Light(
                corner2_vect[0].parse().expect("invalid input"),
                corner2_vect[1].parse().expect("invalid input"),
            );

            grid = toggle2(grid, corner1, corner2)
        } else {
            if instructions[1] == "on" {
                let corner1_vect: Vec<_> = instructions[2].split(",").collect();
                let corner2_vect: Vec<_> = instructions[4].split(",").collect();

                let corner1 = Light(
                    corner1_vect[0].parse().expect("invalid input"),
                    corner1_vect[1].parse().expect("invalid input"),
                );
                let corner2 = Light(
                    corner2_vect[0].parse().expect("invalid input"),
                    corner2_vect[1].parse().expect("invalid input"),
                );

                grid = on2(grid, corner1, corner2)
            } else {
                let corner1_vect: Vec<_> = instructions[2].split(",").collect();
                let corner2_vect: Vec<_> = instructions[4].split(",").collect();

                let corner1 = Light(
                    corner1_vect[0].parse().expect("invalid input"),
                    corner1_vect[1].parse().expect("invalid input"),
                );
                let corner2 = Light(
                    corner2_vect[0].parse().expect("invalid input"),
                    corner2_vect[1].parse().expect("invalid input"),
                );

                grid = off2(grid, corner1, corner2)
            }
        }
    }

    let mut total = 0;

    for (_, brightness) in grid.into_iter() {
        total += brightness
    }

    return format!("{}", total);
}

fn toggle2(mut grid: HashMap<Light, i32>, corner1: Light, corner2: Light) -> HashMap<Light, i32> {
    for x in corner1.0..corner2.0 + 1 {
        for y in corner1.1..corner2.1 + 1 {
            let light = Light(x, y);
            *grid.entry(light).or_insert(0) += 2;
        }
    }
    return grid;
}

fn on2(mut grid: HashMap<Light, i32>, corner1: Light, corner2: Light) -> HashMap<Light, i32> {
    for x in corner1.0..corner2.0 + 1 {
        for y in corner1.1..corner2.1 + 1 {
            let light = Light(x, y);
            *grid.entry(light).or_insert(0) += 1;
        }
    }
    return grid;
}

fn off2(mut grid: HashMap<Light, i32>, corner1: Light, corner2: Light) -> HashMap<Light, i32> {
    for x in corner1.0..corner2.0 + 1 {
        for y in corner1.1..corner2.1 + 1 {
            let light = Light(x, y);
            let brightness = grid.entry(light).or_insert(0);
            if *brightness > 0 {
                *brightness -= 1;
            }
        }
    }
    return grid;
}
