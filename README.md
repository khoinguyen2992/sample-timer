# sample-timer

There are some sample task timers

## Getting Started

1. Install [docker](https://www.docker.com/products/docker#/mac)
1. Clone this repository
1. `cd sample-timer && docker-compose up -d --build`
1. `docker logs -f --tail=50 <container_id>` to see logs

## Problem

We have to timer tasks to run later

## Approach #1

Using cronjob and mongodb.

### Pros

- Keep tasks when server down.

### Cons

- Kseep running every 1s
- Duplication processing if there are two or more instances

### Improvements

- Double check tasks in processing or not

## Approach #2

Using redis pubsub.

### Pros

- Run in demand
- No duplication processing

### Cons

- Tasks eliminated when server down

### Improvements

- Persist tasks to database