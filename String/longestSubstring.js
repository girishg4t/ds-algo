function longestSubstring(s) {
    let str = s.split("")
    let count = 0;
    let dic = {}
    let maxCount = 0;
    for (let i = 0; i < str.length; i++) {
        if (!dic[str[i]]) {
            dic[str[i]] = 1
            count++;
        } else {
            if (count > maxCount) {
                maxCount = count
            }
            count = 0
            dic = {}
            dic[str[i]] = 1
            count++;
        }
    }
    return maxCount
}

console.log(longestSubstring("abcabcbb"))