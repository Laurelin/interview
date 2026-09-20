const std = @import("std");

pub fn twoSum(allocator: std.mem.Allocator, nums: []const i32, target: i32) ![2]i32 {
    var seen = std.AutoHashMap(i32, i32).init(allocator);
    defer seen.deinit();

    for (nums, 0..) |num, i| {
        const complement = target - num;
        if (seen.get(complement)) |j| {
            return .{ j, @intCast(i) };
        }
        try seen.put(num, @intCast(i));
    }
    unreachable;
}
