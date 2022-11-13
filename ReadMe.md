# Golang Game Framework

Framework in Golang to simplify game development with SDL in Go.

Currently, only basic rendering has been implemented.

For an example on how to use it, see `game/game.go` and make your own changes!

I have a strong desire to make this into a more flexible and feature complete system, but I cannot give any guarantees. Hopefully at minimum, this project can be seen as a reference for others!

## Requirements

Fedora:

```bash
sudo dnf install -y golang SDL2{,_image,_mixer,_ttf,_gfx}-devel
```

Ubuntu:

```bash
sudo apt install -y golang-go libsdl2{,-image,-mixer,-ttf,-gfx}-dev
```