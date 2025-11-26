# Configuration
$Repo = "yanklio/courls"
$Tool = "courls.exe"
$InstallDir = "$env:USERPROFILE\courls"

# 1. Get latest tag name automatically (Parses JSON)
$ApiUrl = "https://api.github.com/repos/$Repo/releases/latest"
$Tag = (Invoke-RestMethod -Uri $ApiUrl).tag_name

# 2. Create installation folder
if (!(Test-Path $InstallDir)) { New-Item -ItemType Directory -Force -Path $InstallDir | Out-Null }

# 3. Download the .exe
$Url = "https://github.com/$Repo/releases/download/$Tag/$Tool"
Write-Host "Downloading $Tool version $Tag..."
Invoke-WebRequest -Uri $Url -OutFile "$InstallDir\$Tool"

# 4. Add to User PATH (if not already there)
$UserPath = [Environment]::GetEnvironmentVariable("Path", "User")
if ($UserPath -notlike "*$InstallDir*") {
    Write-Host "Adding to PATH..."
    [Environment]::SetEnvironmentVariable("Path", "$UserPath;$InstallDir", "User")
    Write-Host "Success! Please restart your terminal to use '$Tool'."
} else {
    Write-Host "Success! '$Tool' updated."
}
