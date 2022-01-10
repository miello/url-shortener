import type { AxiosError } from 'axios'
import { writable } from 'svelte/store'

export interface IUpdateStores {
  status: 'error' | 'success'
  message: string | AxiosError<{ error: string }>
}

export interface IAlertStores {
  status: 'error' | 'success'
  message: string
}

export const AlertStores = writable<IAlertStores>({
  status: 'success',
  message: '',
})

export const OpenError = writable<boolean>(false)

let cnt = -1

export const UpdateAlert = (newVal: IUpdateStores, delay = 7000) => {
  if (cnt !== -1) return

  let msgErr = newVal.message
  if (typeof msgErr !== 'string') {
    msgErr = msgErr.response?.data.error || msgErr.message
  }

  AlertStores.set({
    status: newVal.status,
    message: msgErr,
  })
  OpenError.set(true)
  cnt = window.setTimeout(() => {
    OpenError.set(false)
    cnt = -1
  }, delay)
}

export const CancelAlert = () => {
  if (cnt !== -1) {
    window.clearTimeout(cnt)
    OpenError.set(false)
    cnt = -1
  }
}
