//
// Created by yangzhengkun on 2020/7/22.
//

#include <vector>
#include <stack>

using namespace std;

struct TreeNode {
  int val;
  TreeNode *left;
  TreeNode *right;
  explicit TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
};

class Solution {
 public:

  // 递归解法
  vector<int> res;
  vector<int> preorderTraversal(TreeNode *root) {
    if (root) {
      res.push_back(root->val);
      preorderTraversal(root->left);
      preorderTraversal(root->right);
    }
  }

  // 迭代解法
  vector<int> preorderTraversal1(TreeNode *root) {
    vector<int> ret;
    if (root == nullptr) {
      return ret;
    }
    stack<TreeNode *> s;
    s.push(root);
    while (!s.empty()) {
      TreeNode *cur = s.top();
      s.pop();
      ret.push_back(cur->val);
      if (cur->right) {
        s.push(cur->right);
      }
      if (cur->left) {
        s.push(cur->left);
      }
    }
    return ret;
  }
};


