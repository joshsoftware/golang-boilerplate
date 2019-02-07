## Golang Boilerplate
Golang boilerplate to kickstart any go api project

### Setup

1. Run the below command to copy sample config file (application.yml.sample) to application.yml. Please do not forgot to update the DB env variables in it according to your local setup.
```
$ make copy-config
```

2. Run the below commands to install all dependencies of the project.
```
$ make build
```

3. Run the below command to create db and run migrations on it
```
$ make db.setup
```

### Testing

Run test locally
```
$ make test
```
Generate coverage report
```
$ make test-coverage
```
