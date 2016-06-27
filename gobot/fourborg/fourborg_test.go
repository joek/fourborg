package fourborg_test

import (
	. "github.com/joek/fourborg/gobot/fourborg"
	"github.com/joek/picoborgrev/revtesthelpers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("FourBorg", func() {
	var motor *revtesthelpers.FakeRevDriver

	BeforeEach(func() {
		motor = revtesthelpers.NewFakeRevDriver()
	})

	It("Creates a new FourBorgDriver instance", func() {
		var d *FourBorgDriver
		d = NewFourBorgDriver(revtesthelpers.NewI2cTestAdaptor("adaptor"), "Test", motor)
		Ω(d).Should(BeAssignableToTypeOf(&FourBorgDriver{}))
	})

	It("Is starting the robot", func() {
		m1 := false
		epo1 := false
		motor.StartImpl = func() []error {
			m1 = true
			return nil
		}

		motor.ResetEPOImpl = func() error {
			epo1 = true
			return nil
		}

		d := NewFourBorgDriver(revtesthelpers.NewI2cTestAdaptor("adaptor"), "Test", motor)
		d.Start()

		Ω(m1).Should(BeTrue())
		Ω(epo1).Should(BeTrue())

	})

	It("Is stopping the robot", func() {
		stop1 := false
		motor.HaltImpl = func() []error {
			stop1 = true
			return nil
		}

		d := NewFourBorgDriver(revtesthelpers.NewI2cTestAdaptor("adaptor"), "Test", motor)
		d.Halt()

		Ω(stop1).Should(BeTrue())
	})

	It("Is returning name", func() {

		d := NewFourBorgDriver(revtesthelpers.NewI2cTestAdaptor("adaptor"), "Test", motor)
		d.Halt()

		Ω(d.Name()).Should(Equal("Test"))
	})

	It("Is returning connection", func() {
		i := revtesthelpers.NewI2cTestAdaptor("adaptor")
		d := NewFourBorgDriver(i, "Test", motor)
		d.Halt()

		Ω(d.Connection()).Should(Equal(i))
	})

	It("Is setting left Motors", func() {
		var m1 float32
		motor.SetMotorAImpl = func(p float32) error {
			m1 = p
			return nil
		}

		d := NewFourBorgDriver(revtesthelpers.NewI2cTestAdaptor("adaptor"), "Test", motor)
		d.SetMotorLeft(0.32)

		Ω(m1).Should(Equal(float32(0.32)))
	})

	It("Is setting right Motors", func() {
		var m1 float32
		motor.SetMotorBImpl = func(p float32) error {
			m1 = p
			return nil
		}

		d := NewFourBorgDriver(revtesthelpers.NewI2cTestAdaptor("adaptor"), "Test", motor)
		d.SetMotorRight(0.32)

		Ω(m1).Should(Equal(float32(-0.32)))
	})
})
