// 原型继承
// 问题 共享修改 会出现问题
// 无法向父级传递参数

function F() {
    console.log("F")
    this.colors = ['red']
}

F.prototype.name = []
F.prototype.getColors = () => {
    console.log(this.colors)
}



function S() {}

S.prototype = new F()

const si = new S()
const si2 = new S()

si.colors.push('yellow')
si.colors.push('black')

console.log(si.colors)
console.log(si2.colors)

// 借用构造函数
// 重复调用父级 
// 子类也不能访问父类原型上定义的方法

function J(...rest) {
    F.apply(this,rest) // 这样子 无法访问到F的prototype属性
}

const j1 = new J()
const j2 = new J();

j1.colors.push("j1")

console.log(j1.colors)
console.log(j2.colors)
console.log([j1])

// 组合继承
// 原型继承方法
// 构造函数继承属性
// 效率问题 父级 会被调用俩次 

function FN(name) {
    this.name = name
    this.colors = ['red']
}

FN.prototype.getName = function() {
    console.log(this.name)
}

function SN(...rest) {
    FN.apply(this, rest) // 第一次
}
SN.prototype = new FN() // 第二次

const sn1 = new SN("one")
const sn2 = new SN("two")

sn1.colors.push("ye")

console.log(sn1.colors, 'sn1')
console.log(sn2.colors, 'sn2')
sn1.getName()
sn2.getName()


// 原型式继承
// 借用一个中间函数
// Ojbect.create()
// 主要用途方便快捷的扩展一个对象
function cobject (o) {
    function T () {}
    T.prototype = o;
    return new T()
}

// 寄生式继承

function jcobject (o) {
    const F =  cobject(o)
    F.name = () =>{}
    return F
}

//寄生式组合继承
function inheritPrototype(subType, superType) {
    let prototype = cobject(superType.prototype); // 创建对象
    prototype.constructor = subType; // 增强对象
    subType.prototype = prototype; // 赋值对象
   } 

inheritPrototype(F,S)

const is = new S()
const iss = new S()
is.colors.push("xx")
iss.colors.push("kks")
console.log(is.colors)
console.log(iss.colors)