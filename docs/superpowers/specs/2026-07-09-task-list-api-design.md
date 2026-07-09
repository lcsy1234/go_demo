# Task 分页列表 API 设计

## 目标

新建独立 `task` 模块，提供任务分页列表接口：支持按 `userId` 可选过滤，返回 `{ list, total }`。只含任务字段，不聚合用户信息。

## 接口

| 项 | 值 |
|----|-----|
| Method | `GET` |
| Path | `/task/list` |
| Tags | `Task` |
| Summary | 任务分页列表 |

### 请求（query）

| 字段 | 类型 | 默认 | 校验 | 说明 |
|------|------|------|------|------|
| `userId` | `uint` | — | 可选；传了则 `min:1` | 按用户过滤；不传则查全部 |
| `page` | `int` | `1` | `min:1` | 页码 |
| `size` | `int` | `10` | `between:1,100` | 每页条数 |

### 响应

```json
{
  "code": 0,
  "message": "OK",
  "data": {
    "list": [
      {
        "id": 1,
        "userId": 1,
        "title": "完成周报",
        "urgency": "medium",
        "deadline": "2026-07-10 18:00:00",
        "status": 0
      }
    ],
    "total": 6
  }
}
```

- `list`：任务列表（可用 `model.TaskList` 或等价 DTO）
- 无数据时 `list: []`，`total: 0`
- 排序：`id` 降序（与现有 `/user/list` 一致）

## 模块结构

与 `user` 平级，新建完整链路：

| 路径 | 职责 |
|------|------|
| `api/task/v1/task.go` | `GetListReq` / `GetListRes` |
| `api/task/task.go` | `ITaskV1` |
| `internal/controller/task/` | controller |
| `internal/logic/task/` | 查询逻辑 |
| `internal/service/task.go`（及 input 类型） | `ITask` 接口 |
| `internal/cmd/cmd.go` | `Bind(task.NewV1())` |
| `internal/logic/logic.go` | 注册 `_ "demo/internal/logic/task"` |

## 行为约定

- `task` logic 只通过 `dao.Task` 访问任务表
- 不返回用户详情，不嵌入 User 字段
- `userId` 未传：不过滤用户；传入且有效：`Where user_id = ?`
- 分页默认值与校验对齐 `/user/list`

## 范围

- **做**：独立 task 模块 + `GET /task/list` 全链路（api → controller → logic → service → 注册）
- **不做**：
  - 不动现有 `GetTaskDetail`（仍挂在 User）
  - 不加 status / urgency 等筛选
  - 不抽通用 Pagination 包
  - 不做 Task CRUD（创建/更新/删除）
  - 不改前端

## 数据来源

基于已有 `task` 表与 `dao.Task`，字段见 `entity.Task` / `model.TaskList`。
