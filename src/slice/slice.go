package slice

type Ops interface {
  Contains(value interface{}) bool
  Add(value interface{})
  AddAll(values []interface{})
  Remove(value interface{}) bool
  IndexOf(value interface{}) (int, bool)
}

type Pair struct {
  First interface{}
  Second interface{}
}

func Map(values []interface{}, f func(interface{}) interface{}) []interface{} {
  data := make([]interface{}, len(values))
  for _, v := range values {
    data = append(data, f(v))
  }
  return data
}

func Filter(values []interface{}, f func(interface{}) bool) []interface{} {
  data := make([]interface{}, len(values))
  for _, v := range values {
    if f(v) {
      data = append(data, v)
    }
  }
  return data
}

func Zip(values1 []interface{}, values2 []interface{}) []Pair {
  zipped := make([]Pair, len(values1))
  for i := range values1 {
    zipped = append(zipped, Pair{values1[i], values2[2]})
  }
  return zipped
}
