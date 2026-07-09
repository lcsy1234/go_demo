# User Task 表明细设计

## 目标

为 `user` 表新增一对多任务表 `task`，仅通过 goose 迁移建表并插入示例数据，不改动 Go 业务代码。

## 关系

- 一个用户（`user`）对应多条任务（`task`）
- 通过 `task.user_id` 关联 `user.id`
- 不在数据库层强制外键（与现有 `article` 表风格一致）

## 表结构：`task`

| 字段 | 类型 | 说明 |
|------|------|------|
| `id` | INT UNSIGNED PK AI | 任务 ID |
| `user_id` | INT UNSIGNED NOT NULL | 所属用户 ID，索引 `idx_user_id` |
| `title` | VARCHAR(128) NOT NULL | 任务标题 |
| `urgency` | VARCHAR(16) NOT NULL | 紧急程度：`low` / `medium` / `high` |
| `deadline` | DATETIME NOT NULL | 截止时间 |
| `status` | TINYINT NOT NULL DEFAULT 0 | `0` 未完成 / `1` 已完成 |
| `created_at` | DATETIME | 创建时间 |
| `updated_at` | DATETIME | 更新时间 |

## 范围

- **做**：goose 迁移（CREATE + 示例 INSERT）
- **不做**：GoFrame dao 白名单、`gf gen dao`、CRUD API

## 示例数据

为 `user_id` 1、2 各插入若干任务，覆盖不同 `urgency` 与 `status`。

## 执行方式

在 `sever/demo` 目录执行：`make migrate`
