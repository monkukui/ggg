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

|  Flags  |  Description  | default |
| ---- | ---- | ---- |
|  `-i, --indexed`  |  1-indexed であることを指定します。  | true |
|  `-d, --directed`  |  有向でグラフであることを指定します。  | false |
|  `-w, --weighted`  |  重み付きグラフであることを指定します。 | false |


### Case Study

#### 1-indexed、重みなし無向グラフ

```
❯ ggg

（省略）

please input your graph below...
>>> 
3 2
1 2
2 3
```

#### 0-indexed、重みなし有向グラフ

```
❯ ggg --directed

（省略）

please input your graph below...
>>> 
3 2
1 2
2 3
```

#### 1-indexed、重みあり有向グラフ

```
❯ ggg --directed --weighted

（省略）

please input your graph below...
>>> 
3 2
1 2 3
2 3 5
```

