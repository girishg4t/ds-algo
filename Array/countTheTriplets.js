//Input:
// N = 4
// arr[] = {1, 5, 3, 2}
// Output: 2
// Explanation: There are 2 triplets: 
// 1 + 2 = 3 and 3 +2 = 5 

//solution
// [1,2,3,5]
//
function countTheTriplets() {
    let arr = [1,2,2,3,4]
    arr = arr.sort(function(a, b){return a - b});
    let result = 0;
    let pointer = 0
    let j = 2;
    console.log(arr)
    for (let i = 2; i < arr.length; i++) {
        if ((arr[pointer] + arr[pointer + 1]) == arr[i]) {           
            result++
        }
        j = i+1;
        while (arr[j] >= (arr[pointer] + arr[pointer + 1])) {
            if ((arr[pointer] + arr[pointer + 1]) == arr[j]) {
                result++
            }  
            j++     
        }
        pointer++
    }
    console.log(result)
}
countTheTriplets()