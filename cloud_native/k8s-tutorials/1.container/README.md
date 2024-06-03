# Docker Container

    # build images
    docker build . -t xxx/hellok8s:v1

    # start container
    docker run -p 3000:3000 --name hellok8s -d xxx/hellok8s:v1

    # push to docker hub
    docker push xxx/hellok8s:v1
