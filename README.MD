## Simple DevX trial for FF usage

`main.go` - entry point of the program
`ff_sdk/sdk_mock.go` - A simple, type safe flagd FF generation tool
`common/flags.go` - Flag definitions of the app. Type safe, reusable (both program and generation)
`non_build/generator.go` - Invoke FF generation using ff_sdk

### How to

1. Generate FFs: `go generate ./...`. This generates `ff.json`
2. Source `ff.json` to
   flagd: `docker run -p 8013:8013 -v $(pwd)/:/etc/flagd/ -it ghcr.io/open-feature/flagd:latest start --uri /etc/flagd/ff.json`
3. Run app to see outputs: `go run main.go`

## Variants

Two ways to change the output

1. Changing `ff.json` : Change and re-run the app
2. Changing `flags.go` : For example default variant change