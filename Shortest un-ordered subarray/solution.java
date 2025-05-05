class Solution {
    
    public Boolean increasing(int arr[]){
        for(int i=0;i<arr.length-1;i++)
            if(arr[i]>arr[i+1])
                return false;
        return true;
    }
    
    public Boolean decreasing(int arr[]){
        for(int i=0;i<arr.length-1;i++)
            if(arr[i]<arr[i+1])
                return false;
        return true;
    }

    public int shortestUnorderedSubarray(int arr[]) {
       if(arr.length == 0)
            return 0;
       if(increasing(arr) || decreasing(arr))
            return 0;
        return 3;
    }
}