package queue

func (q *Queue) Peek() interface{} {
	if len(q.Value) == 0 {
		return nil
	}
	return q.Value[0]
}
