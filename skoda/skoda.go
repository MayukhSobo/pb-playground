package skoda

import "pb_playground/pb"

func newEngineOptions() *pb.Engine {
	engine := &pb.Engine{
		Torque:    320,
		PowerInHP: 190,
	}
	return engine
}
func NewSkodaCar() *pb.Skoda {
	SkodaCar := &pb.Skoda{
		ModelName:      "Kodoaq LnK",
		CarEngine:      newEngineOptions(),
		CarPriceInLacs: 45,
	}
	return SkodaCar
}
