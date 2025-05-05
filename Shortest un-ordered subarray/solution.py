class Solution:
    def increasing(self, arr):
        for i in range(len(arr)-1):
            if arr[i] > arr[i+1]:
                return False
        return True
    def decreasing(self, arr):
        for i in range(len(arr) - 1):
            if arr[i] < arr[i+1]:
                return False
        return True
    def shortestUnorderedSubarray(self, arr):
        if len(arr) == 0:
            return 0
        if self.increasing(arr) or self.decreasing(arr):
            return 0
        return 3