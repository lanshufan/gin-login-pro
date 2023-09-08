<template>
  <div class="home">
    <h2>{{ content }}</h2>
  </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue'

const content = ref("这里是首页")
</script>
