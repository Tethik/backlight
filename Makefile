
all: backlight

backlight: main.go
	go build 	

clean: 
	rm backlight

install: backlight
	setcap cap_dac_override+ep ./backlight
	mv backlight /usr/local/bin/