# faasify - A framework for building a simple and blazingly fast FaaS server from scratch

## Get the sources

    git clone https://github.com/andygeiss/faasify.git
    mv faasify YOUR_NAME
    cd YOUR_NAME

## Build the server

    ./scripts/build.sh

## Run the server

    FAASIFY_ADDESS=":3000" FAASIFY_TOKEN="YOUR_TOKEN" ./build/serve-http

## Call a function

    curl -H "Authorization: Bearer YOUR_TOKEN" http://127.0.0.1:3000/function/status

## Display the function statistics

    curl http://127.0.0.1:3000/stats

## Display the embedded web content

    curl http://127.0.0.1:3000/static/

