-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `article` (
    `id`         INT UNSIGNED     NOT NULL AUTO_INCREMENT COMMENT '文章ID',
    `user_id`    INT UNSIGNED     NOT NULL DEFAULT 0    COMMENT '作者用户ID',
    `title`      VARCHAR(128)     NOT NULL DEFAULT ''   COMMENT '标题',
    `content`    TEXT             NOT NULL              COMMENT '内容',
    `status`     TINYINT          NOT NULL DEFAULT 1    COMMENT '状态 1:发布 0:草稿',
    `created_at` DATETIME         NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` DATETIME         NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文章表';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `article`;
-- +goose StatementEnd
