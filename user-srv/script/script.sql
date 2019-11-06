create database nezha_chat_user character set utf8mb4;
use nezha_chat_user;
create table user (
  id int(10) unsigned not null auto_increment comment '主键',
  username varchar(20) comment '用户名称',
  password varchar(128) comment '密码',
  image varchar(128) comment '图片地址',
  created_at datetime default now(),
  updated_at datetime default now(),
  PRIMARY KEY (`id`),
  unique key `user_username_uindex` (`username`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_bin COMMENT ='用户表';