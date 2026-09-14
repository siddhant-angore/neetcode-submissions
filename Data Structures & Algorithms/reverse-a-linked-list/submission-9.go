/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
    var prv *ListNode
	cur := head
	nxt := head
	for cur != nil {
		nxt = cur.Next
		cur.Next = prv
		prv = cur
		cur = nxt
	}
	return prv
}
