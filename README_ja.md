[English](./README.md) | [中文](./README_cn.md) | 日本語 | [Contributing](./CONTRIBUTING.md) | [CODE_OF_CONDUCT](./CODE_OF_CONDUCT.md)
## 紹介

alistの詳細については、[alist公式リポジトリ](https://github.com/alist-org/alist)をご覧ください。

このフォークは、公式のalistを改良したもので、ダウンロード統計機能を追加しています。

## 原理

元のalistデータベースに新しいテーブルを追加しました。各ダウンロードごとにデータベースにエントリが追加され、APIを介して統計情報を取得できます。

APIは以下の通りです：

```
/api/admin/counter/get
```

リクエストボディ：

```json
{
  "current_page": 1,
  "page_size": 1,
  "sort_key": "download_time",
  "reverse": false,
  "file_name": "fileName",
  "ip_address": "1.1.1.1",
  "status_code": 200
}
```
| パラメータ            | 型      | 必須  | 説明             |
|------------------|--------|-----|----------------|
| current_page            | int    | y   | 現在のページ番号       |
| page_size       | int    | はい  | ページサイズ         |
| sort_key         | string | はい  | ソートする列名        |
| reverse          | bool   | はい  | 降順または昇順        |
| file_name        | string | いいえ | パスを含まないファイル名   |
| ip_address       | string | いいえ | ダウンロードリクエストのIP |
| status_code | int    | いいえ | HTTPステータスコード   |

## 使い方

![image](https://github.com/MooWantFree/alist/assets/46401523/d32227fc-008b-4017-bad9-dcb1e396d4ac)
管理インターフェースに「Counter」という新しいエントリを追加します。それをクリックすると、ダウンロード統計情報が表示されます。

## 完成計画

- [x] API設計完了
- [ ] フロントエンド管理インターフェースはまだ修正されていません
- [ ] APIに1日以内のダウンロード数などの機能を追加
- [ ] フロントエンドのi18nはまだ適用されていません
- [ ] SQLiteでのみテストされています。他のデータベースではまだテストされていません。

```