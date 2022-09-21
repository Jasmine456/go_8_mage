import {reactive} from "vue";
// //定义一个全局变量
// const username=ref('')
// const password=ref('')
//
// //维护Login变量的状态
//
// export function useLogin(){
//     return {username,password};
// }

//维护Login变量的状态
export var login = reactive({username:"",password:""})