@echo off

REM Exit immediately if a command exits with a non-zero status
SETLOCAL ENABLEEXTENSIONS
SET ERRORLEVEL=0

REM Navigate to the frontend directory
cd frontend || goto :error

REM Run the build command using bun
bun run build || goto :error

REM Navigate to the backend\public directory
cd ..\backend\public || goto :error

REM Delete index.html
del /F /Q index.html

REM Go back to the root directory
cd ..\..\..

REM Move mango.exe from backend to the root directory
move /Y backend\mango.exe . || goto :error

echo Script executed successfully.
goto :eof

:error
echo An error occurred. Exiting script.
exit /b 1
