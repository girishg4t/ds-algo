function largestNumber() {
    let arr = [54, 546, 548, 60]
    let result = ""
    let lessThen10 = ""
    let greaterThen10 = ""
    let lesstrack = 0
    let greateTrack = 0
    for (let i = 0; i < arr.length; i++) {
        if (arr[i] < 10) {
            if (lesstrack < arr[i]) {
                lessThen10 = arr[i].toString() + lessThen10

            } else {
                lesstrack = arr[i]
                lessThen10 = lessThen10 + arr[i].toString()
            }
        } else {
            if (greateTrack < arr[i]) {
                greateTrack = arr[i]
                greaterThen10 = arr[i].toString() + greaterThen10
            } else {
                greaterThen10 = greaterThen10 + arr[i].toString()
            }
        }
    }
    console.log(lessThen10 + greaterThen10)
}

//largestNumber()


function largest1Number(num) {
    return num.sort(function(a, b) {
        return (b + '' + a ) - (a + '' + b);
    })
}

console.log(largest1Number([3, 30, 34, 5, 9]))
