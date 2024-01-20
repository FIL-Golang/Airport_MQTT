#!/bin/bash

if [ -z "$REGISTRY" ]; then
    echo "no registry set use default"
    REGISTRY="docker.io"
fi

if [ -z "$REPOSITORY" ]; then
    echo "no image base name set use default"
    REPOSITORY="airport-mqtt"
fi

registryUrl="$REGISTRY/$REPOSITORY/"
#lower name if not already
registryUrl=$(echo "$registryUrl" | tr '[:upper:]' '[:lower:]')

# For each executable
for dir in cmd/*/
do
    name=$(basename $dir)
    echo "Building $name"

    #lower name if not alread
    name=$(echo "$name" | tr '[:upper:]' '[:lower:]')

    # Get the location of the file containing the main function
    mainLocation="$dir""$name".go
    echo "$mainLocation"

    # Get the location of the Dockerfile
    dockerfileLocation="build/default.Dockerfile"
    if [ -d "build/$name" ]; then
        echo "Using build/$name/Dockerfile"
        dockerfileLocation="build/$name/Dockerfile"
    fi

    docker build -f "$dockerfileLocation" \
            -t "$registryUrl""$name":latest \
            --build-arg MAIN_PATH="$mainLocation" \
            .

    docker push "$registryUrl""$name":latest
done

exit 0
