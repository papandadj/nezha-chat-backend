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


insert into `nezha_chat_common_user_image` (`name`, `url`) values('哪吒', 'nezha1.png');
insert into `nezha_chat_common_user_image` (`name`, `url`) values('哪吒', 'nezha2.png');
insert into `nezha_chat_common_user_image` (`name`, `url`) values('哪吒', 'nezha3.png');
insert into `nezha_chat_common_user_image` (`name`, `url`) values('妲己', 'daji.png');
insert into `nezha_chat_common_user_image` (`name`, `url`) values('二郎神', 'erlangshen.png');
insert into `nezha_chat_common_user_image` (`name`, `url`) values('姬发', 'jifa.png');
insert into `nezha_chat_common_user_image` (`name`, `url`) values('雷震子', 'leizhenzi.png');
insert into `nezha_chat_common_user_image` (`name`, `url`) values('李靖', 'lijing.png');
insert into `nezha_chat_common_user_image` (`name`, `url`) values('女娲', 'nvwa.png');
insert into `nezha_chat_common_user_image` (`name`, `url`) values('石矶', 'shiji.png');
insert into `nezha_chat_common_user_image` (`name`, `url`) values('太乙', 'taiyi.png');
insert into `nezha_chat_common_user_image` (`name`, `url`) values('小猪熊', 'xiaozhuxiong.png');
insert into `nezha_chat_common_user_image` (`name`, `url`) values('殷夫人', 'yinfuren.png');
insert into `nezha_chat_common_user_image` (`name`, `url`) values('小龙女', 'xiaolongnv.png');