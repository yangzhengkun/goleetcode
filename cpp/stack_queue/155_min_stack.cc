//
// Created by yangzhengkun on 2020/7/24.
//

//设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。
//
//push(x) —— 将元素 x 推入栈中。
//pop() —— 删除栈顶的元素。
//top() —— 获取栈顶元素。
//getMin() —— 检索栈中的最小元素。
// 
//
//示例:
//
//输入：
//["MinStack","push","push","push","getMin","pop","top","getMin"]
//[[],[-2],[0],[-3],[],[],[],[]]
//
//输出：
//[null,null,null,null,-3,null,0,-2]
//
//解释：
//MinStack minStack = new MinStack();
//minStack.push(-2);
//minStack.push(0);
//minStack.push(-3);
//minStack.getMin();   --> 返回 -3.
//minStack.pop();
//minStack.top();      --> 返回 0.
//minStack.getMin();   --> 返回 -2.
//

#include <stack>
class MinStack {
 private:
  std::stack<int> stackData;
  std::stack<int> stackMin;
 public:
  /** initialize your data structure here. */
  MinStack() {
    stackMin.push(INT_MAX);
  }

  void push(int x) {
	stackData.push(x);
	stackMin.push(std::min(stackMin.top(), x));
  }

  void pop() {
	stackMin.pop();
	stackData.pop();
  }

  int top() {
	return stackData.top();
  }

  int getMin() {
	return stackMin.top();
  }
};