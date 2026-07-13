class Solution:
    def countStudents(self, students: List[int], sandwiches: List[int]) -> int:
        n = len(students)
        q = deque(students)

        studentsInQueue = n
        for sandwich in sandwiches:
            notHavingSandwich = 0
            while notHavingSandwich < n and q[0] != sandwich:
                currentStudent = q.popleft()
                q.append(currentStudent)
                notHavingSandwich += 1
            
            if q[0] == sandwich:
                q.popleft()
                studentsInQueue -= 1
            else:
                break
        
        return studentsInQueue