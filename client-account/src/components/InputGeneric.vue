<script setup>

import { onMounted, reactive } from 'vue'

defineProps({
    label: {
        type: String,
        required: true
    },
    name: {
        type: String,
        required: true
    },
    type: {
        type: String,
        required: true
    },
    placeholder: {
        type: String,
        required: false,
        default: ''
    },
    modelValue: {
        type: [String, Number],
        required: true,
        default: ''
    }
})

const state = reactive({
    //
})

onMounted(() => {
    //
})

</script>


<template>
    <div class="form-control">
        <label class="label">
            <span class="label-text">{{ label }}</span>
        </label>
        <input
            class="input input-bordered"
            :type="type"
            :placeholder="placeholder"
            :name="name"
            :value="modelValue"
            @input="$emit('update:modelValue', $event.target.value)"
        >
    </div>
</template>