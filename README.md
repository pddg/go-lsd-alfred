# go-lsd-alfred

[![Build Status](https://travis-ci.org/pddg/go-lsd-alfred.svg?branch=master)](https://travis-ci.org/pddg/go-lsd-alfred)

Alfred 3用のWorkfrowで用いる簡単なCLIツール．  

ライフサイエンス辞書を検索する．

## Usage

### 検索方法

* 前方一致

```bash
lsd begin "..."
```

* 後方一致

```bash
lsd end "..."
```

* 部分一致

```bash
lsd in "..."
```

* 完全一致

```bash
lsd eq "..."
```

* シソーラス： 未実装
* コーパス： 未実装

### 検索結果

* ブラウザで開く： `Shift` + `Enter`
* クリップボードにコピー： `Cmd` + `Enter`
* 検索結果の内容を元に再検索： `Fn` + `Enter`
* 検索結果の内容を元にSpotlight検索： `Alt` + `Enter`
* 検索結果の内容を元にブラウザで検索： `Ctrl` + `Enter`

### ショートカット

Workflowのホットキーはデフォルトでは空欄になっている．適当に割り振ると良い（`Shift` + `Cmd` + `L`を使っていた．）  

文字列を選択してショートカットを押すことで部分一致検索できる．

## License

MIT

## Author

pudding