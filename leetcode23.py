class ListNode:
    def __init__(self, x, next):
        self.val = x
        self.next = next

class Solution:
    def mergeKLists(self, lists) -> ListNode:
        n = len(lists)
        best = -1
        best_id = 0
        for i in range(0,n):
            if lists[i] == None:
                continue
            if lists[i].val <= best or best == -1:
                best = lists[i].val
                best_id = i
        if best == -1:
            return None
        result = lists[best_id]
        lists[best_id] = lists[best_id].next
        result.next = self.mergeKLists(lists)
        return result
                

q = [ListNode(2,None), None,ListNode(-1,None)]
a = Solution()
a.mergeKLists(q)