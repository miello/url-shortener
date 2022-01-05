import type { ExpiresTimeType } from './time'

export interface IShortenHistory {
  original: string
  shorten_id: string
  created_at: string
  expires: string
  raw_expires: ExpiresTimeType
}

export interface IHistoryRequest {
  all_pages: number
  count: number
  current: number
  data: IShortenHistory[]
}
