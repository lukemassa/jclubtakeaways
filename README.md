# jclubtakeaways

## Setup

### Production

This repo is located in https://github.com/lukemassa/jclubtakeaways. A github action builds the page into docs/, uploads as an artifact and deploys to github pages (https://github.com/lukemassa/jclubtakeaways/settings/pages).

The page is then visible at https://jclubtakeaways.com

## Development

High level into "main" to deploy prod.

### Propose a change

TODO: Write up docs for local development
- Probably just run like go run cmd/local/main.go
- If you have a secret in env, that part will work, otherwise warn the user it's readonly
