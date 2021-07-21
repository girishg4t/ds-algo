let arr = [1, 2, 3, 4, 5, 7, 8, 9, 10]

function missingNumber(newArr) {
    console.log(newArr)
    let m = Math.round(newArr.length / 2)
    if (m < 0 || m > n) {
        console.log(m)
        console.log("not found")
        return
    }
    if (elem == newArr[m]) {
        console.log("found at index", m);
        return;
    }
    if (elem < newArr[m]) {
        console.log("element less found", newArr[m])
        missingNumber(newArr.splice(0, m-1))
    }
    if (elem > newArr[m]) {
        console.log("element greater found", newArr[m])
        missingNumber(newArr.splice(m+1, n))
    }
   
}
arr = arr.sort(function(a,b){ return a < b})
let n = arr.length-1;

let elem = 16
missingNumber(arr)