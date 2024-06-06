# Docker Container

    # build images
    docker build . -t colorfulgray0/hellok8s:v1

    # start container
    docker run -p 3000:3000 --name hellok8s -d colorfulgray0/hellok8s:v1

    # push to docker hub
    docker push colorfulgray0/hellok8s:v1
