class Solution:
    def removeElement(self, nums: List[int], val: int) -> int:
        tempArray = []                

        for num in nums:
            if num == val:
                continue
            tempArray.append(num)
        
        for i in range(len(tempArray)):
            nums[i] = tempArray[i]
                
        return len(tempArray)