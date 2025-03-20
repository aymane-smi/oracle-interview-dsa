class Solution{
    String modify(String s){
        StringBuilder strBuild = new StringBuilder(s);
        // we assume that s length is greater or equal 1
        if(Character.isUpperCase(strBuild.charAt(0)))
            for(int i=1;i<strBuild.length();i++)
                strBuild.setCharAt(i, Character.toUpperCase(strBuild.charAt(i)));
        else if(Character.isLowerCase((strBuild.charAt(0))))
            for(int i=1;i<strBuild.length();i++)
                strBuild.setCharAt(i, Character.toLowerCase(strBuild.charAt(i)));
        return strBuild.toString();
    }
}