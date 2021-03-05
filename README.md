# pmm-api-tests

[![Build Status](https://travis-ci.com/Percona-Lab/pmm-api-tests.svg?branch=master)](https://travis-ci.com/Percona-Lab/pmm-api-tests)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2FPercona-Lab%2Fpmm-api-tests.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2FPercona-Lab%2Fpmm-api-tests?ref=badge_shield)

API tests for PMM 2.x

# Setup Instructions

Make sure you have Go 1.13.x installed on your systems, execute the following steps
to setup API-tests in your local systems.

1. Fetch the Repo: `go get -u -v github.com/Percona-Lab/pmm-api-tests`
2. Navigate to the tests root folder: `cd ~/go/src/github.com/Percona-Lab/pmm-api-tests`

# Usage

Run the tests using the following command:

```
go test ./... -pmm.server-url **pmm-server-url** -v
```

where `pmm-server-url` should be pointing to pmm-server.

# Docker

Build Docker image using the following command:

```
docker build -t IMAGENAME .
```

Run Docker container using the following command:

```
docker run -e PMM_SERVER_URL=**pmm-server-url** IMAGENAME
```

where `PMM_SERVER_URL` should be pointing to pmm-server.

If pmm-server located locally:

- Use --network=host while running docker container.
- Use non-secure url if you use self generated certificate.
