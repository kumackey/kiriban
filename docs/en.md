# kiriban

| **English** | [日本語](/docs/ja.md) |
|-------------|--------------------|

This is a GitHub Actions that will congratulate you when you reach a milestone number (kiriban) for your issue or pull
request number.
It targets Issues or Pull Requests with numbers greater than 10.
It is under active development, and we welcome your issues and pull requests!

## How to Use

Create a yml file like the following in `.github/workflows/`.

```yml:.github/workflows/kiriban.yml
# .github/workflows/kiriban.yml

name: kiriban

on:
  issues:
    types: [ opened ]
  pull_request:
    types: [ opened ]

jobs:
  kiriban:
    runs-on: ubuntu-latest
    steps:
      - uses: kumackey/kiriban@v1
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
```

### The Comment Provided

![The comment posted by the Action](/img/comment_en.png)

### with Options

| Key Name              | Description                        | Example    |
|-----------------------|------------------------------------|------------|
| LOCALE                | Language setting (default is 'en') | 'ja'       |
| USER_DEFINED_KIRIBANS | User-defined numbers to celebrate  | '101,4649' |

## What is a Kiriban?

https://fanlore.org/wiki/Kiriban