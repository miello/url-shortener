<script lang="ts">
import axios from "axios";
import { onMount,  } from "svelte";


    export let id: string
    onMount(() => {
        window.location.href = `http://localhost:8080/api/short/${id}`
    })
</script>

<main>
    {id}
</main>