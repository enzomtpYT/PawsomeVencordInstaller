# PawsomeVencord Installer

This Vencord Installer allows you to install [PawsomeVencord, a modification of the best Discord Desktop client mod](https://github.com/enzomtpYT/PawsomeVencord)

![image](https://user-images.githubusercontent.com/45497981/226734476-5fb42420-844d-4e27-ae06-4799118e086e.png)

## Usage

Windows

- [GUI](https://github.com/enzomtpYT/PawsomeVencordInstaller/releases/latest/download/PawsomeVencordInstaller.exe)
- [CLI](https://github.com/enzomtpYT/PawsomeVencordInstaller/releases/latest/download/PawsomeVencordInstallerCli.exe)

MacOS

- [X64 GUI](https://github.com/enzomtpYT/PawsomeVencordInstaller/releases/latest/download/PawsomeVencordInstaller-darwin-x64.zip)
- [ARM64 GUI](https://github.com/enzomtpYT/PawsomeVencordInstaller/releases/latest/download/PawsomeVencordInstaller-darwin-arm64.zip)

Linux

- [GUI](https://github.com/enzomtpYT/PawsomeVencordInstaller/releases/latest/download/PawsomeVencordInstaller-x11)
- [CLI](https://github.com/enzomtpYT/PawsomeVencordInstaller/releases/latest/download/PawsomeVencordInstallerCli-Linux)

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
