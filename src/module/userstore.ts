import { defineStore } from "pinia"

type info = {
    exp: number
    username: string
    firstname:string
    lastname:string
    fullname:string
    isadmin: boolean
    groups: string[]
    work: {}
}


export const useUserStore = defineStore('userdata', {
  state: () => ({ info: null as info | null }),
  actions: {
    setInfo(info : info | null){
      this.info = info
    }
  }
})

export default useUserStore;