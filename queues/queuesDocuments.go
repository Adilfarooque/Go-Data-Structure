package queues

const capcity = 100

type Queues struct{
	Data [capcity]int
	Size int
	Front int
	Back int
}

func (q *Queues)Add(value int)bool{
	if q.Size >= capcity{
		return false
	}
	q.Size++
	q.Data[q.Back]=value
	q.Back = (q.Back+1)%(capcity-1)
	return true
}