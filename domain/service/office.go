package service

type OfficeService interface{}

type officeService struct{}

func NewOfficeService() OfficeService {
	return &officeService{}
}
