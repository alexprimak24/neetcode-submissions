class Solution:
    def calPoints(self, operations: List[str]) -> int:
        score_arr = []

        for i, val in enumerate(operations):
            match val:
                case "+":
                    temp = score_arr.pop()
                    new = score_arr[-1] + temp

                    score_arr.append(int(temp))
                    score_arr.append(int(new))

                   
                case "C":
                    score_arr.pop()
                case "D":
                    score_arr.append(int(score_arr[-1] * 2))
                case _:
                    score_arr.append(int(val))

        score = 0
        print(score_arr)
        for i, value in enumerate(score_arr):
            score += value
            
        return score
        