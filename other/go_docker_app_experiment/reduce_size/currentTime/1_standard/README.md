# Docker + golang Web App

### Final image stats
```
PS C:\Users\JackB\Documents\GitHub\go_docker_app_experiment> docker image ls -a
REPOSITORY                 TAG                 IMAGE ID            CREATED             SIZE
jackburdick/curtime        latest              b0b97d6f4e41        About an hour ago   704MB
```

## Creation
### Build
- $`docker image build -t jackburdick/curtime .`
    - where:
        - `-t` tag/name the image
        - `jackburdick` is `<your_username>`
        - `curtime` is a custom `<projectName>`

### Run
  - `docker run -d --rm -p 8000:8080 jackburdick/curtime`
    - where:
        - `-d`
            - 'detaches' the image --> runs in the background
        - `-p`
            - exposes the port
                - '8000:8080' --> '`container`:`local`'
                    - `8080` on the container, `8000` on the host machine
        - `jackburdick/curtime` is `<your_username>/<projectName>`
  - Confirm:
```
PS C:\Users\JackB\Documents\GitHub\go_docker_app_experiment> docker container ls -a
CONTAINER ID        IMAGE                 COMMAND             CREATED             STATUS              PORTS                    NAMES
92e3bbcd36aa        jackburdick/curtime   "go-wrapper run"    About an hour ago   Up About an hour    0.0.0.0:8000->8080/tcp   condescending_aryabhata
```

### Confirm
- Open [localhost:8000](http://localhost:8000/) in your browser

### Stop
1. `docker container stop 92e3`
    - where `92e3` is the container name from above

### Clean up
1. `docker container rm 92e3`
    - where `92e3` is the container name from above
2. confirm clean up with a `docker container ls -a`

#### (Optional) Remove the image
- `docker container ls -a`
```
PS C:\Users\JackB\Documents\GitHub\go_docker_app_experiment> docker image ls -a
REPOSITORY                 TAG                 IMAGE ID            CREATED             SIZE
jackburdick/curtime        latest              b0b97d6f4e41        About an hour ago   704MB
```
- this can be removed with `docker image rm b0b97d6f4e41`
- confirm the removal with `docker image ls -a`
```
PS C:\Users\JackB\Documents\GitHub\go_docker_app_experiment> docker image ls -a
REPOSITORY                 TAG                 IMAGE ID            CREATED             SIZE
```

### Note:
The `--rm` flag "Automatically remove[s] the container when it exits", but the image still exists.

### Resources
- [dockerize-simple-go-app](http://www.nikola-breznjak.com/blog/go/dockerize-simple-go-app/)