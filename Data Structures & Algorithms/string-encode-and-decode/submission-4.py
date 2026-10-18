class Solution:

    def encode(self, strs: List[str]) -> str:
        finalString = ""
        for word in strs:
            finalString += f"{len(word)}#{word}"
        # print(finalString)
        return finalString
            
    def decode(self, s: str) -> List[str]:
        acc = []
        while True:
            if not s:
                return acc
            
            for i, char in enumerate(s):
                if char == "#":
                    wordLen = int(s[:i])
                    # print(wordLen)
                    acc.append(s[i+1:wordLen+i+1])
                    # print(s[i+1:wordLen+i+1])
                    s = s[wordLen + i+1:]
                    break
                    # print(s[wordLen + i+1:])
                    

            
            # print(length, specialChar)
            # acc.append(s[2:int(length)+2])
            # s = s[int(length)+2:]

        return acc
