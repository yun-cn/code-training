class Solution(object):
    def reverse(self, x):
        """
        :type x: int
        :rtype: int
        """
        def func(x,res=0):
            x= x if x>=0 else -x
            if x==0:
                return res
            res = res*10+x%10
            w = func(x//10,res)
            if res > 2**31:
                return 0
            return w
        return func(x) if x>=0 else -1*func(x)

