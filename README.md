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

1. Keep tasks when server down.

### Cons

1. Kseep running every 1s
1. Duplication processing if there are two or more instances

### Improvements

1. Double check tasks in processing or not

## Approach #2

Using redis pubsub.

### Pros

1. Run in demand
1. No duplication processing

### Cons

1. Tasks eliminated when server down

### Improvements

1. Persist tasks to database