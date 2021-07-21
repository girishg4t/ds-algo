function minimumOperation(n) {
    let acNum = n
    let countOpera = 2
    if (isOdd(n)) {
        acNum = n - 1
        countOpera++
    }
    while (acNum != 2 || acNum != 3) {
        if (isOdd(acNum)) {
            if (acNum <= 3) {
                acNum = acNum - 1
                countOpera++
            } else {
                acNum = acNum / 3
                countOpera++
            }
        }
    }
    return countOpera;

}

function isOdd(num) {
    return num % 2
}

console.log(minimumOperation(15))