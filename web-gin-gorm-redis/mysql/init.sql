CREATE DATABASE IF NOT EXISTS web-gin-gorm-redis;
CREATE TABLE IF NOT EXISTS user (
   `id` bigint(20) NOT NULL AUTO_INCREMENT,
   `name` varchar(10)  COMMENT '姓名',
    PRIMARY KEY (`id`),
)  ENGINE=InnoDB COMMENT='用户信息';

