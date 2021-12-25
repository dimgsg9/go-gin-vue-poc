<script setup>

import { onMounted, reactive } from 'vue'
import { string } from 'yup'
import axios from 'axios'
import FormGeneric from '../components/FormGeneric.vue'

const state = reactive({
    username: '',
    password: '',
    loginFormSchema: {
        fields: [
              {
            label: 'Username',
            name: 'username',
            type: 'email',
            as: 'input',
            rules: string().email().required(),
        },
        {
            label: 'Password',
            name: 'password',
            type: 'password',
            as: 'input',
            rules: string().min(6).required(),
        }
        ],
        submitLabel: 'Login'
    },
    validationErrors: 
    {
        username: "wrong username",
        password: "wrong password",
        email: "wrong email",
    }
})

onMounted(() => {
    //
})

/**
 * 
 */
function signin(form) {
    console.log('calling signin() method with values:\n' + JSON.stringify(form.value, null, 2))
    // axios.post('/api/account/signin', form.value)
    //     .then()
    //     .catch()
    // await sleep(2000);
    sleep(2000).then(() => {
        form.actions.setErrors(state.validationErrors)
    });   
}

function sleep(ms) {
  return new Promise(resolve => setTimeout(resolve, ms));
}

 
</script>


<template>
    <div class="text-2xl font-bold text-center my-12">Login view</div>
    <FormGeneric @submit="signin" :schema="state.loginFormSchema"></FormGeneric>
</template>