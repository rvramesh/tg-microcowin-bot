# TgBotGo

The purpose of this project is to create a base for those who want to build a Telegram bot using Golang.


## How to Run

Install Go. I installed go in wsl2 using https://medium.com/@benzbraunstein/how-to-install-and-setup-golang-development-under-wsl-2-4b8ca7720374 go ver 1.16.3

Using VS Code with Remote WSL extension (https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-wsl) and Go extension for VS Code (https://marketplace.visualstudio.com/items?itemName=golang.Go)

To install dependencies use go mod tidy

copy `.envkeys` to `.env` and set the API token from @BotFather after =

## Project Structure

The project contains...


## CI/CD

The project contains some basics configuration for CI/CD. Basically you already have the configuration for the image creation using Docker as the container manager. And for testing purpose the configuration on travis.