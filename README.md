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
  "reverse": false
}
```

| Parameter  | Type   | Description                   |
|------------|--------|-------------------------------|
| index      | int    | Current page number           |
| pagination | int    | Page size                     |
| order_by   | string | Column name for sorting       |
| reverse    | bool   | Descending or ascending order |

## Usage

Add a new entry named "Counter" in the Manage interface. Click it to view the download statistics.

## Completion Plan

- [x] API design completed
- [ ] Frontend Manage interface not yet modified
- [ ] API to add functionality for download counts within one day
- [ ] Frontend i18n not yet adapted
- [ ] Only tested with SQLite, other databases have not been tested yet.