package chronograph

import (
	"time"
)

// NewMonth instantiates a new Month pointer.
func NewMonth(month time.Month) *Month {
	return &Month{
		Month:  month,
		Name:   month.String(),
		Number: int(month)}
}
