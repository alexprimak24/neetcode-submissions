class Solution:
    def isValid(self, s: str) -> bool:
        temp = []
        for ch in s:
            match ch:
                case '(' | '{' | '[':
                    temp.append(ch)
               
                case ')':
                    if len(temp) == 0 or temp[-1] != '(':
                        return False
                    temp.pop()
               
                case '}':
                    if len(temp) == 0 or temp[-1] != '{':
                        return False
                    temp.pop()
                case ']':
                    if len(temp) == 0 or temp[-1] != '[':
                        return False
                    temp.pop()
    
        
        if len(temp) > 0:
            return False
        return True