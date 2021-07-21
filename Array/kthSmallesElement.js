
function poivote(arr, start, end, index) {
    while (start < end) {
        if (arr[start] > arr[index] && arr[index] > arr[end]) {
            let temp = arr[start]
            arr[start] = arr[end]
            arr[end] = temp
            start++
            end--
        }
        if (arr[start] < arr[index]) {
            start++
        }
        if (arr[end] > arr[index]) {
            end--
        }
    }
    let temp = arr[end]
    arr[end] = arr[index]
    arr[index] = temp
    return end
}
let arr = [10, 3, 6, 23, 2, 4, 15, 9]
let po = arr[0]
let start = 1
let end = arr.length - 1
function smallestElement(arr, i, j, p) {
    let index = 0
    if (i < j) {
        index = poivote(arr, i, j, p)
    }
    if (i < index) {
        smallestElement(arr, 0, index - 2, index - 1)
    }
    if (index < j) {
        smallestElement(arr, index + 2, arr.length - 1, index + 1)
    }
}
console.log(smallestElement(arr, start, end, 0))