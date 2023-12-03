use std::collections::HashMap;
use std::fs;

#[derive(Debug)]
struct Rucksack {
    first_part: String,
    second_part: String,
}

#[derive(Debug)]
struct Group {
    first_elf: String,
    second_elf: String,
    third_elf: String,
}

fn priorities() -> HashMap<String, i32> {
    let characters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ";
    let mut priorities: HashMap<String, i32> = HashMap::new();

    characters.chars().enumerate().for_each(|(i, c)| {
        priorities.insert(c.to_string(), (i + 1) as i32);
    });

    priorities
}

fn get_groups(lines: Vec<&str>) -> Vec<Group> {
    let mut groups: Vec<Group> = Vec::new();
    let mut group: Vec<String> = Vec::new();

    lines.iter().enumerate().for_each(|(i, l)| {
        group.push(l.to_string());

        if (i + 1) % 3 == 0 {
            groups.push(Group {
                first_elf: group[0].clone(),
                second_elf: group[1].clone(),
                third_elf: group[2].clone(),
            });
            group = Vec::new();
        }
    });

    groups
}

fn matching_characters(a: &String, b: &String, c: &String) -> String {
    let mut matching: String = String::from("");

    a.chars().for_each(|character| match character {
        _ if b.contains(character) && c.contains(character) && !matching.contains(character) => {
            matching.push_str(&character.to_string())
        }
        _ => (),
    });

    matching
}

fn main() {
    let contents = fs::read_to_string("./input/rucksacks").expect("Could not read file!");
    let lines: Vec<&str> = contents.split("\n").collect();

    let groups = get_groups(lines);
    let priorities = priorities();

    let total_prio: i32 = groups
        .iter()
        .map(|g| matching_characters(&g.first_elf, &g.second_elf, &g.third_elf))
        .fold(0, |acc, curr| {
            let mut total = acc;

            curr.chars().for_each(|character| {
                let priority = match priorities.get(&character.to_string()) {
                    Some(&kek) => kek,
                    _ => panic!("No priority for {}", &character.to_string()),
                };

                total = total + priority
            });

            total
        });

    println!("{}", total_prio);
}
