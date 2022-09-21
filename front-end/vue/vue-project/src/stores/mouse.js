import { ref,onMounted,onUnmounted } from "vue";
// 需要导出的有状态函数，维护当前鼠标的坐标位置
// 返回一个 响应式对象
//  函数有一个命名风格：use+状态名称

export function useMouse(){
//    被组合式函数封装和管理的状态
    const x = ref(0)
    const y = ref(0)

//    这个状态是需要绑定鼠标事件
//    组合式函数可以随时更改其状态
    function update(event){
        x.value = event.pageX
        y.value = event.pageY
    }

//    一个组合式函数也可以挂靠在所属组件的生命周期上
//    来启动和卸载副作用
    onMounted(() => window.addEventListener('mousemove',update))
    onUnmounted(() => window.removeEventListener('mousemove',update))

    //{ x,y } = {x:x,y:y}
    return { x,y };
}