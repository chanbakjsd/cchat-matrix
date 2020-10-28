package session

func (s Session) Listen() {
	for !s.ShouldStop {
		s.Sync()
	}
}
