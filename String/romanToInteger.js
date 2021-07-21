
function romanToInteger(s) {
    let a = {
        "I": 1,
        "V": 5,
        "X": 10,
        "L": 50,
        "C": 100,
        "D": 500,
        "M": 1000
    }
    let newString = s.split("")
    let nValue = 0
    let prev = ""
    for (let i = 0; i < newString.length; i++) {
        switch (prev + newString[i]) {
            case "IV":
                nValue = nValue + 4 - 1
                break;
            case "IX":
                nValue = nValue + 9 - 1
                break;
            case "XL":
                nValue = nValue + 40 - 10
                break;
            case "XC":
                nValue = nValue + 90 - 10
                break;
            case "CD":
                nValue = nValue + 400 - 100
                break;
            case "CM":
                nValue = nValue + 900 - 100
                break;
            default:
                nValue = nValue + (a[newString[i]])
                break
        }
        prev = newString[i]
    }
    return nValue

}

console.log(romanToInteger("III"))