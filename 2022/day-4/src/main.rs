use std::fs;
use std::ops::Range;

#[derive(Debug)]
struct AssignmentPair {
    first: Vec<i32>,
    second: Vec<i32>,
}

impl AssignmentPair {
    fn from_line(line: &str) -> AssignmentPair {
        let pairs: Vec<&str> = line.split(",").collect();
        let nums: Vec<Vec<i32>> = pairs.iter().map(|p| pair_to_nums(p)).collect();

        AssignmentPair {
            first: nums[0].clone(),
            second: nums[1].clone(),
        }
    }
}

fn pair_to_nums(pair: &str) -> Vec<i32> {
    let start_and_end: Vec<i32> = pair.split("-").map(|s| s.parse::<i32>().unwrap()).collect();

    let start = start_and_end[0];
    let end = start_and_end[1] + 1;

    (start..end).map(|n| n).collect()
}

fn get_assignment_pairs() -> Vec<AssignmentPair> {
    let contents = fs::read_to_string("./input/pairs").expect("Could not read file!");
    let lines: Vec<&str> = contents.split("\n").collect();

    lines
        .iter()
        .filter(|l| !l.is_empty())
        .map(|l| AssignmentPair::from_line(l))
        .collect()
}

fn get_pairs_with_overlap(pairs: Vec<AssignmentPair>) -> Vec<AssignmentPair> {
    pairs
        .into_iter()
        .filter(|p| vecs_overlap(&p.first, &p.second) || vecs_overlap(&p.second, &p.first))
        .collect()
}

fn vecs_overlap(a: &Vec<i32>, b: &Vec<i32>) -> bool {
    b.iter().any(|item| a.contains(item))
}

fn main() {
    let assignment_pairs = get_assignment_pairs();
    let overlapping_pairs = get_pairs_with_overlap(assignment_pairs);

    let total_overlapping = overlapping_pairs.len();

    println!("Amount of overlap: {}", total_overlapping);
}
