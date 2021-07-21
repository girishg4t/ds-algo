// N = 5, S = 12
// A[] = {1,2,3,7,5}

//index i = 1, sum = 1
//index i = 2, sum = 3
//i = 3, sum = 6
//i=4, sum = 15 , it is greater then 12 
//15 -1 = 14
//14-2 = 12

//**************
//N = 10, S = 15
//A[] = {1,2,3,4,5,6,7,8,9,10}
//1 = sum=1
//2= sum=3
//3= sum=6
//4= sum=10
//5= sum=15
function subarrayWithGivenSum() {
    let a = [1, 4, 20, 3, 10, 5]
    let sum = 33
    let pointer = 0
    let total = 0
    let endIndex = 0
    for (let i = 0; i < a.length; i++) {
        total = total + a[i]
        if (total > sum) {
            endIndex = i;
            break
        }
        if (total == sum) {
            console.log("1 " + (i + 1))
        }
    }

    for (let i = 0; i < endIndex; i++) {
        total = total - a[i]
        if (total == sum) {
            console.log((i + 2) + " " + (endIndex + 1))
            break
        }
    }
}
subarrayWithGivenSum()