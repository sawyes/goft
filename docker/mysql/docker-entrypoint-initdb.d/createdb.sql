CREATE DATABASE IF NOT EXISTS `test`;

CREATE table test.`user` (
    id int primary key auto_increment,
    user_name varchar(255) not null comment "用户名称"
);

INSERT INTO `test`.`user`(`id`, `user_name`) VALUES (1, 'u1');
INSERT INTO `test`.`user`(`id`, `user_name`) VALUES (2, 'u2');
INSERT INTO `test`.`user`(`id`, `user_name`) VALUES (3, 'u3');
INSERT INTO `test`.`user`(`id`, `user_name`) VALUES (4, 'u4');

CREATE table test.`mynews` (
    id int primary key auto_increment,
    newstitle varchar(255) not null comment "新闻标题",
    newscontent varchar(255) not null comment "新闻内容",
    views int not null default 0,
    addtime datetime default now()
);

INSERT INTO `test`.`mynews`(`id`, `newstitle`, `newscontent`, `views`, `addtime`) VALUES (1, '新闻1', '1111111', 0, '2021-02-14 09:21:11');
INSERT INTO `test`.`mynews`(`id`, `newstitle`, `newscontent`, `views`, `addtime`) VALUES (2, '新闻2', '2222222222', 0, '2021-02-14 09:21:22');
INSERT INTO `test`.`mynews`(`id`, `newstitle`, `newscontent`, `views`, `addtime`) VALUES (3, '新闻3', '3333333333', 0, '2021-02-14 09:21:33');
INSERT INTO `test`.`mynews`(`id`, `newstitle`, `newscontent`, `views`, `addtime`) VALUES (4, '新闻4', '444444444', 0, '2021-02-14 09:21:44');