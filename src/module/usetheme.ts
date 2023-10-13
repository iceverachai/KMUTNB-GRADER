import { defineStore } from "pinia"

type theme = "dark" | "light"

export const useTheme = defineStore('theme', { 
  state: () => ({ theme: "dark" as theme}),
  actions: {
    setTheme(theme : theme){
        if(theme === "dark"){
            document.body.classList.remove('light');
            document.body.classList.add('dark');
        }else if(theme === "light"){
            document.body.classList.remove('dark');
            document.body.classList.add('light');
        }
        localStorage.setItem('theme',theme);
        this.theme = theme
    }
  }
})

export default useTheme;