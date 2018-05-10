# Reduce Image Size using the Alpine Image
### Final image stats
```
PS C:\Users\JackB\Documents\GitHub\go_docker_app_experiment\simple_smallContainer> docker image ls -a
REPOSITORY                 TAG                 IMAGE ID            CREATED             SIZE
jackburdick/smaller        latest              5be988dc6835        6 minutes ago       263MB
```
## Creation
### Build
- $`docker image build -t jackburdick/smaller .`

### Run
- $`docker container run -d --rm -p 8000:8080 jackburdick/smaller`
- Confirm;
```
PS C:\Users\JackB\Documents\GitHub\go_docker_app_experiment\simple_smallContainer> docker container ls
CONTAINER ID        IMAGE                 COMMAND                CREATED              STATUS              PORTS                    NAMES
281e1c2bcd37        jackburdick/smaller   "/bin/sh -c ./goapp"   About a minute ago   Up About a minute   0.0.0.0:8000->8080/tcp   ecstatic_golick
```

### Confirm
- Open [localhost:8000](http://localhost:8000/) in your browser

### Stop
- `docker container stop 281`
    - where `281` is your container id

### Note:
The `--rm` flag "Automatically remove[s] the container when it exits", but the image still exists.

## Resources
-  [create-the-smallest-possible-docker-container](http://blog.xebia.com/create-the-smallest-possible-docker-container/)
- [building-minimal-docker-containers-for-go-applications](https://blog.codeship.com/building-minimal-docker-containers-for-go-applications/)
- [multi-stage-docker-builds-for-creating-tiny-go-images](https://medium.com/travis-on-docker/multi-stage-docker-builds-for-creating-tiny-go-images-e0e1867efe5a)