@echo off
REM Clear previous results
if exist all_results.json del all_results.json

REM Run simulations and analysis
for /L %%i in (1,1,150) do (
    echo Run: %%i
    go run main.go
    timeout /t 10 /nobreak
    python analysis_json.py
    timeout /t 5 /nobreak
)

