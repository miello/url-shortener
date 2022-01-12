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

export const AlertStores = writable<Record<string, IAlertStores>>({})

let key = 1
let timeOut: Record<string, number> = {}

export const UpdateAlert = (newVal: IUpdateStores, delay = 7000) => {
  let msgErr = newVal.message
  if (typeof msgErr !== 'string') {
    msgErr = msgErr.response?.data.error || msgErr.message
  }

  let tmp = key++
  AlertStores.update((prev) => ({
    ...prev,
    [tmp]: {
      status: newVal.status,
      message: msgErr,
    },
  }))

  timeOut[tmp] = window.setTimeout(() => {
    CancelAlert(tmp)
  }, delay)
}

export const CancelAlert = (key: number) => {
  if (timeOut[key]) {
    const keyStr = key.toString()
    window.clearTimeout(timeOut[key])
    AlertStores.update((data) => {
      return Object.keys(data).reduce((prev, val) => {
        if (keyStr !== val) {
          return {
            ...prev,
            [val]: data[val],
          }
        }
        return prev
      }, {})
    })
    timeOut = Object.keys(timeOut).reduce((prev, cur) => {
      if (keyStr !== cur) {
        return {
          ...prev,
          [cur]: timeOut[cur],
        }
      }
      return prev
    }, {} as Record<number, number>)
  }
}
