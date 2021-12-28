<script lang="ts">

export let value = ""
export let label = ""
export let className = ""
export let placeHolder = ""
export let required = false
</script>

<input bind:value label={label} placeholder={placeHolder} class={`px-2.5 py-1 flex-1 rounded-xl ${className}`} required={required}/>