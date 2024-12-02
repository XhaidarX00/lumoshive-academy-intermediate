package repository

import (
	"encoding/json"
	"fmt"
	"golang-chapter-37/implem-gin-gonic/model"
	"net/http"

	"gorm.io/gorm"
)

type ShippingRepository interface {
	CalculateRoute(userID, sellerID int) (float64, error)
	GetShippingRepo(data *[]model.ShippingService) error
}

type shippingRepository struct {
	db           *gorm.DB
	customerRepo CustomerRepository
	sellerRepo   SellerRepository
}

func NewShippingRepository(db *gorm.DB, cs CustomerRepository, slr SellerRepository) ShippingRepository {
	return &shippingRepository{
		db:           db,
		customerRepo: cs,
		sellerRepo:   slr,
	}
}

type OSRMResponse struct {
	Code   string `json:"code"`
	Routes []struct {
		Legs []struct {
			Distance float64 `json:"distance"`
		} `json:"legs"`
	} `json:"routes"`
	Waypoints []struct {
		Distance float64   `json:"distance"`
		Name     string    `json:"name"`
		Location []float64 `json:"location"`
	} `json:"waypoints"`
}

func (r *shippingRepository) CalculateRoute(userID, sellerID int) (float64, error) {
	customerAddresses, err := r.customerRepo.GetCustomerAddressesByUserID(userID)
	if err != nil {
		return 0, fmt.Errorf("error fetching customer address: %w", err)
	}

	if len(customerAddresses) == 0 {
		return 0, fmt.Errorf("no customer address found for userID: %d", userID)
	}

	customerAddress := customerAddresses[0] // Ambil alamat pertama (atau default)

	sellerAddress, err := r.sellerRepo.GetSellerAddressesByUserID(sellerID)
	if err != nil {
		return 0, fmt.Errorf("error fetching seller address: %w", err)
	}

	long1, lat1 := customerAddress.Longitude, customerAddress.Latitude
	long2, lat2 := sellerAddress.Longitude, sellerAddress.Latitude

	url := fmt.Sprintf("https://router.project-osrm.org/route/v1/driving/%f,%f;%f,%f?overview=false", long1, lat1, long2, lat2)

	resp, err := http.Get(url)
	if err != nil {
		return 0, fmt.Errorf("error making request to OSRM API: %w", err)
	}
	defer resp.Body.Close()

	var osrmResponse OSRMResponse
	if err := json.NewDecoder(resp.Body).Decode(&osrmResponse); err != nil {
		return 0, fmt.Errorf("error decoding OSRM response: %w", err)
	}

	if len(osrmResponse.Routes) == 0 || len(osrmResponse.Routes[0].Legs) == 0 {
		return 0, fmt.Errorf("invalid OSRM response: no routes or legs found")
	}

	distance := osrmResponse.Routes[0].Legs[0].Distance

	return distance, nil
}

func (r *shippingRepository) GetShippingRepo(data *[]model.ShippingService) error {

	// Gunakan GORM untuk query data
	err := r.db.Order("courier_name, service_type").Find(&data).Error
	if err != nil {
		return err
	}

	return nil
}
