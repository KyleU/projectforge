@ECHO OFF

rem Starts the app. It doesn't reload on Windows

cd %~dp0\..

echo "Windows doesn't allow reloading... sorry"
@ECHO ON
go.exe build -gcflags "all=-N -l" -o build/debug/projectforge.exe .
build\debug\projectforge.exe -v --addr=0.0.0.0 all
