<script setup lang="ts">
import { reactive, onUnmounted, onMounted } from 'vue';

const Text = reactive<{ Loading: String, Timeout: any }>({
  Loading: "Compiling",
  Timeout: 0,
})

const LoadingText =  ()=> {
    let c = 0;
    return () => {
        if(c<3){
            Text.Loading = Text.Loading+".";
        }else{
            c=0;
            Text.Loading = "Compiling";
        };
        c++;
    }
}

onMounted(() => {
    Text.Timeout = setInterval(LoadingText(), 500);
}),

onUnmounted(() => {
    clearInterval(Text.Timeout);
})
</script>
<template>
    <div class="loading-ring">
        <div></div>
        <div></div>
        <div></div>
        <div></div>
    </div>
    <div>{{Text.Loading }}</div>
</template>

<style scoped>
.loading-ring {
    position: relative;
    width: 5.5rem;
    height: 5.5rem;
}

.loading-ring div {
    position: absolute;
    width: 4rem;
    height: 4rem;
    margin: 0.5rem;
    border: 0.5rem solid var(--editor-LoadingIconB);
    border-radius: 50%;
    animation: loading-ring 1.2s cubic-bezier(0.5, 0, 0.5, 1) infinite;
    border-color: var(--editor-LoadingIconB) transparent transparent transparent;
}

.loading-ring div:nth-child(1) {
    animation-delay: -0.45s;
}

.loading-ring div:nth-child(2) {
    animation-delay: -0.3s;
}

.loading-ring div:nth-child(3) {
    animation-delay: -0.15s;
}

@keyframes loading-ring {
    0% {
        transform: rotate(0deg);
    }

    100% {
        transform: rotate(360deg);
    }
}</style>