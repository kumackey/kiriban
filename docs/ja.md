# kiriban

| [English](/docs/en.md) | **日本語** |
|------------------------|---------|

IssueやPull Requestの番号で、キリ番を引いたら祝ってくれるGitHub Actionsです。<br>
10以上のIssueやPull Requestの番号からが対象です。<br>
鋭意開発中であり、IssueやPull Requestをお待ちしています！

## 使い方

.github/workflows/に以下のようなymlファイルを作成してください。

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
        with:
          LOCALE: 'ja'
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}

```

### なされるコメント

![Actionの結果、コメントされる](/img/comment_ja.png)

### withのオプション

| キー名                   | 説明              | 例          |
|-----------------------|-----------------|------------|
| LOCALE                | 言語設定（デフォルトではen） | 'ja'       |
| USER_DEFINED_KIRIBANS | ユーザ定義で祝う番号      | '101,4649' |

## キリ番とは？

[https://ja.wikipedia.org/wiki/キリバン](https://ja.wikipedia.org/wiki/%E3%82%AD%E3%83%AA%E3%83%90%E3%83%B3)
