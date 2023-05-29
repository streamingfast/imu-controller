package main

import (
	"encoding/hex"
	"fmt"
	"imu-logger/device/iim42652"
	"time"
)

func main() {
	imu := iim42652.NewSpi("/dev/spidev0.0", iim42652.AccelerationSensitivityG16, true)
	err := imu.Init()
	if err != nil {
		panic(fmt.Errorf("initializing IMU: %w", err))
	}

	p, err := imu.ReadRegister(iim42652.RegisterPwrMgmt0)
	if err != nil {
		panic("failed to read pwrmgmt0")
	}
	fmt.Println("PwrMgmt0:", hex.EncodeToString([]byte{p}))

	ac, err := imu.ReadRegister(iim42652.RegisterAccelConfig)
	if err != nil {
		panic("failed to read RegisterAccelConfig")
	}
	fmt.Println("RegisterAccelConfig:", hex.EncodeToString([]byte{ac}))

	//for {
	//	intStatus2, err := imu.ReadRegister(iim42652.RegisterIntStatus2)
	//	if err != nil {
	//		panic("failed to read intstatus2")
	//	}
	//
	//	acc, err := imu.GetAcceleration()
	//	if err != nil {
	//		panic("failed to read acceleration")
	//	}
	//
	//	if err != nil {
	//		panic("failed to read whoami")
	//	}
	//	if intStatus2 > 0 {
	//		fmt.Println("Grrrrrrrrr:", intStatus2)
	//		if intStatus2&0x04 == 0x04 {
	//			fmt.Println("Cam X!", intStatus2&0x04 == 0x04)
	//		}
	//		if intStatus2&0x01 == 0x01 {
	//			fmt.Println("Cam Y!", intStatus2&0x01 == 0x01)
	//		}
	//		if intStatus2&0x02 == 0x02 {
	//			fmt.Println("Cam Z!", intStatus2&0x02 == 0x02)
	//		}
	//		//fmt.Println("Cam X:", intStatus2&0x04 == 0x04, "Y:", intStatus2&0x01 == 0x01, "Z:", intStatus2&0x02 == 0x02, "SMD:", intStatus2&0x08 == 0x08)
	//		fmt.Println("Acceleration:", acc)
	//	}
	//	//if intStatus2&0x08 == 0x08 {
	//	//	fmt.Println("Significant motion detected!")
	//	//}
	//	time.Sleep(10 * time.Millisecond)
	//}
	for {
		acceleration, err := imu.GetAcceleration()
		if err != nil {
			panic(fmt.Errorf("getting acceleration: %w", err))
		}

		fmt.Print("\033[u\033[K")
		imu.Debugln("Acceleration:", acceleration)
		//j, err := json.Marshal(acceleration)
		//if err != nil {
		//	panic(fmt.Errorf("marshaling acceleration: %w", err))
		//}
		//fmt.Printf("Acceleration: %s", string(j))
		//fmt.Println()
		time.Sleep(500 * time.Microsecond)
	}
}
