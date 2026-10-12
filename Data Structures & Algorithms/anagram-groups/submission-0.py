class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        anagrams = defaultdict(list)
        for s in strs:
            tempArr = [0] * 26

            for char in s:
                tempArr[ord(char) - ord("a")] += 1
            
            anagrams[tuple(tempArr)].append(s)

        return list(anagrams.values())

            

        
