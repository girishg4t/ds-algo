function adjusentDuplicates(S) {
    let s = S.split('')
    let index = 1;
    let maxlen = 1
    let prevMax = 0
    for (let i = 0; i < s.length; i++) {
        index = i + 1
        if (s[i] == s[index]) {
            while (s[i] == s[index]) {
                index++
                maxlen++
            }
            if (maxlen == s.length) {
                console.log(s[i])
            }
            s.splice(i, maxlen);
            i = -1
            maxlen = 1
        }

    }
    console.log(s.join(""))
}

adjusentDuplicates("aaaaaaaaa")