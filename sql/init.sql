create database if not exists user;
create database if not exists agency;

use user;

create table if not exists `user_profile_table` ( -- users' basic profile
  `uid` varchar(18) primary key not null, -- e.g. 'user_1615280517432'
  `username` varchar(50) not null, -- initialized as uid=username
  `password` varchar(32),
  `tel` varchar(11) not null,
  `email` varchar(50),
  `sex` tinyint,
  `age` tinyint,
  `address` varchar(256),
  `class_num` int,
  `img` varchar(100)
) engine=innodb default charset=utf8;

create table if not exists `user_login_inf_table` ( -- users' login information
  `uid` varchar(18) primary key not null,
  `username` varchar(50) not null,
  `is_login` tinyint not null,
  `last_login_time` varchar(20) not null, -- format: "yyyy-mm-dd hh:mm:ss"
  `last_login_device` varchar(50)
) engine=innodb default charset=utf8;

create table if not exists `test_user_class_table` ( -- only test table, standard format: [uid]_user_class_table
  `uid` varchar(18) not null,
  `class_id` varchar(19) not null,
  `bought_time` varchar(20) not null,
  `agency_id` varchar(20) not null -- the affiliation
) engine=innodb default charset=utf8; -- user's bought classes

create table if not exists `test_user_chatting_table` ( -- only test table, standard format: [uid]_user_chatting_table
  `chat_id` varchar(18) primary key not null,
  `uid` varchar(18) not null,
  `msg_num` int not null,
  `agency_icon` varchar(60),
  `agency_id` varchar(20) not null,
  `agency_name` varchar(50) not null
) engine=innodb default charset=utf8; -- user's all chatting box

create table if not exists `test_user_evaluations_table` ( -- only test table, standard format: [uid]_user_evaluations_table
  `evaluation_id` varchar(20) primary key not null, -- e.g. 'evalua_1615280517432'
  `favicon` varchar(60),
  `rating` float not null,
  `username` varchar(50) not null,
  `agency_id` varchar(20) not null,
  `class_id` varchar(19) not null,
  `detail` varchar(10000),
  `pics` varchar(700), -- maximum 20 photos
) engine=innodb default charset=utf8; -- user's all evaluations

create table if not exists `user_agency_chatting_contents` ( -- only test table, standard format: [uid]_[agency_id]_chatting_contents
  `content_id` varchar(21) primary key not null,
  `uid` varchar(18) not null,
  `agency_id` varchar(20) not null,
  `time` varchar(20) not null,
  `content` varchar(10000) not null
) engine=innodb default charset=utf8; -- the content of the specific dialogue between the user and an agency

use agency;

create table if not exists `agency_profile_table` ( -- agencies' basic profile
  `agency_id` varchar(20) primary key not null, -- e.g. 'agency_1615280518432'
  `name` varchar(50) not null, -- initialized as uid=agency_name
  `tel` varchar(11) not null,
  `rating` float not null, -- 0-10
  `comments` int not null,
  `order` int not null,
  `tags` varchar(120), -- maximum 6 tags, each tag max 8 characters
  `address` varchar(256) not null,
  `address_detail` varchar(256) not null,
  `icon` varchar(60) not null,
  `photos` varchar(700), -- maximum 20 photos, save each photo's filename(a hash)
  `brand_history` varchar(5000),
  `characteristics` varchar(250) -- maximum 8 words, each word maximum 30 characters. eg. 'excellent teacher,professional course,flexible time'
) engine=innodb default charset=utf8;

create table if not exists `agency_login_inf_table` ( -- agencies' login information
    `agency_id` varchar(18) primary key not null,
    `name` varchar(50) not null,
    `is_login` tinyint not null,
    `last_login_time` varchar(20) not null, -- format: "yyyy-mm-dd hh:mm:ss"
    `last_login_device` varchar(50)
) engine=innodb default charset=utf8;

create table if not exists `test_agency_chatting_table` ( -- only test table, standard format: [agency_id]_user_chatting_table
    `chat_id` varchar(18) primary key not null,
    `agency_id` varchar(18) not null,
    `msg_num` int not null,
    `user_icon` varchar(60),
    `uid` varchar(20) not null,
    `username` varchar(50) not null
) engine=innodb default charset=utf8; -- agency's all chatting box

create table if not exists `test_agency_class_table` ( -- only test table, standard format: [agency_id]_agency_class_table
  `agency_id` varchar(20) not null,
  `class_id` varchar(19) not null,
  `price` float,
  `name` varchar(50) not null,
  `stu_number` int,
  `age` varchar(10), -- eg. '4-14'; '20-30'
  `level` varchar(20), -- beginner, intermediate, advanced, expert
  `sales` int, -- number of sold
  `create_time` varchar(19), -- eg. '2006-07-01 12:24:58'
  `last_update_time` varchar(19) -- same as `create_time`
) engine=innodb default charset=utf8; -- agency's all classes

create table if not exists `test_agency_teacher_table` ( -- only test table, standard format: [agency_id]_agency_teacher_table
    `agency_id` varchar(20) not null,
    `teacher_id` varchar(21) not null,
    `name` varchar(50) not null,
    `pic`  varchar(60),
    `tag`  varchar(120),
    `tel`  varchar(11) not null,
    `description` varchar(400)
) engine=innodb default charset=utf8;

create table if not exists `test_agency_evaluations_table` ( -- only test table, standard format: [agency_id]_agency_evaluations_table
    `evaluation_id` varchar(20) primary key not null, -- e.g. 'evalua_1615280517432'
    `favicon` varchar(60),
    `rating` float not null, -- 0-10
    `username` varchar(50) not null,
    `agency_id` varchar(20) not null,
    `class_id` varchar(19) not null,
    `detail` varchar(10000),
    `pics` varchar(700), -- maximum 20 photos
) engine=innodb default charset=utf8;