//
// Created by yangzhengkun on 2020/7/22.
//

#include <vector>

using namespace std;

class Solution {
 public:
  bool searchMatrix(const vector<vector<int>> &matrix, int target) {
    //边界值
    if (matrix.size() == 0 || matrix[0].size() == 0) {
      return false;
    }
    int row = matrix.size();
    int i = 0, j = matrix[0].size() - 1;
    while (i < row && j >= 0) {
      if (matrix[i][j] == target) {
        return true;
      } else if (matrix[i][j] < target) {
        ++i;
      } else if (matrix[i][j] > target) {
        --j;
      }
    }
    return false;
  }
};