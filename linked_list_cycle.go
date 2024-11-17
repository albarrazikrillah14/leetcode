/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func hasCycle(head *ListNode) bool {
	if head == nil {
			return false
	}

	slow, fast := head, head

	for fast != nil && fast.Next != nil {
			slow = slow.Next
			fast = fast.Next.Next

			if slow == fast {
					return true // Cycle detected
			}
	}

	return false // No cycle found
}

slow 3 2 0
fast 3 0 -4