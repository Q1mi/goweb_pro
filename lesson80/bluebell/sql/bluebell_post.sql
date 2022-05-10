create table post
(
    id           bigint auto_increment
        primary key,
    post_id      bigint                              not null comment '帖子id',
    title        varchar(128)                        not null comment '标题',
    content      varchar(8192)                       not null comment '内容',
    author_id    bigint                              not null comment '作者的用户id',
    community_id bigint                              not null comment '所属社区',
    status       tinyint   default 1                 not null comment '帖子状态',
    create_time  timestamp default CURRENT_TIMESTAMP null comment '创建时间',
    update_time  timestamp default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP comment '更新时间',
    constraint idx_post_id
        unique (post_id)
)
    collate = utf8mb4_general_ci;

create index idx_author_id
    on post (author_id);

create index idx_community_id
    on post (community_id);

INSERT INTO bluebell.post (id, post_id, title, content, author_id, community_id, status, create_time, update_time) VALUES (1, 14283784123846656, '学习使我快乐', '只有学习才能变得更强', 28018727488323585, 1, 1, '2020-08-09 09:58:39', '2020-08-09 09:58:39');
INSERT INTO bluebell.post (id, post_id, title, content, author_id, community_id, status, create_time, update_time) VALUES (2, 14373128436191232, 'CSGO开箱子好上瘾', '花了钱不出金，我好气啊', 28018727488323585, 2, 1, '2020-08-09 15:53:40', '2020-08-09 15:53:40');
INSERT INTO bluebell.post (id, post_id, title, content, author_id, community_id, status, create_time, update_time) VALUES (3, 14373246019309568, 'IG牛逼', '打得好啊。。。', 28018727488323585, 3, 1, '2020-08-09 15:54:08', '2020-08-09 15:54:08');
INSERT INTO bluebell.post (id, post_id, title, content, author_id, community_id, status, create_time, update_time) VALUES (4, 19432670719119360, '投票功能真好玩', '12345', 28018727488323585, 2, 1, '2020-08-23 14:58:29', '2020-08-23 14:58:29');
INSERT INTO bluebell.post (id, post_id, title, content, author_id, community_id, status, create_time, update_time) VALUES (5, 19433711036534784, '投票功能真好玩2', '12345', 28018727488323585, 2, 1, '2020-08-23 15:02:37', '2020-08-23 15:02:37');
INSERT INTO bluebell.post (id, post_id, title, content, author_id, community_id, status, create_time, update_time) VALUES (6, 19434165682311168, '投票功能真好玩2', '12345', 28018727488323585, 2, 1, '2020-08-23 15:04:26', '2020-08-23 15:04:26');
INSERT INTO bluebell.post (id, post_id, title, content, author_id, community_id, status, create_time, update_time) VALUES (7, 21810561880690688, '看图说话', '4321', 28018727488323585, 2, 1, '2020-08-30 04:27:23', '2020-08-30 04:27:23');
INSERT INTO bluebell.post (id, post_id, title, content, author_id, community_id, status, create_time, update_time) VALUES (8, 21810685746876416, '永远不要高估自己', '做个普通人也挺难', 28018727488323585, 3, 1, '2020-08-30 04:27:52', '2020-08-30 04:27:52');
INSERT INTO bluebell.post (id, post_id, title, content, author_id, community_id, status, create_time, update_time) VALUES (9, 21810865955147776, '你知道泛型是什么吗？', '不知道泛型是什么却一直在问泛型什么时候出', 28018727488323585, 1, 1, '2020-08-30 04:28:35', '2020-08-30 04:28:35');
INSERT INTO bluebell.post (id, post_id, title, content, author_id, community_id, status, create_time, update_time) VALUES (10, 21810938202034176, '国庆假期哪里玩？', '走遍四海，还是威海。', 28018727488323585, 1, 1, '2020-08-30 04:28:52', '2020-08-30 04:28:52');
INSERT INTO bluebell.post (id, post_id, title, content, author_id, community_id, status, create_time, update_time) VALUES (11, 1, 'test', 'just for test', 1, 1, 1, '2020-09-12 14:03:18', '2020-09-12 14:03:18');
INSERT INTO bluebell.post (id, post_id, title, content, author_id, community_id, status, create_time, update_time) VALUES (12, 92636388033302528, 'test', 'just a test', 1, 1, 1, '2020-09-12 15:03:56', '2020-09-12 15:03:56');
INSERT INTO bluebell.post (id, post_id, title, content, author_id, community_id, status, create_time, update_time) VALUES (13, 92636388142354432, 'test', 'just a test', 1, 1, 1, '2020-09-12 15:03:56', '2020-09-12 15:03:56');
INSERT INTO bluebell.post (id, post_id, title, content, author_id, community_id, status, create_time, update_time) VALUES (15, 123, 'test', 'just a test', 1, 1, 1, '2020-09-13 03:31:50', '2020-09-13 03:31:50');
INSERT INTO bluebell.post (id, post_id, title, content, author_id, community_id, status, create_time, update_time) VALUES (16, 10, 'test', 'just a test', 123, 1, 1, '2020-09-13 04:12:44', '2020-09-13 04:12:44');