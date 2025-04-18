# 🌚 backlight 🌝

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
wget https://github.com/Tethik/backlight/releases/latest/download/backlight
wget https://github.com/Tethik/backlight/releases/latest/download/backlight-completion.sh
chmod +x backlight
sudo mv backlight /usr/local/bin/backlight
sudo setcap cap_dac_override+ep /usr/local/bin/backlight
sudo mv backlight-completion.sh /etc/bash_completion.d/backlight
```

### Build from source

Requires that you have golang installed.

```
git clone https://github.com/Tethik/backlight
make
sudo make install
```

_Why sudo?_ To allow the program access to sysfs from a normal user the install script sets `cap_dac_override`.

### i3 configuration

I originally built this for my own i3 setup. The following is the configuration I use to
bind the brightness keys on my laptop to run this backlight program.

```
# Sreen brightness controls
bindsym XF86MonBrightnessUp exec backlight inc intel_backlight 50
bindsym XF86MonBrightnessDown exec backlight dec intel_backlight 50
```
