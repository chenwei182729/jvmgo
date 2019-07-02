@echo off

setlocal

if exist install.bat goto ok
echo install.bat must be run from its folder
goto end

: ok

set OLDGOPATH=%GOPATH%
set GOPATH=%cd%

gofmt -w src
go run main -version

@rem go install main
@rem 运行exe文件
@rem %GOPATH%\bin\main.exe

set GOPATH=OLDGOPATH

:end
echo finished