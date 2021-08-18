package campaign

import (
	"gorm.io/gorm"
)

type Repository interface {
	Save(campaign Campaign) (Campaign, error)
	Update(campaign Campaign) (Campaign, error)
	FindAll() ([]Campaign, error)
	FindByUserID(userID int) ([]Campaign, error)
	FindByID(campaignID int) (Campaign, error)
	CreateImage(campaignImage CampaignImage)(CampaignImage, error)
	MarkAllImagesAsNonPrimary(CampaignID int) (bool, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Campaign, error) {
	var campaigns []Campaign
	err := r.db.Preload("CampaignImages", "campaign_images.is_primary=1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (r *repository) FindByUserID(userID int) ([]Campaign, error) {
	var campaigns []Campaign
	err := r.db.Where("user_id=?", userID).Preload("CampaignImages", "campaign_images.is_primary=1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (r *repository) FindByID(campaignID int) (Campaign, error) {
	var campaign Campaign
	err := r.db.Preload("User").Preload("CampaignImages").Where("id=?", campaignID).Find(&campaign).Error
	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (r *repository) Save(campaign Campaign) (Campaign, error) {
	err := r.db.Create(&campaign).Error
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}

func (r *repository) Update(campaign Campaign) (Campaign, error) {
	err := r.db.Save(&campaign).Error
	if err != nil {
		return campaign, err
	}
	return campaign, nil 
}

func (r *repository) CreateImage(campaignImage CampaignImage)(CampaignImage, error) {
	err := r.db.Create(&campaignImage).Error
	if err != nil {
		return campaignImage, err
	}

	return campaignImage, err
}

func (r *repository) MarkAllImagesAsNonPrimary(CampaignID int) (bool, error) {
	// UPDATE campaign_images SET is_primary = false WHERE campaign_id = 1
	err := r.db.Model(&CampaignImage{}).Where("campaign_id=?", CampaignID).Update("is_primary", false).Error
	if err != nil {
		return false, err
	}

	return false, nil
}