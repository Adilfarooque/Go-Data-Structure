package stacks

type Stacks []string

func (s *Stacks) isEmpty() bool {
	return len(*s) == 0
}

func (s *Stacks) Push(str string){
	*s = append(*s, str)
}

