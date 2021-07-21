let arr = [0, 1, 3]

function missingNumber(newArr) {
    newArr = newArr.sort(function (a, b) { return a < b })
    let start = 1;
    let end = newArr.length - 2;
    let m = Math.round()
    while (start <= end) {
        let elm = newArr[start]
        if ((elm + 1 == newArr[start + 1]) && (elm - 1 == newArr[start - 1])) {
            start = start + 2
        } else {
            if (elm + 1 != newArr[start + 1]) {
                console.log("missing found", elm + 1)
            } else {
                console.log("missing found", elm - 1)
            }
            break
        }
        elm = newArr[end]
        if ((elm + 1 == newArr[end + 1]) && (elm - 1 == newArr[end - 1])) {
            end = end - 2
        } else {
            if (elm + 1 != newArr[end + 1]) {
                console.log("missing found", elm + 1)
            } else {
                console.log("missing found", elm - 1)
            }
            break
        }

    }

}

missingNumber(arr)

// function simple() {
//     let n = arr.length + 1
//     let total = n * (n + 1) / 2
//     let sum = arr.reduce((prev, cur) => prev + cur, 0)
//     console.log(total - sum)
// }

// simple()

var missingNumber1 = function (nums) {
    nums = nums.sort(function (a, b) { return a < b ? -1 : 0 })
    console.log(nums)
    let start = 1;
    let end = nums.length - 2;
    
    let found = ""
    while (start <= end) {
        let elm = nums[start]
        if ((elm + 1 == nums[start + 1]) && (elm - 1 == nums[start - 1])) {
            start = start + 2
        } else {
            found = elm - 1
            if (elm + 1 != nums[start + 1]) {
                found = elm + 1
            }
            return found
        }
        elm = nums[end]
        if ((elm + 1 == nums[end + 1]) && (elm - 1 == nums[end - 1])) {
            end = end - 2
        } else {
            found = elm - 1
            if (elm + 1 != nums[end + 1]) {
                found = elm + 1
            }
            return found
        }

    }

}

console.log(missingNumber1([3, 0, 1]))