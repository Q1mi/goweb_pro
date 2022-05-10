create table user
(
    id          bigint auto_increment
        primary key,
    user_id     bigint                              not null,
    username    varchar(64)                         not null,
    password    varchar(64)                         not null,
    email       varchar(64)                         null,
    gender      tinyint   default 0                 not null,
    create_time timestamp default CURRENT_TIMESTAMP null,
    update_time timestamp default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP,
    constraint idx_user_id
        unique (user_id),
    constraint idx_username
        unique (username)
)
    collate = utf8mb4_general_ci;

INSERT INTO bluebell.user (id, user_id, username, password, email, gender, create_time, update_time) VALUES (1, 28018727488323585, 'q1mi', '313233343536639a9119599647d841b1bef6ce5ea293', null, 0, '2020-07-12 07:01:03', '2020-07-12 07:01:03');
INSERT INTO bluebell.user (id, user_id, username, password, email, gender, create_time, update_time) VALUES (2, 4183532125556736, '七米', '313233639a9119599647d841b1bef6ce5ea293', null, 0, '2020-07-12 13:03:51', '2020-07-12 13:03:51');