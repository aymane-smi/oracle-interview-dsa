class Solution {
    increasing(arr){
        for(let i=0;i<arr.length-1;i++)
            if(arr[i]>arr[i+1])
                return false;
        return true;
    }
    decreasing(arr){
        for(let i=0;i<arr.length-1;i++)
            if(arr[i]<arr[i+1])
                return false;
        return true;
    }
    // Function to find the length of the shortest unordered subarray
    // in a given array.
    shortestUnorderedSubarray(arr) {
        if(!arr.length)
            return 0;
        if(this.increasing(arr) || this.decreasing(arr))
            return 0;
        return 3;
    }
}