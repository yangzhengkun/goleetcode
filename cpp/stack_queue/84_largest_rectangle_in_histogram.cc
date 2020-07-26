//
// Created by yangzhengkun on 2020/7/25.
//

//给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
//
//求在该柱状图中，能够勾勒出来的矩形的最大面积。给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
//

//单调栈
#include <vector>
#include <stack>
#include <iostream>
using namespace std;

class Solution {
 public:
  int largestRectangleArea(vector<int> &heights) {
	int res;
	std::stack<int> s;
	for (int i = 0; i < heights.size(); ++i) {
	  while (!s.empty() && heights[s.top()] >= heights[i]) {
		int top = s.top();
		s.pop();
		int left = s.empty() ? -1 : s.top();
		res = std::max(res, (i - left - 1) * heights[top]);
	  }
	  s.push(i);
	}
	while (!s.empty()) {
	  int top = s.top();
	  s.pop();
	  int left = s.empty() ? -1 : s.top();
	  res = std::max(res, ((int)heights.size() - left - 1) * heights[top]);
	}
	return res;
  }
};

class Solution1 {
 public:
  int largestRectangleArea(vector<int> &heights) {
	int n = heights.size();
	vector<int> left(n), right(n, n);
	stack<int> s;
	for (int i = 0; i < n; ++i) {
	  while (!s.empty() && heights[s.top()] >= heights[i]) {
		right[s.top()] = i;
		s.pop();
	  }
	  left[i] = (s.empty() ? -1 : s.top());
	  s.push(i);
	}

	int ans = 0;
	for (int i = 0; i < n; ++i) {
	  ans = max(ans, (right[i] - left[i] - 1) * heights[i]);
	}
	return ans;
  }
};

int main() {
  vector<int> s{2, 1, 5, 6, 2, 3};
  std::cout << Solution().largestRectangleArea(s) << std::endl;
  return 0;
}