log-request-body
==

Answers all HTTP requests with OK (200) and logs the request body to stdout.

When the `PORT` environment variable is specified, the application listens on this port. Otherwise, it uses port 80.


Usage
--

```sh
docker build . -t corneliusweig/log-request-body
docker run --rm -ti -p 8080:80 corneliusweig/log-request-body

# test with
curl -XPOST -d'Hello World!' localhost:8080
```

Deploy to CloudRun
--

[![Run on Google Cloud](https://deploy.cloud.run/button.svg)](https://deploy.cloud.run)
