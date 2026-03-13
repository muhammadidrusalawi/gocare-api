package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhammadidrusalawi/gocare-api/internal/model"
	"github.com/muhammadidrusalawi/gocare-api/internal/repository"
	"github.com/muhammadidrusalawi/gocare-api/internal/request"
	"github.com/muhammadidrusalawi/gocare-api/provider/database"
)

func GetAllAddressesByOwner(userID string) []*model.Address {
	addressRepo := repository.NewAddressRepository(database.DB)

	addresses, err := addressRepo.FindAll(userID)
	if err != nil {
		return []*model.Address{}
	}
	return addresses
}

func GetAddressByIDByOwner(userID, id string) (*model.Address, error) {
	addressRepo := repository.NewAddressRepository(database.DB)

	address, err := addressRepo.FindByID(userID, id)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, "Address not found")
	}

	return address, nil
}

func CreateAddressByOwner(userID string, req request.CreateAddressRequest) (*model.Address, error) {
	addressRepo := repository.NewAddressRepository(database.DB)

	address := &model.Address{
		UserID:       userID,
		ReceiverName: req.ReceiverName,
		Phone:        req.Phone,
		Province:     req.Province,
		City:         req.City,
		District:     req.District,
		PostalCode:   req.PostalCode,
		FullAddress:  req.FullAddress,
		Label:        req.Label,
		IsDefault:    false,
	}

	existing, err := addressRepo.FindAll(userID)
	if err != nil {
		return nil, err
	}
	if len(existing) == 0 {
		address.IsDefault = true
	}

	if err := addressRepo.Create(userID, address); err != nil {
		return nil, err
	}

	return address, nil
}

func UpdateAddressByOwner(userID, id string, req request.UpdateAddressRequest) (*model.Address, error) {
	addressRepo := repository.NewAddressRepository(database.DB)

	address, err := addressRepo.FindByID(userID, id)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, "Address not found")
	}

	if req.ReceiverName != nil {
		address.ReceiverName = *req.ReceiverName
	}

	if req.Phone != nil {
		address.Phone = *req.Phone
	}

	if req.Province != nil {
		address.Province = *req.Province
	}

	if req.City != nil {
		address.City = *req.City
	}

	if req.District != nil {
		address.District = *req.District
	}

	if req.PostalCode != nil {
		address.PostalCode = *req.PostalCode
	}

	if req.FullAddress != nil {
		address.FullAddress = *req.FullAddress
	}

	if req.Label != nil {
		address.Label = *req.Label
	}

	if err := addressRepo.Update(userID, address); err != nil {
		return nil, err
	}

	return address, nil
}

func DeleteAddressByOwner(userID, id string) error {
	addressRepo := repository.NewAddressRepository(database.DB)

	_, err := addressRepo.FindByID(userID, id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Address not found")
	}

	if err := addressRepo.Delete(userID, id); err != nil {
		return err
	}

	return nil
}

func SetDefaultAddressByOwner(userID, id string) error {
	addressRepo := repository.NewAddressRepository(database.DB)

	_, err := addressRepo.FindByID(userID, id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Address not found")
	}

	if err := addressRepo.SetDefault(userID, id); err != nil {
		return err
	}

	return nil
}
