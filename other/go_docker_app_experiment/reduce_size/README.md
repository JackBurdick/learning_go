# Reducing the size Docker images for golang apps
### Overview
Build (small) Docker images for a golang app.

### Approach
1. Use default image
2. Use alpine
3. Multistage build

### Results
#### CurrentTime
Method | Size (MB) | Directory | Dockerfile
:--- | :---: | :---: | :---: |
standard | **704** | [link](./currentTime/1_standard) | [link](./currentTime/1_standard/Dockerfile) |
alpine | **263** | [link](./currentTime/2_alpine) | [link](./currentTime/2_alpine/Dockerfile) |
multi-stage | **9.85** | [link](./currentTime/3_multiStage) | [link](./currentTime/3_multiStage/Dockerfile) |

### Resources
- [dockerize-simple-go-app](http://www.nikola-breznjak.com/blog/go/dockerize-simple-go-app/)
- [create-the-smallest-possible-docker-container](http://blog.xebia.com/create-the-smallest-possible-docker-container/)
- [building-minimal-docker-containers-for-go-applications](https://blog.codeship.com/building-minimal-docker-containers-for-go-applications/)
- [multi-stage-docker-builds-for-creating-tiny-go-images](https://medium.com/travis-on-docker/multi-stage-docker-builds-for-creating-tiny-go-images-e0e1867efe5a)