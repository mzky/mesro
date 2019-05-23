package main

import (
	//"fmt"
	"github.com/mzky/configparser"
)

func main() {
	m := "[\"stdout\", \"metrics.out\"]"
	c, _ := configparser.NewConfigInstance("test.conf")
	c.SetItemValue("NewModule", "TestStr", "a new item")
	c.SetItemValue("NewModule", "TestInt", 123)
	c.SetItemValue("NewModule", "TestFloat", 456.78)
	c.SetItemValue("NewModule", "Testbool", true)
	c.SetItemValue("NewModule", "Test[]", m)
	//cfile.SaveToFile("example.conf")		//auto rename the save file name is example.conf.new
	c.SaveFile("a.ini") //auto add suffix -> newfile.conf
}
