export class ListNode {
    val: number
    next: ListNode | null
    constructor(val?: number, next?: ListNode | null) {
        this.val = (val===undefined ? 0 : val)
        this.next = (next===undefined ? null : next)
    }
}
export function addTwoNumbers(l1: ListNode | null, l2: ListNode | null): ListNode | null {
    let dummy = new ListNode(0)
    let tail = dummy
    let carry = 0

    while (!done(l1, l2, carry)) {
        let [digit, nextCarry] = addDigits(getVal(l1), getVal(l2), carry)
        tail = appendDigit(tail, digit)
        carry = nextCarry
        l1 = advance(l1)
        l2 = advance(l2)
    }
    return dummy.next
};

function done(l1: ListNode | null, l2: ListNode | null, carry: number): boolean {
    return l1 === null && l2 === null && carry === 0
}

function getVal(node: ListNode | null): number {
    return node?.val ?? 0
}

function advance(node: ListNode | null): ListNode | null {
    return node?.next ?? null
}

function addDigits(a: number, b: number, carry: number): [number, number] {
    let sum = a + b + carry
    return [sum % 10, Math.floor(sum / 10)]
}

function appendDigit(tail: ListNode, digit: number): ListNode {
    tail.next = new ListNode(digit)
    return tail.next
}