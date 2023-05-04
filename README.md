# Go Chess - A Chess Website Powered by Go and React

### Table of Contents

  * [About](#about)
  * [Features](#features)
  * [Technologies](#technologies)
  * [Requirements](#requirements)
  * [Installation](#installation)
  * [Running the application](#running-the-application)
  * [Usage](#usage)

---

## About 

This is a fullstack web application developed with Go in the backend and React in the frontend.

It works by opening a websocket connection that is used to pass messages back and forth from both parts of the application in order to create real time chess matches.

## Features

- Play chess online with other players in real-time

- See the movement options for your pieces

- Choose between light and dark modes

## Technologies

- **Go** - The programming language used for the backend
- **React** - The library used for the frontend
- **WebSockets** - The communication protocol used between the frontend and the backend

## Requirements 

- Go 1.18 or higher

- Gnu Make

- Node.js

- NPM 

## Installation

1. Clone the repo

```bash
git clone https://github.com/pedro-git-projects/projeto-integrado-frontend
```

2. Run the Makefile 

```bash
make install 
```

## Running the application

1. Start the backend chess server

```bash
make run-chess
```

2. Start the backend blog server

```bash
make run-blog
```

3. Start the frontend server

```bash
make dev
```

Now you can access the application at **http://localhost:5173/**


## Usage

1. Go to "/play"

2. Create a new room

3. Send the room code which is in the url to your friends

4. Play togheter
