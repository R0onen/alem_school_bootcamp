package queue

type Queue struct {
	Value []interface{}
}

func NewQueue() *Queue {
	return &Queue{
		Value: make([]interface{}, 0),
	}
}

func (q *Queue) Push(v interface{}) {
	q.Value = append(q.Value, v)
}

func (q *Queue) Pop() interface{} {
	if len(q.Value) == 0 {
		return nil
	}
	firstItem := q.Value[0]
	q.Value = q.Value[1:]
	return firstItem
}
