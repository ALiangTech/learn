function quickSort(data = []) {
    if(data.length < 2) {
        return data
    }
    const dataLen = data.length - 2;
    const baseIndex = Math.floor(dataLen / 2)
    const base = data.splice(baseIndex, 1)[0]
    let left = [];
    let right = [];
    for (let i = 0; i <= dataLen; i++) {
        if (data[i] >= base) {
            right.push(data[i])
        } else {
            left.push(data[i])
        }
    }
    return quickSort(left).concat([base, ...quickSort(right)])
}
// 分治法 + 递归
const test = [2, 33, 45, 62, 93, 31, 55, 6,1]
const datatest = quickSort(test)
console.log(datatest, "Tes")