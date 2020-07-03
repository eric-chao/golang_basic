package stucture

type Stack struct {
	N     int
	Count int
	Items []string
}

func NewStack(n int) *Stack {
	items := make([]string, 0, n)
	return &Stack{N: n, Count: 0, Items: items}
}

func (s *Stack) Push(item string) bool {
	if s.Count == s.N {
		return false
	}
	s.Items = append(s.Items, item)
	s.Count = s.Count + 1
	return true
}

func (s *Stack) Pop() string {
	if s.Count == 0 {
		return ""
	}
	s.Count = s.Count - 1
	return s.Items[s.Count]
}
