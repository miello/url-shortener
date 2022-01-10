<script lang="ts">
  import type { AxiosError } from 'axios'
  import { useNavigate } from 'svelte-navigator'
  import Button from '../components/common/Button.svelte'
  import Container from '../components/common/Container.svelte'
  import Input from '../components/common/Input.svelte'
  import { UpdateAlert } from '../stores/AlertStores'
  import type { EventInput } from '../types/Event'
  import { apiClient } from '../utils/apiClient'

  export let handleName = ''
  export let username = ''
  export let password = ''
  export let confirmPassword = ''

  let loading = false

  const navigate = useNavigate()

  const onSubmit = async (e: EventInput) => {
    e.preventDefault()
    try {
      if (password !== confirmPassword) {
        UpdateAlert({
          status: 'error',
          message: 'Confirm Password is not matching',
        })
        return
      }
      loading = true
      await apiClient.post('/auth/register', {
        handle: handleName,
        username,
        password,
      })
      navigate('/login')
      UpdateAlert({
        status: 'success',
        message: 'Register Successfully',
      })
    } catch (err) {
      const axiosErr = err as AxiosError
      UpdateAlert({
        status: 'error',
        message: axiosErr,
      })
    }
    loading = false
  }
</script>

<div class="flex justify-center items-center w-full">
  <Container className="h-fit w-full max-w-[500px] mx-5">
    <form on:submit={onSubmit} class="flex flex-col ">
      <h3 class="font-display text-xl md:text-3xl text-center mb-5 font-bold">
        Register
      </h3>
      <div id="container" class="mb-5">
        <span class="font-display md:text-base text-sm"
          >Handle Name<span class="text-red-500 md:text-base text-sm">*</span
          ></span
        >
        <Input label="handle name" bind:value={handleName} required={true} />
        <span class="font-display md:text-base text-sm"
          >Username<span class="text-red-500 md:text-base text-sm">*</span
          ></span
        >
        <Input label="username" bind:value={username} required={true} />
        <span class="font-display md:text-base text-sm"
          >Password<span class="text-red-500 md:text-base text-sm">*</span
          ></span
        >
        <Input
          type="password"
          label="password"
          bind:value={password}
          required={true}
        />
        <span class="font-display md:text-base text-sm"
          >Confirm Password<span class="text-red-500 md:text-base text-sm"
            >*</span
          ></span
        >
        <Input
          type="password"
          label="confirm password"
          bind:value={confirmPassword}
          required={true}
        />
      </div>
      <Button
        disabled={loading}
        type="submit"
        className="bg-lime-600 text-md font-display mb-5 text-white"
        >Register</Button
      >
      <hr class="h-[2px] bg-black mb-5" />
      <h6 class="font-display text-md text-center mb-5 font-normal">
        Already have account ?
      </h6>
      <Button
        className="bg-yellow-500 text-md font-display"
        on:click={() => navigate('/login')}>Login</Button
      >
    </form>
  </Container>
</div>

<style>
  #container {
    display: grid;
    grid-template-columns: auto auto;
    row-gap: 0.75rem;
    column-gap: 0.75rem;
    align-items: center;
    min-height: 0;
  }
</style>
