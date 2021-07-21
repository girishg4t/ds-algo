function AllElement() {
    let arr = [4, 2, 5, 7]
    let leftMax = []
    let rightMin = 100
    leftMax[0] = arr[0]
    for(let i = 1; i < arr.length-1; i++){
        leftMax[i] = Math.max(arr[i], leftMax[i-1])
    }
    for(let i = arr.length-1; i > 0; i--){
        if(leftMax[i-1] < arr[i] && rightMin > arr[i]){
                console.log(arr[i])
        }
        rightMin = Math.min(rightMin, arr[i])
    }
}

AllElement()