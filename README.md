# Counter-Strike Global Offensive Demo File Parser


## Installation

#### 1.Install Docker

The project contains a dockerfile with all the dependencies.

On macos: `brew install boot2docker`

Create docker image: `docker build -t parser .`

## Generate Go wrappers

```bash

docker run -ti -v `pwd`:/go/src/app parser /bin/bash

```

Inside the docker container:

```bash

protoc --go_out=. --proto_path=./ --proto_path=/usr/local/include/ netmessages/netmessages_public.proto

protoc --go_out=. --proto_path=./ --proto_path=/usr/local/include/ usermessages/cstrike15_usermessages_public.proto

```

