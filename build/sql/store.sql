CREATE TABLE `store`
(
    `id`              int(10) NOT NULL AUTO_INCREMENT,
    `name`            varchar(10)  NOT NULL DEFAULT '' COMMENT '校区名称',
    `created_at`      timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`      timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '业务更新时间',
    `data_updated_at` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `creator`         varchar(32)  NOT NULL DEFAULT '' COMMENT '创建者',
    `updater`         varchar(32)  NOT NULL DEFAULT '' COMMENT '更新者',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='校区表';