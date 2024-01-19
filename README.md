# 概要

DDDの各層の責務となるバリデーションを行う際のエラー定義の実装です。

# 構成
```
.
├── application // TODO: 
│   ├── error.go
│   └── sample_a.go
├── application_clean // Applicationレイヤの出力にOutputPortを使用し依存関係を逆転させた実装
│   ├── error.go
│   └── sample_a.go
├── common            // エンジニアのためのエラー。主にスタックトレースを出力
│   └── error.go
├── domain
│   ├── error.go     // Domainレイヤのエラー
│   ├── models
│   │   └── sample_a.go
│   └── repositories
│       └── sample_a.go
├── go.mod
├── graph_clean       // Applicationレイヤの出力にOutputPortを使用し依存関係を逆転させた実装
│   ├── models.go
│   └── sample_a.go
└── main.go          // 実行ファイル
```
