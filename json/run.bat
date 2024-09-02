@echo off
REM バイナリのパスを設定
set BINARY_PATH=createJson.exe

REM ルートフォルダのパスを設定
set ROOT_FOLDER=.

REM バイナリを実行
%BINARY_PATH% %ROOT_FOLDER%

REM 実行結果を表示
echo create json done
pause