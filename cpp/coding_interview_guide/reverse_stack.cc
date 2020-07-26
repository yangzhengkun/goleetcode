//
// Created by yangzhengkun on 2020/7/24.
//

//仅用递归函数和栈操作逆序一个栈
//一个栈依次压入1,2,3,4,5、那么从栈顶到栈底分别为5,4,3,2,1。将这个栈转置后，从栈顶到栈底为1,2,3,4,5. 也就是实现栈中元素的逆序，但是只能用递归函数来实现，不能用其他数据结构。
#include <stack>

int getAndRemoveLastElement(std::stack<int> &s) {
  int res = s.top();
  s.pop();
  if (s.empty()) {
	return res;
  } else {
	int r = getAndRemoveLastElement(s);
	s.push(r);
	return r;
  }
}

void reverse(std::stack<int> &s) {
  if (s.empty()) {
	return;
  }

  int i = getAndRemoveLastElement(s);
  reverse(s);
  s.push(i);
}


