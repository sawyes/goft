CREATE DATABASE IF NOT EXISTS `test`;

CREATE table test.`user` (
    id int primary key auto_increment,
    user_name varchar(255) not null comment "用户名称"
);

INSERT INTO `test`.`user`(`id`, `user_name`) VALUES (1, 'u1');
INSERT INTO `test`.`user`(`id`, `user_name`) VALUES (2, 'u2');
INSERT INTO `test`.`user`(`id`, `user_name`) VALUES (3, 'u3');
INSERT INTO `test`.`user`(`id`, `user_name`) VALUES (4, 'u4');