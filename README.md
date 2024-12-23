# PawsomeVencord Installer

This Vencord Installer allows you to install [PawsomeVencord, a modification of the best Discord Desktop client mod](https://github.com/enzomtpYT/PawsomeVencord)

![image](https://user-images.githubusercontent.com/45497981/226734476-5fb42420-844d-4e27-ae06-4799118e086e.png)

## Usage

### Windows

> **Warning**
**Do not** run the installer as Admin

Download [PawsomeVencordInstaller.exe](https://github.com/enzomtpYT/PawsomeVencordInstaller/releases/latest/download/PawsomeVencordInstaller.exe) and run it

If the above doesn't work/open, for example because you're using Windows 7, 32 bit, or have a bad GPU, you can instead use our terminal based installer.

To do so, open Powershell, run the following command, then follow along with the instructions/prompts

```ps1
iwr "https://raw.githubusercontent.com/enzomtpYT/PawsomeVencordInstaller/main/install.ps1" -UseBasicParsing | iex
```

### Linux

**PLEASE MAKE SURE YOU HAVE ANY OTHER VERSION OF VENCORD REMOVED BEFORE INSTALLING PAWSOME, IT CAN CAUSE SOFT CRASHES, IF YOUR DISCORD DOESNT START REINSTALL IT THEN INSTALL PAWSOME AGAIN**
(this is unconfirmed as it has only happened once so far, but I put it here incase you get the issue)

Run the following command in your terminal and follow along with the instructions/prompts

```sh
sh -c "$(curl -sS https://raw.githubusercontent.com/enzomtpYT/PawsomeVencordInstaller/main/install.sh)"
```

### MacOs

Download the latest [MacOs build](https://github.com/enzomtpYT/PawsomeVencordInstaller/releases/latest/download/PawsomeVencordInstaller.MacOS.zip), unzip it, and run `PawsomeVencordInstaller.app` 

If you get a `PawsomeVencordInstaller can't be opened` warning, right-click `PawsomeVencordInstaller.app` and click open.

This warning shows because the app isn't signed since I'm not willing to pay 100 bucks a year for an Apple Developer license.

___

## Building from source

### Prerequisites 

You need to install the [Go programming language](https://go.dev/doc/install) and GCC, the GNU Compiler Collection (MinGW on Windows)

<details>
<summary>Additionally, if you're using Linux, you have to install some additional dependencies:</summary>

#### Base dependencies
```sh
apt install -y pkg-config libsdl2-dev libglx-dev libgl1-mesa-dev
```

#### X11 dependencies
```sh
apt install -y xorg-dev
```

#### Wayland dependencies
```sh
apt install -y libwayland-dev libxkbcommon-dev wayland-protocols extra-cmake-modules
```

</details>

### Building

#### Install dependencies

```sh
go mod tidy
```

#### Build the GUI

##### Windows / Mac / Linux X11
```sh
go build
```

##### Linux Wayland
```sh
go build --tags wayland
```

#### Build the CLI
```
go build --tags cli
```

You might want to pass some flags to this command to get a better build.
See [the GitHub workflow](https://github.com/enzomtpYT/PawsomeVencordInstaller/blob/main/.github/workflows/release.yml) for what flags I pass or if you want more precise instructions
