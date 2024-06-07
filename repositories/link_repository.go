package repositories

import (
	model "example.com/prac02/models"
	"gorm.io/gorm"
)

type LinktreeRepository interface {
	CreateLink(link *model.Link) error
	GetLinksByUser(userID uint) ([]model.Link, error)
	UpdateLink(link *model.Link) error
	DeleteLink(link *model.Link) error
	FindLinkByID(linkId uint) (*model.Link, error)
}

type linktreeRepository struct {
	db *gorm.DB
}

func NewLinktreeRepository(db *gorm.DB) LinktreeRepository {
	return &linktreeRepository{db: db}
}

func (r *linktreeRepository) CreateLink(link *model.Link) error {
	return r.db.Create(link).Error
}

func (r *linktreeRepository) GetLinksByUser(userId uint) ([]model.Link, error) {
	links := []model.Link{}
	err := r.db.Where("user_id = ?", userId).Order("type ASC").Find(&links).Error
	return links, err
}

func (r *linktreeRepository) UpdateLink(link *model.Link) error {
	return r.db.Save(link).Error
}

func (r *linktreeRepository) DeleteLink(link *model.Link) error {
	return r.db.Delete(link).Error
}

func (r *linktreeRepository) FindLinkByID(linkID uint) (*model.Link, error) {
	link := model.Link{}
	result := r.db.Where("id = ? ", linkID).First(&link)
	if result.Error != nil {
		return nil, result.Error
	}
	return &link, nil
}
