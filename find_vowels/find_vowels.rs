use regex::Regex;

fn main() {
    let test = "This is test string";
    println!("Count by normal fn {}", count_vowels(test));
    println!("Count by regex fn {}", count_by_regex(test));
}

fn count_by_regex(str: &str) -> usize {
    let regex = Regex::new(r"(?i)[aeiou]").unwrap();
    let matches = regex.find_iter(str).count();
    matches
}

fn count_vowels(str: &str) -> usize {
    let vowels = ['a', 'e', 'i', 'o', 'u'];
    let mut count = 0;

    for c in str.chars() {
        let lower = c.to_ascii_lowercase();
        if vowels.contains(&lower) {
            count += 1;
        }
    }

    count
}

// Rust's `Vec` has a built-in `contains` method
fn contains<T: PartialEq>(slice: &[T], elem: &T) -> bool {
    slice.contains(elem)
}
