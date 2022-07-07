create table blog_admin
(
    id       int unsigned auto_increment
        primary key,
    username varchar(50) default '' null comment '账号',
    password varchar(50) default '' null comment '密码',
    image    varchar(15)            not null comment '管理员头像'
);

create table blog_article
(
    id            int unsigned auto_increment
        primary key,
    tag_id        int unsigned     default '0' null comment '标签ID',
    title         varchar(100)     default ''  null comment '文章标题',
    `desc`        varchar(255)     default ''  null comment '简述',
    content       text                         not null comment '文章内容',
    created_on    datetime                     not null comment '创建时间',
    created_by    varchar(100)     default ''  null comment '创建人',
    modified_on   datetime                     null comment '修改时间',
    modified_by   varchar(255)     default ''  null comment '修改人',
    state         tinyint unsigned default '1' null comment '状态 0为禁用1为启用',
    article_click int                          not null comment '查看人数',
    article_ip    varchar(15)                  not null comment '发布ip'
)
    comment '文章管理';

create table blog_click
(
    id           int auto_increment comment '点击的id'
        primary key,
    article_id   int         not null comment '文章id
',
    article_name varchar(15) not null comment '点击文章的名字',
    topic        varchar(15) not null comment '文章的主题',
    click_time   datetime    not null comment '点击的时间',
    click_by     varchar(15) not null comment '点击人的用户名
'
);

create table blog_comment
(
    c_id            int auto_increment comment '评论自增id'
        primary key,
    user_id         int          not null comment '收到评论的id',
    type_id         int          not null comment '收到评论文章的类型',
    comment_id      int          not null comment '评论内容的id',
    comment_content varchar(255) not null comment '评论内容',
    comment_user_id int          not null comment '评论者id',
    comment_time    datetime     not null comment '评论时间',
    comment_ip      varchar(15)  not null comment '评论地址'
);

create table blog_power_list
(
    p_id       int auto_increment comment '自增id'
        primary key,
    power_id   int          not null comment '权限ID',
    power_desc varchar(255) not null comment '权限描述'
);

create table blog_tag
(
    id          int unsigned auto_increment
        primary key,
    name        varchar(100)     default ''  null comment '标签名称',
    created_on  datetime                     null comment '创建时间',
    created_by  varchar(100)     default ''  null comment '创建人',
    modified_on datetime                     null comment '修改时间',
    modified_by varchar(100)     default ''  null comment '修改人',
    state       tinyint unsigned default '1' null comment '状态 0为禁用、1为启用'
)
    comment '文章标签管理';

create table blog_user
(
    user_id            int auto_increment comment '用户id'
        primary key,
    user_name          varchar(32)  not null comment '用户名',
    user_password      varchar(32)  not null comment '用户密码',
    user_phone         int          not null comment '用户手机号',
    user_desc          varchar(32)  not null comment '用户自我描述',
    user_image         varchar(255) not null comment '用户头像路径',
    user_register_time datetime     not null comment '用户注册时间',
    user_last_login_ip varchar(15)  not null comment '用户上一次登陆IP',
    user_register_ip   varchar(15)  not null comment '用户注册ip',
    user_power         int          not null comment '用户权限'
)
    engine = MyISAM;


