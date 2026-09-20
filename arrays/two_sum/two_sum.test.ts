import { readFileSync } from "node:fs";
import path from "node:path";
import { describe, expect, it } from "vitest";
import { twoSum } from "./two_sum.ts";

type LiteralCase = {
  name: string;
  nums: number[];
  target: number;
  expected: [number, number];
};

type PairAt = "start" | "end";

type ScaleSpec = {
  name: string;
  n: number;
  fill: number;
  left: number;
  right: number;
  target: number;
  pair_at: PairAt;
};

const dir = import.meta.dirname;
const cases: LiteralCase[] = JSON.parse(
  readFileSync(path.join(dir, "cases.json"), "utf8"),
);
const scaleSpecs: ScaleSpec[] = JSON.parse(
  readFileSync(path.join(dir, "scale.json"), "utf8"),
);

function expandScale(spec: ScaleSpec): LiteralCase {
  if (spec.pair_at !== "start" && spec.pair_at !== "end") {
    throw new Error(`unknown pair_at: ${spec.pair_at}`);
  }

  const nums = Array<number>(spec.n).fill(spec.fill);
  const expected: [number, number] =
    spec.pair_at === "start" ? [0, 1] : [spec.n - 2, spec.n - 1];
  nums[expected[0]] = spec.left;
  nums[expected[1]] = spec.right;

  return { name: spec.name, nums, target: spec.target, expected };
}

function sameIndices(actual: number[], expected: [number, number]): boolean {
  if (actual.length !== 2) {
    return false;
  }
  const [a, b] = actual;
  const [i, j] = expected;
  return (a === i && b === j) || (a === j && b === i);
}

describe("twoSum", () => {
  it.each(cases)("$name", ({ nums, target, expected }) => {
    expect(sameIndices(twoSum(nums, target), expected)).toBe(true);
  });
});

describe("twoSum scale", () => {
  it.each(scaleSpecs.map(expandScale))("$name", ({ name, nums, target, expected }) => {
    const started = performance.now();
    const actual = twoSum(nums, target);
    const elapsedMs = performance.now() - started;

    expect(sameIndices(actual, expected)).toBe(true);
    console.log(`${name}: ${elapsedMs.toFixed(2)} ms`);
  });
});
