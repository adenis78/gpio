package gpio

import "log"
import "io/ioutil"
import "strconv"
import "fmt"

const LOW uint8 = 0
const HIGH uint8 = 1
const IN string = "in"
const OUT string = "out"

type GPIO_Pin struct {
	number    uint8
	direction string
	value     uint8
}

func (pin *GPIO_Pin) exportPin() {
	/* For export pin for work with it need write to

	/sys/class/gpio/export

	number of the pin

	For example in bash (export 124 pin)

	echo 124 > /sys/class/gpio/export

	After this will be created file with this pin

	/sys/class/gpio/gpio124/
	*/

	/* Set to buffer number of the pin */
	s := strconv.FormatUint(uint64(pin.number), 10)
	buffer := []byte(s)

	/* Write pin number storing in buffer to file */
	err := ioutil.WriteFile("/sys/class/gpio/export", buffer, 0644)
	if err != nil {
		log.Println(err)
	}
}

func (pin *GPIO_Pin) setDirection() {
	/* For setting direction of the pin need write to

	/sys/class/gpio/gpio<number>/direction

		- "in" for input (in this package for it there is IN const)
		- "out" for output (in this package for it there is OUT const)

	For example in bash (set 124 pin to out)

	echo out > /sys/class/gpio/gpio124/direction
	*/

	/* Set to buffer direction of the pin */
	buffer := []byte(pin.direction)

	/* Forming name of file accodring to sysfs */
	pinFile := fmt.Sprintf("/sys/class/gpio/gpio%s/direction", strconv.Itoa(int(pin.number)))

	/* Write direction stored in buffer to file */
	err := ioutil.WriteFile(pinFile, buffer, 0644)
	if err != nil {
		log.Println(err)
	}
}

func (pin *GPIO_Pin) init() {
	/* For init of GPIO pin need:
	1. Export the desired pin
	2. Set direction (in or out)
	*/

	/* Export pin */
	pin.exportPin()
	/* Set direction if the pin*/
	pin.setDirection()
}

func (pin *GPIO_Pin) deinit() {

}

func (pin *GPIO_Pin) setOutput() {
	/* For set out value of the pin need write to

	/sys/class/gpio/gpio<number>/value

		- 1 for high level (in this package for it there is HIGH const)
		- 0 for low level (in this package for it there is LOW const)

	For example in bash (set high value on 124 pin)

	echo 1 > /sys/class/gpio/gpio124/value
	*/

	/* Set to buffer direction of the pin */
	s := strconv.FormatUint(uint64(pin.value), 10)
	buffer := []byte(s)

	/* Forming name of file accodring to sysfs */
	pinFile := fmt.Sprintf("/sys/class/gpio/gpio%s/value", strconv.Itoa(int(pin.number)))

	/* Write direction stored in buffer to file */
	err := ioutil.WriteFile(pinFile, buffer, 0644)
	if err != nil {
		log.Println(err)
	}
}

func (pin *GPIO_Pin) Set() {
	pin.value = HIGH
	pin.setOutput()
}

func (pin *GPIO_Pin) Clear() {
	pin.value = LOW
	pin.setOutput()
}

func (pin *GPIO_Pin) Toggle() {
	if pin.value == LOW {
		pin.value = HIGH
	} else {
		pin.value = LOW
	}
	pin.setOutput()
}

func (pin *GPIO_Pin) GetState() uint8 {
	/* For get in value of the pin need read 2 bytes to string from file

	/sys/class/gpio/gpio<number>/value

		- 1 for high level (in this package for it there is HIGH const)
		- 0 for low level (in this package for it there is LOW const)

	For example in bash (get value on 124 pin)

	cat /sys/class/gpio/gpio124/value
	*/

	/* Construct pin filename */
	pinFile := fmt.Sprintf("/sys/class/gpio/gpio%s/value", strconv.Itoa(int(pin.number)))

	/* Read data from pin file */
	buffer, err := ioutil.ReadFile(pinFile)
	if err != nil {
		log.Println(err)
	}

	/* buffer represent ASCII code of the value
	Reading value may be:
	- 0x30 - equal 0
	- 0x31 - equal 1

	For convert ASCII value to uint8 need to subtract 0x30 from code
	*/
	result := uint8(buffer[0] - 0x30)
	/* Return read value */
	return result
}

func NewPin(number uint8, direction string) *GPIO_Pin {
	p := new(GPIO_Pin)
	p.number = number
	p.direction = direction

	p.init()

	return p
}
func NewPinNoInit(number uint8, direction string) *GPIO_Pin {
	p := new(GPIO_Pin)
	p.number = number
	p.direction = direction

	//p.init()

	return p
}
