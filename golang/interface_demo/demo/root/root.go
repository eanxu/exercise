package root

import "fmt"

// type DistrictRepo struct {
//	db *gorm.DB
//}
//
//func NewDistrictRepository(db *gorm.DB) *DistrictRepo {
//	return &DistrictRepo{db}
//}
//
//var _ repository.DistrictRepository = &DistrictRepo{}

type Root struct {
	Name string
	Age  int
}

type RootInterface interface {
	Save(name string, age int)
}

func (r *Root) Save(name string, age int)  {
	r.Age = age
	r.Name = name
	fmt.Println("root: ", r)
}
