package Users

type Repository interface {
	ReadUser(id int) (User, error)
}

type Service interface {
	ReadUser(id int) (User, error)
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) ReadUser(id int) (User, error) {
	u := User{}
	u, err := s.r.ReadUser(id)
	if err != nil {
		return User{}, err
	}
	return u, nil
}
