package structure

import "fmt"

/**
适配器模式
 */
type computer interface {
	insertIntoLightningPort()
}

type client struct {

}

func (c *client)insertLightningConnectorIntoComputer(com computer)  {
	fmt.Println("client inserts Lightning connector into computer.")
	com.insertIntoLightningPort()
}

type mac struct {
	
}

func (c *mac)insertIntoLightningPort()  {
	fmt.Println("Lightning connector is plugged into mac machine.")
}

type windows struct {

}
func (c *windows)insertIntoUSBPort()  {
	fmt.Println("USB connector is plugged into windows machine.")
}

type windowsAdapter struct {
	windows *windows
}

func (w windowsAdapter) insertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	w.windows.insertIntoUSBPort()
}

