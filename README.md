**THIS SDK IS IN EARLY DEVELOPMENT AND NOT MEANT FOR PUBLIC RELEASE, BREAKING CHANGES ARE GUARANTEED***

# NexGen Hyperstack Go SDK (gosdk)

Hyperstack Go SDK (gosdk) is a package that allows users to interact with the Infrahub API by Hyperstack. It is autogenerated from the available OpenAPI specifications and provides type-safe, easy-to-use methods for calling the API endpoints.

## Development

### Dependencies

The SDK uses a few dependencies which can be installed using the `task init` command. These dependencies include package(s) like `github.com/deepmap/oapi-codegen/cmd/oapi-codegen` that helps with the SDK auto-generation.

### Updating the SDK

The SDK pulls its latest `api.json` from the server by running `task pull-api`.

To auto-generate the SDK from the latest `api.json` file, use `task build`.

The SDK can be generated and tested by simply running `task`.

### Testing

To run tests for the SDK, use `task test`.

## Taskfile

The `Taskfile.yaml` is a task runner / simpler Make alternative that is written in Go. It is used to define and run routine development and build tasks. It's designed to be simple and easy to use, while still being flexible and powerful.

To run a task, use the `task` command followed by the task name. For example, to run the `init` task, you would use `task init`.

## GitHub Actions

To automate the build and test process, the repository is integrated with GitHub Actions.

The workflows are designed to perform the following activities whenever a push or pull request is made to the `main` branch:

- Set up the Go environment
- Install the necessary dependencies
- Pull the latest OpenAPI specifications (api.json)
- Generate the SDK using the latest specifications
- Run the unit tests

These workflows ensure that any new changes or updates to the API specifications are reflected in the Go SDK with each push or pull request, thus keeping the SDK up-to-date.

## About Hyperstack Cloud Go SDK (gosdk)

The Hyperstack Cloud SDK Generator leverages the power of automatic generation to provide developers with easy-to-use methods to interact with the latest NexGen Cloud API. It adapts to the most current API specifications and ensures that users always have the latest, most accurate SDK for use.

Just run `task` and get your SDK ready for the Hyperstack Cloud API!

_Built on Hyperstack Cloud, built for you._

## TL;DR :
1. Install and generate :

   brew install go && echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.zshrc && source ~/.zshrc && task

2.[optional] Enable full test client and re-run tests to be sure
export NEXGEN_CLOUD_API_KEY = <your nexgen API KEY> && task test
