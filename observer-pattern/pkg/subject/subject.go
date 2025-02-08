package subject

type Observer struct {
	url string
}

type Subject struct {
	observers []Observer
}

func (s *Subject) Add(observer Observer) {
	s.observers = append(s.observers, observer)
}

func (s *Subject) Remove(observer Observer) {
	s.observers = remove(s.observers, observer)
}

func (s *Subject) LatestState() string {
	return s.state
}
