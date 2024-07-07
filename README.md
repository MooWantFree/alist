English | [中文](./README_cn.md) | [日本語](./README_ja.md) | [Contributing](./CONTRIBUTING.md) | [CODE_OF_CONDUCT](./CODE_OF_CONDUCT.md)
## Introduction

For detailed information about alist, please refer to
the [alist official repository](https://github.com/alist-org/alist).

This fork is a modified version of the official alist, adding download statistics functionality.

## Principle

A new table is added to the original alist database. Each download adds an entry to the database, and the statistics can
be retrieved via an API.

The API is:

```
/api/admin/counter/get
```

Request body:

```json
{
  "index": 1,
  "pagination": 1,
  "order_by": "download_time",
  "reverse": false,
  "file_name": "fileName",
  "ip_address": "1.1.1.1",
  "http_status_code": 200
}
```

| Parameter        | Type   | Required | Description                   |
| ---------------- | ------ | -------- | ----------------------------- |
| index            | int    | y        | Current page number           |
| pagination       | int    | y        | Page size                     |
| order_by         | string | y        | Column name for sorting       |
| reverse          | bool   | y        | Descending or ascending order |
| file_name        | string | n        | File name without path        |
| ip_address       | string | n        | download request IP           |
| http_status_code | int    | n        | http status code              |

## Usage
![image](https://github.com/MooWantFree/alist/assets/46401523/d32227fc-008b-4017-bad9-dcb1e396d4ac)

Add a new entry named "Counter" in the Manage interface. Click it to view the download statistics.

## Completion Plan

- [x] API design completed
- [ ] Frontend Manage interface not yet modified
- [ ] API to add functionality for download counts within one day
- [ ] Frontend i18n not yet adapted
- [ ] Only tested with SQLite, other databases have not been tested yet.
