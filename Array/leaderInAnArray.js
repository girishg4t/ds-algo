function leaders() {
    let arr = [1,2,3,4,0]
    let max = arr[arr.length - 1]
    for (let i = arr.length - 1; i > 0; i--) {
        if (max <= arr[i]) {
            max = arr[i]
            console.log(max)
        }
    }


}

leaders()