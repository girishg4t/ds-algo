function longestPlain() {
    let s = "babad"
    let r = "dabab"
    let res = ""
    let start = 0
    let end = s.length - 1
    let maxlength = 0
    for (let i = 0; i < s.length; i++) {
        start = 0
        end = s.length - 1
        let j = i
        while (start < end) {
            if (s[j] == r[start]) {
                res = res + s[j]
                maxlength++
            } else {
                res = ""
            }
            j++
            start++
        }
        console.log(res)
        maxlength = 0
        res = ""
    }
}
longestPlain()