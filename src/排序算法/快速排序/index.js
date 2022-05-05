function quickSort(data = []) {
    const dataLen = data.length - 1;
    if(dataLen > 1) {
        const baseIndex = Math.floor(dataLen / 2)
        const base = data[baseIndex]
        console.log(base, "base")
        let left = [];
        let right = [];
        for(let i = 0; i < dataLen; i++) {
            if( data[i] < base) {
                left.push(data[i])
            }else {
                right.push(data[i])
            }
        }
        left.push(base)
        left = quickSort(left)
        right = quickSort(right)
        console.log(left, 'left')
        console.log(right, 'right')
    }
    return data
}

const test = [2,3,45,6,9,3,5,6]
const datatest = quickSort(test)
console.log(datatest, "Tes")