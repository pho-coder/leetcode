#!/usr/bin/env python3

class Solution(object):
    def reverseWords(self, s):
        """
        :type s: str
        :rtype: str
        """
        s = s.strip().split(" ")
        a = []
        for i in range(0, len(s)):
            if s[i].strip() != "":
                a.append(s[i])
        s = a
        for i in range(0, int(len(s)/2)):
            tmp = s[i]
            s[i] = s[len(s)-i-1]
            s[len(s)-i-1] = tmp
        return " ".join(x.strip() for x in s)

if __name__ == "__main__":
    s = Solution().reverseWords(" the    sky is blue ")
    print(s)