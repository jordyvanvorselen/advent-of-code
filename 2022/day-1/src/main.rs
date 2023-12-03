use std::fs;

struct Elf {
    calories: Vec<i32>,
}

impl Elf {
    fn total_calories(&self) -> i32 {
        self.calories.iter().sum()
    }
}

fn main() {
    let contents = fs::read_to_string("./input/calories").expect("Could not read file!");
    let elves = get_elves(contents);

    let mut calories: Vec<i32> = elves.iter().map(|elf| elf.total_calories()).collect();
    calories.sort();

    let index = calories.len() - 4;
    let mut total = 0;

    for (i, c) in calories.iter().enumerate() {
        if i > calories.len() - 4 {
            total = total + c;
        }
    }

    println!("{}", total);
}

fn get_elves(contents: String) -> Vec<Elf> {
    let mut elves: Vec<Elf> = Vec::new();
    let mut calories: Vec<i32> = Vec::new();

    let lines: Vec<&str> = contents.split("\n").collect();

    for line in lines {
        if line.is_empty() {
            elves.push(Elf { calories });
            calories = Vec::new();
            continue;
        }

        let line = line.parse::<i32>().unwrap();
        calories.push(line)
    }

    return elves;
}
