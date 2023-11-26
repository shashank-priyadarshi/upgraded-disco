# Upgraded-Disco with hexagonal architecture

mode := dev or prod
base_dir := root directory of application
image := container image name for the application
version := version for the $image
config_source := config file to read config data for the application
config_path := path of $config_source

To run this application, use the following command:

```shell
 docker-compose -f ./build/dev/docker-compose.yml up -d
 export CONFIG_SOURCE=config.yaml
 export CONFIG_PATH=./build/dev
 air -c .air.toml
```

To remove dependencies after stopping the application, use the following command:

```shell
docker-compose -f ./build/dev/docker-compose.yml down
```

To run this application in containerized environment, use the following command:

```shell
./run.sh $mode $base_dir $image $version $config_source $config_path
```

e.g.

```shell
 export CONFIG_SOURCE=config.yaml
 export CONFIG_PATH=./build/dev
 ./run.sh dev . upgraded-disco v0.0.1 config.yaml ./build/dev
```

To stop containerized setup, use the following command:

```shell
./stop.sh $mode $base_dir 
```

e.g.

```shell
./stop.sh dev .
```
