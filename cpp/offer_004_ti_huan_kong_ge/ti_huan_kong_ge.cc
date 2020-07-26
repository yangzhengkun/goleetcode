//
// Created by yangzhengkun on 2020/7/22.
//

#include <string>
using namespace std;
class Solution {
 public:
  string replaceString (string s)
  {
	int len = s.length () - 1;
	// 先扩展空间
	for (int i = 0; i <= len; i++)
	  {
		if (s[i] == ' ')
		  {
			s += "00";
		  }
	  }
	int nl = s.length () - 1;
	if (nl <= len)
	  {
		return s;
	  }
	// 从原字符串的尾部开始替换
	for (int i = len; i > 0; i--)
	  {
		char c = s[i];
		if (c == ' ')
		  {
			s[nl--] = '0';
			s[nl--] = '2';
			s[nl--] = '%';
		  }
		else
		  {
			s[nl--] = c;
		  }
	  }
	return s;
  }

  string replaceString1 (string s)
  {
	string res;
	for (auto c:s)
	  {
		if (c == ' ')
		  {
			res += "%20";
		  }
		else
		  {
			res += c;
		  }
	  }
	return res;
  }
};