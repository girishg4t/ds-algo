
function nextGreateElement() {
    let arr = [5, 7, 2, 12]
    let stack = [];
    let pointer = 1;
    stack.push(arr[0]);
    for (let i = 1; i < arr.length; i++) {
        let elem = stack.pop();
        while (elem != null) {
            if (arr[i] > elem) {
                console.log(elem + "=>" + arr[i])
            } else {
                stack.push(arr[i])
            }
            elem = null
        }
    }
}
nextGreateElement()
//stack = 5 , 7 => 5
//stacl = 7, 
//stack 7, 2