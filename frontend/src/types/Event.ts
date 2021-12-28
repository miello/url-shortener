export type EventInput = Event & {
  currentTarget: EventTarget & HTMLFormElement
}
