#
# @lc app=leetcode.cn id=232 lang=python3
#
# [232] 用栈实现队列
#

# @lc code=start
class MyQueue:

    def __init__(self):
        self.instack = []
        self.outstack = []

    def push(self, x: int) -> None:
        self.instack.append(x)

    def in2out(self) -> None:
        while self.instack:
            self.outstack.append(self.instack.pop())

    def pop(self) -> int:
        if len(self.outstack) == 0:
            self.in2out()
        return self.outstack.pop()

    def peek(self) -> int:
        if len(self.outstack) == 0:
            self.in2out()
        return self.outstack[-1]

    def empty(self) -> bool:
        return not self.instack and not self.outstack


# Your MyQueue object will be instantiated and called as such:
# obj = MyQueue()
# obj.push(x)
# param_2 = obj.pop()
# param_3 = obj.peek()
# param_4 = obj.empty()
# @lc code=end
