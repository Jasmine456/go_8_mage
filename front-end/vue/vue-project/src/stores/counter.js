import { ref } from "vue";
import { defineStore } from "pinia";

// export const useCounterStore = defineStore('counter', {
//   state: () => ({ count: 0, name: 'Eduardo' }),
//   getters: {
//     doubleCount: (state) => state.count * 2,
//   },
//   actions: {
//     increment() {
//       this.count++
//     },
//   },
// })


//保存用户状态的一个Store
export const useUserStore = defineStore("user", () => {
  const username = ref("");
  const password = ref("")
  // const doubleCount = computed(() => count.value * 2);
  function login(user,pass) {
    username.value=user;
    password.value=pass;
  }


  return { username,password,login };
},{persist: {enabled:true,
  strategies: [
    // { storage: sessionStorage, paths: ['firstName', 'lastName'] }, // firstName 和 lastName字段用sessionStorage存储
    { storage: localStorage, paths: ["username","password"] } // accessToken字段用 localstorage存储
  ]}}
);
