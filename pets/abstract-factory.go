package pets

import (
	"errors"
	"fmt"
	"go-breeders/internal/cat"
	"go-breeders/internal/dog"
)

type AnimalInterface interface {
	Show() string
}

type DogFromFactory struct {
	Pet *dog.Dog
}

func (dff *DogFromFactory) Show() string {
	return fmt.Sprintf("this animal is a dog")
}

type CatFromFactory struct {
	Pet *cat.Cat
}

func (cff *CatFromFactory) Show() string {
	return fmt.Sprintf("this animal is a cat")
}

type PetFactoryInterface interface {
	newPet() AnimalInterface
}

type DogAbstractFactory struct{}

func (df *DogAbstractFactory) newPet() AnimalInterface {
	return &DogFromFactory{
		Pet: &dog.Dog{},
	}
}

type CatAbstractFactory struct{}

func (cf *CatAbstractFactory) newPet() AnimalInterface {
	return &CatFromFactory{
		Pet: &cat.Cat{},
	}
}

func NewPetFromAbstractFactory(species string) (AnimalInterface, error) {
	switch species {
	case "dog":
		var dogFactory DogAbstractFactory
		dog := dogFactory.newPet()
		return dog, nil
	case "cat":
		var catFactory CatAbstractFactory
		cat := catFactory.newPet()
		return cat, nil
	default:
		return nil, errors.New("invalid species supplied")
	}
}
