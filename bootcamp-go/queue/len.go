package queue

func (q *Queue) Len() int {
	return len(q.Value)
}
