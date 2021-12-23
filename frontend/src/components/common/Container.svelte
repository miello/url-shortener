<script lang="ts">
    export let _class: string = ""
</script>

<section class={`p-4 backdrop-blur-md drop-shadow-md bg-white rounded-xl bg-opacity-50 ${_class}`}>
    <slot />
</section>