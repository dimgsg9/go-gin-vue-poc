<script setup>

import HelloWorld from './components/HelloWorld.vue'
// import Login from './views/Login.vue'
// import Register from './views/Register.vue'
import { onMounted, reactive } from 'vue'
import axios from 'axios'

const state = reactive({
    errorCode: '',
    errorMessage: ''
})

onMounted(() => {
    axios.get('/api/account/me').then().catch(
        error => {
            state.errorCode = error.response.status
            state.errorMessage = error.response.data.error.message
        }
    )
})

</script>


<template>
    <div class="text-4xl font-bold text-center my-12">Account vue app</div>
    <div class="text-xl text-center" v-if="state.errorCode">Error: {{ state.errorCode }}</div>
    <HelloWorld :msg="state.errorMessage" />
    <div class="flex justify-around my-4">
        <router-link :to="{ name: 'login' }">Login</router-link>
        <router-link to="/">Home</router-link>
        <router-link :to="{ name: 'register' }">Register</router-link>
    </div>
    <div class="container mx-auto px-4">
        <router-view></router-view>
    </div>
</template>