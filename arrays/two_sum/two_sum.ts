export function twoSum(nums: number[], target: number): number[] {
    let seen = new Map<number, number>();
    for (let i = 0; i< nums.length; i++) {
        let complement = target - nums[i];
        if (seen.has(complement)) {
            return [seen.get(complement)!, i];
        }
        seen.set(nums[i], i);
    }
    return [];
}