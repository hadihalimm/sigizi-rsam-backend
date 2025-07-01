package repo

import (
	"github.com/hadihalimm/sigizi-rsam/internal/config"
	"github.com/hadihalimm/sigizi-rsam/internal/model"
)

type SnackRepo interface {
	Create(snack *model.Snack) (*model.Snack, error)
	FindAll() ([]model.Snack, error)
	FindByID(id uint) (*model.Snack, error)
	Update(snack *model.Snack) (*model.Snack, error)
	Delete(id uint) error

	CreateVariant(variant *model.SnackVariant) (*model.SnackVariant, error)
	FindAllVariant(snackID uint) ([]model.SnackVariant, error)
	FindVariantByID(id uint) (*model.SnackVariant, error)
	UpdateVariant(variant *model.SnackVariant) (*model.SnackVariant, error)
	DeleteVariant(id uint) error
	ReplaceMealTypes(variant *model.SnackVariant, mealTypeIDs []uint) error
	ReplaceDiets(variant *model.SnackVariant, dietIDs []uint) error
}

type snackRepo struct {
	db *config.Database
}

func NewSnackRepo(db *config.Database) SnackRepo {
	return &snackRepo{db: db}
}

func (r *snackRepo) Create(snack *model.Snack) (*model.Snack, error) {
	tx := r.db.Gorm.Create(&snack)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var created model.Snack
	tx = r.db.Gorm.Preload("SnackVariants").
		Preload("SnackVariants.MealTypes").
		Preload("SnackVariants.Diets").
		Preload("SnackVariants.SnackVariantMaterialUsages").
		Preload("SnackVariants.SnackVariantMaterialUsages.FoodMaterial").First(&created, snack.ID)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &created, nil
}

func (r *snackRepo) FindAll() ([]model.Snack, error) {
	var snacks []model.Snack
	tx := r.db.Gorm.Preload("SnackVariants").
		Preload("SnackVariants.MealTypes").
		Preload("SnackVariants.Diets").
		Preload("SnackVariants.SnackVariantMaterialUsages").
		Preload("SnackVariants.SnackVariantMaterialUsages.FoodMaterial").Find(&snacks)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return snacks, nil
}

func (r *snackRepo) FindByID(id uint) (*model.Snack, error) {
	var snack model.Snack
	tx := r.db.Gorm.Preload("SnackVariants").
		Preload("SnackVariants.MealTypes").
		Preload("SnackVariants.Diets").
		Preload("SnackVariants.SnackVariantMaterialUsages").
		Preload("SnackVariants.SnackVariantMaterialUsages.FoodMaterial").First(&snack, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &snack, nil
}

func (r *snackRepo) Update(snack *model.Snack) (*model.Snack, error) {
	if err := r.db.Gorm.Save(&snack).Error; err != nil {
		return nil, err
	}

	updated, err := r.FindByID(snack.ID)
	if err != nil {
		return nil, err
	}

	return updated, nil
}

func (r *snackRepo) Delete(id uint) error {
	var variantIDs []uint
	if err := r.db.Gorm.Model(&model.SnackVariant{}).
		Where("snack_id = ?", id).
		Pluck("id", &variantIDs).Error; err != nil {
		return err
	}

	for _, variantID := range variantIDs {
		var variant model.SnackVariant
		if err := r.db.Gorm.First(&variant, variantID).Error; err != nil {
			continue
		}
		// Clear many-to-many associations
		if err := r.db.Gorm.Model(&variant).Association("MealTypes").Clear(); err != nil {
			return err
		}
		if err := r.db.Gorm.Model(&variant).Association("Diets").Clear(); err != nil {
			return err
		}
		if err := r.db.Gorm.Model(&variant).Association("SnackVariantMaterialUsages").Clear(); err != nil {
			return err
		}
	}
	if err := r.db.Gorm.Where("snack_id = ?", id).Delete(&model.SnackVariant{}).Error; err != nil {
		return err
	}
	if err := r.db.Gorm.Delete(&model.Snack{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *snackRepo) CreateVariant(variant *model.SnackVariant) (*model.SnackVariant, error) {
	tx := r.db.Gorm.Create(&variant)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var created model.SnackVariant
	tx = r.db.Gorm.
		Preload("MealTypes").
		Preload("Diets").
		Preload("SnackVariantMaterialUsages").
		Preload("SnackVariantMaterialUsages.FoodMaterial").First(&created, variant.ID)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &created, nil
}

func (r *snackRepo) FindAllVariant(snackID uint) ([]model.SnackVariant, error) {
	var variants []model.SnackVariant

	tx := r.db.Gorm.
		Preload("MealTypes").
		Preload("Diets").
		Preload("SnackVariantMaterialUsages").
		Preload("SnackVariantMaterialUsages.FoodMaterial").
		Where("snack_id = ?", snackID).Find(&variants)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return variants, nil
}

func (r *snackRepo) FindVariantByID(id uint) (*model.SnackVariant, error) {
	var variant model.SnackVariant
	tx := r.db.Gorm.
		Preload("MealTypes").
		Preload("Diets").
		Preload("SnackVariantMaterialUsages").
		Preload("SnackVariantMaterialUsages.FoodMaterial").First(&variant, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &variant, nil
}

func (r *snackRepo) UpdateVariant(variant *model.SnackVariant) (*model.SnackVariant, error) {
	if err := r.db.Gorm.
		Where("snack_variant_id = ?", variant.ID).
		Delete(&model.SnackVariantMaterialUsage{}).Error; err != nil {
		return nil, err
	}

	if err := r.db.Gorm.Save(&variant).Error; err != nil {
		return nil, err
	}

	updated, err := r.FindVariantByID(variant.ID)
	if err != nil {
		return nil, err
	}
	return updated, nil
}

func (r *snackRepo) ReplaceDiets(variant *model.SnackVariant, dietIDs []uint) error {
	var diets []model.Diet
	if len(dietIDs) == 0 {
		return r.db.Gorm.Model(&variant).Association("Diets").Clear()
	}
	tx := r.db.Gorm.Where("id IN ?", dietIDs).Find(&diets)
	if tx.Error != nil {
		return tx.Error
	}
	return r.db.Gorm.Model(&variant).Association("Diets").Replace(diets)
}

func (r *snackRepo) ReplaceMealTypes(variant *model.SnackVariant, mealTypeIDs []uint) error {
	var mealTypes []model.MealType
	if len(mealTypeIDs) == 0 {
		return r.db.Gorm.Model(&variant).Association("MealTypes").Clear()
	}
	tx := r.db.Gorm.Where("id IN ?", mealTypeIDs).Find(&mealTypes)
	if tx.Error != nil {
		return tx.Error
	}
	return r.db.Gorm.Model(&variant).Association("MealTypes").Replace(mealTypes)
}

func (r *snackRepo) DeleteVariant(id uint) error {
	if err := r.db.Gorm.
		Where("snack_variant_id = ?", id).
		Delete(&model.SnackVariantMaterialUsage{}).Error; err != nil {
		return err
	}

	if err := r.db.Gorm.Delete(&model.SnackVariant{}, id).Error; err != nil {
		return err
	}
	return nil
}
