# AppLauncher
Simple application launcher as alternative to (not existing) dynamic schortcuts in Windows

## icons
https://hjr265.me/blog/adding-icons-for-go-built-windows-executable/

go get github.com/akavel/rsrc
go install github.com/akavel/rsrc
rsrc -arch amd64 -ico icon.ico

Run this tool multiple times, once for each architecture you are targetting.
It will create a new .syso file every time. Store these files inside your main package.

## compile
go build -ldflags="-H=windowsgui" -o AppLauncher.exe