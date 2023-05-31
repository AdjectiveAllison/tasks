# tasks

* I'm going to eventually have my public tasks use this repo, for now it's just code.
* This is powered by cloudflare workers and the d1 database(uses sqlite).
* This is possible with [SyumAI's](https://github.com/syumai) **phenomenal** [golang workers package](https://github.com/syumai/workers).
* Currently TinyGO is used to compile to reduce the web assembly package size. TinyGO has issues with the relfection package currently which is why easyjson is used instead of encoding/json. Version 28 of tinygo(coming soon) will make golang in web assembly even more powerful. 
* I'm a noob, but this tech is cool :)
* This was built [Live on twitch!](https://twitch.tv/AdjectiveAllison)

# Work in Progress

## Example

* [tasks.adjective.workers.dev](https://tasks.adjective.workers.dev)


### List tasks

Depending on the `Accept` header, the response can be returned in JSON or HTML format.

#### For JSON Response:
```bash
$ curl 'https://tasks.adjective.workers.dev' -H "Accept: application/json"
```
Sample response:
```json
{
  "tasks": [
    {
      "id": 1,
      "title": "Improve task writing process",
      "description": "Need to develop a more efficient method for adding tasks to the database.",
      "links": [
        "https://tasks.adjective.workers.dev",
        "https://github.com/AdjectiveAllison/tasks",
        "https://AdjectiveAllison.com"
      ],
      "updatedAt": 1677383758,
      "completed": false
    },
    {
      "id": 2,
      "title": "Write better code",
      "description": "Who wrote this?!",
      "links": ["https://github.com/AdjectiveAllison"],
      "updatedAt": 1677175342,
      "completed": false
    }
  ]
}
```

For HTML Response:

[Use your browser!](https://tasks.adjective.workers.dev)(or just don't pass an `Accept` header)

## Inspiration

[SyumAI's d1-blog-server example](https://github.com/syumai/workers/tree/main/_examples/d1-blog-server) was used to bootstrap this.

That means that I was lazy and have the same problem that the example has at the time of me using it(not using `wrangler d1 migraions` but instead utilizing raw SQL.)

## Development

### Requirements

This project requires these tools to be installed globally.

* wrangler
* tinygo
* [easyjson](https://github.com/mailru/easyjson)
  - `go install github.com/mailru/easyjson/...@latest`

### Commands

```
# development
make init-db-preview # initialize preview DB (remove all rows)
make generate        # generate easyjson models
make dev             # run dev server
make build           # build Go Wasm binary

# production
make init-db # initialize production DB (This also currently executes tasks.sql)
make deploy # deploy worker
```

* Notice: This example uses raw SQL commands to initialize the DB for simplicity, but in general you should use `wrangler d1 migraions` for your application.


**TODO**: Figure out how d1 migrations work and implement them here.

**TODO**: Get local script to execute and add tasks to the database automatically.
