# jclubtakeaways

## Setup

### Production

This repo is located in https://github.com/lukemassa/jclubtakeaways. A github action builds the page into docs/, uploads as an artifact and deploys to github pages (https://github.com/lukemassa/jclubtakeaways/settings/pages).

The page is then visible at https://jclubtakeaways.com

### Development

The repo is mirrored in lukefrederickmassa@gmail.com's AWS account, then AWS Amplify is used to deploy the "development" branch.

The page is then visible at https://development.d7s2t12hepi6g.amplifyapp.com

## Development

High level, submit PRs/merge directly to "development" to develop, then merge development into "master" to deploy prod.

### Propose a change

1. Go to the development branch in git (https://github.com/lukemassa/jclubtakeaways/tree/development)
2. Create a PR with target branch "development"
3. Have Luke review and merge
4. Visit https://development.d7s2t12hepi6g.amplifyapp.com to view change (should be live in a few minutes)
5. Luke merges development into master

### Directly update development

Same as above, but commit directly to development, instead of creating an MR.

After this (and a few minutes) your changes will be live immediately on development

Then Luke will merge development into master
