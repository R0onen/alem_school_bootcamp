package stack

type Stack struct {
	Value []interface{}
}

func (s *Stack) Peek() interface{} {
	if len(s.Value) == 0 {
		return nil
	}
	return s.Value[len(s.Value)-1]
}
