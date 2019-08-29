
# `http-redirect` docker image

This is helper Docker image, which redirects everything it receives. 

## Environment variables:

 * `REDIRECT_TARGET` - (required). target
 * `REDIRECT_PERMANENT` - (optional). when defined, redirects are permanent


## Usage

    docker run --env REDIRECT_TARGET=http://www.google.com/ -p 8080:80 wftech/http-redirect



# License

Apache || WTFPL
