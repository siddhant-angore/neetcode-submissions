class ListNode:
    def __init__(self, value):
        self.value = value
        self.next = None
        self.previous = None

class MyLinkedList:

    def __init__(self):
        self.head = ListNode(-1)
        self.tail = ListNode(-1)
        self.head.next = self.tail
        self.tail.previous = self.head

    def get(self, index: int) -> int:
        current = self.head.next
        while current and index > 0:
            current = current.next
            index -= 1
        if current and current != self.tail and index == 0:
            return current.value
        return -1

    def addAtHead(self, val: int) -> None:
        newNode, next, previous = ListNode(val), self.head.next, self.head
        previous.next = newNode
        next.previous = newNode
        newNode.next = next
        newNode.previous = previous

    def addAtTail(self, val: int) -> None:
        newNode, next, previous = ListNode(val), self.tail, self.tail.previous
        previous.next = newNode 
        next.previous = newNode        
        newNode.next = next
        newNode.previous = previous

    def addAtIndex(self, index: int, val: int) -> None:
        current = self.head.next
        while current and index > 0:
            current = current.next
            index -= 1
        if current and index == 0:
            newNode, next, previous = ListNode(val), current, current.previous
            previous.next = next.previous = newNode
            newNode.next, newNode.previous = next, previous

    def deleteAtIndex(self, index: int) -> None:
        current = self.head.next
        while current and index > 0:
            current = current.next
            index -= 1
        if current and current != self.tail and index == 0:
            next, previous = current.next, current.previous
            next.previous = previous
            previous.next = next


# Your MyLinkedList object will be instantiated and called as such:
# obj = MyLinkedList()
# param_1 = obj.get(index)
# obj.addAtHead(val)
# obj.addAtTail(val)
# obj.addAtIndex(index,val)
# obj.deleteAtIndex(index)