# Docker with Go

## Prerequisite

* [Docker 🐳](https://docker.com) installed on host machine, that's all 😒

## Setup

To build and run our `Dockerfile` simply run:

```command
docker-compose up
```

Trust me that's all you need to do.

## Usage

Visit `0.0.0.0:8888` on your web browser

![Hello World](https://res.cloudinary.com/ichtrojan/image/upload/v1566000812/Screenshot_2019-08-17_at_1.11.41_AM_rbteun.png)

Also, if you visit `0.0.0.0:8888/greet/{your-name}` you will get a greeting with your name attached

![greet](https://res.cloudinary.com/ichtrojan/image/upload/v1566000813/Screenshot_2019-08-17_at_1.11.51_AM_mrvj2j.png)

All running within a docker container.
