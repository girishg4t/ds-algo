function sortArray(nums) {
    if (nums.length <= 1) return nums
    let middle = Math.floor(nums.length / 2)
    const left = nums.slice(0, middle)
    const right = nums.slice(middle)
    return merge(sortArray(left), sortArray(right))
}

function merge(left, right) {
    let result = []
    while (left.length && right.length) {
        if (left[0] < right[0]) {
            result.push(left.shift())
        } else {
            result.push(right.shift())
        }
    }
    return [...result, ...left, ...right]
}

console.log(sortArray([3, 6, 1, 2, 7, 8, 1]))