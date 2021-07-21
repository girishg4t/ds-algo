function nextLargerElement() {
    let gret = [1, 3, 2, 4]
    //let gret = [6 ,8, 0, 1, 3]
    let newArr = [-1];
    let maxGret = gret.pop()
    let prevElem = maxGret
    while (gret.length > 0) {
        let newElem = gret.pop();
        if (prevElem > newElem) {
            newArr.push(prevElem)
        } else {          
            if(maxGret > newElem){
                newArr.push(maxGret)
                maxGret = newElem         
            }else{
                newArr.push(-1)
            }          
        }

        prevElem = newElem
    }
    console.log(newArr.reverse())
}

nextLargerElement()