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
        submitLabel: 'Login',
        loading: false
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
    state.loginFormSchema.loading = true
    axios.post('/api/account/signin', form.value)
        .then(response => {
            console.log('inside .then()\n' + response)
        })
        .catch(error => {
            // check for failed server side field validation
            if (error.response.status == 400) {
                state.validationErrors.username = error.response.data.invalidArgs.email
            } else {
                //state.loginFormSchema.loading = false
                // show modal?
            }        
            sleep(2000).then(() => {
                form.actions.setErrors(state.validationErrors)
                // state.loginFormSchema.loading = false
            })
        })
        .finally(() => {
            // state.loginFormSchema.loading = false
        })
}

// TODO: remove temporary function sleep()
function sleep(ms) {
  return new Promise(resolve => setTimeout(resolve, ms));
}

 
</script>


<template>
    <div class="text-2xl font-bold text-center my-12">Login view</div>
    <FormGeneric @submit="signin" :schema="state.loginFormSchema"></FormGeneric>
</template>