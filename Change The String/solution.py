class Solution:
    def modify(self, s):
        result = s[0]
        for i in range(1, len(s)):
            if s[0].islower():
                result += s[i].lower()
            elif s[0].isupper():
                result += s[i].upper()
        return result
