function callback() {
    console.log('Down')
}

// console.log('Start doWork');
// // 如何实现sleep，Go time.Sleep()
// // 5s后执行 doWork函数
// setTimeout(doWork, 5000)
// console.log('doWork Down')

//同步编程模式
//Start doWork
//Down
//doWork Down

//异步编程模式
//Start doWork
//doWork Down
//Down

//执行一个功能,通过网络IO获取数据
//feachData向服务端请求数据，服务端相应数据需要5s
function feachData(resolve,reject) {
//    可能会出现网络IO,可能会阻塞
//    需要把阻塞的逻辑抽象成回调
//    如何实现
//    调用后立即返回，不会阻塞
//    IO操作执行完成后，通过回调函数，返回执行后的结果
//    执行的结果：
//      执行成功(resolve(resp): 执行后的数据返回回去)，
//      执行失败(reject(error): 把执行过程的报错返回回去)
//    使用setTimeout(doWork,5000)
    setTimeout(()=>{
        console.log('Feach Data Down')
        //模拟网络IO请求
        try{
            //手动注入异常
            // throw new Error('Not Found')
            resolve({code:0,data:[{name:'博客名称'}]})
        }catch (error){
            reject(new Error('Feach Data Failed'+error))
        }

    },5000)

}

//如何实现回调函数
// feachData(
//     (resp)=>{
//         console.log(resp)
// },(error)=>{
//         console.log(error)
// })
//


// Promise版本
// var p1 = new Promise(feachData)
// p1.then((resp)=>{
// //    执行成功的回调处理
//     console.log(resp)
// }).catch((error)=>{
// //    处理执行失败的逻辑
//     console.log(error)
// }).finally(()=>{
// //    无论成功或失败 都需要的处理
//     console.log('finally')
// })
// console.log('is work down')


//async/await版本
async function feachDataAsync() {
// 回调风格的异步执行
    var p1=new Promise(feachData)
// 同步风格的异步执行
    try{
        var resp=await p1
        console.log(resp)
    }catch (error){
        console.log(error)
    }
}

//调用异步函数
feachDataAsync()