_backlight_completions() {
    local cur prev commands
    cur="${COMP_WORDS[COMP_CWORD]}"
    prev="${COMP_WORDS[COMP_CWORD-1]}"
    commands="ls set dec inc help"

    # Fetch devices into an array
    mapfile -t devices < <(ls /sys/class/backlight 2>/dev/null)

    case "${prev}" in
        backlight)
            # Suggest subcommands
            COMPREPLY=( $(compgen -W "${commands}" -- "${cur}") )
            ;;
        set|dec|inc)
            # Suggest device names
            COMPREPLY=( $(compgen -W "${devices[*]}" -- "${cur}") )
            ;;
        *)
            for dev in "${devices[@]}"; do
                if [[ "$prev" == "$dev" ]]; then
                    # Suggest brightness levels from 10 to 100, in steps of 10
                    brightness_values=$(seq 10 10 100)
                    COMPREPLY=( $(compgen -W "${brightness_values}" -- "${cur}") )
                    return 0
                fi
            done
            COMPREPLY=()
            ;;
    esac
}

# Register the completion function
complete -F _backlight_completions backlight
