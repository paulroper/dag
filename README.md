# dag

A tool for building a DAG for a repo containing multiple apps. [Based on an earlier Node implementation](https://github.com/paulroper/dotnet-monorepo/tree/main/tools/dag).

Can be easily passed into something like `docker buildx` to only build what's changed in the repo ([example](https://github.com/paulroper/dotnet-monorepo/blob/main/build.ps1#L20)).

## Setup

📓 Make sure you have Docker

- Add a `deps.json` to each module you want to build in your Git repository
- Run `./scripts/build.sh`
- Run `docker run dag:latest --repository <PATH_TO_REPOSITORY>` to generate a DAG
