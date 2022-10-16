@ECHO OFF
 
:: Written by Cory Dolphin (@wcdolphin www.corydolphin.com)
:: Simple bat script to start a simple Python server in the current directory
:: Usage: server [port <default = 8000>]
 
 if "%1"  =="" ( :: no parameter, default port
  python -m SimpleHTTPServer 8000
) else ( :: use specified port
    python -m SimpleHTTPServer %1
)
