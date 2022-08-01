create table if not exists user(
                                   id int not null auto_increment comment '主键自增id',
                                   uid int not null comment '用户id',
                                   keywords text not null comment '索引词',
                                   degree char(2) not null comment '学历',
    gender char(1) not null comment '性别',
    city char(2) not null comment '城市',
    primary key (id),
    unique key idx_uid (uid)
    )engine=innodb default charset=utf8 comment '用户信息表';