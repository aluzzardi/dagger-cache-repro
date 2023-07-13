# cache repro

## usage

Run the concurrent repro in loop:

```sh
_EXPERIMENTAL_DAGGER_JOURNAL=journal.log go run . repro
```

Check suspicious activity:

```sh
go run . check journal.log
```
