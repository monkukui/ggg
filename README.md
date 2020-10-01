# ggg(go GRAPH × GRAPH)

<img width="264" alt="スクリーンショット 2020-09-30 22 50 16" src="https://user-images.githubusercontent.com/47474057/94693994-55c29200-036f-11eb-827c-1aa72a166f64.png">


`ggg` is a CLI tool for visualizing graph

`ggg` は、グラフ理論可視化サイト「GRAPH × GRAPH」をターミナル上で起動するための CLI ツールです。

**注意**

## Requirement

- `go1.15.1`

## Install（Update）
```
> go get -u github.com/monkukui/ggg
```

```
ggg --help
```
を叩いて、起動すれば成功です。

## Usage（簡易版）

各種オプションで、
- 有向・無向
- 重み付き・重みなし
- 1-indexed・0-indexed

を選択できます。

デフォルトでは、重みなし無向グラフで 1-indexed を扱います。

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

<img width="301" alt="スクリーンショット 2020-09-30 22 49 33" src="https://user-images.githubusercontent.com/47474057/94693893-3b88b400-036f-11eb-87ef-a64ad4c00f63.png">



## Usage（詳細）

### 起動方法
```
❯ ggg
```
で起動します。

### ヘルプコマンド

```
❯ ggg --help
```
でヘルプを表示します。

### オプション
Command line flag syntax を用いて、グラフの形式を変更できます。

|  Flags  |  Description  | default | Usage |
| ---- | ---- | ---- | ---- |
|  `-i, --indexed`  |  1-indexed であることを指定します。  | true | `-i 0`, `-i 1`, `--indexed=0`, `--indexed=1` |
|  `-d, --directed`  |  有向でグラフであることを指定します。  | false | `-d` |
|  `-w, --weighted`  |  重み付きグラフであることを指定します。 | false | `-w` |


### Case Study

#### 1-indexed、重みなし無向グラフ

```
❯ ggg

Options:
   indexed:  1
   directed:  false
   weighted:  false

Example:

   [1] ----- [2] ----- [3]

Format:
   3 2    (the number of nodes, the number of edges)
   1 2    (edge informations)
   2 3

Please input your graph.
>>> 
3 3
1 2
2 3
3 1
```

<img width="202" alt="スクリーンショット 2020-10-01 17 21 57" src="https://user-images.githubusercontent.com/47474057/94785469-a935ee00-040a-11eb-9d53-e4e71062fd2e.png">


#### 0-indexed、重みなし有向グラフ

```
❯ ggg --indexed 0 --directed

Options:
   indexed:  0
   directed:  true
   weighted:  false

Example:

   [0] ----> [1] ----> [2]

Format:
   3 2    (the number of nodes, the number of edges)
   0 1    (edge informations)
   1 2

Please input your graph.
>>> 
3 3
0 1
1 2
2 0
```

<img width="202" alt="スクリーンショット 2020-10-01 17 23 22" src="https://user-images.githubusercontent.com/47474057/94785633-e1d5c780-040a-11eb-97af-7e71ee0a64b8.png">

#### 1-indexed、重みあり有向グラフ

```
❯ ggg --weighted --directed

Options:
   indexed:  1
   directed:  true
   weighted:  true

Example:

   [1] --5-> [2] --7-> [3]

Format:
   3 2    (the number of nodes, the number of edges)
   1 2 5  (edge informations)
   2 3 7

Please input your graph.
>>> 
3 2
1 2 5
2 3 7
```

<img width="202" alt="スクリーンショット 2020-10-01 17 24 03" src="https://user-images.githubusercontent.com/47474057/94785659-eb5f2f80-040a-11eb-8a91-692c94972c1b.png">
