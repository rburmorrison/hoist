# Hoist

A CLI to explore private Docker Registries.

### Building

To build, run `go get -u github.com/rburmorrison/hoist/cmd/hoist`. The binary will be located in `$GOPATH/bin`.

### Usage

Hoist is a small front end for v2 Docker registries. The default registry that hoist sends requests to is `http://localhost:5000`. To change the registry that requests are being sent to, you can run the `hoist config` sub-command. You can see examples with `hoist config set --help`.

Images stored in a private Docker registry are stored as "repositories". To display all repositories in a registry, run `hoist repos`. To see the tags associated with a repository, run `hoist tags REPOSITORY`.
