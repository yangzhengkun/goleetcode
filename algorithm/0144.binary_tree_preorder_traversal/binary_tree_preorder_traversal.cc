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

    TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
};

// 递归解法
class Solution {
public:
    std::vector<int> result;

    std::vector<int> preorderTraversal(TreeNode *root) {
        if (root) {
            result.push_back(root->val);
            preorderTraversal(root->left);
            preorderTraversal(root->right);
        }
        return result;
    }
};


// 使用栈迭代
class Solution {
public:
    std::vector<int> preorderTraversal(TreeNode *root) {
        std::vector<int> res;
        if (root == nullptr) {
            return result;
        }
        std::stack<TreeNode *> s;
        s.push(root);

        return result;
    }
};

