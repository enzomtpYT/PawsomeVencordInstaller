/*
 * SPDX-License-Identifier: GPL-3.0
 * Vencord Installer, a cross platform gui/cli app for installing Vencord
 * Copyright (c) 2023 Vendicated and Vencord contributors
 */

package main

import (
	"PawsomeVencordInstaller/buildinfo"
	"image/color"
)

const ReleaseUrl = "https://api.github.com/repos/enzomtpYT/PawsomeVencord/releases/latest"
const ReleaseUrlFallback = "https://equicord.org/releases/equicord"
const InstallerReleaseUrl = "https://api.github.com/repos/enzomtpYT/PawsomeVencordInstaller/releases/latest"
const InstallerReleaseUrlFallback = "https://equicord.org/releases/PawsomeVencordInstaller"

var UserAgent = "PawsomeVencordInstaller/" + buildinfo.InstallerGitHash + " (https://github.com/enzomtpYT/PawsomeVencordInstaller)"

const SupportUrl = "https://github.com/enzomtpYT/PawsomeVencordInstaller/issues"

var (
	DiscordGreen        = color.RGBA{0, 133, 69, 0xff}
	DiscordGreenHovered = color.RGBA{0, 108, 55, 0xff}
	DiscordRed          = color.RGBA{210, 45, 57, 0xff}
	DiscordRedHovered   = color.RGBA{169, 35, 46, 0xff}
	DiscordBlue         = color.RGBA{88, 101, 242, 0xff}
	DiscordBlueHovered  = color.RGBA{68, 82, 187, 0xff}
	DiscordYellow       = color.RGBA{0xfe, 0xe7, 0x5c, 0xff}
	PawsomeINFO 		= color.RGBA{R: 0xc2, G: 0x7f, B: 0xbe, A: 0xff}
	PawsomeINFOHovered 	= color.RGBA{R: 0x9c, G: 0x5f, B: 0x8e, A: 0xff}
)

var LinuxDiscordNames = []string{
	"Discord",
	"DiscordPTB",
	"DiscordCanary",
	"DiscordDevelopment",
	"discord",
	"discordptb",
	"discordcanary",
	"discorddevelopment",
	"discord-ptb",
	"discord-canary",
	"discord-development",
	// Flatpak
	"com.discordapp.Discord",
	"com.discordapp.DiscordPTB",
	"com.discordapp.DiscordCanary",
	"com.discordapp.DiscordDevelopment",
}