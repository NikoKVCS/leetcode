
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
    def depth(self, root: 'TreeNode', p: 'TreeNode', q: 'TreeNode') -> ('TreeNode', 'bool'):
        if root == None:
            return None, False
        found_root = False
        n1, found_left = self.depth(root.left,p,q)
        n2, found_right = self.depth(root.right,p,q)
        if n1 != None:
            return n1, True
        if n2 != None:
            return n2, True
        if root == p or root == q :
            found_root = True
        if (found_root and (found_left or found_right)) or (found_left and found_right):
            return root, True
        elif found_root or found_left or found_right:
            return None, True
        return None, False
    
    def lowestCommonAncestor(self, root: 'TreeNode', p: 'TreeNode', q: 'TreeNode') -> 'TreeNode':
        node, q = self.depth(root,p,q)
        return node

b0 = TreeNode(0)
b1 = TreeNode(1)
b2 = TreeNode(2)
b3 = TreeNode(3)
b4 = TreeNode(4)
b5 = TreeNode(5)
b6 = TreeNode(6)
b7 = TreeNode(7)
b8 = TreeNode(8)
b5.left = b6
b5.right = b2
b2.left = b7
b2.right = b4
b1.left = b0
b1.right = b8
b3.left = b5
b3.right = b1
a = Solution()
q = a.lowestCommonAncestor(b3,b5,b1)
print(q.val)