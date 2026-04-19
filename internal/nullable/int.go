package nullable

import (
	"database/sql/driver"
	"encoding/json"
)

// Int SQL type that can retrieve NULL value
type Int struct {
	RealValue int
	IsValid   bool
}

func FromInt(value int) Int {
	return Int{
		RealValue: value,
		IsValid:   true,
	}
}

// NewInt creates a new nullable integer
func NewInt(value *int) Int {
	if value == nil {
		return Int{
			RealValue: 0,
			IsValid:   false,
		}
	}
	return Int{
		RealValue: *value,
		IsValid:   true,
	}
}

// Get either nil or integer
func (n Int) Get() *int {
	if !n.IsValid {
		return nil
	}
	return &n.RealValue
}

// Set either nil or integer
func (n *Int) Set(value *int) {
	n.IsValid = (value != nil)
	if n.IsValid {
		n.RealValue = *value
	} else {
		n.RealValue = 0
	}
}

// MarshalJSON converts current value to JSON
func (n Int) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.Get())
}

// UnmarshalJSON writes JSON to this type
func (n *Int) UnmarshalJSON(data []byte) error {
	dataString := string(data)
	if len(dataString) == 0 || dataString == "null" {
		n.IsValid = false
		n.RealValue = 0
		return nil
	}

	var parsed int
	if err := json.Unmarshal(data, &parsed); err != nil {
		return err
	}

	n.IsValid = true
	n.RealValue = parsed
	return nil
}

// Scan implements scanner interface
func (n *Int) Scan(value interface{}) error {
	if value == nil {
		n.RealValue, n.IsValid = 0, false
		return nil
	}
	n.IsValid = true
	return convertAssign(&n.RealValue, value)
}

// Value implements the driver Valuer interface.
func (n Int) Value() (driver.Value, error) {
	if !n.IsValid {
		return nil, nil
	}
	return int64(n.RealValue), nil
}
