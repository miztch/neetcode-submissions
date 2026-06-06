import re

class Solution:
	def isPalindrome(self, s: str) -> bool:
		t = re.sub("[^a-zA-Z0-9]","",s).lower()
		return t == t[::-1]