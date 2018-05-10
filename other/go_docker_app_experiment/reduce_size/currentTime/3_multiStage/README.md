# Creating a smaller image
### Final image stats
```
PS C:\Users\JackB\Documents\GitHub\go_docker_app_experiment\simple_multiStage> docker image ls -a
REPOSITORY                 TAG                 IMAGE ID            CREATED              SIZE
jackburdick/multistage     latest              fba5963ab482        57 seconds ago       9.85MB
```
## Creation
### Build
- $`docker image build -t jackburdick/multistage .`

### Run
- $`docker container run -d --rm -p 8000:8080 jackburdick/multistage`
- Confirm;
```
PS C:\Users\JackB\Documents\GitHub\go_docker_app_experiment\simple_multiStage> docker container ls
CONTAINER ID        IMAGE                    COMMAND                CREATED             STATUS              PORTS                    NAMES
50c7c0d04c3c        jackburdick/multistage   "/bin/sh -c ./goapp"   29 seconds ago      Up 28 seconds       0.0.0.0:8000->8080/tcp   serene_shockley
```

### Confirm
- Open [localhost:8000](http://localhost:8000/) in your browser

### Stop
- `docker container stop 50c7`
    - where `50c7` is your container id

### Note:
The `--rm` flag "Automatically remove[s] the container when it exits", but the image still exists.

## Resources
- [multi-stage-docker-builds-for-creating-tiny-go-images](https://medium.com/travis-on-docker/multi-stage-docker-builds-for-creating-tiny-go-images-e0e1867efe5a)