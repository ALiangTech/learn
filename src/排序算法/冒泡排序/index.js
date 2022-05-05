function bubbleSort(data = []) {
    const dataLen = data.length - 1
    for(let i = 0 ; i < dataLen; i++) {
        for(let j = 0; j < dataLen - i ; j++) {
            if(data[j] > data[j+1]) { // 大于交换
                const temp = data[j+1]
                data[j+1] = data[j]
                data[j] = temp
            }
        }
    }
    console.log(data)
}

const data = [23,3,4,5,6,344,42,3,323]
bubbleSort(data)
const datas = [1,3,4,5,6,344,42,3,0,2]

bubbleSort(datas)