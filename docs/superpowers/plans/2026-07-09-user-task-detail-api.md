# 用户任务明细 API Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** 在 User API 下新增 `GET /user/task_detail/{userId}`，仅定义 Req/Res 与接口声明，不实现业务逻辑。

**Architecture:** 挂在现有 `api/user` 模块，与 `GetAllList` 同级；响应复用已有 `entity.Task`。`api/user/user.go` 为 GoFrame 生成文件，手改 Req/Res 后用 `gf gen ctrl` 同步接口声明（若项目惯例是手改则直接改）。

**Tech Stack:** GoFrame API 定义（`g.Meta`）、`demo/internal/model/entity.Task`

**Spec:** `docs/superpowers/specs/2026-07-09-user-task-detail-api-design.md`

---

## File Map

| 文件 | 职责 |
|------|------|
| `sever/demo/api/user/v1/user.go` | 新增 `GetUserTaskDetailReq` / `GetUserTaskDetailRes` |
| `sever/demo/api/user/user.go` | `IUserV1` 增加方法（生成或手改） |

本次不改：controller、service、dao。

---

### Task 1: 在 `api/user/v1/user.go` 增加 Req/Res

**Files:**
- Modify: `sever/demo/api/user/v1/user.go`

- [ ] **Step 1: 在文件末尾追加任务明细 API 定义**

在 `GetAllListRes` 之后追加：

```go
type GetUserTaskDetailReq struct {
	g.Meta `path:"/user/task_detail/{userId}" method:"get" tags:"User" summary:"按用户获取任务明细"`
	UserId uint `json:"userId" in:"path" v:"required|min:1#用户ID不能为空|用户ID无效"`
}

type GetUserTaskDetailRes struct {
	List []*entity.Task `json:"list" dc:"该用户下的全部任务"`
}
```

说明：`entity` 已在同文件 import，无需新增 import。

- [ ] **Step 2: 确认文件可编译**

Run（在 `sever/demo` 目录）：

```bash
go build ./api/user/v1/
```

Expected: 无错误输出，exit 0

- [ ] **Step 3: Commit**

```bash
git add sever/demo/api/user/v1/user.go
git commit -m "$(cat <<'EOF'
feat: 定义按用户获取任务明细 API

EOF
)"
```

---

### Task 2: 同步 `IUserV1` 接口声明

**Files:**
- Modify: `sever/demo/api/user/user.go`

- [ ] **Step 1: 更新接口方法**

在 `IUserV1` 中追加：

```go
GetUserTaskDetail(ctx context.Context, req *v1.GetUserTaskDetailReq) (res *v1.GetUserTaskDetailRes, err error)
```

若项目用 `make ctrl` / `gf gen ctrl` 生成该文件：在 `sever/demo` 执行生成命令，确认新方法出现即可，不要手写覆盖其它生成内容。

注意：本次**不**要求 controller 实现可编译通过；若 `gf gen ctrl` 生成了空 controller 方法可保留，但不要写 service 查询逻辑。若生成导致现有 controller 编译失败，只保留 `api/user/user.go` 的接口声明变更，controller 留给后续任务。

- [ ] **Step 2: Commit**

```bash
git add sever/demo/api/user/user.go
# 若 gen ctrl 还改了其它仅声明相关文件，一并加入；不要加入业务实现
git commit -m "$(cat <<'EOF'
feat: IUserV1 增加 GetUserTaskDetail

EOF
)"
```

---

## Spec Coverage Check

| Spec 要求 | Task |
|-----------|------|
| `GET /user/task_detail/{userId}` | Task 1 |
| path 参数 `userId` + 校验 | Task 1 |
| 响应 `list []*entity.Task` | Task 1 |
| 挂在 User API | Task 1–2 |
| 不做 controller/service | 全计划遵守 |
