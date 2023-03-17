use std::collections::HashMap;
use std::fs;

#[derive(Clone)]
struct Rule {
    judge: [i32; 3],
    chosen: i32,
}

fn first_part(
    content: &str,
    rules: &HashMap<char, Rule>,
    opponent_pos: &HashMap<char, usize>,
) -> i32 {
    let mut result: i32 = 0;

    for line in content.lines() {
        let items: Vec<&str> = line.split_whitespace().collect();
        let opp: &char = &items[0].chars().next().unwrap();
        let me: &char = &items[1].chars().next().unwrap();
        let opp_pos = *opponent_pos.get(opp).unwrap_or(&3);
        let me_pos = rules.get(me).unwrap();
        result += me_pos.chosen + me_pos.judge[opp_pos];
    }

    result
}

fn second_part(
    content: &str,
    rules: &HashMap<char, Rule>,
    second_strategy_score: &HashMap<char, i32>,
    second_strategy_guide: &HashMap<char, HashMap<char, char>>,
) -> i32 {
    let mut result: i32 = 0;

    for line in content.lines() {
        let items: Vec<&str> = line.split_whitespace().collect();
        let opp: &char = &items[0].chars().next().unwrap();
        let me: &char = &items[1].chars().next().unwrap();
        let me_pos = second_strategy_score.get(me).unwrap();
        let opp_guide = second_strategy_guide.get(opp).unwrap();
        let opp_pos = opp_guide.get(me).unwrap();
        result += me_pos + rules.get(opp_pos).unwrap().chosen;
    }

    result
}

fn main() {
    let rules: HashMap<char, Rule> = [
        (
            'X',
            Rule {
                judge: [3, 0, 6],
                chosen: 1,
            },
        ),
        (
            'Y',
            Rule {
                judge: [6, 3, 0],
                chosen: 2,
            },
        ),
        (
            'Z',
            Rule {
                judge: [0, 6, 3],
                chosen: 3,
            },
        ),
    ]
    .iter()
    .cloned()
    .collect();

    let opponent_pos: HashMap<char, usize> =
        [('A', 0), ('B', 1), ('C', 2)].iter().cloned().collect();

    let second_strategy_score: HashMap<char, i32> =
        [('X', 0), ('Y', 3), ('Z', 6)].iter().cloned().collect();

    let second_strategy_guide: HashMap<char, HashMap<char, char>> = [
        (
            'A',
            [('X', 'Z'), ('Y', 'X'), ('Z', 'Y')]
                .iter()
                .cloned()
                .collect(),
        ),
        (
            'B',
            [('X', 'X'), ('Y', 'Y'), ('Z', 'Z')]
                .iter()
                .cloned()
                .collect(),
        ),
        (
            'C',
            [('X', 'Y'), ('Y', 'Z'), ('Z', 'X')]
                .iter()
                .cloned()
                .collect(),
        ),
    ]
    .iter()
    .cloned()
    .collect();

    let content = fs::read_to_string("../input.txt").expect("should able to read the file");
    let n = first_part(&content, &rules, &opponent_pos);
    let m = second_part(
        &content,
        &rules,
        &second_strategy_score,
        &second_strategy_guide,
    );
    println!("The result of first part is {}", n);
    println!("The result of second part is {}", m);
}
