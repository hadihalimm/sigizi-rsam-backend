package repo

import (
	"fmt"
	"time"

	"github.com/hadihalimm/sigizi-rsam/internal/config"
	"github.com/hadihalimm/sigizi-rsam/internal/model"
)

type MealMenuRepo interface {
	Create(menu *model.MealMenu) (*model.MealMenu, error)
	FindAll() ([]model.MealMenu, error)
	FindByID(id uint) (*model.MealMenu, error)
	Update(menu *model.MealMenu) (*model.MealMenu, error)
	Delete(id uint) error

	CreateNewMealMenuTemplate(template *model.MealMenuTemplate) error
	FindAllMealMenuTemplate() ([]model.MealMenuTemplate, error)
	FindByIDMealMenuTemplate(id uint) (*model.MealMenuTemplate, error)
	UpdateMealMenuTemplate(template *model.MealMenuTemplate) (*model.MealMenuTemplate, error)
	DeleteMealMenuTemplate(id uint) error

	CreateMenuTemplateSchedule(schedule *model.MenuTemplateSchedule) (*model.MenuTemplateSchedule, error)
	FindMenuTemplateScheduleByID(id uint) (*model.MenuTemplateSchedule, error)
	FilterMenuTemplateScheduleByDate(
		date time.Time) (*model.MenuTemplateSchedule, error)
	UpdateMenuTemplateSchedule(
		schedule *model.MenuTemplateSchedule) (*model.MenuTemplateSchedule, error)
}

type mealMenuRepo struct {
	db *config.Database
}

func NewMealMenuRepo(db *config.Database) MealMenuRepo {
	return &mealMenuRepo{db: db}
}

func (r *mealMenuRepo) Create(menu *model.MealMenu) (*model.MealMenu, error) {
	tx := r.db.Gorm.Create(&menu)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var created model.MealMenu
	tx = r.db.Gorm.Preload("Foods").
		Preload("Foods.FoodMaterialUsages").
		Preload("Foods.FoodMaterialUsages.FoodMaterial").
		First(&created, menu.ID)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &created, nil
}

func (r *mealMenuRepo) FindAll() ([]model.MealMenu, error) {
	var menus []model.MealMenu
	tx := r.db.Gorm.Preload("Foods").
		Preload("Foods.FoodMaterialUsages").
		Preload("Foods.FoodMaterialUsages.FoodMaterial").
		Preload("MealType").
		Find(&menus)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return menus, nil
}

func (r *mealMenuRepo) FindByID(id uint) (*model.MealMenu, error) {
	var menu model.MealMenu
	tx := r.db.Gorm.Preload("Foods").
		Preload("Foods.FoodMaterialUsages").
		Preload("Foods.FoodMaterialUsages.FoodMaterial").
		First(&menu, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &menu, nil
}

func (r *mealMenuRepo) Update(menu *model.MealMenu) (*model.MealMenu, error) {
	if err := r.db.Gorm.Model(menu).Association("Foods").Replace(menu.Foods); err != nil {
		return nil, err
	}
	tx := r.db.Gorm.Save(menu)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var updated model.MealMenu
	tx = r.db.Gorm.Preload("Foods").
		Preload("Foods.FoodMaterialUsages").
		Preload("Foods.FoodMaterialUsages.FoodMaterial").
		First(&updated, menu.ID)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &updated, nil
}

func (r *mealMenuRepo) Delete(id uint) error {
	menu, err := r.FindByID(id)
	if err != nil {
		return err
	}
	err = r.db.Gorm.Model(&menu).Association("Foods").Clear()
	if err != nil {
		return err
	}
	tx := r.db.Gorm.Delete(&model.MealMenu{}, id)
	return tx.Error
}

func (r *mealMenuRepo) CreateNewMealMenuTemplate(template *model.MealMenuTemplate) error {
	tx := r.db.Gorm.Create(&template)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (r *mealMenuRepo) FindAllMealMenuTemplate() ([]model.MealMenuTemplate, error) {
	var templates []model.MealMenuTemplate
	tx := r.db.Gorm.Find(&templates)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return templates, nil
}

func (r *mealMenuRepo) FindByIDMealMenuTemplate(id uint) (*model.MealMenuTemplate, error) {
	var template model.MealMenuTemplate
	tx := r.db.Gorm.Preload("MealMenus").
		Preload("MealMenus.Foods").
		Preload("MealMenus.MealType").
		First(&template, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &template, nil
}

func (r *mealMenuRepo) UpdateMealMenuTemplate(template *model.MealMenuTemplate) (*model.MealMenuTemplate, error) {
	tx := r.db.Gorm.Save(template)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var updated *model.MealMenuTemplate
	updated, err := r.FindByIDMealMenuTemplate(template.ID)
	if err != nil {
		return nil, err
	}
	return updated, nil
}

func (r *mealMenuRepo) DeleteMealMenuTemplate(id uint) error {
	template, err := r.FindByIDMealMenuTemplate(id)
	if err != nil {
		return err
	}
	err = r.db.Gorm.Model(&template).Association("MealMenus").Clear()
	if err != nil {
		return err
	}

	tx := r.db.Gorm.Delete(&model.MealMenuTemplate{}, id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (r *mealMenuRepo) CreateMenuTemplateSchedule(
	schedule *model.MenuTemplateSchedule) (*model.MenuTemplateSchedule, error) {
	tx := r.db.Gorm.Create(&schedule)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return schedule, nil
}

func (r *mealMenuRepo) FindMenuTemplateScheduleByID(id uint) (*model.MenuTemplateSchedule, error) {
	var schedule model.MenuTemplateSchedule
	tx := r.db.Gorm.Preload("MealMenuTemplate").
		Preload("MealMenuTemplate.MealMenus").
		Preload("MealMenuTemplate.MealMenus.Foods").
		Preload("MealMenuTemplate.MealMenus.Foods.FoodMaterialUsages").
		Preload("MealMenuTemplate.MealMenus.Foods.FoodMaterialUsages.FoodMaterial").
		Preload("MealMenuTemplate.MealMenus.MealType").
		First(&schedule, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &schedule, nil
}

func (r *mealMenuRepo) FilterMenuTemplateScheduleByDate(
	date time.Time) (*model.MenuTemplateSchedule, error) {
	var schedule model.MenuTemplateSchedule
	tx := r.db.Gorm.Preload("MealMenuTemplate").
		Preload("MealMenuTemplate.MealMenus").
		Preload("MealMenuTemplate.MealMenus.Foods").
		Preload("MealMenuTemplate.MealMenus.Foods.FoodMaterialUsages").
		Preload("MealMenuTemplate.MealMenus.Foods.FoodMaterialUsages.FoodMaterial").
		Preload("MealMenuTemplate.MealMenus.MealType").
		Where("DATE(menu_template_schedules.date) = ?", date.Format("2006-01-02")).
		First(&schedule)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &schedule, nil
}

func (r *mealMenuRepo) UpdateMenuTemplateSchedule(
	schedule *model.MenuTemplateSchedule) (*model.MenuTemplateSchedule, error) {
	tx := r.db.Gorm.Model(&model.MenuTemplateSchedule{}).
		Where("id = ?", schedule.ID).
		Updates(map[string]interface{}{
			"meal_menu_template_id": schedule.MealMenuTemplateID,
		})
	if tx.Error != nil {
		return nil, tx.Error
	}
	var updated *model.MenuTemplateSchedule
	updated, err := r.FindMenuTemplateScheduleByID(schedule.ID)
	fmt.Println(updated.MealMenuTemplateID)
	if err != nil {
		return nil, err
	}
	return updated, nil
}
