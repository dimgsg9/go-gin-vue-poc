<script setup>
import { nextTick } from 'vue'
import { Form, Field, ErrorMessage } from 'vee-validate'

const emits = defineEmits(['submit'])

defineProps({
    schema: {
        type: Object,
        required: true,
        default() {
            return {
                fields: [],
                submitLabel: 'submit',
                loading: false,
            }
      }
    }
})

/**
 * Handle submission of vee. Form
 * @param {Object} values Form values
 * @param {Object} actions Form actions
 */
function submitForm(values, actions) {
    const args = {
        value: values,
        actions: actions,
    }
    emits('submit', args)
}

/**
 * Handle failed validation of vee. Form
 * @param {Object} errors Form validation errors
 */
function onInvalidSubmit({ errors }) {
    const fieldName = Object.keys(errors)[0]
    const el = document.querySelector(`input[name="${fieldName}"]`)
    nextTick(() => {
        el?.scrollIntoView()
        el?.focus()
    })
}

</script>

<template>
    <div class="text-xl font-bold text-center my-12">Form</div>
    <Form @submit="submitForm" @invalid-submit="onInvalidSubmit" ref="FormRef" v-slot="{ isSubmitting, errors, meta }">
        <div class="form-control" v-for="field in schema.fields" :key="field.name">
            <label :for="field.name" class="label">
                 <span class="label-text">{{ field.label }}</span>
            </label>
            <Field v-bind="field" :id="field.name" :name="field.name" :type="field.type" class="input input-bordered" :class="{ 'input-error': errors[field.name] }"/>
            <label class="label">
                <ErrorMessage :name="field.name" class="label-text-alt" />
            </label>
        </div>

        <div class="form-control pt-5">
            <button class="btn btn-primary" :class="{'loading': isSubmitting || schema.loading}" :disabled="!meta.dirty">{{ schema.loading ? 'loading' : schema.submitLabel }}</button>
        </div>
    </Form>
</template>

/* TODO:

[done] Validation styles
Server non validation errors
Select
Checkbox
[done] Scroll to the input with error
[done] disabled button and load spinner

*/