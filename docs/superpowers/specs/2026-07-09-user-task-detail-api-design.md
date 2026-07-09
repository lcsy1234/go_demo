# 用户任务明细 API 设计

## 目标

为前端「查看任务详情」提供接口：按用户 ID 返回该用户下全部任务列表（不分页）。本次只定义 API（Req/Res + 接口声明），不实现 controller / service。

## 接口

| 项 | 值 |
|----|-----|
| Method | `GET` |
| Path | `/user/task_detail/{userId}` |
| Tags | `User` |
| Summary | 按用户获取任务明细 |

### 请求

| 字段 | 位置 | 类型 | 校验 |
|------|------|------|------|
| `userId` | path | `uint` | `required\|min:1` |

### 响应

```json
{
  "list": [
    {
      "id": 1,
      "userId": 1,
      "title": "完成周报",
      "urgency": "medium",
      "deadline": "2026-07-10 18:00:00",
      "status": 0,
      "createdAt": "...",
      "updatedAt": "..."
    }
  ]
}
```

- `list`：`[]*entity.Task`，该用户下全部任务
- 无任务时返回空数组 `[]`，不报错

## 放置位置

挂在现有 **User API** 下（与 `GetUser`、`GetList` 同级），不新建独立 `task` 模块：

- `api/user/v1/user.go` — 增加 `GetUserTaskDetailReq` / `GetUserTaskDetailRes`
- `api/user/user.go` — `IUserV1` 增加对应方法（若由 `gf gen` 维护则按项目惯例生成）

原因：路径以 `/user/...` 开头，且入口来自用户管理页的「查看任务详情」。

## 范围

- **做**：API 定义（Req/Res）
- **不做**：controller、service、dao 查询逻辑、前端对接

## 数据来源

基于已有 `task` 表（`user_id` 关联用户），字段见 `entity.Task`。
