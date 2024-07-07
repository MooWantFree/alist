[English](./README.md) | 中文 | [日本語](./README_ja.md) | [Contributing](./CONTRIBUTING.md) | [CODE_OF_CONDUCT](./CODE_OF_CONDUCT.md)
## 介绍

具体alist介绍请见[alist官方仓库](https://github.com/alist-org/alist)

本fork为alist官方修改版, 添加了下载统计功能

## 原理

在alist原有的数据库中新增了一个table, 每次下载会向数据库添加条目, 并通过api获取, 实现下载统计

api为

```
/api/admin/counter/get
```

请求体

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

| 参数               | 类型     | 必填 | 描述        |
|------------------|--------|----|-----------|
| current_page            | int    | 是  | 当前页码      |
| page_size       | int    | 是  | 页面大小      |
| sort_key         | string | 是  | 排序的列名     |
| reverse          | bool   | 是  | 降序或升序     |
| file_name        | string | 否  | 不含路径的文件名  |
| ip_address       | string | 否  | 下载请求的IP地址 |
| status_code | int    | 否  | HTTP状态码   |

## 使用方法

![image](https://github.com/MooWantFree/alist/assets/46401523/f58ff682-d247-49f3-a069-aade17b1f60b)

在Manage界面新增一个条目为Counter, 点击后即可查看下载统计信息

## 完成计划

- [x] api设计完毕
- [ ] 前端Manage界面尚未修改完成
- [ ] api新增1天内下载量等功能
- [ ] 前端i18n尚未适配
- [ ] 仅仅测试了sqlite, 其余数据库尚未测试

