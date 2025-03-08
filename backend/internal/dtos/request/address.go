package request

import (
	"github.com/google/uuid"
)

type CreateAddressRequest struct {
	ProvinceCode    int    `json:"province_code" validate:"required"`
	DistrictCode    int    `json:"district_code" validate:"required"`
	SubDistrictCode int    `json:"sub_district_code" validate:"required"`
	AddressDetail   string `json:"address_detail"`
}

type UpdateAddressRequest struct {
	ID              uuid.UUID `json:"id,omitempty"` 
	ProvinceCode    int       `json:"province_code" validate:"required"`
	DistrictCode    int       `json:"district_code" validate:"required"`
	SubDistrictCode int       `json:"sub_district_code" validate:"required"`
	AddressDetail   string    `json:"address_detail"`
}
