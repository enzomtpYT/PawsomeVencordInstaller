/*
 * SPDX-License-Identifier: GPL-3.0
 * Vencord Installer, a cross platform gui/cli app for installing Vencord
 * Copyright (c) 2023 Vendicated and Vencord contributors
 */

package main

import "image/color"

// these are replaced by the linker

var InstallerGitHash = "Unknown"
var InstallerTag = "Unknown"

const ReleaseUrl = "https://api.github.com/repos/enzomtpYT/PawsomeVencord/releases/latest"
const ReleaseUrlFallback = "https://vencord.dev/releases/vencord"
const InstallerReleaseUrl = "https://api.github.com/repos/enzomtpYT/PawsomeVencordInstaller/releases/latest"
const InstallerReleaseUrlFallback = "https://vencord.dev/releases/installer"

var UserAgent = "PawsomeVencordInstaller/" + InstallerGitHash + " (https://github.com/Vencord/Installer)"

var (
	DiscordBlue   = color.RGBA{R: 0x58, G: 0x65, B: 0xF2, A: 0xFF}
	PawsomeINFO = color.RGBA{R: 0x6f, G: 0x3e, B: 0x6d, A: 0xff}
	PawsomeRed = color.RGBA{R: 0x74, G: 0x03, B: 0x00, A: 0xFF}
	PawsomeInstall = color.RGBA{R: 0xa9, G: 0x80, B: 0xbb, A: 0xFF}
	PawsomeGreen = color.RGBA{R: 0x00, G: 0x55, B: 0x09, A: 0xFF}
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