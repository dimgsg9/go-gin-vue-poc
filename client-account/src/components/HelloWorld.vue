<script setup>
import { ref } from 'vue'

defineProps({
    msg: String
})

const count = ref(0)
</script>

<template>
    <div class="text-center">
        <h1>{{ msg }}</h1>
    </div>
</template>