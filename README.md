# Blog Aggregator (Gator)

## Prerequisites

- Go v1.24>, look [here](https://go.dev/doc/install) for installation.
- Postgres v15>, look [here](https://www.postgresql.org/download/) for installation.

## Setup

1. Get the repo

You can install the cli with go install:
```bash
go install github.com/e-mar404/blog-aggregator
```

2. Set up config file

The CLI needs a config file at `~/.gatorconfig.json` and needs the following:

```json
{"db_url": "postgres://[user_name]:@localhost:5432/gator?sslmode=disable"}
```

Note: `[user_name]` will be substituted for your username and in case you have a
specific port for your postgres installation then you replace 5432 for that
custom port.

3. Usage

The purpose of the CLI is to have one instance running on the background and
fetches blogs/any rss feeds and then you can browse the newest content.


To start you need to register a user:

```bash
gator register emar
```

You then can add feeds like so:

```bash
gator addfeed "Haker News" "https://news.ycombinator.com/rss"
```

And check all the feeds that you are following:

```bash
gator following
```

The process that will will run in the background is `gator agg <time-duration>`.
The time duration is in the format of 1s, 5m, 2hr. The service will fetch the
feeds and save the posts every time-duration.


To look at the newest posts you can use the browse command that defaults to 2
posts but you can put an optional parameter for how many posts you want to see.
This example you can see 10 posts (if there are that many to view).

```bash
gator browse 10
```
