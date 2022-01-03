<script lang="ts">
  import PersonIcon from '../components/icons/PersonIcon.svelte'
  import HistoryIcon from '../components/icons/HistoryIcon.svelte'
  import ExitIcon from '../components/icons/ExitIcon.svelte'
  import { Logout, UserStores } from '../stores/UserStores'
  import { scale } from 'svelte/transition'
  import { useNavigate } from 'svelte-navigator'
  import { createEventDispatcher, onMount } from 'svelte'

  export let className: string = ''

  let dispatchEvent = createEventDispatcher()
  const navigate = useNavigate()
  const ICON_LIST = [
    {
      Icon: PersonIcon,
      text: 'Profile',
      navigateTo: '/user/profile',
    },
    {
      Icon: HistoryIcon,
      text: 'History',
      navigateTo: '/user/history',
    },
    {
      Icon: ExitIcon,
      text: 'Logout',
      navigateTo: 'logout',
    },
  ]

  const handleClick = async (navigateTo: string) => {
    dispatchEvent('close')
    if (navigateTo === 'logout') {
      await Logout()
      return
    }

    navigate(navigateTo)
  }
</script>

<div
  transition:scale={{ duration: 200 }}
  class={`${className} bg-black w-[220px] flex flex-col rounded-2xl py-2 origin-top-right`}
>
  <h5 class="text-white font-display pl-4 font-bold">
    Signed in as {$UserStores.handle || ''}
  </h5>
  <hr class="my-2" />
  <div class="grid grid-cols-1 gap-2">
    {#each ICON_LIST as { Icon, text, navigateTo }}
      <button
        class="text-white flex ml-4 items-center"
        on:click={() => handleClick(navigateTo)}
      >
        <Icon />
        <span class="font-display ml-2">{text}</span>
      </button>
    {/each}
  </div>
</div>
