# dockerGo
Alta3 Research | https://alta3.com | created by @RZFeeser | A Simple Go Webservice

The purpose of this repository is to provide, "A Simple Go Webservice" for the purposes of studying Go Lang, software design, as well as container testing within Docker and Kubernetes.

This service mimics the Simple Python Flask Service available @:
https://gitlab.com/alta3/simpleflaskservice


The following endpoints are currently available:
| `- /endpoint` | Description of what is returned |
|---------------|---------------------------------|
| `/`  | HTTP + HTML - root info |
| `/ping` | HTTP + JSON - pong! |
| `/spock` | HTTP + JSON - Live long and prosper |
| `/env` | HTTP + JSON - Dump environmental variables |
| `/alta3` | HTTP + JSON - Exciting news relating to Alta3 Research! |
| `/health` | HTTP + JSON - response may be delayed by setting HEALTH_DELAY |
| `/info` | HTTP + JSON - including the host IP and Port, as well as the souce (client) IP |   

Some behaviors of the webservice may be modified by setting the following environmental variables.

*Default Vaules:*
- `ENV_VAR_TO_SET` - Description
- `PORT` - The default port to run the webservice on. The default is 9876
- `HEALTH_DELAY` - The time in seconds to delay an HTTP response sent to /health. The default is 0. The only reason to change this might be mimicing a failure within a Kubernetes health check.
- `VERSION` - The version returned by the server. The default is 0.1.


Builds may be accomplished using the`Dockerfile`and Dockerfile.multistage`.

## Getting started

### Option 1 - Pull the Pre-built Image from the Image Repo and Run the Project

Run the following command and you're done.

`docker run -d  -p 9876:9876 registry.gitlab.com/alta3/simplegoservice:latest`

### Option 2 - Clone the Code, and Build the Container Image Yourself

The second option to using the code would be to clone the project, reivew the code, then build an image yourself. The advantage here is that you can review the code (there is security in this). 

`git clone https://github.com/alta3/simplegoservice`

#### Running Dockerfile (building a container)

Now you need to build the image. The typical way to do this is with instructions contained in a `Dockerfile`. This file is included in the repo. To use it, run the following command.

`sudo docker build --tag dockergo .`

#### Running Dockerfile.multistage (improving container builds)

Alternatively, you could use an *improved* version of the `Dockerfile`, called a "multistage" Dockerfile. The advantage here is that we can use multiple images in our build process. In this case, we are compiling Go code. Therefore, we need the Go programming langauge to compile the source code, but once we've compiled a binary, we no longer need the Go programming langage. We just need to export our binary into a near-empty container for use. The result is a smaller, more light weight, and more secure service.

`(sudo) docker build -t dockergo:multistage -f Dockerfile.multistage .`

#### Run the image

`(sudo) docker run -d  -p 9876:9876 dockergo:latest`  
`(sudo) docker run -d  -p 9876:9876 dockergo:multistage`

#### Stop all containers

`sudo docker stop $(sudo docker ps -aq)`

#### Delete all running containers

`sudo docker container prune`

#### Author(s)
Russell Zachary Feeser - @RZFeeser
