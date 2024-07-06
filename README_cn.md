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
  "index": 1,
  "pagination": 1,
  "order_by": "download_time",
  "reverse": false
}
```

| 参数名        | 类型     | 说明       |
|------------|--------|----------|
| index      | int    | 当前页数     |
| pagination | int    | 分页大小     |
| order_by   | string | 用来排序的表头名 |
| reverse    | bool   | 降序or升序   |

## 使用方法

在Manage界面新增一个条目为Counter, 点击后即可查看下载统计信息

## 完成计划

- [x] api设计完毕
- [ ] 前端Manage界面尚未修改完成
- [ ] api新增1天内下载量等功能
- [ ] 前端i18n尚未适配
- [ ] 仅仅测试了sqlite, 其余数据库尚未测试

