<script setup lang="ts" >

</script>

<template>
    <div class="Container">
        <div class="icon-wrap">
            <svg class="checkmark" viewBox="0 0 52 52">
                <circle class="checkmark__circle" cx="26" cy="26" r="25" fill="none" />
                <path class="checkmark__check" fill="none" d="M14.1 27.2l7.1 7.2 16.7-16.8" />
            </svg>
            <div >
                Success
            </div>
        </div>
    </div>
</template>

<style scoped>
.Container {
    width: 100%;
    height: 100vh;
    display: flex;
    align-items: center;
    justify-content: center;
    position: fixed;
    top: 0;
    left: 0;
    z-index: 100;
    backdrop-filter: blur(0.5rem);
}
.icon-wrap{
    padding: 7rem;
    background-color: white;
    border-radius: 2rem;
    text-align: center;
    font-size: 2rem;
}
.checkmark {
    width: 8rem;
    height: 8rem;
  border-radius: 50%;
  display: block;
  stroke-width: 2;
  stroke: #fff;
  stroke-miterlimit: 10;
  box-shadow: inset 0px 0px 0px rgb(76, 204, 95);
  animation: fill .4s ease-in-out .4s forwards, scale .3s ease-in-out .9s both;
  z-index: 999;
}
.checkmark__circle {
  stroke-dasharray: 166;
  stroke-dashoffset: 166;
  stroke-width: 2;
  stroke-miterlimit: 10;
  stroke: rgb(76, 204, 95);
  fill: none;
  animation: stroke .6s cubic-bezier(0.650, 0.000, 0.450, 1.000) forwards;
}
.checkmark__check {
  transform-origin: 50% 50%;
  stroke-dasharray: 48;
  stroke-dashoffset: 48;
  animation: stroke .3s cubic-bezier(0.650, 0.000, 0.450, 1.000) .8s forwards;
}
@keyframes scale{
  0%, 100% {
    transform: none;
  }
  50% {
    transform: scale3d(1.1, 1.1, 1);
  }
}

@keyframes fill{
  100% {
    box-shadow: inset 0px 0px 0px 8rem rgb(76, 204, 95);
  }
}
@keyframes stroke {
  100% {
    stroke-dashoffset: 0;
  }
}
</style>