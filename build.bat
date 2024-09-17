@echo off

REM Exit immediately if a command exits with a non-zero status
SETLOCAL ENABLEEXTENSIONS
SET ERRORLEVEL=0

echo Starting build script...

REM Navigate to the frontend directory
cd frontend || goto :error
echo In frontend directory: %CD%

REM Run the build command using bun
bun run build || goto :error

REM Navigate to the backend directory
cd ..\backend || goto :error
echo In backend directory: %CD%

REM Build the Go backend
go build -o mango.exe || goto :error

REM Navigate to the public directory
cd public || goto :error
echo In public directory: %CD%

REM Delete index.html
del /F /Q index.html

REM Go back to the root directory
cd ..\..
echo Back to root directory: %CD%

REM Move mango.exe from backend to the root directory
move /Y backend\mango.exe . || goto :error

echo Script executed successfully.
goto :eof

:error
echo An error occurred at %CD%. Exiting script.
pause
exit /b 1
