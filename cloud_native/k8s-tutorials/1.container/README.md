# Docker Container

    # build images
    docker build . -t colorfulgray0/hellok8s:v1
    docker push colorfulgray0/hellok8s:v1   # push是因为默认拉远程镜像

    # start container
    docker run -p 3000:3000 --name hellok8s -d colorfulgray0/hellok8s:v1
