# jclubtakeaways

## Setup

### UI

A GitHub Action builds the page in `src/` into `docs/` and deploys it via GitHub Pages.

The page is then visible at https://jclubtakeaways.com

### API

To call into the google sheets API, it needs a token. There's a lambda + URL deployed to an aws account 917497682277: https://e7s2nudakzfoy5cio723wzlr440pfkfx.lambda-url.us-east-1.on.aws. The infra is managed by terraform/, and the code is built by a different github workflow and uploaded.

## Development

- Run `go run cmd/server/main.go`
- If you have a secret in `WEB_CLIENT_KEY`, that part will work, otherwise warn the user it's readonly
  - Hint: `WEB_CLIENT_KEY="$(gpg -d  ~/creds/webclientaccount.key.asc)" go run cmd/server/main.go`
