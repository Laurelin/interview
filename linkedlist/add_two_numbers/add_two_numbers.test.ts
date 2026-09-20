import { readFileSync } from "node:fs";
import path from "node:path";
import { describe, expect, it } from "vitest";
import { addTwoNumbers, ListNode } from "./add_two_numbers.ts";

type LiteralCase = {
  name: string;
  l1: number[];
  l2: number[];
  expected: number[];
};

type ScaleSpec = {
  name: string;
  l1_len: number;
  l1_fill: number;
  l2_len: number;
  l2_fill: number;
};

const dir = import.meta.dirname;
const cases: LiteralCase[] = JSON.parse(
  readFileSync(path.join(dir, "cases.json"), "utf8"),
);
const scaleSpecs: ScaleSpec[] = JSON.parse(
  readFileSync(path.join(dir, "scale.json"), "utf8"),
);

function filled(n: number, fill: number): number[] {
  return Array<number>(n).fill(fill);
}

function schoolbookAdd(l1: number[], l2: number[]): number[] {
  const out: number[] = [];
  let carry = 0;
  let i = 0;
  while (i < l1.length || i < l2.length || carry > 0) {
    let sum = carry;
    if (i < l1.length) {
      sum += l1[i];
    }
    if (i < l2.length) {
      sum += l2[i];
    }
    out.push(sum % 10);
    carry = Math.floor(sum / 10);
    i++;
  }
  return out;
}

function expandScale(spec: ScaleSpec): LiteralCase {
  const l1 = filled(spec.l1_len, spec.l1_fill);
  const l2 = filled(spec.l2_len, spec.l2_fill);
  return {
    name: spec.name,
    l1,
    l2,
    expected: schoolbookAdd(l1, l2),
  };
}

function toList(digits: number[]): ListNode | null {
  const dummy = new ListNode(0);
  let tail = dummy;
  for (const d of digits) {
    tail.next = new ListNode(d);
    tail = tail.next;
  }
  return dummy.next;
}

function fromList(node: ListNode | null): number[] {
  const out: number[] = [];
  while (node !== null) {
    out.push(node.val);
    node = node.next;
  }
  return out;
}

describe("addTwoNumbers", () => {
  it.each(cases)("$name", ({ l1, l2, expected }) => {
    expect(fromList(addTwoNumbers(toList(l1), toList(l2)))).toEqual(expected);
  });
});

describe("addTwoNumbers scale", () => {
  it.each(scaleSpecs.map(expandScale))("$name", ({ name, l1, l2, expected }) => {
    const started = performance.now();
    const actual = fromList(addTwoNumbers(toList(l1), toList(l2)));
    const elapsedMs = performance.now() - started;

    expect(actual).toEqual(expected);
    console.log(`${name}: ${elapsedMs.toFixed(2)} ms`);
  });
});
