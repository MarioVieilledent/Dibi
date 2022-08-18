# Dibi

![GitHub top language](https://img.shields.io/github/languages/top/MarioVieilledent/Dibi)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/MarioVieilledent/Dibi)
![GitHub repo size](https://img.shields.io/github/repo-size/MarioVieilledent/Dibi)

A go application for multiple use about Dibi language.

## Deploy

This app needs to be be compiled with Go. [Download Go](https://go.dev/dl/).

> `go build`

## Features

### API
An API is served by default on `127.0.0.1:3001`

- `GET` | `/api` | Get all words (similar to [dibi-dictionary](https://dibi-dictionary.herokuapp.com/dictionary/getWords/all))
- `GET` | `/api/:query` | Get all words matching the query

### Web application
A native web application is served on `127.0.0.1:3000`

The application is a graphic interface for the API.

## Resources

- [Official Discord of Dibi](https://discord.com/invite/xSk3RMpEXB)
- [Main Dibi Dictionary with grammar manual](https://dibi-dictionary.herokuapp.com/)
- [Learn Dibi](https://sites.google.com/view/apprendre-le-dibi-avec-blatha/accueil)
- [Leksiro, lightweight dictionary](https://leksiro.disly.fr/)
- [Clikari, dibi browser](https://clikari.disly.fr/)