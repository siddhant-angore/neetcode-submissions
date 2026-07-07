class Solution:
    def replaceElements(self, arr: List[int]) -> List[int]:
        length = len(arr)
        currentMax = arr[-1]
        arr[-1] = -1

        for i in range(length - 2, -1, -1):
            tempValue = arr[i]            
            arr[i] = currentMax
            currentMax = max(currentMax, tempValue)
        
        return arr
                

