
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
    def maxDepth(self, root: TreeNode) -> int:
        if root == None:
            return 0
        return self.depth(root, 0)
        
    def depth(self, node:TreeNode, h:int):
        if node == None:
            return h
        h = h+1
        if node.left == None and node.right == None:
            return h
        h1 = self.depth(node.left, h)
        h2 = self.depth(node.right, h)
        return max(h1,h2)
        
