package volvo

import "pb_playground/pb"

func newAdasFeatues() *pb.Adas {
	adas := &pb.Adas{
		AdaptiveCruiseControl: true,
		LaneKeepAssist:        true,
		BlindSoptMontior:      true,
	}
	return adas
}
func NewVolvoCar() *pb.Volvo {
	volvoCar := &pb.Volvo{
		ModelName:      "XC60",
		AdasFeatures:   newAdasFeatues(),
		CarPriceInLacs: 78,
	}
	return volvoCar
}
