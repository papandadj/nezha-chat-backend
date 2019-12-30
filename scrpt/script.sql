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

create database nezha_chat_common character set utf8mb4;
use nezha_chat_common;
create table nezha_chat_common_user_image (
  id int(10) unsigned not null auto_increment comment '主键',
  name varchar(20) not null comment '中文名',
  url varchar(200) not null comment 'url',
  created_at datetime default now(),
  updated_at datetime default now(),
  PRIMARY KEY (`id`),
  unique key `url` (`url`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_bin COMMENT ='默认人物表';

create database nezha_chat_friend character set utf8mb4;
use nezha_chat_friend;
create table friend (
  id int(10) unsigned not null auto_increment comment '主键',
  user_id1 varchar(12) not null comment '第一个用户id',
  user_id2 varchar(12) not null comment '第二个用户id',
  created_at datetime default now(),
  updated_at datetime default now(),
  PRIMARY KEY (`id`),
  unique key `userid_1_userid_2` (`user_id1`, `user_id2`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_bin COMMENT ='用户表';
