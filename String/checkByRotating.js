function checkByRotating() {
    let a = "amazon"
    let b = "onamazonamaz"
    let newA = a
    for (let i = 0; i < a.length; i++) {
        newA = newA.split("")
        newA.splice(0, 1)
        newA = newA.join("") + a[i]
        if (b == newA) {
            return true
        }
    }
}

console.log(checkByRotating())