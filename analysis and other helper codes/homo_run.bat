@echo off
REM Clear previous results
if exist all_results.json del all_results.json

REM Run simulations and analysis
for /L %%i in (1,1,140) do (
    echo Run: %%i
    go run main.go
    timeout /t 10 /nobreak
    python analyse_json_homo.py
    timeout /t 5 /nobreak
)

