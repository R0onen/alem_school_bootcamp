package stack

func NewStack() *Stack {
	return &Stack{
		Value: make([]interface{}, 0),
	}
}

func (s *Stack) Push(v interface{}) {
	s.Value = append(s.Value, v)
}

func (s *Stack) Pop() interface{} {
	if len(s.Value) == 0 {
		return nil
	}
	topIndex := len(s.Value) - 1
	topItem := s.Value[topIndex]
	s.Value = s.Value[:topIndex]
	return topItem
}
