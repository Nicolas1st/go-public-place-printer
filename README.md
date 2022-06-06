# go-printer

## Description

Web-application for setting up and managing a printer

in public places like universities and offices

## Contents
  - [Structure](#structure)
  - [How To Run](#how-to-run)

## Structure

- [cmd/](./cmd/)

  Contains the executable

- [handlers/](./handlers/) contains [REST API](./handlers/api/) and [view functions](./handlers/views/) for server side rendering

- [persistence/](./persistence/) contains data storage packages like
    
  - [db](./persistence/db/) for storing stats and user accounts

  - [filer](./persistence/filer/) for storing files to be printed

  - [session](./persistence/session/) for storing user sessions used for authentication

  - [jobq](./persistence/session/) used as wrapper around CUPS server which is used for printing on UNIX systems

- [web/](./web/) contains static files like [js](./web/js/), [css](./web/css/), and [html templates](./web/html/)

- [pkg/](./pkg) contains helper function(s)


## How To Run

1) Install CUPS and Docker + docker-compose on the target machine

2) Connect a printer to the computer if it hasn't been done yet

3) Configure the environment variables in [docker-compose.yml](./docker-compose.yml) if needed

4) Launch the app

```sh
docker-compose up
```

5) Connect to the app using your browser of choice

   with the credentials and app port that were specified in [docker-compose.yml](./docker-compose.yml)
