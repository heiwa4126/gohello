# gohello

- go modulesを使ったプロジェクトの練習
- [GoReleaser](https://goreleaser.com/)の練習

内容自体は、hello worldに毛が生えたようなもの。
(少しだけ引数を取る、とかにしていく予定)


# プロジェクトの場所

- GitHub : [heiwa4126/gohello: GoReleaserの練習](https://github.com/heiwa4126/gohello)


# install

```
go get github.com/heiwa4126/gohello
```

# GitHub Actionsでビルドするメモ

1. git commit & push
2. git tag "v.0.0.0" (0.0.0は適切なバージョンに置き換え)
3. git push --tag
4. [Actions · heiwa4126/gohello](https://github.com/heiwa4126/gohello/actions)で進行結果を見る。
5. うまくいけば[Releases · heiwa4126/gohello](https://github.com/heiwa4126/gohello/releases)にchecksumやtarballが出来る。


# 参考

- [Go で書いた CLI ツールのリリースは GoReleaser と GitHub Actions で個人的には決まり | tellme.tokyo](https://tellme.tokyo/post/2020/02/04/release-go-cli-tool/)
- [GoReleaser](https://goreleaser.com/)
- [goreleaser+go-github-selfupdateでお手軽自動リリース&amp;アップデート - Qiita](https://qiita.com/mpppk/items/ab328356ca14938a1208)

# 履歴

## v0.0.1

1st release

## v0.0.2
- バージョン・リビジョンが出ない →直した
- Windowsがビルドされない →直した
- Mac用の386版はいらない →直した

## v0.0.3
- モジュールの例と、ユニットテストを追加してみた
