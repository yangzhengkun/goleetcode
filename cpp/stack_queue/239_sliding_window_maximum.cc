//
// Created by yangzhengkun on 2020/7/24.
//

#include <vector>
#include <deque>
using namespace std;
class Solution {
 public:
  vector<int> maxSlidingWindow(vector<int> &nums, int k) {
	vector<int> res;
	if (k == 0)return {};
	deque<int> window;
	int right = 0;
	while (right < nums.size()) {
	  if (!window.empty() && window.front() <= right - k) {
		//窗口中的索引过期
		window.pop_front();
	  }
	  while (!window.empty() && nums[window.back()] < nums[right]) {
		window.pop_back();
	  }
	  window.push_back(right);
	  right++;
	  if (right >= k) {
		res.push_back(nums[window.front()]);
	  }
	}
	return res;
  }
};