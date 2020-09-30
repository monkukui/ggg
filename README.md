# ggg(go GRAPH × GRAPH)
`ggg` is a CLI tool for visualizing graph

`ggg` は、グラフ理論可視化サイト「GRAPH × GRAPH」をターミナル上で起動するための CLI ツールです。

**注意**
このツールはβ版です。バグがある場合があります。

## Requirement

- `go1.15.1`

## Install
```
go get -u github.com/monkukui/ggg
```

## Usage

各種オプションで、
- 有向・無向
- 重み付き・重みなし
- 1-indexed・0-indexes

を選択できます。

デフォルトでは、重みなし無向グラフを扱います。

詳細は
```
ggg --help
```
で確認してください。

### Start
```
> ggg
```
で起動します。

### Copy & Paste
```
4 6
1 2
2 3
3 4
1 3
2 4
1 4
```
をコピーし、ターミナルにペーストします。

### Visualize
グラフが可視化されます。


