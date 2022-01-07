import type { AxiosError } from 'axios'
import { writable } from 'svelte/store'
import { apiClient } from '../utils/apiClient'
import { UpdateAlert } from './AlertStores'

export interface IRequestUser {
  handle: string
}

export interface IUserStores extends Partial<IRequestUser> {
  ready: boolean
}

export const UserStores = writable<IUserStores>({
  ready: false,
})

export const GetUser = async () => {
  try {
    const res = await apiClient.get<IRequestUser>('/auth/user')
    UserStores.set({
      ...res.data,
      ready: true,
    })
  } catch (err) {
    const axiosErr = err as AxiosError<{ error: string }>
    UserStores.set({
      ready: true,
    })
    if (axiosErr.response?.data?.error === 'Please Login First') {
      return
    }
    await Logout()
    UpdateAlert({
      status: 'error',
      message: axiosErr.response?.data?.error || axiosErr.message,
    })
  }
}

export const Logout = async () => {
  await apiClient.patch('/auth/logout')
  UserStores.set({
    ready: true,
  })
}
