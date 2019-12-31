CREATE USER 'nezha'@'%' IDENTIFIED BY 'test';
revoke ALL ON *.* from 'nezha'@'%';
grant ALL on nezha_chat_user.user to 'nezha'@'%' with grant option;
grant ALL on nezha_chat_friend.friend to 'nezha'@'%' with grant option;
grant ALL on nezha_chat_common.nezha_chat_common_user_image to 'nezha'@'%' with grant option;
flush privileges;
