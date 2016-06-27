package fourborg

import (
	"sync"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/i2c"
	"github.com/joek/picoborgrev"
)

var _ gobot.Driver = (*FourBorgDriver)(nil)

// FourBorg driver interace
type FourBorg interface {
	Name() string
	Connection() gobot.Connection
	Start() []error
	Halt() []error
	SetMotorLeft(float32) error
	SetMotorRight(float32) error
}

// FourBorgDriver struct
type FourBorgDriver struct {
	name       string
	connection i2c.I2c
	motor      picoborgrev.RevDriver
	lock       sync.Mutex
}

// NewFourBorgDriver creates a new beerbot driver with specified name and i2c interface and MotorController adresses
func NewFourBorgDriver(a i2c.I2c, name string, motor picoborgrev.RevDriver) *FourBorgDriver {
	return &FourBorgDriver{
		name:       name,
		connection: a,
		motor:      motor,
		lock:       sync.Mutex{},
	}
}

// Name is giving the robot name
func (d *FourBorgDriver) Name() string {
	return d.name
}

// Connection is returning the i2c connection
func (d *FourBorgDriver) Connection() gobot.Connection {
	return d.connection
}

// Start is starting the robot
func (d *FourBorgDriver) Start() []error {
	d.lock.Lock()
	defer d.lock.Unlock()

	errors := d.motor.Start()
	if errors != nil {
		return errors
	}

	err := d.motor.ResetEPO()
	if err != nil {
		return []error{err}
	}

	return nil
}

// Halt is stopping the robot
func (d *FourBorgDriver) Halt() []error {
	d.lock.Lock()
	defer d.lock.Unlock()

	errors := d.motor.Halt()
	if errors != nil {
		return errors
	}

	return nil
}

// SetMotorLeft is setting motor speed of left motor
func (d *FourBorgDriver) SetMotorLeft(p float32) error {
	d.lock.Lock()
	defer d.lock.Unlock()

	err := d.motor.SetMotorA(p)
	if err != nil {
		return err
	}

	return nil
}

// SetMotorRight is setting motor speed of right motor
func (d *FourBorgDriver) SetMotorRight(p float32) error {
	d.lock.Lock()
	defer d.lock.Unlock()

	p = p * (-1)

	err := d.motor.SetMotorB(p)
	if err != nil {
		return err
	}

	return nil
}
