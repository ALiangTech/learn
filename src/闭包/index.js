function fn1 () {
    const a = 1;
    const b = 2;
    const c = 3;
    function fn2() {
      
    }
    function fn3() {
      console.log(a);
      console.log(c);
    }
    function fn4() {
      console.log(b);
    }
    return fn2;
  }
  const clo = fn1();
  console.dir(clo);
  debugger
