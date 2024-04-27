package services

type Svc struct {
	Name string `json:"name"`
	Email string `json:"email" binding:"required"`
}



func (s *Svc) GetDataService(data Svc) Svc{
	return data
}

func (s *Svc) PostDataService(data Svc) Svc{
	return data
}

func (s *Svc) DeleteDataService(data Svc) Svc{
	data = Svc{}
	return data

}