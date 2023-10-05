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
	q.Data[q.Back] = value
	q.Back = (q.Back + 1 ) % (capcity - 1)
	return true
}

func (q *Queues)Remove()(int,bool){
	value:=0
	if q.Size <= 0{
		return 0,false
	}
	q.Size--
	value = q.Data[q.Front]
	q.Front = (q.Front + 1) % (capcity - 1)
	return value,true
}