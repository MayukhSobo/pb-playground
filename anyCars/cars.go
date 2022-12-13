package anycars

import (
	"pb_playground/pb"
)

func NewAnyCar(CarType string, carData []byte) *pb.Car {
	return &pb.Car{
		CarType: CarType,
		CarData: carData,
	}
}
