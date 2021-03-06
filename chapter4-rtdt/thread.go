package chapter4_rtdt

import (
	"GoVM/chapter6-obj/heap"
)

type Thread struct {
	// pc 寄存器中存放当前正在执行的java虚拟机指令的 地址
	pc    int
	// java虚拟机栈指针
	stack *Stack
}

func NewThread() *Thread {
	return &Thread{
		stack:        newStack(1024),
	}
}

func (self *Thread) ClearStack() {
	self.stack.clear()
}

func (self *Thread) PC() int {
	return self.pc
}

func (self *Thread) GetFrames() []*Frame {
	return self.stack.getFrames()
}

func (self *Thread) SetPC(pc int) {
	self.pc = pc
}

func (self *Thread) IsStackEmpty() bool {
	return self.stack.IsEmpty()
}

/**
	frame指虚拟机中线程栈的栈帧
 */
func (self *Thread) PushFrame(frame *Frame) {
	self.stack.push(frame)
}

func (self *Thread) PopFrame() *Frame {
	return self.stack.pop()
}

func (self *Thread) TopFrame() *Frame {
	return self.stack._top
}

func (self *Thread) CurrentFrame() *Frame {
	return self.stack.top()
}

func (self *Thread) NewFrame(method *heap.Method) *Frame {
	return newFrame(self, method)
}
