// import {MyAPP as profile} from './profile.mjs'
import profile from './profile.mjs'
// import {age, firstName, lastName} from './profile.mjs'

//默认导出
//export default
//import {default as profile}

//import {default as profile} from './profile.mjs'
//简写
//import profile from './profile.mjs'

// var obj = {firstName: firstName, lastName: lastName, age: age}
// obj = {firstName, lastName, age}

//对象的结构赋值
//{firstName,lastName,age} = {firstName:firstName,lastName:lastName,age:age}

console.log(profile.age,profile.firstName,profile.lastName)