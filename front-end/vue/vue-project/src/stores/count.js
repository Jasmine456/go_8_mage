import { ref } from "vue";



// 维护Count变量的状态
export function useCount(){
    // 定义一个全局变量
    const count = ref(0);
    return count
}
