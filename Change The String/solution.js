class Solution{
    modify(s){
        let first = s.charAt(0);
        //using new string to conccat since js string are immutable
        let result = first;
        if(first == first.toUpperCase())
            for(let i=1;i<s.length;i++)
                //you can use str.concat in large submission range ofr optimal results
                result += s[i].toUpperCase();
        else if(first == first.toLowerCase())
            for(let i=1;i<s.length;i++)
                result += s[i].toLowerCase()
        return result;
    }
}