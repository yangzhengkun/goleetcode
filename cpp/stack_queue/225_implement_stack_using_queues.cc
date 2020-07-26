//
// Created by yangzhengkun on 2020/7/24.
//

#include <queue>

class MyStack {
 private:
  std::queue<int> q;
 public:
  MyStack() = default;
  void push(int x) {
	q.push(x);
	for (int i = 0; i < q.size() - 1; ++i) {
	  q.push(q.front());
	  q.pop();
	}
  }

  int pop() {
	int r = q.front();
	q.pop();
	return r;
  }

  int top() {
	return q.front();
  }

  bool empty() {
	return q.empty();
  }
};