all: backlight

backlight: main.go
	go build

clean:
	rm backlight

install: backlight
	cp backlight /usr/local/bin/
	setcap cap_dac_override+ep /usr/local/bin/backlight
	cp backlight-completion.sh /etc/bash_completion.d/backlight

# install-completion:
# 	cp backlight-completion.sh /etc/bash_completion.d/backlight
# 	echo "Bash completion script installed. Restart your shell or run 'source /etc/bash_completion.d/backlight' to enable it."

