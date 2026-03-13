package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/muhammadidrusalawi/gocare-api/internal/helper"
	"github.com/muhammadidrusalawi/gocare-api/internal/request"
	"github.com/muhammadidrusalawi/gocare-api/internal/response"
	"github.com/muhammadidrusalawi/gocare-api/internal/service"
)

func GetAllAddressesByOwnerHandler(c *fiber.Ctx) error {
	claims := c.Locals("claims").(jwt.MapClaims)
	userID := claims["user_id"].(string)

	addresses := service.GetAllAddressesByOwner(userID)

	var result []response.AddressResponse

	for _, addr := range addresses {
		result = append(result, response.AddressResponse{
			ID:           addr.ID,
			ReceiverName: addr.ReceiverName,
			Phone:        addr.Phone,
			Province:     addr.Province,
			City:         addr.City,
			District:     addr.District,
			PostalCode:   addr.PostalCode,
			FullAddress:  addr.FullAddress,
			IsDefault:    addr.IsDefault,
			Label:        addr.Label,
		})
	}

	return c.Status(fiber.StatusOK).JSON(helper.ApiSuccess("Shipping addresses data successfully retrieved", result))
}

func GetAddressByIDByOwnerHandler(c *fiber.Ctx) error {
	claims := c.Locals("claims").(jwt.MapClaims)
	userID := claims["user_id"].(string)
	id := c.Params("id")

	address, err := service.GetAddressByIDByOwner(userID, id)

	if err != nil {
		return err
	}

	res := response.AddressResponse{
		ID:           address.ID,
		ReceiverName: address.ReceiverName,
		Phone:        address.Phone,
		Province:     address.Province,
		City:         address.City,
		District:     address.District,
		PostalCode:   address.PostalCode,
		FullAddress:  address.FullAddress,
		IsDefault:    address.IsDefault,
		Label:        address.Label,
		CreatedAt:    &address.CreatedAt,
		UpdatedAt:    &address.UpdatedAt,
	}

	return c.Status(fiber.StatusOK).JSON(helper.ApiSuccess("Shipping address detail successfully retrieved", res))
}

func CreateAddressByOwnerHandler(c *fiber.Ctx) error {
	claims := c.Locals("claims").(jwt.MapClaims)
	userID := claims["user_id"].(string)
	var req request.CreateAddressRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(helper.ApiError("Invalid JSON"))
	}

	if err := helper.ValidateStruct(req); err != "" {
		return c.Status(fiber.StatusBadRequest).
			JSON(helper.ApiError(err))
	}

	address, err := service.CreateAddressByOwner(userID, req)
	if err != nil {
		return err
	}

	res := response.AddressResponse{
		ID:           address.ID,
		ReceiverName: address.ReceiverName,
		Phone:        address.Phone,
		Province:     address.Province,
		City:         address.City,
		District:     address.District,
		PostalCode:   address.PostalCode,
		FullAddress:  address.FullAddress,
		IsDefault:    address.IsDefault,
		Label:        address.Label,
		CreatedAt:    &address.CreatedAt,
		UpdatedAt:    &address.UpdatedAt,
	}

	return c.Status(fiber.StatusCreated).JSON(helper.ApiSuccess("Shipping address created", res))
}

func UpdateAddressByOwner(c *fiber.Ctx) error {
	claims := c.Locals("claims").(jwt.MapClaims)
	userID := claims["user_id"].(string)
	id := c.Params("id")
	var req request.UpdateAddressRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(helper.ApiError("Invalid JSON"))
	}

	if err := helper.ValidateStruct(req); err != "" {
		return c.Status(fiber.StatusBadRequest).
			JSON(helper.ApiError(err))
	}

	address, err := service.UpdateAddressByOwner(userID, id, req)
	if err != nil {
		return err
	}

	res := response.AddressResponse{
		ID:           address.ID,
		ReceiverName: address.ReceiverName,
		Phone:        address.Phone,
		Province:     address.Province,
		City:         address.City,
		District:     address.District,
		PostalCode:   address.PostalCode,
		FullAddress:  address.FullAddress,
		IsDefault:    address.IsDefault,
		Label:        address.Label,
		CreatedAt:    &address.CreatedAt,
		UpdatedAt:    &address.UpdatedAt,
	}

	return c.Status(fiber.StatusOK).JSON(helper.ApiSuccess("Shipping address updated", res))
}

func DeleteAddressByOwnerHandler(c *fiber.Ctx) error {
	claims := c.Locals("claims").(jwt.MapClaims)
	userID := claims["user_id"].(string)
	id := c.Params("id")

	if err := service.DeleteAddressByOwner(userID, id); err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(helper.ApiSuccess("Shipping address deleted successfully", nil))
}

func SetDefaultAddressByOwnerHandler(c *fiber.Ctx) error {
	claims := c.Locals("claims").(jwt.MapClaims)
	userID := claims["user_id"].(string)
	id := c.Params("id")

	if err := service.SetDefaultAddressByOwner(userID, id); err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(helper.ApiSuccess("Shipping address set as default successfully", nil))
}
