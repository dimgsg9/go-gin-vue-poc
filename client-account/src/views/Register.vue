<script setup>

import { onMounted, reactive } from 'vue'
import { Field, ErrorMessage, Form, defineRule } from 'vee-validate'
import { required, email } from '@vee-validate/rules'
import axios from 'axios'

defineRule('required', required)
defineRule('email', email)

const state = reactive({
    //
})

onMounted(() => {
    //
})

function onSubmit(values) {
    console.log(JSON.stringify(values, null, 2))
}

</script>


<template>
    <div class="text-2xl font-bold text-center my-12">Register view</div>
    <Form v-slot="{ meta, isSubmitting }" @submit="onSubmit">
        <div class="form-control">
            <label for="email" class="label">
                <span class="label-text">Email</span>
            </label>
            <Field name="email" type="email" rules="required|email" class="input input-bordered" />
            <label class="label">
                <ErrorMessage name="email" class="label-text-alt" />
            </label>
        </div>
        <div class="form-control">
            <label for="password" class="label">
                <span class="label-text">Password</span>
            </label>
            <Field name="password" type="password" rules="required" class="input input-bordered" />
            <label class="label">
                <ErrorMessage name="password" class="label-text-alt" />
            </label>
        </div>
        <div class="form-control">
            <label for="repeat-password" class="label">
                <span class="label-text">Repeat password</span>
            </label>
            <Field name="repeat-password" type="password" rules="required" class="input input-bordered" />
            <label class="label">
                <ErrorMessage name="repeat-password" class="label-text-alt" />
            </label>
        </div>

        <div class="form-control">
            <button class="btn btn-primary">register</button>
        </div>
    </Form>
</template>