package transaction

import (
	"time"

	"crowdfund.com/campaign"
	"crowdfund.com/user"
)

type Transaction struct {
	ID         int    `db:"id"`
	CampaignID int    `db:"campaign_id"`
	UserID     int    `db:"user_id"`
	Amount     int    `db:"amount"`
	Status     string `db:"status"`
	Code       string `db:"code"`
	PaymentURL string `db:"payment_url"`
	User       user.User
	Campaign   campaign.Campaign
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}
