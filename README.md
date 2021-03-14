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

## Managing rules

### What are rules

Rules are how `gitlab-lint` knows what to look for on each processed project.

Take for an example the `Empty Repository` rule: its goal is to check if the
current project houses an empty repository.

### Creating new rules

Rules must implement the [Ruler][file.rules.ruler] interface and they must have at least the
following fields on their struct:

```go
type Ruler interface {
	Run(client *gitlab.Client, p *gitlab.Project) bool
	GetSlug() string
	GetLevel() string
}
```
```go
type MyAwesomeRule struct {
	Description string `json:"description"`
	ID          string `json:"ruleId"`
	Level       string `json:"level"`
	Name        string `json:"name"`
}
```

A good practice is to also have a `NewMyAwesomeRule()` function that returns an
instatiaded rule's struct.

Notice that `ID` and `GetSlug()` should return a unique value to identify your
rule.

Also, there's already a couple of pre-determined Levels on
[levels][file.rules.levels]. We must use those instead of random strings.

### Registering rules

After creating the rule itself, we must register it so it's considered when we
parse the projects. In order to do it, we should just add it to the `init()`
function on [my_registry][file.rules.my_registry], just like so:

```go
func init() {
	MyRegistry.AddRule(NewMyAwesomeRule())
	...
}
```

Then, you should be able to save (or recompile, if running a binary) and check
your new rule being returned by running a GET to `/api/v1/rules`:

```json
[
  {
    "description": "",
    "ruleId": "my-awesome-rule",
    "level": "error",
    "name": "My Awesome Rule"
  }
]
```

## Contribute

Fork the repository and send your pull-requests.


[personal_access_tokens]: https://docs.gitlab.com/ce/user/profile/personal_access_tokens.html
[file.rules.ruler]: ./rules/ruler.go
[file.rules.levels]: ./rules/levels.go
[file.rules.my_registry]: ./rules/my_registry.go
