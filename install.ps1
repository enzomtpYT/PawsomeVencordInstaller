$link = "https://github.com/enzomtpYT/PawsomeVencordInstaller/releases/latest/download/PawsomeVencordInstallerCli.exe"

$outfile = "$env:TEMP\PawsomeVencordInstallerCli.exe"

Write-Output "Downloading installer to $outfile"

Invoke-WebRequest -Uri "$link" -OutFile "$outfile"

Write-Output ""

Start-Process -Wait -NoNewWindow -FilePath "$outfile"

# Cleanup
Remove-Item -Force "$outfile"
