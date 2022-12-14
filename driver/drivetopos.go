package driver

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"github.com/xXNurioXx/go-gta-sa-driver/driver/movement"
	"github.com/xXNurioXx/go-gta-sa-driver/math"
)

func (driver *Driver) DriveToPos(x, y float64) {
	destination := math.Location{X: x, Y: y}
	fmt.Printf("Driving to %f, %f...\n", destination.X, destination.Y)

	for driver.getDistanceTo(destination) >= 5 {
		robotgo.MilliSleep(100)
		driver.DriveTo(destination)
	}
	movement.StopMoving()
}
