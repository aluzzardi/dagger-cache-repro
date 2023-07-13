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

## what it does

- `repro()` is called every 5 minutes, each time spawning 10 goroutines with a random sleep of up to a second, to trigger concurrency

- each repro pulls alpine, execs a `sleep 5` (which is a cached operation), then echo's a UUID (therefore not cached)

- `check` is a suspicious checking tool. it's pretty dumb, scans the journal and sees if a vertex started timestamp stands out (something that's at least 10 minutes older than other vertices)
