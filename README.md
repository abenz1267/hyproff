# Hyproff = Helper for Hyprland and Kickoff

This tool is still in it's infancy. It'll output a list in the proper format for [Kickoff](https://github.com/j0ru/kickoff).

Modules:

- `hyprland`: windows to focus
- `path`: executables in $PATH
- `desktop`: desktop entries
- `vim`: vim/neovim sessions
- `custom`: custom entries

## Install

`go install github.com/abenz1267/hyproff@latest`

## Usage

`hyproff | kickoff --from-stdin`

## Config

Example:

```json
{
  "terminal": "kitty",
  "modules": ["hyprland", "path", "desktop", "vim"],
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
