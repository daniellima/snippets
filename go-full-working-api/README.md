# Counting API

## TODO

- ~~Creating a routing layer that supports~~
  - ~~GET and POST being handled by different handlers (use configured middleware?)~~
  - ~~Using native handler~~
  - ~~match routes exactly (must test that the current one does not)~~
- ~~Handle generic errors (500) in a sane way~~
- ~~Find a good log library to output log in JSON easily~~
  - ~~Listening and Exiting should be in JSON!~~
- ~~Add environment variables (generic config handling)~~
- ~~Create an abstraction for Redis~~
- ~~Send metrics to Jaeger~~
  - ~~Send metrics correctly~~
  - ~~Send metrics using a middleware~~
- ~~Share the trace id between traces and logs~~
- ~~Integrate with Datadog to see metrics there~~
- ~~Add a TLS certificate~~
- ~~Call handlers with their Middlewares active for good end-to-end tests~~
- ~~Change the pipeline to run tests~~
- ~~Add a DI container to handle dependencies~~
- use multistage build
- build only when source code files change
- add validation to environment variables and request parameters?
- start prometheus handler on a separated port
- try to move infra files to the infrastructure repo
