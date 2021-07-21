function reverseArray() {
    let arr = [1, 2, 3, 4, 5, 6, 7, 8]
    let k = 3
    let i = 0
    while (i < arr.length) {
        let start = i
        let end = i + (k - 1)
        if (end > arr.length - 1) {
            end = arr.length - 1
        }
        while (start < end) {
            swap(arr, start, end)
            start++
            end--
        }
        i += k
    }
    console.log(arr)
}
function swap(arr, i, j) {
    let temp = arr[i]
    arr[i] = arr[j]
    arr[j] = temp
}
reverseArray()
