# Task 分页列表 API Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** 新建独立 `task` 模块，实现 `GET /task/list` 分页列表（可选 `userId` 过滤）。

**Architecture:** 与 `user` 平级完整链路：`api/task` → `controller/task` → `logic/task` → `service.ITask` → `dao.Task`。响应复用 `model.TaskList`，风格对齐现有 `/user/list`。

**Tech Stack:** GoFrame（`g.Meta`、dao Page/Count）、现有 `dao.Task`

**Spec:** `docs/superpowers/specs/2026-07-09-task-list-api-design.md`

---

## File Map

| 文件 | 职责 |
|------|------|
| `sever/demo/api/task/v1/task.go` | `GetListReq` / `GetListRes` |
| `sever/demo/api/task/task.go` | `ITaskV1` |
| `sever/demo/internal/service/task.go` | `ITask` + Register |
| `sever/demo/internal/service/task_input.go` | GetList Input/Output |
| `sever/demo/internal/logic/task/task.go` | 分页查询实现 |
| `sever/demo/internal/logic/logic.go` | 注册 task logic |
| `sever/demo/internal/controller/task/task.go` | 包占位 |
| `sever/demo/internal/controller/task/task_new.go` | `NewV1` |
| `sever/demo/internal/controller/task/task_v1_get_list.go` | GetList handler |
| `sever/demo/internal/cmd/cmd.go` | Bind task controller |

复用已有：`internal/model/task.go`（`TaskList`）、`dao.Task`。

---

### Task 1: API 定义

**Files:**
- Create: `sever/demo/api/task/v1/task.go`
- Create: `sever/demo/api/task/task.go`

- [ ] **Step 1: 创建 `api/task/v1/task.go`**

```go
package v1

import (
	"demo/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type GetListReq struct {
	g.Meta `path:"/task/list" method:"get" tags:"Task" summary:"任务分页列表"`
	UserId uint `json:"userId" in:"query" v:"omitempty|min:1#用户ID无效"`
	Page   int  `json:"page"   in:"query" d:"1"  v:"min:1#页码最小为1"`
	Size   int  `json:"size"   in:"query" d:"10" v:"between:1,100#每页数量范围为1-100"`
}

type GetListRes struct {
	List  []*model.TaskList `json:"list"  dc:"任务列表"`
	Total int               `json:"total" dc:"总数"`
}
```

- [ ] **Step 2: 创建 `api/task/task.go`**

```go
package task

import (
	"context"

	"demo/api/task/v1"
)

type ITaskV1 interface {
	GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error)
}
```

- [ ] **Step 3: 编译 API 包**

Run（在 `sever/demo`）：

```bash
go build ./api/task/...
```

Expected: exit 0

- [ ] **Step 4: Commit**

```bash
git add sever/demo/api/task/
git commit -m "$(cat <<'EOF'
feat: 定义 Task 分页列表 API

EOF
)"
```

---

### Task 2: Service 接口与 Input/Output

**Files:**
- Create: `sever/demo/internal/service/task_input.go`
- Create: `sever/demo/internal/service/task.go`

- [ ] **Step 1: 创建 `task_input.go`**

```go
package service

import "demo/internal/model"

type (
	TaskGetListInput struct {
		UserId uint
		Page   int
		Size   int
	}
	TaskGetListOutput struct {
		List  []*model.TaskList
		Total int
	}
)
```

- [ ] **Step 2: 创建 `task.go`（ITask）**

```go
package service

import "context"

type (
	ITask interface {
		GetList(ctx context.Context, in TaskGetListInput) (out TaskGetListOutput, err error)
	}
)

var (
	localTask ITask
)

func Task() ITask {
	if localTask == nil {
		panic("implement not found for interface ITask, forgot register?")
	}
	return localTask
}

func RegisterTask(i ITask) {
	localTask = i
}
```

- [ ] **Step 3: Commit**

```bash
git add sever/demo/internal/service/task.go sever/demo/internal/service/task_input.go
git commit -m "$(cat <<'EOF'
feat: 增加 ITask 服务接口

EOF
)"
```

---

### Task 3: Logic 实现与注册

**Files:**
- Create: `sever/demo/internal/logic/task/task.go`
- Modify: `sever/demo/internal/logic/logic.go`

- [ ] **Step 1: 创建 logic**

```go
package task

import (
	"context"
	"demo/internal/dao"
	"demo/internal/model"
	"demo/internal/service"
)

type sTask struct{}

func init() {
	service.RegisterTask(New())
}

func New() *sTask {
	return &sTask{}
}

func (s *sTask) GetList(ctx context.Context, in service.TaskGetListInput) (out service.TaskGetListOutput, err error) {
	page, size := in.Page, in.Size
	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 10
	}
	m := dao.Task.Ctx(ctx)
	if in.UserId > 0 {
		m = m.Where(dao.Task.Columns().UserId, in.UserId)
	}
	total, err := m.Count()
	if err != nil {
		return out, err
	}
	var list []*model.TaskList
	err = m.Page(page, size).
		OrderDesc(dao.Task.Columns().Id).
		Scan(&list)
	if err != nil {
		return out, err
	}
	if list == nil {
		list = make([]*model.TaskList, 0)
	}
	return service.TaskGetListOutput{
		List:  list,
		Total: total,
	}, nil
}
```

- [ ] **Step 2: 注册到 `logic.go`**

```go
package logic

import (
	_ "demo/internal/logic/task"
	_ "demo/internal/logic/test1"
	_ "demo/internal/logic/user"
)
```

- [ ] **Step 3: 编译**

```bash
go build ./internal/logic/...
```

Expected: exit 0

- [ ] **Step 4: Commit**

```bash
git add sever/demo/internal/logic/task/ sever/demo/internal/logic/logic.go sever/demo/internal/model/task.go
git commit -m "$(cat <<'EOF'
feat: 实现 Task 分页查询 logic

EOF
)"
```

说明：若 `model/task.go` 尚未入库，一并加入。

---

### Task 4: Controller + 路由绑定

**Files:**
- Create: `sever/demo/internal/controller/task/task.go`
- Create: `sever/demo/internal/controller/task/task_new.go`
- Create: `sever/demo/internal/controller/task/task_v1_get_list.go`
- Modify: `sever/demo/internal/cmd/cmd.go`

- [ ] **Step 1: controller 包文件**

`task.go`:

```go
package task
```

`task_new.go`:

```go
package task

import (
	"demo/api/task"
)

type ControllerV1 struct{}

func NewV1() task.ITaskV1 {
	return &ControllerV1{}
}
```

`task_v1_get_list.go`:

```go
package task

import (
	"context"

	v1 "demo/api/task/v1"
	"demo/internal/service"
)

func (c *ControllerV1) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	out, err := service.Task().GetList(ctx, service.TaskGetListInput{
		UserId: req.UserId,
		Page:   req.Page,
		Size:   req.Size,
	})
	if err != nil {
		return nil, err
	}
	return &v1.GetListRes{
		List:  out.List,
		Total: out.Total,
	}, nil
}
```

- [ ] **Step 2: `cmd.go` 绑定**

在 import 增加 `"demo/internal/controller/task"`，在 `Bind` 中增加 `task.NewV1()`。

- [ ] **Step 3: 全量编译**

```bash
go build -o /dev/null .
```

Expected: exit 0

- [ ] **Step 4: Commit**

```bash
git add sever/demo/internal/controller/task/ sever/demo/internal/cmd/cmd.go
git commit -m "$(cat <<'EOF'
feat: 注册 Task 列表接口路由

EOF
)"
```

---

## Spec Coverage

| Spec 要求 | Task |
|-----------|------|
| `GET /task/list` | Task 1 |
| userId 可选 + page/size | Task 1 + 3 |
| `{ list, total }` | Task 1 + 3 + 4 |
| 独立 task 模块全链路 | Task 1–4 |
| 不动 GetTaskDetail | 无改动 |
| 不抽 Pagination / 不做 CRUD | 无相关任务 |
