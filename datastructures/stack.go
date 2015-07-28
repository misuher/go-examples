package stack

type stack []int

func (s stack) isEmpty() bool { return len(s) == 0 }
func (s stack) Peek() int     { return s[len(s)-1] }
func (s *stack) Put(i int)    { (*s) = append((*s), i) }
func (s *stack) Pop() int {
	d := (*s)[len(*s)-1]
	(*s) = (*s)[:len(*s)-1]
	return d
}
