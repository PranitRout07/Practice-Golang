package services


type DataService struct{}

type Svc struct {
	Name string `json:"name"`
	Email string `json:"email" binding:"required"`
}



func (s *DataService) GetDataService(data Svc) Svc{
	return data
}

func (s *DataService) PostDataService(data Svc) Svc{
	return data
}

func (s *DataService) DeleteDataService(data Svc) Svc{

	return data

}