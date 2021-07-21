function longestPreFix(strs) {
    let firstStr = strs[0].split("")
    let foundPrevChar = ""
    let newStr = strs.slice(1, strs.length)
    for (let i = 0; i < firstStr.length; i++) {
        let strsLen = 0
        newStr.forEach((str, index) => {
            if (str.indexOf(firstStr[i]) == 0) {
                newStr[index] = str.substring(1, str.length)
                strsLen++
            }
        })
        if (strsLen != strs.length - 1) {
            break;
        }
        foundPrevChar = foundPrevChar + firstStr[i]
    }
    return foundPrevChar
}

console.log(longestPreFix(["flower","flow","flight"]))