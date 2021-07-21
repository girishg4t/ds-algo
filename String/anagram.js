function anagram(s, t) {
    let newA = s.split("")
    let newB = t.split("")
    if (newA.length != newB.length) {
        return false
    }
    for (let i = 0; i < newA.length; i++) {
        let index = newB.indexOf(newA[i])
        if (index > -1) {
            newB.splice(index, 1)
        } else {
            break
        }
    }
    if (newB == "") {
        return true
    }
    return false
}

console.log(anagram("aaaa", "aaaa"))