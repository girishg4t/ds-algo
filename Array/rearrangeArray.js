let arr = [1,2,3,4,5,6]

function rearrange() {
    let start = 0
    let end = arr.length - 1
    let newArr = []
    while (start < end) {
        newArr.push(arr[end])
        end--
        newArr.push(arr[start])
        start++
    }
    if (start == end) {
        newArr.push(arr[end])
    }
    console.log(newArr)
}
rearrange()