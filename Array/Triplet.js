function Triplet() {
    let arr = [3, 7, 2, 24, 12, 1, 5, 6, 4, 9, 13, 10, 25]
    for (let i = 0; i < arr.length; i++) {
        arr[i] = arr[i] * arr[i]
    }
    arr = arr.sort(function(a, b){return a - b});
    console.log(arr)
    // let double = [9, 4, 16, 36, 25]
    //let arr = [9, 25, 64]
    let n = arr.length - 1
    for (let i = n; i > -1; i--) {
        let diff = arr[i] - arr[i - 1]
        let j = i - 2
        while (diff<=arr[j]) {
            if (diff == arr[j]) {
                console.log("c=>", Math.sqrt(arr[i]))
                console.log("a+b=>" + Math.sqrt(arr[i - 1]) + " " + Math.sqrt(arr[j]))
            }
            j--
        }
    }
}

Triplet()