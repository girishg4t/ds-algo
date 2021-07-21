function makePalindrome(s) {
    let str = s.split("")
    let i = 0;
    let count = 0;
    let j = str.length - 1
    while (i <= j) {
        if (str[i] == str[j]) {
            i++
            j--
        } else {
            count++
            i++
        }
    }
   return count
}

console.log(makePalindrome("JAVAJ"))