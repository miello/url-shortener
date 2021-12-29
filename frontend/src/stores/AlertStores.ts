import { writable } from "svelte/store"

export interface IErrorStores {
  status: "error" | "success"
  message: string
}

export const AlertStores = writable<IErrorStores>({
  status: "success",
  message: "",
})

export const OpenError = writable<boolean>(false)

let cnt = -1

export const UpdateAlert = (newVal: IErrorStores, delay: number = 7000) => {
  if (cnt !== -1) return
  AlertStores.set(newVal)
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
