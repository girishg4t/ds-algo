function zigZagFassion() {
    let stack = [];
    let arr = [1, 4, 3, 2]
    let n = arr.length
    let i = 0
    let j = 1
    while (i < n) {
        if (arr[i] > arr[j]) {
            swap(arr, i, j)
        } else {
            k = j + 1
            while (arr[j] < arr[k]) {
                swap(arr, j, k)
            }
            i = i + 2
            j = i + 1
            console.log(arr)
        }
    }
    console.log(arr)
}

function swap(arr, i, j) {
    let temp = arr[i]
    arr[i] = arr[j]
    arr[j] = temp
}

zigZagFassion()