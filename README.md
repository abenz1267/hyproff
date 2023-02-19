# Hyproff = Helper for Hyprland and Kickoff

This tool is still in it's infancy. It'll output a list in the proper format for [Kickoff](https://github.com/j0ru/kickoff).

Modules:

- `hyprland`: windows to focus
- `path`: executables in $PATH
- `desktop`: desktop entries
- `vim`: vim/neovim sessions
- `custom`: custom entries

## Install

### Download binary

Simply download a pre-built binary from the [release page](https://github.com/abenz1267/hyproff/releases)

### With Go

Regular `GOBIN` folder (make sure it's in your `PATH`):
`go install github.com/abenz1267/hyproff@latest`

Install to custom location:
`GOBIN=<custom location> go install github.com/abenz1267/hyproff@latest`

f.e. `sudo GOBIN=/usr/bin/ go install github.com/abenz1267/hyproff@latest`

## Usage

`hyproff | kickoff --from-stdin`

## Config

Example:

```json
{
  "terminal": "kitty",
  "modules": ["custom", "hyprland", "vim", "desktop", "path"], // also acts as order for priority from higher to lowest.
  "vim": {
    "session_dir": "/home/andrej/.local/share/nvim/sessions",
    "editor": "nvim"
  },
  "custom": {
    "label": "Custom",
    "entries": [
      {
        "name": "MyCustomEntry",
        "exec": "dosomething"
      }
    ]
  }
}
```
