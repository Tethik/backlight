
all: backlight

backlight: main.go
	go build 	

clean: 
	rm backlight

install: backlight
	cp backlight /usr/local/bin/
	setcap cap_dac_override+ep /usr/local/bin/backlight
	
