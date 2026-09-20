use std::collections::HashMap;

pub struct Solution;

impl Solution {
    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        let mut seen = HashMap::new();
        for (i, num) in nums.iter().enumerate() {
            let complement = target - num;
            if let Some(&index) = seen.get(&complement) {
                return vec![index as i32, i as i32];
            }
            seen.insert(num, i as i32);
        }
        vec![]
    }
}