package services

import (
	"fmt"
	"tentativa/datamodels"
	"tentativa/repo"
)

type AddressService interface {
	List(m map[string]interface{}) (response datamodels.Response)
	Save(address datamodels.Address) (response datamodels.Response)
	GetByID(id string) (response datamodels.Response)
	GetByName(Usename string) (response datamodels.Response)
	DeleteByID(id string) (response datamodels.Response)
}

type addressService struct {
	repo repo.AddressesRepository
}

var addressRepo = repo.NewAddressesRepository()

func NewAddressService() AddressService {
	return &addressService{
		repo: addressRepo,
	}
}

func (g *addressService) List(m map[string]interface{}) (response datamodels.Response) {
	addresses, _ := g.repo.List(nil)
	response.Code = 20000
	response.Msg = "success"
	response.Data = addresses
	return
}

func (g *addressService) Save(address datamodels.Address) (response datamodels.Response) {
	err := g.repo.Save(address)
	if err != nil {
		response.Code = 30001
		response.Msg = fmt.Sprintf("保存数据失败：%v", err)
	}
	response.Code = 20000
	response.Msg = "success"

	return
}

func (g *addressService) GetByID(id string) (response datamodels.Response) {
	address, err := g.repo.GetByID(id)
	if err != nil {
		response.Code = 30002
		response.Msg = fmt.Sprintf("查询数据失败：%v", err)
	}
	response.Code = 20000
	response.Msg = "success"
	response.Data = address
	return
}

func (g *addressService) GetByName(Usename string) (response datamodels.Response) {
	address, err := g.repo.GetByName(Usename)
	if err != nil {
		response.Code = 30002
		response.Msg = fmt.Sprintf("查询数据失败：%v", err)
	}
	response.Code = 20000
	response.Msg = "success"
	response.Data = address
	return
}

func (g *addressService) DeleteByID(id string) (response datamodels.Response) {
	err := g.repo.DeleteByID(id)
	if err != nil {
		response.Code = 30003
		response.Msg = fmt.Sprintf("删除数据失败：%v", err)
	}
	response.Code = 20000
	response.Msg = "success"
	return
}
