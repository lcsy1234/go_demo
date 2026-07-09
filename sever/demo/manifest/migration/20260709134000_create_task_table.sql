-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `task` (
    `id`         INT UNSIGNED     NOT NULL AUTO_INCREMENT COMMENT '任务ID',
    `user_id`    INT UNSIGNED     NOT NULL DEFAULT 0    COMMENT '所属用户ID',
    `title`      VARCHAR(128)     NOT NULL DEFAULT ''   COMMENT '任务标题',
    `urgency`    VARCHAR(16)      NOT NULL DEFAULT 'low' COMMENT '紧急程度 low/medium/high',
    `deadline`   DATETIME         NOT NULL              COMMENT '截止时间',
    `status`     TINYINT          NOT NULL DEFAULT 0    COMMENT '状态 0:未完成 1:已完成',
    `created_at` DATETIME         NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` DATETIME         NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户任务表';

INSERT INTO `task` (`user_id`, `title`, `urgency`, `deadline`, `status`) VALUES
(1, '完成周报', 'medium', '2026-07-10 18:00:00', 0),
(1, '修复登录页样式', 'high', '2026-07-09 20:00:00', 1),
(1, '整理读书笔记', 'low', '2026-07-15 23:59:59', 0),
(2, '准备项目演示', 'high', '2026-07-11 12:00:00', 0),
(2, '更新接口文档', 'medium', '2026-07-12 18:00:00', 0),
(3, '代码评审', 'low', '2026-07-14 17:00:00', 1);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `task`;
-- +goose StatementEnd
