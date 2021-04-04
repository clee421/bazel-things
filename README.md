# How to do things

## Use gazelle to create the build files

```
./bazel run :gazelle
```

## Use gazelle to update build files

```
./bazel run //:gazelle -- update-repos -from_file=go.mod
```

## Executing a the hello command

```
./bazel run //go/cmd/hello
```