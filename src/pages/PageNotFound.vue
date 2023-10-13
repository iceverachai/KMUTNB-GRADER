<script setup lang="ts">

</script>

<template>
    <div class="alert-wrapper">
        <div class="alert-message">
            ERROR : 404 page not found
        </div>
    </div>
</template>

<style scoped>
.alert-wrapper{
    display: flex;
    width: 100%;
    height: 100vh;
    justify-content: center;
    align-items: center;
}
.alert-message{
    padding: 1rem 2rem;
    background: var(--alerterrorpage-background);
    color: var(--alerterrorpage-fontcolor);
    font-size: 1.25rem;
    border: var(--alerterrorpage-bordercolor) solid 0.2rem;
}
</style>
