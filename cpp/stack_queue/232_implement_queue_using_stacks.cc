//
// Created by yangzhengkun on 2020/7/24.
//

//使用栈实现队列的下列操作：
//
//push(x) -- 将一个元素放入队列的尾部。
//pop() -- 从队列首部移除元素。
//peek() -- 返回队列首部的元素。
//empty() -- 返回队列是否为空。
// 
//
//示例:
//
//MyQueue queue = new MyQueue();
//
//queue.push(1);
//queue.push(2);
//queue.peek();  // 返回 1
//queue.pop();   // 返回 1
//queue.empty(); // 返回 false
//


#include <stack>
class MyQueue {
 private:
  std::stack<int> inStack;
  std::stack<int> outStack;
 public:
  /** Initialize your data structure here. */
  MyQueue() {
  }

  /** Push element x to the back of queue. */
  void push(int x) {
	inStack.push(x);
  }

  /** Removes the element from in front of queue and returns that element. */
  int pop() {
	trans();
	int r = outStack.top();
	outStack.pop();
	return r;
  }

  /** Get the front element. */
  int peek() {
	trans();
	return outStack.top();
  }

  /** Returns whether the queue is empty. */
  bool empty() {
	return inStack.empty() && outStack.empty();
  }
	/** 从inStack栈倒输入到outStack栈*/
  void trans() {
	if (outStack.empty()) {
	  while (!inStack.empty()) {
		int v = inStack.top();
		inStack.pop();
		outStack.push(v);
	  }
	}
  }
};