function mapforeach(array, cb) {
    const maparr = [];
    array.forEach((element, i) => {
        maparr.push(cb(element, i))
    });
    return maparr
}

var result = [{ id: 1, name: 'w' }, { id: 2, name: 'b' }].map((val, i) => {
    return { d: val.id, p: val.name }
})

const d = mapforeach([{ id: 1, name: 'w' }, { id: 2, name: 'b' }],(val, i) => {
    return { d: val.id, p: val.name }
});
console.log(result)
console.log(d)