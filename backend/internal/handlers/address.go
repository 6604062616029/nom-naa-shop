package handlers

import (
	"fmt"
	"strconv"

	"github.com/baimhons/nom-naa-shop.git/internal/dtos/request"
	"github.com/baimhons/nom-naa-shop.git/internal/dtos/response"
	"github.com/baimhons/nom-naa-shop.git/internal/models"
	"github.com/baimhons/nom-naa-shop.git/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type AddressHandler struct {
	addressService services.AddressService
}

func NewAddressHandler(addressService services.AddressService) *AddressHandler {
	return &AddressHandler{addressService: addressService}
}

func (h *AddressHandler) GetProvince(c *fiber.Ctx) error {
	provinceCode := c.Params("province_code")
	provinceCodeInt, err := strconv.Atoi(provinceCode)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Message: "Invalid province ID",
		})
	}
	province, err := h.addressService.GetProvinceByCode(provinceCodeInt)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Message: "Failed to get province",
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.SuccessResponse{
		Message: "Province fetched successfully",
		Data:    province,
	})
}

func (h *AddressHandler) GetDistrict(c *fiber.Ctx) error {
	districtCode := c.Params("district_code")
	districtCodeInt, err := strconv.Atoi(districtCode)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Message: "Invalid district ID",
		})
	}
	district, err := h.addressService.GetDistrictByCode(districtCodeInt)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Message: "Failed to get district",
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.SuccessResponse{
		Message: "District fetched successfully",
		Data:    district,
	})
}

func (h *AddressHandler) GetSubDistrict(c *fiber.Ctx) error {
	subDistrictCode := c.Params("sub_district_code")
	subDistrictCodeInt, err := strconv.Atoi(subDistrictCode)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Message: "Invalid sub district ID",
		})
	}
	subDistrict, err := h.addressService.GetSubDistrictByCode(subDistrictCodeInt)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Message: "Failed to get sub district",
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.SuccessResponse{
		Message: "Sub district fetched successfully",
		Data:    subDistrict,
	})
}

func (h *AddressHandler) CreateAddress(c *fiber.Ctx) error {
	userContext := c.Locals("userContext").(models.UserContext)
	if userContext.ID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponse{
			Message: "Unauthorized user",
		})
	}

	var req request.CreateAddressRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Message: "Invalid request",
		})
	}

	address, err := h.addressService.CreateAddress(req, userContext)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(address)
}

func (h *AddressHandler) UpdateAddress(c *fiber.Ctx) error {
	userContext := c.Locals("userContext").(models.UserContext)
	if userContext.ID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponse{
			Message: "Unauthorized user",
		})
	}

	var req request.UpdateAddressRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Message: "Invalid request",
		})
	}

	address, err := h.addressService.UpdateAddress(req, userContext)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(address)
}

func (h *AddressHandler) DeleteAddress(c *fiber.Ctx) error {
	userContext := c.Locals("userContext").(models.UserContext)
	if userContext.ID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponse{
			Message: "Unauthorized user",
		})
	}

	addressID := c.Params("id")
	addressIDUUID, err := uuid.Parse(addressID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Message: "Invalid address ID",
		})
	}

	err = h.addressService.DeleteAddress(addressIDUUID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Message: "Failed to delete address",
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.SuccessResponse{
		Message: "Address deleted successfully",
	})
}

func (h *AddressHandler) GetAllAddress(c *fiber.Ctx) error {
	userContext, ok := c.Locals("userContext").(models.UserContext)
	if !ok || userContext.ID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponse{
			Message: "Unauthorized user",
		})
	}

	userUUID, err := uuid.Parse(userContext.ID)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponse{
			Message: "Invalid user context",
		})
	}

	addresses, err := h.addressService.GetAllAddressByUserID(userUUID, userContext)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Message: "Failed to get addresses",
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.SuccessResponse{
		Message: "Addresses fetched successfully",
		Data:    addresses,
	})
}

func (h *AddressHandler) GetAddressByID(c *fiber.Ctx) error {
	fmt.Println("test4==")
	userContext := c.Locals("userContext").(models.UserContext)
	if userContext.ID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponse{
			Message: "Unauthorized user",
		})
	}

	addressID := c.Params("id")
	addressIDUUID, err := uuid.Parse(addressID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Message: "Invalid address ID",
		})
	}

	address, err := h.addressService.GetAddressByID(addressIDUUID, userContext)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Message: "Failed to get address",
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.SuccessResponse{
		Message: "Address fetched successfully",
		Data:    address,
	})
}

func (h *AddressHandler) GetAllProvinces(c *fiber.Ctx) error {
	provinces, err := h.addressService.GetAllProvince()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Message: "Failed to get provinces",
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.SuccessResponse{
		Message: "Provinces fetched successfully",
		Data:    provinces,
	})
}

func (h *AddressHandler) GetAllDistrictsByProvinceCode(c *fiber.Ctx) error {
	provinceCode := c.Params("province_code")
	provinceCodeInt, err := strconv.Atoi(provinceCode)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Message: "Invalid province code",
		})
	}

	districts, err := h.addressService.GetAllDistrictsByProvinceCode(provinceCodeInt)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Message: "Failed to get districts",
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.SuccessResponse{
		Message: "Districts fetched successfully",
		Data:    districts,
	})
}

func (h *AddressHandler) GetAllSubDistrictsByDistrictCode(c *fiber.Ctx) error {
	districtCode := c.Params("district_code")
	districtCodeInt, err := strconv.Atoi(districtCode)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Message: "Invalid district code",
		})
	}

	subDistricts, err := h.addressService.GetAllSubDistrictsByDistrictCode(districtCodeInt)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Message: "Failed to get sub districts",
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.SuccessResponse{
		Message: "Sub districts fetched successfully",
		Data:    subDistricts,
	})
}
