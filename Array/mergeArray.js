let arr1 = [1,2,3,0,0,0]
let arr2 =  [2,5,6]

// [0, 3, 5, 7]
// [1, 2, 6, 8, 9]

// [0, 1, 5, 7]
// [2, 3, 6, 8, 9]

// [0, 1, 2, 7]
// [3, 5, 6, 8, 9]

// [0, 1, 2, 3]
// [5, 6, 7, 8, 9]

function merge() {
    for (let i = 0; i < arr1.length; i++) {
        for (let j = 0; j < arr2.length; j++) {
            if (arr1[i] > arr2[j]) {
                swap(arr1, i, arr2, j)
                let pointer = j + 1
                while (true) {
                    if (pointer > arr2.length) {
                        break
                    }
                    if (arr2[j] > arr2[pointer]) {
                        swap(arr2, j, arr2, pointer)
                        j++
                    }
                    pointer++
                }
            }
        }
    }
    console.log(arr1)

    console.log(arr2)
}

function swap(f, i, s, j) {
    let temp = f[i]
    f[i] = s[j]
    s[j] = temp
}

merge();

