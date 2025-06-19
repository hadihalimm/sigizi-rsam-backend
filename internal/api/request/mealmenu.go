package request

type CreateMealMenu struct {
	Name       string `json:"name" binding:"required" validate:"required"`
	Day        string `json:"day" binding:"required" validate:"required,oneof=senin selasa rabu kamis jumat sabtu minggu"`
	Time       string `json:"time" binding:"required" validate:"required,oneof=pagi siang sore"`
	MealTypeID uint   `json:"mealTypeID" binding:"required" validate:"required"`
	FoodIDs    []uint `json:"foodIDs" validate:"required,min=1,dive,gt=0"`
}

type UpdateMealMenu struct {
	Name       string `json:"name" binding:"required" validate:"required"`
	Day        string `json:"day" binding:"required" validate:"required,oneof=senin selasa rabu kamis jumat sabtu minggu"`
	Time       string `json:"time" binding:"required" validate:"required,oneof=pagi siang sore"`
	MealTypeID uint   `json:"mealTypeID" binding:"required" validate:"required"`
	FoodIDs    []uint `json:"foodIDs" validate:"required,min=1,dive,gt=0"`
}
