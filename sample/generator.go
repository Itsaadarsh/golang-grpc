package sample

import (
	"grpc/project/protopackage"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func NewKeyboard() *protopackage.Keyboard {
	keyboard := &protopackage.Keyboard{
		Layout:  randomKeyboardLayout(),
		Backlit: randomBool(),
	}

	return keyboard
}

func NewCPU() *protopackage.CPU {
	brand := randonCPUBrand()
	name := randomCPUName(brand)

	numberCores := randomInt(2, 8)
	numberThreads := randomInt(numberCores, 12)

	minGhz := randomFloat64(2.0, 3.5)
	maxGhz := randomFloat64(minGhz, 5.0)

	cpu := &protopackage.CPU{
		Brand:         brand,
		Name:          name,
		NumberCores:   uint32(numberCores),
		NumberThreads: uint32(numberThreads),
		MinGhz:        minGhz,
		MaxGhz:        maxGhz,
	}
	return cpu
}

func NewGPU() *protopackage.GPU {
	brand := randomGPUBrand()
	name := randomGPUName(brand)

	minGhz := randomFloat64(1.0, 1.5)
	maxGhz := randomFloat64(minGhz, 2.0)
	memGB := randomInt(2, 6)

	gpu := &protopackage.GPU{
		Brand:  brand,
		Name:   name,
		MinGhz: minGhz,
		MaxGhz: maxGhz,
		Memory: &protopackage.Memory{
			Value: uint32(memGB),
			Unit:  protopackage.Memory_GIGABYTE,
		},
	}

	return gpu
}

func NewRAM() *protopackage.Memory {
	memGB := randomInt(4, 64)

	ram := &protopackage.Memory{
		Value: uint32(memGB),
		Unit:  protopackage.Memory_GIGABYTE,
	}

	return ram
}

func NewSSD() *protopackage.Storage {
	memGB := randomInt(128, 1024)

	ssd := &protopackage.Storage{
		Driver: protopackage.Storage_SSD,
		Memory: &protopackage.Memory{
			Value: uint32(memGB),
			Unit:  protopackage.Memory_GIGABYTE,
		},
	}

	return ssd
}

func NewHDD() *protopackage.Storage {
	memTB := randomInt(1, 6)

	hdd := &protopackage.Storage{
		Driver: protopackage.Storage_HDD,
		Memory: &protopackage.Memory{
			Value: uint32(memTB),
			Unit:  protopackage.Memory_TERABYTE,
		},
	}

	return hdd
}

func NewScreen() *protopackage.Screen {
	screen := &protopackage.Screen{
		SizeInch:   randomFloat32(13, 17),
		Resolution: randomScreenResolution(),
		Panel:      randomScreenPanel(),
		Multitouch: randomBool(),
	}

	return screen
}

func NewLaptop() *protopackage.Laptop {
	brand := randomLaptopBrand()
	name := randomLaptopName(brand)

	laptop := &protopackage.Laptop{
		Id:       randomID(),
		Brand:    brand,
		Name:     name,
		Cpu:      NewCPU(),
		Ram:      NewRAM(),
		Gpus:     []*protopackage.GPU{NewGPU()},
		Storages: []*protopackage.Storage{NewSSD(), NewHDD()},
		Screen:   NewScreen(),
		Keyboard: NewKeyboard(),
		Weight: &protopackage.Laptop_WeightKg{
			WeightKg: randomFloat64(1.0, 3.0),
		},
		PriceUsd:    randomFloat64(1500, 3500),
		ReleaseYear: uint32(randomInt(2015, 2019)),
		UpdatedAt:   timestamppb.Now(),
	}

	return laptop
}
