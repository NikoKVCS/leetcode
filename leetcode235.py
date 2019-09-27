# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def lowestCommonAncestor(self, root: 'TreeNode', p: 'TreeNode', q: 'TreeNode') -> 'TreeNode':
        queue = []
        queue.append(root)
        while len(queue) != 0 :
            item = queue[0]
            queue.remove(item)
            if ( p.val <= item.val and item.val <= q.val) or ( q.val <= item.val and item.val <= p.val):
                return item
            if item.left != None:
                queue.append(item.left)
            if item.right != None:
                queue.append(item.right)
            
        