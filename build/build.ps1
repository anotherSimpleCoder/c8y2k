$dest = "<your destination goes here>\c8y2k.exe"

go build main.go
Move-Item -Path "main.exe" -Destination $dest -Force