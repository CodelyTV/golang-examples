package beerscli

// Store representation of store into data struct
type Store struct {
	StoreID        int     `json:"store_id"`
	Name           string  `json:"name"`
	Address        string  `json:"address"`
	City           string  `json:"city"`
	PostalCode     string  `json:"postal_code"`
	Latitude       float32 `json:"latitude"`
	Longitude      float32 `json:"longitude"`
	MondayOpen     string  `json:"monday_open"`
	MondayClose    string  `json:"monday_close"`
	TuesdayOpen    string  `json:"tuesday_open"`
	TuesdayClose   string  `json:"tuesday_close"`
	WednesdayOpen  string  `json:"wednesday_open"`
	WednesdayClose string  `json:"wednesday_close"`
	ThursdayOpen   string  `json:"thusday_open"`
	ThursdayClose  string  `json:"thursday_close"`
	FridayOpen     string  `json:"friday_open"`
	FridayClose    string  `json:"friday_close"`
	SaturdayOpen   string  `json:"saturday_open"`
	SaturdayClose  string  `json:"saturday_close"`
	SundayOpen     string  `json:"sunday_open"`
	SundayClose    string  `json:"sunday_close"`
}

// StoreRepo definiton of methods to access a data store
type StoreRepo interface {
	GetStores() ([]Store, error)
}

// NewStore initialize struct store
func NewStore(storeID int,
	name, address, city, postalCode string,
	latitude, longitude float32,
	mondayOpen, mondayClose, tuesdayOpen,
	tuesdayClose, wednesdayOpen, wednesdayClose,
	thursdayOpen, thursdayClose, fridayOpen,
	fridayClose, saturdayOpen, saturdayClose,
	sundayOpen, sundayClose string) (s Store) {

	s = Store{
		StoreID:        storeID,
		Name:           name,
		Address:        address,
		City:           city,
		PostalCode:     postalCode,
		Latitude:       latitude,
		Longitude:      longitude,
		MondayOpen:     mondayOpen,
		MondayClose:    mondayClose,
		TuesdayOpen:    tuesdayOpen,
		TuesdayClose:   tuesdayClose,
		WednesdayOpen:  wednesdayOpen,
		WednesdayClose: wednesdayClose,
		ThursdayOpen:   thursdayOpen,
		ThursdayClose:  thursdayClose,
		FridayOpen:     fridayOpen,
		FridayClose:    fridayClose,
		SaturdayOpen:   saturdayOpen,
		SaturdayClose:  saturdayClose,
		SundayOpen:     sundayOpen,
		SundayClose:    sundayClose,
	}
	return
}
