function removeDuplicates() {
    let a = "bcabc"
    let newA = a.split("")
    let dicCount = []
    for (let i = 0; i < newA.length; i++) {
        if (!dicCount[newA[i]]) {
            dicCount[newA[i]] = 1
        } else {
            dicCount[newA[i]] =  dicCount[newA[i]] + 1
        }
    }
    console.log(dicCount)
    let result = ""
    Object.keys(dicCount).forEach((k)=>{
        result = result + k
    })           
    console.log(result)
}
removeDuplicates()