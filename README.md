# brightness

Program to manipulate screen brightness on Linux using sysfs.

```
backlight. Simple program to control your backlights.

Usage:
	backlight ls                    Lists all devices and their values
	backlight set <device> <val>    Sets the brightness of a device to a specific value
	backlight dec <device> <val>    Decreases the brightness of a device by a specific value
	backlight inc <device> <val>    Increases the brightness of a device by a specific value
	backlight help					Displays this help message :)
```

## Install

Requires golang. No binary releases at this point.

```
git clone https://github.com/Tethik/backlight
make
sudo make install
```

_Why sudo?_ To allow the program access to sysfs from a normal user the install script sets `cap_dac_override`.
