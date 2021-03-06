package main

import "errors"

type callStack struct {
    stack []uint16
    tos int
    capacity int
}

func newStack(size int) *callStack {
    cstack := new(callStack)
    cstack.capacity = size
    cstack.tos = -1
    cstack.stack = make([]uint16, size)
    return cstack
}

func (stack *callStack) push(instr uint16) error {
    if stack.full() {
        return errors.New("stack is full")
    }

    stack.tos++
    stack.stack[stack.tos] = instr
    return nil
}

func (stack *callStack) pop() (uint16, error) {
    if stack.empty() {
        return 0, errors.New("stack is empty")
    }

    retval := stack.stack[stack.tos]
    stack.tos--
    return retval, nil
}

func (stack *callStack) clear() {
    stack.tos = -1
}

func (stack *callStack) empty() bool {
    return stack.tos == -1
}

func (stack *callStack) full() bool {
    return stack.tos == stack.capacity - 1
}
