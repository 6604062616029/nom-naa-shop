package repositories

import (
	"github.com/baimhons/nom-naa-shop.git/internal/models"
	addressModel "github.com/baimhons/nom-naa-shop.git/internal/models/address_models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AddressRepository interface {
	BaseRepository[models.Address]
	GetAllByUserID(userID uuid.UUID, items *[]models.Address) error

	GetProvinceByCode(provinceCode int, province *addressModel.Province) error
	GetDistrictByCode(districtCode int, district *addressModel.Districts) error
	GetSubDistrictByCode(subDistrictCode int, subDistrict *addressModel.SubDistricts) error
	GetAllDistrictsByProvinceCode(provinceCode int, districts *[]addressModel.Districts) error
	GetAllSubDistrictsByDistrictCode(districtCode int, subDistricts *[]addressModel.SubDistricts) error
}

type addressRepository struct {
	BaseRepository[models.Address]
	DB *gorm.DB
}

func NewAddressRepository(db *gorm.DB) AddressRepository {
	return &addressRepository{
		BaseRepository: NewBaseRepository[models.Address](db),
		DB:             db,
	}
}

func (ar *addressRepository) GetAllByUserID(userID uuid.UUID, items *[]models.Address) error {
	return ar.DB.Where("user_id = ?", userID).Find(items).Error
}

func (ar *addressRepository) GetProvinceByCode(provinceCode int, province *addressModel.Province) error {
	return ar.DB.Where("code = ?", provinceCode).First(province).Error
}

func (ar *addressRepository) GetDistrictByCode(districtCode int, district *addressModel.Districts) error {
	return ar.DB.Where("code = ?", districtCode).First(district).Error
}

func (ar *addressRepository) GetSubDistrictByCode(subDistrictCode int, subDistrict *addressModel.SubDistricts) error {
	return ar.DB.Where("code = ?", subDistrictCode).First(subDistrict).Error
}

func (ar *addressRepository) GetAllDistrictsByProvinceCode(provinceCode int, districts *[]addressModel.Districts) error {
	return ar.DB.Where("province_code = ?", provinceCode).Find(districts).Error
}

func (ar *addressRepository) GetAllSubDistrictsByDistrictCode(districtCode int, subDistricts *[]addressModel.SubDistricts) error {
	return ar.DB.Where("district_code = ?", districtCode).Find(subDistricts).Error
}
