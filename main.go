package main

import (
	"fmt"
	anycars "pb_playground/anyCars"
	"pb_playground/pb"
	skoda "pb_playground/skoda"
	volvo "pb_playground/volvo"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

func GetGenericCars() [2]*pb.Car {
	/////// For Car 1 data /////////
	car1 := volvo.NewVolvoCar()
	car1Bytes, err := proto.Marshal(car1)
	if err != nil {
		panic(err)
	}
	// fmt.Println(car1Bytes)
	// fmt.Println(car1)
	///////////////////////////////

	////// For Car 2 data ////////
	car2 := skoda.NewSkodaCar()
	car2Bytes, err := proto.Marshal(car2)
	if err != nil {
		panic(err)
	}
	// fmt.Println(car2Bytes)
	// fmt.Println(car2)

	/////////////////////////////
	anyCar1 := anycars.NewAnyCar("Volvo", car1Bytes)
	anyCar2 := anycars.NewAnyCar("Skoda", car2Bytes)

	// fmt.Println(anyCar1)
	// fmt.Println(anyCar2)
	return [2]*pb.Car{anyCar1, anyCar2}
}

func UnmarshalFromByteIntoMessage(packageName, messageName string, byteData []byte) (proto.Message, error) {
	messageFullName := protoreflect.FullName(packageName + "." + messageName)
	pbType, err := protoregistry.GlobalTypes.FindMessageByName(messageFullName)
	if err != nil {
		return nil, err
	}
	msg := pbType.New().Interface()
	err = proto.Unmarshal(byteData, msg)
	if err != nil {
		return nil, err
	}
	return msg, nil
}

func main() {
	anyCars := GetGenericCars()
	anyCar1 := anyCars[0]
	anyCar2 := anyCars[1]
	msg, err := UnmarshalFromByteIntoMessage("cars", anyCar1.GetCarType(), anyCar1.GetCarData())
	if err != nil {
		panic(err)
	}
	fmt.Println("=============")
	fmt.Println(msg)

	msg, err = UnmarshalFromByteIntoMessage("cars", anyCar2.GetCarType(), anyCar2.GetCarData())
	if err != nil {
		panic(err)
	}
	fmt.Println("=============")
	fmt.Println(msg)
}
