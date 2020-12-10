create database course default character set utf8mb4 collate utf8mb4_general_ci;

-- 课程
drop table if exists course;
create table course
(
    id         char(8)     not null default '' comment 'id',
    name       varchar(50) not null comment '名称',
    summary    varchar(2000) comment '概述',
    time       int                  default 0 comment '时长|单位秒',
    price      decimal(8, 2)        default 0.00 comment '价格（元）',
    image      varchar(100) comment '封面',
    level      char(1) comment '级别|枚举[CourseLevelEnum]：ONE("1", "初级"),TWO("2", "中级"),THREE("3", "高级")',
    charge     char(1) comment '收费|枚举[CourseChargeEnum]：CHARGE("C", "收费"),FREE("F", "免费")',
    status     char(1) comment '状态|枚举[CourseStatusEnum]：PUBLISH("P", "发布"),DRAFT("D", "草稿")',
    enroll     integer              default 0 comment '报名数',
    sort       int comment '顺序',
    created_at datetime(3) comment '创建时间',
    updated_at datetime(3) comment '修改时间',
    primary key (id)
) engine = innodb
  default charset = utf8mb4 comment ='课程';

insert into course (id, name, summary, time, price, image, level, charge, status, enroll, sort, created_at, updated_at)
values ('00000001', '测试课程01', '这是一门测试课程', 7200, 19.9, '', 1, 'C', 'P', 100, 0, now(), now());

alter table `course`
    add column (`teacher_id` char(8) comment '讲师|teacher.id');

-- 大章
drop table if exists `chapter`;
create table `chapter`
(
    `id`        char(8) not null comment 'id',
    `course_id` char(8) comment '课程id',
    `name`      varchar(50) comment '名称',
    primary key (`id`)
) engine = innodb
  default charset = utf8mb4 comment ='大章';

insert into `chapter` (id, course_id, name)
values ('00000001', '00000000', '测试大章01');
insert into `chapter` (id, course_id, name)
values ('00000002', '00000000', '测试大章02');
insert into `chapter` (id, course_id, name)
values ('00000003', '00000000', '测试大章03');
insert into `chapter` (id, course_id, name)
values ('00000004', '00000000', '测试大章04');
insert into `chapter` (id, course_id, name)
values ('00000005', '00000000', '测试大章05');
insert into `chapter` (id, course_id, name)
values ('00000006', '00000000', '测试大章06');
insert into `chapter` (id, course_id, name)
values ('00000007', '00000000', '测试大章07');
insert into `chapter` (id, course_id, name)
values ('00000008', '00000000', '测试大章08');
insert into `chapter` (id, course_id, name)
values ('00000009', '00000000', '测试大章09');
insert into `chapter` (id, course_id, name)
values ('00000010', '00000000', '测试大章10');
insert into `chapter` (id, course_id, name)
values ('00000011', '00000000', '测试大章11');
insert into `chapter` (id, course_id, name)
values ('00000012', '00000000', '测试大章12');
insert into `chapter` (id, course_id, name)
values ('00000013', '00000000', '测试大章13');
insert into `chapter` (id, course_id, name)
values ('00000014', '00000000', '测试大章14');

-- 小节
drop table if exists `section`;
create table `section`
(
    `id`         char(8)     not null default '' comment 'id',
    `title`      varchar(50) not null comment '标题',
    `course_id`  char(8) comment '课程|course.id',
    `chapter_id` char(8) comment '大章|chapter.id',
    `video`      varchar(200) comment '视频',
    `time`       int comment '时长|单位秒',
    `charge`     char(1) comment '收费|枚举[SectionChargeEnum]：CHARGE("C", "收费"),FREE("F", "免费")',
    `sort`       int comment '顺序',
    `created_at` datetime(3) comment '创建时间',
    `updated_at` datetime(3) comment '修改时间',
    primary key (`id`)
) engine = innodb
  default charset = utf8mb4 comment ='小节';

alter table `section`
    add column (`vod` char(32) comment 'vod|阿里云vod');

insert into `section` (id, title, course_id, chapter_id, video, time, charge, sort, created_at, updated_at)
values ('00000001', '测试小节01', '00000001', 'dHV8gubi', '', 500, 'f', 1, now(), now());

-- 分类
drop table if exists `category`;
create table `category`
(
    `id`     char(8)     not null default '' comment 'id',
    `parent` char(8)     not null default '' comment '父id',
    `name`   varchar(50) not null comment '名称',
    `sort`   int comment '顺序',
    primary key (`id`)
) engine = innodb
  default charset = utf8mb4 comment ='分类';

insert into `category` (id, parent, name, sort)
values ('00000100', '00000000', '前端技术', 100);
insert into `category` (id, parent, name, sort)
values ('00000101', '00000100', 'html/css', 101);
insert into `category` (id, parent, name, sort)
values ('00000102', '00000100', 'javascript', 102);
insert into `category` (id, parent, name, sort)
values ('00000103', '00000100', 'vue.js', 103);
insert into `category` (id, parent, name, sort)
values ('00000104', '00000100', 'react.js', 104);
insert into `category` (id, parent, name, sort)
values ('00000105', '00000100', 'angular', 105);
insert into `category` (id, parent, name, sort)
values ('00000106', '00000100', 'node.js', 106);
insert into `category` (id, parent, name, sort)
values ('00000107', '00000100', 'jquery', 107);
insert into `category` (id, parent, name, sort)
values ('00000108', '00000100', '小程序', 108);
insert into `category` (id, parent, name, sort)
values ('00000200', '00000000', '后端技术', 200);
insert into `category` (id, parent, name, sort)
values ('00000201', '00000200', 'java', 201);
insert into `category` (id, parent, name, sort)
values ('00000202', '00000200', 'springboot', 202);
insert into `category` (id, parent, name, sort)
values ('00000203', '00000200', 'spring cloud', 203);
insert into `category` (id, parent, name, sort)
values ('00000204', '00000200', 'ssm', 204);
insert into `category` (id, parent, name, sort)
values ('00000205', '00000200', 'python', 205);
insert into `category` (id, parent, name, sort)
values ('00000206', '00000200', '爬虫', 206);
insert into `category` (id, parent, name, sort)
values ('00000300', '00000000', '移动开发', 300);
insert into `category` (id, parent, name, sort)
values ('00000301', '00000300', 'android', 301);
insert into `category` (id, parent, name, sort)
values ('00000302', '00000300', 'ios', 302);
insert into `category` (id, parent, name, sort)
values ('00000303', '00000300', 'react native', 303);
insert into `category` (id, parent, name, sort)
values ('00000304', '00000300', 'ionic', 304);
insert into `category` (id, parent, name, sort)
values ('00000400', '00000000', '前沿技术', 400);
insert into `category` (id, parent, name, sort)
values ('00000401', '00000400', '微服务', 401);
insert into `category` (id, parent, name, sort)
values ('00000402', '00000400', '区块链', 402);
insert into `category` (id, parent, name, sort)
values ('00000403', '00000400', '机器学习', 403);
insert into `category` (id, parent, name, sort)
values ('00000404', '00000400', '深度学习', 404);
insert into `category` (id, parent, name, sort)
values ('00000405', '00000400', '数据分析&挖掘', 405);
insert into `category` (id, parent, name, sort)
values ('00000500', '00000000', '云计算&大数据', 500);
insert into `category` (id, parent, name, sort)
values ('00000501', '00000500', '大数据', 501);
insert into `category` (id, parent, name, sort)
values ('00000502', '00000500', 'hadoop', 502);
insert into `category` (id, parent, name, sort)
values ('00000503', '00000500', 'spark', 503);
insert into `category` (id, parent, name, sort)
values ('00000504', '00000500', 'hbase', 504);
insert into `category` (id, parent, name, sort)
values ('00000505', '00000500', '阿里云', 505);
insert into `category` (id, parent, name, sort)
values ('00000506', '00000500', 'docker', 506);
insert into `category` (id, parent, name, sort)
values ('00000507', '00000500', 'kubernetes', 507);
insert into `category` (id, parent, name, sort)
values ('00000600', '00000000', '运维&测试', 600);
insert into `category` (id, parent, name, sort)
values ('00000601', '00000600', '运维', 601);
insert into `category` (id, parent, name, sort)
values ('00000602', '00000600', '自动化运维', 602);
insert into `category` (id, parent, name, sort)
values ('00000603', '00000600', '中间件', 603);
insert into `category` (id, parent, name, sort)
values ('00000604', '00000600', 'linux', 604);
insert into `category` (id, parent, name, sort)
values ('00000605', '00000600', '测试', 605);
insert into `category` (id, parent, name, sort)
values ('00000606', '00000600', '自动化测试', 606);
insert into `category` (id, parent, name, sort)
values ('00000700', '00000000', '数据库', 700);
insert into `category` (id, parent, name, sort)
values ('00000701', '00000700', 'mysql', 701);
insert into `category` (id, parent, name, sort)
values ('00000702', '00000700', 'redis', 702);
insert into `category` (id, parent, name, sort)
values ('00000703', '00000700', 'mongodb', 703);

# 课程分类
drop table if exists `course_category`;
create table `course_category`
(
    `id`          char(8) not null default '' comment 'id',
    `course_id`   char(8) comment '课程|course.id',
    `category_id` char(8) comment '分类|course.id',
    primary key (`id`)
) engine = innodb
  default charset = utf8mb4 comment ='课程分类';

-- 课程内容
drop table if exists `course_content`;
create table `course_content`
(
    `id`      char(8)    not null default '' comment '课程id',
    `content` mediumtext not null comment '课程内容',
    primary key (`id`)
) engine = innodb
  default charset = utf8mb4 comment ='课程内容';

-- 课程文件
drop table if exists course_file;
create table `course_file`
(
    `id`        char(8) not null default '' comment 'id',
    `course_id` char(8) not null comment '课程id',
    `url`       varchar(100) comment '地址',
    `name`      varchar(100) comment '文件名',
    `size`      int comment '大小|字节b',
    primary key (`id`)
) engine = innodb
  default charset = utf8mb4 comment ='课程文件';

-- 讲师
drop table if exists `teacher`;
create table `teacher`
(
    `id`       char(8)     not null default '' comment 'id',
    `name`     varchar(50) not null comment '姓名',
    `nickname` varchar(50) comment '昵称',
    `image`    varchar(100) comment '头像',
    `position` varchar(50) comment '职位',
    `motto`    varchar(50) comment '座右铭',
    `intro`    varchar(500) comment '简介',
    primary key (`id`)
) engine = innodb
  default charset = utf8mb4 comment ='讲师';

-- 文件
drop table if exists `file`;
create table `file`
(
    `id`         char(8)      not null default '' comment 'id',
    `path`       varchar(100) not null comment '相对路径',
    `name`       varchar(100) comment '文件名',
    `suffix`     varchar(10) comment '后缀',
    `size`       int comment '大小|字节B',
    `use`        char(1) comment '用途|枚举[FileUseEnum]：COURSE("C", "讲师"), TEACHER("T", "课程")',
    `created_at` datetime(3) comment '创建时间',
    `updated_at` datetime(3) comment '修改时间',
    primary key (`id`),
    unique key `path_unique` (`path`)
) engine = innodb
  default charset = utf8mb4 comment ='文件';

alter table `file`
    add column (`shard_index` int comment '已上传分片');
alter table `file`
    add column (`shard_size` int comment '分片大小|B');
alter table `file`
    add column (`shard_total` int comment '分片总数');
alter table `file`
    add column (`key` varchar(32) comment '文件标识');
alter table `file`
    add unique key key_unique (`key`);
alter table `file`
    add column (`vod` char(32) comment 'vod|阿里云vod');

drop table if exists `user`;
create table `user`
(
    `id`         char(8)     not null default '' comment 'id',
    `login_name` varchar(50) not null comment '登陆名',
    `name`       varchar(50) comment '昵称',
    `password`   char(32)    not null comment '密码',
    primary key (`id`),
    unique key `login_name_unique` (`login_name`)
) engine = innodb
  default charset = utf8mb4 comment ='用户';

# 初始test/test
insert into `user` (id, login_name, name, password)
values ('10000000', 'test', '测试', 'e70e2222a9d67c4f2eae107533359aa4');

-- 资源
drop table if exists `menu`;
create table `menu`
(
    `id`          char(6)      not null default '' comment 'id',
    `title`       varchar(255)          DEFAULT NULL COMMENT '菜单标题',
    `name`        varchar(255)          DEFAULT NULL COMMENT '组件名称',
    `component`   varchar(255)          DEFAULT NULL COMMENT '组件',
    `path`        varchar(200) null comment '路由地址',
    `icon`        varchar(255)          DEFAULT NULL COMMENT '图标',
    `parent`      int(11) comment '父id',
    `request`     varchar(255)          DEFAULT NULL comment '请求接口',
    `create_by`   varchar(255)          DEFAULT NULL COMMENT '创建者',
    `update_by`   varchar(255)          DEFAULT NULL COMMENT '更新者',
    `create_time` datetime              DEFAULT NULL COMMENT '创建日期',
    `update_time` datetime              DEFAULT NULL COMMENT '更新时间',
    primary key (`id`)
) engine = innodb
  default charset = utf8mb4 comment ='菜单';

insert into `menu`
values ('1', '系统管理', null, null, 'system', 'system', null, null, 'admin', 'admin', now(), null);
insert into `menu`
values ('2', '用户管理', 'User', 'system/user/index', 'user', 'peoples', 1, '["/system/user/list"]', 'admin', 'admin', now(), null);
insert into `menu`
values ('3', '权限管理', 'Menu', 'system/menu/index', 'menu', 'menu', 1, '["/system/menu/list"]', 'admin', 'admin', now(), null);
insert into `menu`
values ('4', '角色管理', 'Role', 'system/role/index', 'role', 'role', 1, '["/system/role/list"]', 'admin', 'admin', now(), null);


insert into `menu`
values ('5', '业务管理', null, null, 'business', 'app', null, null, 'admin', 'admin', now(), null);
insert into `menu`
values ('6', '分类管理', 'Category', 'business/category/index', 'category', 'category', 5, '["/business/category/list"]', 'admin', 'admin', now(), null);
insert into `menu`
values ('7', '课程管理', 'Course', 'business/course/index', 'course', 'course', 5, '["/business/course/list"]', 'admin', 'admin', now(), null);
insert into `menu`
values ('8', '讲师管理', 'Teacher', 'business/teacher/index', 'teacher', 'teacher', 5, '["/business/teacher/list"]', 'admin', 'admin', now(), null);

insert into `menu`
values ('9', '保存用户', null, '', '', '', 2, '["/system/user/save","/system/user/save-password"]', 'admin', 'admin', now(), null);
insert into `menu`
values ('10', '删除用户', null, '', '', '', 2, '["/system/user/delete"]', 'admin', 'admin', now(), null);

insert into `menu`
values ('11', '新增权限', null, '', '', '', 3, '["/system/menu/save"]', 'admin', 'admin', now(), null);
insert into `menu`
values ('12', '删除权限', null, '', '', '', 3, '["/system/menu/delete"]', 'admin', 'admin', now(), null);

insert into `menu`
values ('13', '保存角色', null, '', '', '', 4, '["/system/role/save"]', 'admin', 'admin', now(), null);
insert into `menu`
values ('14', '删除角色', null, '', '', '', 4, '["/system/role/delete"]', 'admin', 'admin', now(), null);


insert into `menu`
values ('15', '保存分类', null, '', '', '', 6, '["/business/category/save"]', 'admin', 'admin', now(), null);
insert into `menu`
values ('16', '删除分类', null, '', '', '', 6, '["/business/category/delete"]', 'admin', 'admin', now(), null);

insert into `menu`
values ('17', '新增课程', null, '', '', '', 7, '["/business/course/save"]', 'admin', 'admin', now(), null);
insert into `menu`
values ('18', '删除课程', null, '', '', '', 7, '["/business/course/delete"]', 'admin', 'admin', now(), null);

insert into `menu`
values ('19', '保存讲师', null, '', '', '', 8, '["/business/teacher/save"]', 'admin', 'admin', now(), null);
insert into `menu`
values ('20', '删除讲师', null, '', '', '', 8, '["/business/teacher/delete"]', 'admin', 'admin', now(), null);


drop table if exists `role`;
create table `role`
(
    `id`   char(8)      not null default '' comment 'id',
    `name` varchar(50)  not null comment '角色',
    `desc` varchar(100) not null comment '描述',
    primary key (`id`)
) engine = innodb
  default charset = utf8mb4 comment ='角色';

insert into `role`
values ('00000000', '系统管理员', '管理用户、角色权限');
insert into `role`
values ('00000001', '开发', '维护资源');
insert into `role`
values ('00000002', '业务管理员', '负责业务管理');

drop table if exists `role_menu`;
create table `role_menu`
(
    `id`          char(8) not null default '' comment 'id',
    `role_id`     char(8) not null comment '角色|id',
    `resource_id` char(6) not null comment '资源|id',
    primary key (`id`)
) engine = innodb
  default charset = utf8mb4 comment ='角色资源关联';

insert into `role_menu`
values ('00000000', '00000000', '1');
insert into `role_menu`
values ('00000001', '00000000', '2');
insert into `role_menu`
values ('00000002', '00000000', '3');
insert into `role_menu`
values ('00000003', '00000000', '4');
insert into `role_menu`
values ('00000004', '00000000', '5');
insert into `role_menu`
values ('00000005', '00000000', '6');
insert into `role_menu`
values ('00000006', '00000000', '7');
insert into `role_menu`
values ('00000007', '00000000', '8');
insert into `role_menu`
values ('00000008', '00000000', '9');
insert into `role_menu`
values ('00000009', '00000000', '10');
insert into `role_menu`
values ('00000010', '00000000', '11');
insert into `role_menu`
values ('00000011', '00000000', '12');
insert into `role_menu`
values ('00000012', '00000000', '13');
insert into `role_menu`
values ('00000013', '00000000', '14');
insert into `role_menu`
values ('00000014', '00000000', '15');
insert into `role_menu`
values ('00000015', '00000000', '16');
insert into `role_menu`
values ('00000016', '00000000', '17');
insert into `role_menu`
values ('00000017', '00000000', '18');
insert into `role_menu`
values ('00000018', '00000000', '19');

drop table if exists `role_user`;
create table `role_user`
(
    `id`      char(8) not null default '' comment 'id',
    `role_id` char(8) not null comment '角色|id',
    `user_id` char(8) not null comment '用户|id',
    primary key (`id`)
) engine = innodb
  default charset = utf8mb4 comment ='角色用户关联';

insert into `role_user`
values ('00000000', '00000000', '10000000');

-- 会员
drop table if exists `member`;
create table `member`
(
    `id`            char(8)     not null default '' comment 'id',
    `mobile`        varchar(11) not null comment '手机号',
    `password`      char(32)    not null comment '密码',
    `name`          varchar(50) comment '昵称',
    `photo`         varchar(200) comment '头像url',
    `register_time` datetime(3) comment '注册时间',
    primary key (`id`),
    unique key `mobile_unique` (`mobile`)
) engine = innodb
  default charset = utf8mb4 comment ='会员';

# 初始test/test
insert into `member` (id, mobile, password, name, photo, register_time)
values ('00000000', '12345678901', 'e70e2222a9d67c4f2eae107533359aa4', '测试', null, now());

# 短信验证码
drop table if exists `sms`;
create table `sms`
(
    `id`     char(8)     not null default '' comment 'id',
    `mobile` varchar(50) not null comment '手机号',
    `code`   char(6)     not null comment '验证码',
    `use`    char(1)     not null comment '用途|枚举[SmsUseEnum]：REGISTER("R", "注册"), FORGET("F", "忘记密码")',
    `at`     datetime(3) not null comment '生成时间',
    `status` char(1)     not null comment '用途|枚举[SmsStatusEnum]：USED("U", "已使用"), NOT_USED("N", "未使用")',
    primary key (`id`)
) engine = innodb
  default charset = utf8mb4 comment ='短信验证码';

insert into `sms` (id, mobile, code, `use`, at, status)
values ('00000000', '12345678901', '123456', 'R', now(), 'N');

# 会员课程报名
drop table if exists `member_course`;
create table `member_course`
(
    `id`        char(8)     not null default '' comment 'id',
    `member_id` char(8)     not null comment '会员id',
    `course_id` char(8)     not null comment '课程id',
    `at`        datetime(3) not null comment '报名时间',
    primary key (`id`),
    unique key `member_course_unique` (`member_id`, `course_id`)
) engine = innodb
  default charset = utf8mb4 comment ='会员课程报名';

# ---------------------- 测试

drop table if exists `test`;
create table `test`
(
    `id`   char(8) not null default '' comment 'id',
    `name` varchar(50) comment '名称',
    primary key (`id`)
) engine = innodb
  default charset = utf8mb4 comment ='测试';

insert into `test` (id, name)
values (1, '测试');
insert into `test` (id, name)
values (2, '测试2');