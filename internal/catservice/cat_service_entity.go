package catservice

type CatServiceEntity struct {
	ID           uint    `gorm:"column:id"`
	Name         string  `gorm:"column:name"`
	Registration *string `gorm:"column:registration"`
	Microchip    *string `gorm:"column:microchip"`
	Gender       string  `json:"gender"`
	BreedID      uint    `gorm:"column:breed_id"`
	ColorID      uint    `gorm:"column:color_id"`
	MotherID     uint   `gorm:"column:mother_id"`
	FatherID     uint   `gorm:"column:father_id"`
	OwnerID      uint   `gorm:"column:owner_id"`
}
