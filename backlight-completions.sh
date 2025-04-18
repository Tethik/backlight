_backlight_completions() {
    local cur prev commands devices
    cur="${COMP_WORDS[COMP_CWORD]}"
    prev="${COMP_WORDS[COMP_CWORD-1]}"
    commands="ls set dec inc help"

    # Dynamically fetch devices from /sys/class/backlight
    devices=$(ls /sys/class/backlight 2>/dev/null)

    case "${prev}" in
        backlight)
            # Suggest commands
            COMPREPLY=( $(compgen -W "${commands}" -- "${cur}") )
            ;;
        set|dec|inc)
            # Suggest devices for commands that require a device
            COMPREPLY=( $(compgen -W "${devices}" -- "${cur}") )
            ;;
        *)
            COMPREPLY=()
            ;;
    esac
}

# Register the completion function for the `backlight` command
complete -F _backlight_completions backlight