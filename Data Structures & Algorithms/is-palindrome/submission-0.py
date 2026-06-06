import re

class Solution:
	def isPalindrome(self, s: str) -> bool:
		t = re.sub("[^a-zA-Z0-9]","",s)
		for i,x in enumerate(t):
			if x.lower() != t[::-1][i].lower():
				return False
				break
			
		return True
        