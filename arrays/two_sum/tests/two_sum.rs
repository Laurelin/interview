use std::time::Instant;

use serde::Deserialize;
use two_sum::Solution;

#[derive(Deserialize)]
struct LiteralCase {
    name: String,
    nums: Vec<i32>,
    target: i32,
    expected: [i32; 2],
}

#[derive(Deserialize)]
struct ScaleSpec {
    name: String,
    n: usize,
    fill: i32,
    left: i32,
    right: i32,
    target: i32,
    pair_at: PairAt,
}

#[derive(Deserialize, PartialEq, Eq)]
#[serde(rename_all = "lowercase")]
enum PairAt {
    Start,
    End,
}

fn expand_scale(spec: ScaleSpec) -> LiteralCase {
    let mut nums = vec![spec.fill; spec.n];
    let expected = match spec.pair_at {
        PairAt::Start => [0, 1],
        PairAt::End => [spec.n as i32 - 2, spec.n as i32 - 1],
    };
    nums[expected[0] as usize] = spec.left;
    nums[expected[1] as usize] = spec.right;
    LiteralCase {
        name: spec.name,
        nums,
        target: spec.target,
        expected,
    }
}

fn same_indices(actual: &[i32], expected: [i32; 2]) -> bool {
    matches!(actual, [a, b] if (*a == expected[0] && *b == expected[1]) || (*a == expected[1] && *b == expected[0]))
}

#[test]
fn two_sum_cases() {
    let cases: Vec<LiteralCase> =
        serde_json::from_str(include_str!("../cases.json")).expect("cases.json");

    for case in cases {
        let got = Solution::two_sum(case.nums, case.target);
        assert!(
            same_indices(&got, case.expected),
            "{}: got {:?}, want {:?} (either order)",
            case.name,
            got,
            case.expected
        );
    }
}

#[test]
fn two_sum_scale() {
    let specs: Vec<ScaleSpec> =
        serde_json::from_str(include_str!("../scale.json")).expect("scale.json");

    for spec in specs {
        let case = expand_scale(spec);
        let started = Instant::now();
        let got = Solution::two_sum(case.nums, case.target);
        let elapsed = started.elapsed();

        assert!(
            same_indices(&got, case.expected),
            "{}: got {:?}, want {:?} (either order)",
            case.name,
            got,
            case.expected
        );
        println!("{}: {:?}", case.name, elapsed);
    }
}
