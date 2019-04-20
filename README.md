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

For x64 you can download and install the released binary with the following script.

```bash
wget https://github.com/Tethik/backlight/releases/download/1.0.0/backlight.0-x64
chmod +x backlight.0-x64
sudo mv backlight.0-x64 /usr/local/bin/backlight
sudo setcap cap_dac_override+ep /usr/local/bin/backlight
```

### Build from source

Requires that you have golang installed.

```
git clone https://github.com/Tethik/backlight
make
sudo make install
```

_Why sudo?_ To allow the program access to sysfs from a normal user the install script sets `cap_dac_override`.
