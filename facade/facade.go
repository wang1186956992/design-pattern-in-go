package facade

import (
	"fmt"
)

const (
	SET_CAR_MODEL  = "CarModel SetCarModel"
	SET_CAR_ENGINE = "CarEngine SetCarEngine"
	SET_CAR_BODY   = "CarBody SetCarBody"
	MAKE_CAR_OK    = "Make Car Ok"
)

type CarModel struct{}

func NewCarModel() *CarModel {
	return &CarModel{}
}

func (m *CarModel) SetCatModel() string {
	fmt.Println(SET_CAR_MODEL)
	return SET_CAR_MODEL
}

type CarEngine struct{}

func NewCarEngine() *CarEngine {
	return &CarEngine{}
}

func (e *CarEngine) SetCarEngine() string {
	fmt.Println(SET_CAR_ENGINE)
	return SET_CAR_ENGINE
}

type CarBody struct{}

func NewCarBody() *CarBody {
	return &CarBody{}
}

func (b *CarBody) SetCarBody() string {
	fmt.Println(SET_CAR_BODY)
	return SET_CAR_BODY
}

type CarMaker struct {
	carModel  *CarModel
	carEngine *CarEngine
	carBody   *CarBody
}

func NewMaker() *CarMaker {
	return &CarMaker{
		carModel:  &CarModel{},
		carEngine: &CarEngine{},
		carBody:   &CarBody{},
	}
}

func (c *CarMaker) MakeCar() string {
	c.carModel.SetCatModel()
	c.carEngine.SetCarEngine()
	c.carBody.SetCarBody()
	fmt.Println(MAKE_CAR_OK)
	return MAKE_CAR_OK
}
