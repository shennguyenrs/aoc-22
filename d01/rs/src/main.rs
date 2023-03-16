use std::fs;

// fn first_part_a(content: &String) -> i32 {
//     let parts = content.split("\n\n");
//     let mut results: Vec<i32> = vec![];
//     let mut max: i32 = 0;
//
//     for part in parts {
//         let mut sum: i32 = 0;
//         let lines = part.split("\n");
//
//         for line in lines {
//             if line != "" {
//                 let num: i32 = line.parse().expect("should able to parse string to i32");
//                 sum += num
//             }
//         }
//
//         results.push(sum)
//     }
//
//     for i in results {
//         if i > max {
//             max = i
//         }
//     }
//
//     return max;
// }

fn first_part_b(content: &str) -> i32 {
    content
        .split("\n\n")
        .map(|part| {
            part.split("\n")
                .filter(|line| !line.is_empty())
                .map(|n| {
                    n.parse::<i32>()
                        .expect("should able to parse string to i32")
                })
                .sum()
        })
        .max()
        .unwrap_or_default()
}

// fn second_part_a(content: &str) -> i32 {
//     let mut results = vec![];
//     let mut top_three: [i32; 3] = [0, 0, 0];
//
//     for part in content.split("\n\n") {
//         let value: i32 = part
//             .lines()
//             .filter(|line| !line.is_empty())
//             .map(|n| n.parse::<i32>().unwrap())
//             .sum::<i32>();
//
//         results.push(value);
//     }
//
//     results.sort_by(|a, b| b.cmp(a));
//
//     for i in 0..3 {
//         top_three[i] = results[i]
//     }
//
//     return top_three.iter().sum();
// }

fn second_part_b(content: &str) -> i32 {
    let mut top_three: [i32; 3] = [0, 0, 0];

    for part in content.split("\n\n") {
        let value: i32 = part
            .lines()
            .filter(|line| !line.is_empty())
            .map(|n| n.parse::<i32>().unwrap())
            .sum::<i32>();

        for i in 0..3 {
            if value > top_three[i] {
                top_three.swap(i, 2);
                top_three.swap(1, 2);
                top_three.swap(0, 1);
                top_three[i] = value;
                break;
            }
        }
    }

    return top_three.iter().sum();
}

fn main() {
    let content = fs::read_to_string("../input.txt").expect("should able to read the file");
    let n = first_part_b(&content);
    let m = second_part_b(&content);
    println!("The maxium calories is {}", n);
    println!("The sum of top three calories is {}", m);
}
