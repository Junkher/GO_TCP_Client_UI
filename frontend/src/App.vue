<script setup>
import Client from './components/Client.vue'
</script>

<template>
  <!-- <img id="logo" alt="Wails logo" src="./assets/images/logo-universal.png"/> -->
  <Client/>
</template>

<style>

:root {
  background: #202124;
  border-radius: 6px;
  /* --el-color-primary: #b2cd4f !important; */
}



#logo {
  display: block;
  width: 50%;
  height: 50%;
  margin: auto;
  padding: 10% 0 0;
  background-position: center;
  background-repeat: no-repeat;
  background-size: 100% 100%;
  background-origin: content-box;
}
</style>
