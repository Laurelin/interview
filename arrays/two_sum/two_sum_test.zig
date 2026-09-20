const std = @import("std");
const twoSum = @import("two_sum.zig").twoSum;

const LiteralCase = struct {
    name: []const u8,
    nums: []const i32,
    target: i32,
    expected: [2]i32,
};

const PairAt = enum { start, end };

const ScaleSpec = struct {
    name: []const u8,
    n: usize,
    fill: i32,
    left: i32,
    right: i32,
    target: i32,
    pair_at: PairAt,
};

fn sameIndices(actual: [2]i32, expected: [2]i32) bool {
    return (actual[0] == expected[0] and actual[1] == expected[1]) or
        (actual[0] == expected[1] and actual[1] == expected[0]);
}

fn expandScale(allocator: std.mem.Allocator, spec: ScaleSpec) !LiteralCase {
    const nums = try allocator.alloc(i32, spec.n);
    @memset(nums, spec.fill);

    const expected: [2]i32 = switch (spec.pair_at) {
        .start => .{ 0, 1 },
        .end => .{ @intCast(spec.n - 2), @intCast(spec.n - 1) },
    };
    nums[@intCast(expected[0])] = spec.left;
    nums[@intCast(expected[1])] = spec.right;

    return .{
        .name = spec.name,
        .nums = nums,
        .target = spec.target,
        .expected = expected,
    };
}

fn expectTwoSum(case: LiteralCase) !void {
    const got = try twoSum(std.testing.allocator, case.nums, case.target);
    if (!sameIndices(got, case.expected)) {
        std.debug.print("{s}: got {any}, want {any} (either order)\n", .{
            case.name,
            got,
            case.expected,
        });
        return error.TestExpectedEqual;
    }
}

fn expectLiteral(name: []const u8) !void {
    const allocator = std.testing.allocator;
    const parsed = try std.json.parseFromSlice(
        []LiteralCase,
        allocator,
        @embedFile("cases.json"),
        .{},
    );
    defer parsed.deinit();

    const case = for (parsed.value) |c| {
        if (std.mem.eql(u8, c.name, name)) break c;
    } else return error.UnknownCase;

    try expectTwoSum(case);
}

fn expectScale(name: []const u8) !void {
    const allocator = std.testing.allocator;
    const parsed = try std.json.parseFromSlice(
        []ScaleSpec,
        allocator,
        @embedFile("scale.json"),
        .{},
    );
    defer parsed.deinit();

    const spec = for (parsed.value) |s| {
        if (std.mem.eql(u8, s.name, name)) break s;
    } else return error.UnknownCase;

    const case = try expandScale(allocator, spec);
    defer allocator.free(case.nums);

    const io = std.testing.io;
    const started = std.Io.Clock.awake.now(io);
    try expectTwoSum(case);
    const elapsed = started.durationTo(std.Io.Clock.awake.now(io));
    std.debug.print("{s}: {d} ms\n", .{ spec.name, elapsed.toMilliseconds() });
}

test "basic" {
    try expectLiteral("basic");
}
test "middle pair" {
    try expectLiteral("middle pair");
}
test "duplicate values" {
    try expectLiteral("duplicate values");
}
test "negative numbers" {
    try expectLiteral("negative numbers");
}
test "minimum length" {
    try expectLiteral("minimum length");
}
test "minimum length negatives" {
    try expectLiteral("minimum length negatives");
}
test "minimum length mixed signs" {
    try expectLiteral("minimum length mixed signs");
}
test "pair at start" {
    try expectLiteral("pair at start");
}
test "first and last" {
    try expectLiteral("first and last");
}
test "last two" {
    try expectLiteral("last two");
}
test "three elements first two" {
    try expectLiteral("three elements first two");
}
test "three elements last two" {
    try expectLiteral("three elements last two");
}
test "three elements first and last" {
    try expectLiteral("three elements first and last");
}
test "pair with large index gap" {
    try expectLiteral("pair with large index gap");
}
test "strictly descending" {
    try expectLiteral("strictly descending");
}
test "strictly ascending" {
    try expectLiteral("strictly ascending");
}
test "descending official values" {
    try expectLiteral("descending official values");
}
test "unsorted official values" {
    try expectLiteral("unsorted official values");
}
test "both zeros" {
    try expectLiteral("both zeros");
}
test "two zeros among others" {
    try expectLiteral("two zeros among others");
}
test "one zero" {
    try expectLiteral("one zero");
}
test "complement is zero" {
    try expectLiteral("complement is zero");
}
test "zeros with a negative" {
    try expectLiteral("zeros with a negative");
}
test "all negatives" {
    try expectLiteral("all negatives");
}
test "all negatives first and last" {
    try expectLiteral("all negatives first and last");
}
test "mixed signs" {
    try expectLiteral("mixed signs");
}
test "symmetric around zero" {
    try expectLiteral("symmetric around zero");
}
test "negative plus positive" {
    try expectLiteral("negative plus positive");
}
test "positive pair amid negatives" {
    try expectLiteral("positive pair amid negatives");
}
test "opposite extremes" {
    try expectLiteral("opposite extremes");
}
test "opposite extremes with distractor" {
    try expectLiteral("opposite extremes with distractor");
}
test "target at max bound" {
    try expectLiteral("target at max bound");
}
test "target at min bound" {
    try expectLiteral("target at min bound");
}
test "max value plus negative" {
    try expectLiteral("max value plus negative");
}
test "min value plus positive" {
    try expectLiteral("min value plus positive");
}
test "bound filled pair at start" {
    try expectLiteral("bound filled pair at start");
}
test "bound filled pair at end" {
    try expectLiteral("bound filled pair at end");
}
test "duplicate pair with distractor" {
    try expectLiteral("duplicate pair with distractor");
}
test "duplicate pair at start" {
    try expectLiteral("duplicate pair at start");
}
test "identical pair among others" {
    try expectLiteral("identical pair among others");
}
test "single half target not reused" {
    try expectLiteral("single half target not reused");
}
test "target exists as element" {
    try expectLiteral("target exists as element");
}
test "many identical distractors" {
    try expectLiteral("many identical distractors");
}
test "half target repeated not the pair" {
    try expectLiteral("half target repeated not the pair");
}
test "large complement pair" {
    try expectLiteral("large complement pair");
}
test "pair after large distractors" {
    try expectLiteral("pair after large distractors");
}
test "max length pair at end" {
    try expectScale("max length pair at end");
}
test "max length pair at start" {
    try expectScale("max length pair at start");
}
