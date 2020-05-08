
# timebook数据库建库

### 1. 资料

* [MySQL UPDATE 更新 | 菜鸟教程](https://www.runoob.com/mysql/mysql-update-query.html)
* [MySQL 列属性修改操作_数据库_u012149181的博客-CSDN博客](https://blog.csdn.net/u012149181/article/details/80336793)
* [胶囊日记](http://timepill.net/)
* [GORM入门指南 | 李文周的博客](https://www.liwenzhou.com/posts/Go/gorm/)

### 2. 建立数据库

#### 2.1 数据库整体规划

![](https://tva1.sinaimg.cn/large/0082zybpgy1gbydkwmd72j31l40k4tb4.jpg)

**time-book.sql:**


```sql
DROP DATABASE IF EXISTS `time book`;

CREATE DATABASE `timebook` DEFAULT CHARSET utf8;

USE `timebook`;

DROP TABLE IF EXISTS `user`;

CREATE TABLE  IF NOT EXISTS `user` (
  `user_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '用户id',
  `email` varchar(128) NOT NULL COMMENT '邮箱',
  `password` varchar(256) NOT NULL COMMENT '密码',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '注册时间',
  `update_time` timestamp NULL DEFAULT NULL COMMENT '更新时间',
   `status` int(10) NOT NULL DEFAULT '1' COMMENT '状态，正常为1，会员为2，试用3，过期4',
    PRIMARY KEY ( `user_id` )
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `userinfo`;

CREATE TABLE `userinfo` (
  `user_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '用户 ID',
  `username` varchar(256) NOT NULL COMMENT '用户昵称',
  `header_img` varchar(256) NOT NULL COMMENT '用户头像',
  `bio` longtext NOT NULL COMMENT 'bio内容',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '发布时间',
  `update_time` timestamp NULL DEFAULT NULL COMMENT '更新时间',
   PRIMARY KEY ( `user_id` )
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `note`;

CREATE TABLE `note` (
  `note_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'note ID',
  `category_id` bigint(20) NOT NULL COMMENT '分类 ID',
  `content` longtext NOT NULL COMMENT 'note内容',
  `img` varchar(1024) NOT NULL COMMENT 'note图片',
  `view_count` int(255) NOT NULL COMMENT '阅读次数',
  `comment_count` int(255) NOT NULL COMMENT '评论次数',
  `user_id` bigint(20) NOT NULL  COMMENT '用户id',
  `status` int(10) NOT NULL DEFAULT '1' COMMENT '状态，正常为1，删除2，和谐3',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '发布时间',
  `update_time` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`note_id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `comment`;

CREATE TABLE `comment` (
  `comment_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '评论 ID',
  `note_id` bigint(20) NOT NULL  COMMENT 'note ID',
  `user_id` bigint(20) NOT NULL  COMMENT '作者 用户 ID',
  `feed_user_id` bigint(20) NOT NULL  COMMENT '被艾特的 用户 ID',
  `content` text NOT NULL COMMENT '评论内容',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '评论发布时间',
  `status` int(255) NOT NULL DEFAULT '1' COMMENT '评论状态: 正常1, 删除2',
  PRIMARY KEY (`comment_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
```

1. user 表，可以做登录注册了。
2. note 表，可以做 note的 新建 修改 删除了。
3. comment 表， 可以做评论了。
4. userinfo 表，可以个人主页了。

