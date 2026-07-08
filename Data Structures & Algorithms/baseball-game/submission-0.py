class Solution:
    def calPoints(self, operations: List[str]) -> int:
        recordBoard = []

        for op in operations:
            if op == "D":
                recordBoard.append(2 * int(recordBoard[-1]))
            elif op == "C":
                recordBoard.pop()
            elif op == "+":
                recordBoard.append(int(recordBoard[-1]) + int(recordBoard[-2]))
            else:
                recordBoard.append(int(op))

        return self.sumOfRecords(recordBoard)
    
    def sumOfRecords(self, records: List[str]) -> int:
        finalSum = 0
        for record in records:
            finalSum += int(record)
        return finalSum