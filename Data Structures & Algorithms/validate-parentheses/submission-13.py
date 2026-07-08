class Solution:
    def isValid(self, s: str) -> bool:
        ourStack = []
        openingParentheses = {'(': ')', '{': '}', '[': ']'}

        for c in s:
            # If c is opening bracket then push it's closing bracket
            if c in openingParentheses:
                ourStack.append(openingParentheses[c])
            # If c is closing bracket see if it already exists at the top
            else:
                # If exists pop it
                if len(ourStack) != 0 and ourStack[-1] == c:
                    ourStack.pop()
                else:
                    return False                
        
        return len(ourStack) == 0