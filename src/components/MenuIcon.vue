<template>
    <div class="MenuIcon">
      <div class="MenuBar"></div>
      <div class="MenuBar"></div>
      <div class="MenuBar"></div>
      <div class="MenuBar"></div>
    </div>
</template>

<style scoped>
.MenuIcon {
  border: 3px solid var(--editor-menuIconB);
  border-radius: 5px;
  width: 2rem;
  height: 2rem;
  display: flex;
  flex-wrap: wrap;
  text-align: center;
  align-items: center;
  cursor: pointer;
  padding: 0.2rem;

}

.MenuBar {
  background-color: var(--editor-menuIconB);
  margin: 0 auto;
  width: 90%;
  height: 0.15rem;
  border-radius: 3rem;
}

</style>