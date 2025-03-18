largest(arr) {
    let max = arr[0];
    arr.forEach((num)=>{
        if(max <= num)
            max = num;
    });
    return max;
}