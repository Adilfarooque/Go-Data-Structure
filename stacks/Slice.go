package stacks

type Stacks []string

func (s *Stacks)isEmpty()bool{
	return len(*s) == 0
}

func (s *Stacks)Push(str string){
	*s = append(*s, str)
}

func (s *Stacks)Pop()(string,bool){
	if s.isEmpty() {
		return "",false
	}else {
		index:=len(*s)-1
		value:=(*s)[index]
		*s = (*s)[:index]
		return value,true
	}
}