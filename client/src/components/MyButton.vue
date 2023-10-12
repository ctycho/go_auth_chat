<template>
    <button class="p-3 mt-2 text-white bg-blue-600 rounded-md font-bold text-lg">{{ name }}</button>
</template>

<script>
export default {
    name: 'MyButton',
    props: {
        name: String
    },
    components: {
        
    }
}
</script>

<style>

</style>