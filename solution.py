def largest(self, arr):
        # code here
        max = arr[0]
        for num in arr:
            if max <= num:
                max = num
        return max