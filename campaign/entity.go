package campaign

import (
	"time"

	"crowdfund.com/user"
)

type Campaign struct {
	ID               int       `db:"id"`
	UserID           int       `db:"user_id"`
	Name             string    `db:"name"`
	ShortDescription string    `db:"short_description"`
	Description      string    `db:"description"`
	Perks            string    `db:"perks"`
	BackerCount      int       `db:"backer_count"`
	GoalAmount       int       `db:"goal_amount"`
	CurrentAmount    int       `db:"current_amount"`
	Slug             string    `db:"slug"`
	CreatedAt        time.Time `db:"created_at"`
	UpdatedAt        time.Time `db:"updated_at"`
	CampaignImages   []CampaignImage
	User             user.User
}

type CampaignImage struct {
	ID         int       `db:"id"`
	CampaignID int       `db:"campaign_id"`
	FileName   string    `db:"file_name"`
	IsPrimary  int       `db:"is_primary"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}
