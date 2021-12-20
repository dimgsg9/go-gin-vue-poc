<script setup>

import { onMounted, reactive } from 'vue'
import InputGeneric from './InputGeneric.vue'

const props = defineProps({
    inputs: {
        type: Array,
        required: true,
    }
})

const state = reactive({
    fields: props.inputs
})

onMounted(() => {
    //
})

// v-model="state.username"
 
</script>


<template>
    <div class="text-xl font-bold text-center my-12">Form</div>
    <form action="">
        <InputGeneric v-for="(item, index) in inputs" :key="item.id" :name="item.name" :type="item.type" :label="item.label" :placeholder="item.placeholder" v-model="state.fields[index].value"></InputGeneric>
        <button type="button">submit</button>
    </form>
</template>