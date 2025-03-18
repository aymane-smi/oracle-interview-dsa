class Solution {
    public static int largest(int[] arr) {
        //since arr length is >= 1
        int max = arr[0];
        for(int i=0;i<arr.length;i++)
            if(max <= arr[i])
                max = arr[i];
        return max;
    }
}