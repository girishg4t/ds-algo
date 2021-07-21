function spiralOrder(matrix) {
    let res = [];
    while (matrix.length) {
        res.push(...matrix.shift())
        for(const m of matrix){
            let val = m.pop()
            if(val){
                res.push(val)
            }
            m.reverse()
        }
        matrix.reverse()
    }
    console.log(res)
};

spiralOrder([[1, 2, 3, 4],
[5, 6, 7, 8],
[9, 10, 11, 12]])