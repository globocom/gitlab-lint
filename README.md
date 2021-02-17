# gitlab-lint API and collector

An open source gitlab linting utility

## Frontend

https://github.com/globocom/gitlab-lint-react

## How to install

### Install dependencies

* Golang
* Docker
* pre-commit
* golangci-lint

### Dev dependencies:

```bash
make setup
```

## Create Gitlab access token

You should create a personal access token:

1. Sign in to GitLab.
1. In the top-right corner, select your avatar.
1. Select Edit profile.
1. In the left sidebar, select Access Tokens.
1. Choose a name and optional expiry date for the token.
1. Choose   the `read_api` scope.
1. Select Create personal access token.
1. Save the personal access token somewhere safe. If you navigate away or
   refresh your page, and you did not save the token, you must create a new
   one.

Set the following environment variable with your token:

More info at [Personal access tokens][personal_access_tokens]

```bash
export GITLAB_TOKEN="token"
```

## Run it

```bash
make run-docker
```

## Collecting the data

```bash
make collector
```

## Contribute

Fork the repository and send your pull-requests.


[personal_access_tokens]: https://docs.gitlab.com/ce/user/profile/personal_access_tokens.html
