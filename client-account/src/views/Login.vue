<script setup>

import { onMounted, reactive } from 'vue'
import axios from 'axios'
import FormGeneric from '../components/FormGeneric.vue'

const state = reactive({
    username: '',
    password: '',
    inputs: [
        {
            id: 1,
            type: 'text',
            name: 'username',
            placecholder: 'username',
            label: 'username',
            value: '',
        },
        {
            id: 2,
            type: 'password',
            name: 'password',
            placecholder: 'password',
            label: 'password',
            value: ''
        }
    ]
})

onMounted(() => {
    //
})

function signin() {
    console.log('calling signin() method')
}
 
</script>


<template>
    <div class="text-2xl font-bold text-center my-12">Login view</div>
    <FormGeneric :inputs="state.inputs" @submit="signin"></FormGeneric>
</template>