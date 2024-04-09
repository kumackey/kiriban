# kiriban

PR番号でキリ番を引いたら祝ってくれるGitHub Actionsです。
100以上のPRからが対象です。

## 使い方

.github/workflowsに以下のようなymlファイルを作成してください。

```yml:.github/workflows/kiriban.yml
name: kiriban

on:
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

## キリ番とは？

https://ja.wikipedia.org/wiki/%E3%82%AD%E3%83%AA%E3%83%90%E3%83%B3