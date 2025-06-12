# Users microservice

---

## Table of contents

- [Users microservice](#users-microservice)
  - [Table of contents](#table-of-contents)
  - [Environment variables](#environment-variables)
    - [Generating a working `.env` file](#generating-a-working-env-file)
  - [Modes](#modes)
  - [Database migrations](#database-migrations)
    - [Create a new migration](#create-a-new-migration)
    - [Apply migrations](#apply-migrations)
    - [Rollback migrations](#rollback-migrations)
  - [Building the project](#building-the-project)
    - [Local build](#local-build)
      - [1. Prerequisites for local build](#1-prerequisites-for-local-build)
      - [2. Local build instructions](#2-local-build-instructions)
    - [Building with docker](#building-with-docker)
    - [Partially clean enviroment](#partially-clean-enviroment)
    - [Fully clean environment](#fully-clean-environment)
      - [Building modes in docker](#building-modes-in-docker)
  - [Running the project](#running-the-project)
    - [Local execution](#local-execution)
    - [Using Docker to run](#using-docker-to-run)
      - [Running without rebuilding](#running-without-rebuilding)
      - [Running while rebuilding images](#running-while-rebuilding-images)

---

## Environment variables

Many configurations can be set configuring a `.env` file. If you want
to see and example about it, please go to [this file](./.env.example).

> [!NOTE]
> Once you have [created](#generating-a-working-env-file) your `.env`
> file, you can run any docker command for
> [building](#building-with-docker) or [running](#using-docker-to-run)
> the project without having to do any extra work.

### Generating a working `.env` file

To generate a file based on a working template run this command:

```sh
cp .env.example .env
```

## Modes

> [!NOTE]
> Modes are controlled by the
> [environment variable](#environment-variables) `mu_users_ms_MODE`

The api can run in one these `MODE`s:

- `release`: Production mode with optimized performance and minimal logging
- `debug`: Development mode with detailed logging and error traces
- `test`: Testing mode, useful to implement mock services and/or
test-specific configurations

## Database migrations

The project uses [migrate](https://github.com/golang-migrate/migrate/). Here
are some useful commands for working with migrations:

### Create a new migration

To create a new migration, run the following command:

```sh
migrate create -ext sql -dir db/migrations -seq <migration_name>
```

Replace `<migration_name>` with a descriptive name for your migration. This
command will create a new migration file in the `db/migrations` directory with
the specified name. The `-ext sql` flag specifies that the migration files
should have the `.sql` extension, and the `-seq` flag specifies that the
migration should be created with a sequential number.

### Apply migrations

To apply all pending migrations, run the following command:

```sh
migrate -path db/migrations -database "postgres://<username>:<password>@<host>:<port>/<database>?sslmode=disable" -verbose up
```

### Rollback migrations

To rollback the last applied migration, run the following command:

```sh
migrate -path db/migrations -database "postgres://<username>:<password>@<host>:<port>/<database>?sslmode=disable" down
```

## Building the project

The recommended way to build the project is to use Docker
[(click here)](#building-with-docker), as it ensures that the build environment
is consistent and eliminates any potential issues with dependencies on your
local machine. However, if you prefer to build the project locally, you can do so
by following [these instructions](#local-build)

### Local build

#### 1. Prerequisites for local build

> [!TIP]
> You may not need to install all of the dependencies listed below if you are
> using Docker to build the project.

- Go 1.24.3 or later
- sqlc (for generating SQL code)

1. Install Go from the [official website](https://golang.org/dl/).
2. Install sqlc by running the following command. Or
visit [this page](https://docs.sqlc.dev/en/latest/overview/install.html#installing-sqlc):

    ```sh
    go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
    ```

3. Install the required Go modules by running the following command in the
project directory:

    ```sh
    go mod tidy
    ```

4. Generate the required SQLC modules by running the following command in the
project directory. Make sure you have
[applied all your migrations](#apply-migrations):

    ```sh
    sqlc generate
    ```

   This command will generate the necessary go code based on the SQL files
   in the `db` directory. The generated code will be placed in the `./repository`
   directory.

#### 2. Local build instructions

> [!IMPORTANT]
> In order to build and run the project, make sure you meet all the
> [prerequisites](#1-prerequisites-for-local-build).

After that you can build and run the project. Like this:

```sh
go build -v -o main .
```

This command compiles the current directory and all its subdirectories, and produces
an executable file named `main` in the current directory. After running this
command, you may [run the project](#local-execution) using the generated executable.

---

### Building with docker

> [!IMPORTANT]
> Make sure your `.dockerignore` file is set up correctly to exclude any
> unnecessary files. Like so:

```sh
cat .gitignore .prodignore > .dockerignore
```

> [!TIP]
> You can modify the behavior of the software system like
> ports, hostnames and more by using a `.env` file. Please refer to
> [this section](#environment-variables) for more information.

To build the whole project using Docker, you can use the provided
[docker-compose.yml](./docker-compose.yml). This will take care of everything related
to the building and running processes, including each needed service to run the
project locally. You can start from a
[partially](#partially-clean-enviroment) clean environment or from a
[clean environmentment](#fully-clean-environment).

### Partially clean enviroment

> [!TIP]
> If you want to have a _almost_ clean build you need to stop
> and remove containers, networks by running:

```sh
docker compose down --remove-orphans
```

### Fully clean environment

> [!WARNING]
> The following command gives you a clean slate to start from, but it
> remove the volumes too. So any data that you may have, it will be
> removed as well.

```sh
docker compose down --remove-orphans --volumes
```

#### Building modes in docker

By default the API will be built in `release` [mode](#modes) and is built like this:

```sh
docker compose build
```

In case you need to build it in other [`<mode>`](#modes). Run the
following command replacing `<mode>` with the one you need. (Or use a
[`.env` file](#environment-variables))

```sh
docker compose build --build-arg MODE=<mode>
```

This command will create the docker images the needed images for each service
based on the [docker compose](docker-compose.yml) configuration file. And you can
run the Docker image using the provided [run](#using-docker-to-run) command.

## Running the project

### Local execution

> [!IMPORTANT]
> In order to build and run the project, make sure you meet all the
> [prerequisites](#1-prerequisites-for-local-build).
<!-- This fixes renderization issues -->
> [!NOTE]
> These commands will not run your database or any other dependency/service
> that you may need

If you have built the project using the [build](#2-local-build-instructions)
command, you can run the generated executable:

```sh
./main
```

To run the project locally without building, you can use the following command:

```sh
go run main.go
```

### Using Docker to run

> [!TIP]
> You can modify the behavior of the software system like
> ports, hostnames and more by using a `.env` file. Please refer to
> [this section](#environment-variables) for more information.

We personally recommend [this option](#running-while-rebuilding-images).

#### Running without rebuilding

It doesn't contemplate the current state of your files but the most recent built
images. Make sure it is in the correct [`mode`](#modes)

```sh
docker compose up
```

#### Running while rebuilding images

This will allow you to start from the current version of your repository.
This means that if you modified any file included in Docker, a new image will be
built in the [default mode](#building-modes-in-docker) and run after the
building process has been completed

```sh
docker compose up --build
```
