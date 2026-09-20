# 1. Two Sum

**Difficulty:** Easy

Given an array of integers `nums` and an integer `target`, return the indices of the two numbers such that they add up to `target`.

You may assume that each input has **exactly one** solution, and you may not use the same element twice.

You can return the answer in any order.

## Examples

### Example 1

```
Input: nums = [2, 7, 11, 15], target = 9
Output: [0, 1]
```

Because `nums[0] + nums[1] == 9`, we return `[0, 1]`.

### Example 2

```
Input: nums = [3, 2, 4], target = 6
Output: [1, 2]
```

### Example 3

```
Input: nums = [3, 3], target = 6
Output: [0, 1]
```

## Constraints

- `2 <= nums.length <= 10^4`
- `-10^9 <= nums[i] <= 10^9`
- `-10^9 <= target <= 10^9`
- Only one valid answer exists.

## Follow-up

Can you come up with an algorithm that is less than $O(n^2)$ time complexity?

## Test fixtures

- `cases.json` — literal cases. Each entry is `{ name, nums, target, expected }`.
- `scale.json` — generation specs for constraint-max inputs. Any language expands a spec before running:

Given `{ n, fill, left, right, target, pair_at }`:

1. Build `nums` of length `n`, every slot set to `fill`.
2. If `pair_at` is `"start"`, set `nums[0] = left` and `nums[1] = right`. Expected indices are `[0, 1]`.
3. If `pair_at` is `"end"`, set `nums[n - 2] = left` and `nums[n - 1] = right`. Expected indices are `[n - 2, n - 1]`.
4. Assert the solution returns those two indices, in either order.
5. Timing on scale cases is report-only. Do not fail on elapsed time.

These specs are unique by construction: `left + right == target`, and `fill` does not pair with itself or with either pair value.
