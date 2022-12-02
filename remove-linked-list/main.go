// Singly-linked lists are already defined with this interface:
// type ListNode struct {
//   Value interface{}
//   Next *ListNode
// }
// I can delete anywhere
// I need to know who is my previous node and my current node
//  current is l
// for every current that is different from nil
// if current value == k, and my previous is not nil
// l takes the place of next
// if my previous is nil, then my previous.next becomes my current
// if k != current value, my previous becomes my current

func solution(l *ListNode, k int) *ListNode {
	var prev *ListNode
	curr := l // current is my head

	for curr != nil { // is my current null? no
		if curr.Value == k { // is the value of my current like the one i want to remove? yes
			if prev == nil { // is the previous from my current null? yes, it's the head
				l = curr.Next // head is now the next one
			} else { // no, the previous from my current null is not null
				prev.Next = curr.Next // so my previous will be the one after my current
			}
		} else { //the value of my current is not like the one I want to remove
			prev = curr // my previous is now my current
		}
		curr = curr.Next // move to the next one

	} // if my current is nil, return my list

	return l
}

