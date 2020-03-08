package slice

import (
  "fmt"
  "strings"
  "sync"
)

type ConcurrentSlice struct {
  Data []interface{}
  mux  sync.Mutex
}

func (cs *ConcurrentSlice) Contains(value interface{}) bool {
  cs.mux.Lock()
  defer cs.mux.Unlock()
  for _, d := range cs.Data {
    if d == value {
      return true
    }
  }
  return false
}

func (cs *ConcurrentSlice) Add(value interface{}) {
  cs.mux.Lock()
  cs.Data = append(cs.Data, value)
  cs.mux.Unlock()
}

func (cs *ConcurrentSlice) AddAll(values []interface{}) {
  cs.mux.Lock()
  cs.Data = append(cs.Data, values)
  cs.mux.Unlock()
}

func (cs *ConcurrentSlice) Remove(value interface{}) bool {
  cs.mux.Lock()
  defer cs.mux.Unlock()
  idx, isPresent := cs.IndexOf(value)
  if !isPresent {
    return false
  }
  updated := make([]interface{}, 0, len(cs.Data) - 1)
  for i, d := range cs.Data {
    if i != idx {
      updated = append(updated, d)
    }
  }

  cs.Data = updated

  return false
}

func (cs *ConcurrentSlice) IndexOf(value interface{}) (int, bool) {
  for i, d := range cs.Data {
    if d == value {
      return i, true
    }
  }
  return -1, false
}

func (cs ConcurrentSlice) String() string {
  data := make([]string, len(cs.Data))
  for i := range data {
    data[i] = fmt.Sprintf("%v", cs.Data[i])
  }
  return fmt.Sprintf("[ %v ]", strings.Join(data, ", "))
}
