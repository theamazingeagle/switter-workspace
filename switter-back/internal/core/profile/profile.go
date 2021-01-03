package profile

type Storage interface {
}

type ProfileDispatcher struct {
}

func NewProfileDispatcher(s *Storage) ProfileDispatcher {
	return ProfileDispatcher{}
}
