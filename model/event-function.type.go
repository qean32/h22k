package model

type EventFunction func(e Event)
type FnRerutnEvent func(arr []string) (e Event, _error bool)
