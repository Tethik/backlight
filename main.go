package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	sysfsDir = "/sys/class/backlight"
	usage    = `backlight. Simple program to control your backlights.

Usage:
	backlight ls                    Lists all devices and their values
	backlight set <device> <val>    Sets the brightness of a device to a specific value
	backlight dec <device> <val>    Decreases the brightness of a device by a specific value
	backlight inc <device> <val>    Increases the brightness of a device by a specific value
	backlight help					Displays this help message :)
`
)

func deviceError(device string, err error) {
	fmt.Println(fmt.Sprintf("Failed to read device, did you type the device name correctly? (you supplied %s) \n", device))
	log.Fatal(err)
}

func setBrightness(device string, val int) {
	max := getMaxBrightness(device)

	if val > max {
		val = max
	}
	if val < 0 {
		val = 0
	}

	filename := fmt.Sprintf("%s/%s/brightness", sysfsDir, device)
	err := ioutil.WriteFile(filename, []byte(strconv.Itoa(val)), 0777)
	if err != nil {
		fmt.Print("Failed to write to sysfs, do you have the correct permissions? \n\n")
		panic(err)
	}
}

func getMaxBrightness(device string) int {
	filename := fmt.Sprintf("%s/%s/max_brightness", sysfsDir, device)
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		deviceError(device, err)
	}

	val, err := strconv.Atoi(strings.Trim(string(bytes), "\n "))
	if err != nil {
		log.Fatal("Non-integer value in max_brightness read from system", err)
	}
	return val
}

func getBrightness(device string) int {
	filename := fmt.Sprintf("%s/%s/brightness", sysfsDir, device)
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		deviceError(device, err)
	}

	val, err := strconv.Atoi(strings.Trim(string(bytes), "\n "))
	if err != nil {
		log.Fatal("Non-integer value in brightness read from system", err)
	}
	return val
}

/*
Commands
*/

func setCommand() {
	if len(os.Args) < 4 {
		fmt.Println(`Missing arguments.`)
		os.Exit(1)
	}

	device := os.Args[2]

	val, err := strconv.Atoi(strings.Trim(os.Args[3], "\n "))
	if err != nil {
		log.Fatalln("Invalid value given, must be an integer.")
	}

	setBrightness(device, val)
}

func incCommand() {
	if len(os.Args) < 4 {
		fmt.Println(`Missing arguments.`)
		os.Exit(1)
	}

	device := os.Args[2]

	val, err := strconv.Atoi(strings.Trim(os.Args[3], "\n "))
	if err != nil {
		log.Fatalln("Invalid value given, must be an integer.")
	}
	current := getBrightness(device)
	newVal := current + val

	setBrightness(device, newVal)
}

func decCommand() {
	if len(os.Args) < 4 {
		fmt.Println(`Missing arguments.`)
		os.Exit(1)
	}

	device := os.Args[2]

	val, err := strconv.Atoi(strings.Trim(os.Args[3], "\n "))
	if err != nil {
		log.Fatalln("Invalid value given, must be an integer.")
	}
	current := getBrightness(device)
	newVal := current - val

	setBrightness(device, newVal)
}

func listCommand() {
	files, err := ioutil.ReadDir(sysfsDir)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		brightness := getBrightness(f.Name())
		max := getMaxBrightness(f.Name())
		fmt.Println(f.Name(), brightness, "/", max)
	}
}

func main() {
	cmd := ""
	if len(os.Args) > 1 {
		cmd = os.Args[1]
	}

	switch cmd {
	case "help":
		fmt.Println(usage)
	case "--help":
		fmt.Println(usage)
	case "set":
		setCommand()
	case "inc":
		incCommand()
	case "dec":
		decCommand()
	case "ls":
		listCommand()
	default:
		listCommand()
	}
}
