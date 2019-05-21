@echo off
rem 准备编译......
cd %~dp0
C:\Windows\Microsoft.NET\Framework64\v3.5\msbuild.exe ../src/明源整体解决方案/明源整体解决方案.sln  /t:Build /p:Configuration="Debug"
if errorlevel 1 goto errorDone

exit /b 0

:errorDone
exit /b 2