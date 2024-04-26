package services

type Svc struct {
	Name string `json:"name"`
	Email string `json:"email" binding:"required"`
}

func (s *Svc) getDataService(n Svc) Svc{
	return n
}