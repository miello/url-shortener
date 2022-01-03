<script lang="ts">

import Button from "../components/common/Button.svelte";
import Container from "../components/common/Container.svelte";

const TABLE_HEADER = ['Date Created', 'Shorten ID', 'Original', 'Expires']
</script>

<div class="w-full h-full flex flex-col justify-center items-center">
  <Container className="flex flex-col items-center">
    <h4 class="font-bold font-display text-2xl mb-4">History</h4>
    <table class="m-0 p-0 border-4 border-solid border-black">
      <tr class="border-2 border-solid border-black font-semibold">
        <th class="border-2 border-solid border-black p-2">Date Created</th>
        <th class="border-2 border-solid border-black p-2">Shorten ID</th>
        <th class="border-2 border-solid border-black p-2">Original</th>
        <th class="border-2 border-solid border-black p-2">Expires</th>
      </tr>
      {#each [1, 2, 3, 4, 5, 6, 7] as x}
      <tr>
        <td class="border-2 border-solid border-black px-2">20/11/2564 16:30</td>
        <td class="border-2 border-solid border-black px-2">DEDOS</td>
        <td id="original-url" class="border-2 border-solid border-black px-2 overflow-auto max-w-[150px] ">www.google.com</td>
        <td class="border-2 border-solid border-black px-2">20/11/2564 16:30</td>
        <td class="border-2 border-solid border-black px-2 py-2">
          <Button className="bg-[#BDFF00] py-1 font-semibold">Edit</Button>
          <Button className="bg-red-500 py-1 font-semibold">Delete</Button>
        </td>
      </tr>
      {/each}
    </table>
  </Container>
</div>

<style>
  #original-url::-webkit-scrollbar {
    height: 2px;
  }

  #original-url::-webkit-scrollbar-thumb {
    background: #808080;
  }

  #original-url::-webkit-scrollbar {
    height: 2px;
  }

  #original-url::-webkit-scrollbar {
    height: 2px;
  }
</style>