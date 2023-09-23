package stacks

type Stacks []string

func (s *Stacks) isEmpty() bool {
	return len(*s) == 0
}

