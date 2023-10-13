<script setup lang="ts">
import useTheme from "./module/usetheme"
import { onMounted } from "vue";
const themedata = useTheme();
onMounted(() => {
  const theme = localStorage.getItem('theme');
  if(theme === 'light'){
    themedata.setTheme('light');
  }else{
    themedata.setTheme('dark');
  }
})
</script>

<template>
  <main :data-theme="themedata.theme">
    <router-view></router-view>
  </main>
</template>

<style>
main * { 
  transition: background .2s ease-in-out;
}

main[data-theme=light]{

  --editor-conMainB:rgb(16, 16, 16);
  --editor-conMainF:#fff;
  --editor_conShadow:#000000ab;
  --editor-conMainB-H:#fff;
  --editor-conMainF-H:rgb(16, 16, 16);
  --editor_conShadow-H:#ffffffab;
  --edit-menuB:#b5b5b5;
  --editor-midBarB:#b5b5b5;
  --editor-menuIconB:#000;
  --editor-PbarB:#b5b5b5;
  --editor-outF:#000;
  --editor-LoadingIconB:#000;
  --editor-zoomB:#b5b5b5;
  --editor-zoomI-B:#0079d5;
  --editor-PbarSuccessF:#008820;
  --editor-PbarF-H:#0084ff;
  --editor-stdinF:#000;
  --editor-zoomIcon-C:#000;
  
  --ico-Edited:#fff;
  --ad-UserHB:#656565;
  --ad-menuconB:#656565;
  --ad-menuButB:#c5c5c5;
  --ad-tableHeadB:#656565;
  --ad-tableRowOddB:#b5b5b5;
  --ad-tableRowEvenB:#a4a4a4;
  --ad-tableRow-H:#fff;
  --ad-addGroupButB:#ffa500;
  --ad-asm-link:#000;
  --authB:#818181;
  --ad-modalIconWrapB:lightgray;

  --nav-B:#656565;
  --nav-DropDownB:#706f6f;
  --nav-DropDownB-H:#a4a4a4;
  --nav-DropDownF-H:#fff;
  --nav-F:#000;
  --nav-signin:#fff;
  --nav-btn-H:#fff;

  --footer-B:#b5b5b5;
}

main[data-theme=dark]{
  --alerterrorpage-fontcolor: #fff;
  --alerterrorpage-background: #1d2125;
  --alerterrorpage-bordercolor: #fff;
  
  --edit-menuB:#1d2125;
  --editor-F:#fff;
  --editor-conB:#1d2125;
  --editor-zoomB:#d3d3d3;
  --editor-zoomI-B:#04AA6D;
  --editor-conF: #fff;
  --editor-conMainB:rgb(16, 16, 16);
  --editor-conMainF:#fff;
  --editor_conShadow:#000000ab;
  --editor-conMainB-H:#fff;
  --editor-conMainF-H:rgb(16, 16, 16);
  --editor_conShadow-H:#ffffffab;
  --editor-conF-H:#000;
  --editor-conB-H:#000;
  --editor-midBarB:#1d2125;
  --editor-ResF:#fff;
  --editor-ResB:#282c34;
  --editor-errResF:rgb(255, 76, 76);
  --editor-stdinF:#fff;
  --editor-stdinB:#1d2125;
  --editor-inputF: #fff;
  --editor-inputF-H: #000;
  --editor-PbarB:#1d2125;
  --editor-PbarF:#fff;
  --editor-PbarSuccessF:#30d556;
  --editor-PbarF-H:#ffa500;
  --editor-menuIconB:#fff;
  --editor-LoadingIconB:#fff;
  --editor-outF:#fff;
  --editor-zoomIcon-C:#fff;

  --ad-asm-link:#fff;
  --ad-menuconB:#18191e;
  --ad-menuButB:#1d2125;
  --ad-tableF:#fff;
  --ad-tableHeadB:#18191e;
  --ad-tableRowOddB:#222227;
  --ad-tableRow-H:#18191e;
  --ad-conContainB:#fff;
  --ad-UserHB:#18191e;
  --ad-addGroupButB:#ff8800;
  --ad-DropDownConB:#fff;
  --ad-DropDownmB-H:rgb(202, 202, 202);
  --ad-DropDownmF-H:#fff;
  --ad-modalIconWrapB:#fff;
  --ad-modalOutline:#ffa500;
  --ad-modalFormButB:#ffa500;
  --ad-modalFormButF:#fff;
  --ad-InputFocusB:#0066ff;
  --ad-tableGroupF:#fff;
  --ad-tableGroupDelB:rgb(244, 85, 85);

  --authB:#1d2125;

  --ico-B:#000;
  --ico-Edited:#585858;
  --ico-Display-H:#787878;
  --ico-Display:#fff;

  --nav-B:#18191e;
  --nav-accButB:#fff;
  --nav-accButB-H:#ffd700;
  --nav-F:#fff;
  --nav-DropDownB:#fff;
  --nav-DropDownB-H:rgb(202, 202, 202);
  --nav-DropDownF-H:#fff;
  --nav-schemeB:#fff;
  --nav-signin:#fff;

  --footer-B:#1d2125;
}
</style>
