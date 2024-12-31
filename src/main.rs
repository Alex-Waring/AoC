use std::iter::empty;
use std::time::{Duration, Instant};

use aoc::year2015;
use clap::Parser;

#[derive(Parser)]
#[command(version, about, long_about = None)]
struct Args {
    #[clap(default_value = "All")]
    year: String,
    #[clap(default_value = "All")]
    day: String,
}

fn main() {
    let args = Args::parse();

    let solutions: Vec<_> = empty()
        .chain(year2015())
        .filter(|s| args.year == "All" || s.year == args.year)
        .filter(|s| args.day == "All" || s.day == args.day)
        .collect();

    let mut duration = Duration::ZERO;

    for Solution { year, day, wrapper } in &solutions {
        let instant = Instant::now();
        let (part1, part2) = wrapper();
        let time_taken = instant.elapsed();
        duration += time_taken;

        println!("{BOLD}{YELLOW}{year} {day:02}{RESET}");
        println!("    Part 1: {part1}");
        println!("    Part 2: {part2}");
        println!("    Time: {:?}μs", time_taken.as_micros());
    }

    println!("{BOLD}{YELLOW}⭐ {}{RESET}", 2 * solutions.len());
    println!("{BOLD}{WHITE}🕓 {} ms{RESET}", duration.as_millis());
}

struct Solution {
    year: String,
    day: String,
    wrapper: fn() -> (String, String),
}

macro_rules! run {
    ($year:tt $($day:tt),*)  => {
        fn $year() -> Vec<Solution> {
            vec![$({
                let year = stringify!($year);
                let day = stringify!($day);

                let wrapper = || {
                    use $year::$day::*;

                    let part1 = part1();
                    let part2 = part2();

                    (part1.to_string(), part2.to_string())
                };

                Solution {
                    year: year.parse().unwrap(),
                    day: day.parse().unwrap(),
                    wrapper,
                }
            },)*]
        }
    };
}

run!(year2015
    day01, day02, day03, day04, day05, day06
);

pub const RESET: &str = "\x1b[0m";
pub const BOLD: &str = "\x1b[1m";
pub const RED: &str = "\x1b[31m";
pub const GREEN: &str = "\x1b[32m";
pub const YELLOW: &str = "\x1b[33m";
pub const BLUE: &str = "\x1b[94m";
pub const WHITE: &str = "\x1b[97m";
pub const HOME: &str = "\x1b[H";
pub const CLEAR: &str = "\x1b[J";
