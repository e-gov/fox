#!/bin/bash

if [ $(dpkg-query -W -f='${Status}' golang-go 2>/dev/null | grep -c "ok installed") -eq 0 ]; then
    read -p "go required but not installed. Install? [Y/n]: " answer
    answer=${answer:=Y}
    [[ $answer =~ [Yy] ]] && sudo apt-get install golang-go
fi


if [ $(dpkg-query -W -f='${Status}' node-js 2>/dev/null | grep -c "ok installed") -eq 0 ]; then
    read -p "node-js required but not installed. Install? [Y/n]: " answer
    answer=${answer:=Y}
    [[ $answer =~ [Yy] ]] && sudo apt-get install node-js
fi

if [ $(dpkg-query -W -f='${Status}' npm 2>/dev/null | grep -c "ok installed") -eq 0 ]; then
    read -p "npm required but not installed. Install? [Y/n]: " answer
    answer=${answer:=Y}
    [[ $answer =~ [Yy] ]] && sudo apt-get install npm
fi

gnome-terminal --tab -e "./bash_scripts/build_frontend.sh" --tab -e "./bash_scripts/build_backend.sh"
xterm -hold -e "./bash_scripts/build_frontend.sh" & xterm -hold -e "./bash_scripts/build_backend.sh"
konsole --noclose -e "./bash_scripts/build_frontend.sh" & konsole --noclose -e "./bash_scripts/build_backend.sh"
