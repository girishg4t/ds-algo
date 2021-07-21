function doCompar() {
    let arr = [0, 1, 1, 0, 1, 2, 1, 2, 0, 0, 0, 1]
    let m = 0
    let l = 0
    let h = arr.length - 1

    while (m <= h) {
        if (arr[m] == 0) {
            let temp = arr[l]
            arr[l] = arr[m]
            arr[m] = temp
            m++
            l++
            continue
        }
        if (arr[m] == 1) {
            m++
        } else {
            let temp = arr[m]
            arr[m] = arr[h]
            arr[h] = temp
            h--
        }       
    }
    console.log(arr)
}
doCompar()
function compare() {
    let arr = [0, 2, 1, 2, 0]
    let zerocount = 0;
    let onecount = 0
    let twocount = 0
    for (let i = 0; i < arr.length; i++) {
        if (arr[i] == 0) {
            zerocount++
        }
        if (arr[i] == 1) {
            onecount++
        }
        if (arr[i] == 2) {
            twocount++
        }
    }

}
function swap(arr, i, j) {
    let temp = arr[i]
    arr[i] = arr[j]
    arr[j] = temp
}