-- ----------------------------
-- Table structure for device
-- ----------------------------
DROP TABLE IF EXISTS `device`;
CREATE TABLE `device`
(
    `id`             bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'Auto-increment primary key',
    `user_id`        bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT 'Account ID',
    `type`           tinyint(3) NOT NULL COMMENT 'Device type, 1: Android; 2: iOS; 3: Windows; 4: MacOS; 5: Web',
    `brand`          varchar(20) NOT NULL COMMENT 'Device manufacturer',
    `model`          varchar(20) NOT NULL COMMENT 'Device model',
    `system_version` varchar(10) NOT NULL COMMENT 'Operating system version',
    `sdk_version`    varchar(10) NOT NULL COMMENT 'App version',
    `status`         tinyint(3) NOT NULL DEFAULT '0' COMMENT 'Online status, 0: Offline; 1: Online',
    `conn_addr`      varchar(25) NOT NULL COMMENT 'Connection server address',
    `client_addr`    varchar(25) NOT NULL COMMENT 'Client address',
    `create_time`    datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Creation time',
    `update_time`    datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Update time',
    PRIMARY KEY (`id`),
    KEY              `idx_user_id` (`user_id`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='Device';

-- ----------------------------
-- Table structure for friend
-- ----------------------------
DROP TABLE IF EXISTS `friend`;
CREATE TABLE `friend`
(
    `id`          bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'Auto-increment primary key',
    `user_id`     bigint(20) unsigned NOT NULL COMMENT 'User ID',
    `friend_id`   bigint(20) unsigned NOT NULL COMMENT 'Friend ID',
    `remarks`     varchar(20)   NOT NULL COMMENT 'Remarks',
    `extra`       varchar(1024) NOT NULL COMMENT 'Additional attributes',
    `status`      tinyint(4) NOT NULL COMMENT 'Status, 1: Request; 2: Accepted',
    `create_time` datetime      NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Creation time',
    `update_time` datetime      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Update time',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_user_id_friend_id` (`user_id`, `friend_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='Friend';

-- ----------------------------
-- Table structure for group
-- ----------------------------
DROP TABLE IF EXISTS `group`;
CREATE TABLE `group`
(
    `id`           bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'Auto-increment primary key',
    `name`         varchar(50)   NOT NULL COMMENT 'Group name',
    `avatar_url`   varchar(255)  NOT NULL COMMENT 'Group avatar URL',
    `introduction` varchar(255)  NOT NULL COMMENT 'Group introduction',
    `user_num`     int(11) NOT NULL DEFAULT '0' COMMENT 'Number of group members',
    `extra`        varchar(1024) NOT NULL COMMENT 'Additional attributes',
    `create_time`  datetime      NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Creation time',
    `update_time`  datetime      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Update time',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='Group';

-- ----------------------------
-- Table structure for group_user
-- ----------------------------
DROP TABLE IF EXISTS `group_user`;
CREATE TABLE `group_user`
(
    `id`          bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'Auto-increment primary key',
    `group_id`    bigint(20) unsigned NOT NULL COMMENT 'Group ID',
    `user_id`     bigint(20) unsigned NOT NULL COMMENT 'User ID',
    `member_type` tinyint(4) NOT NULL COMMENT 'Member type, 1: Admin; 2: Regular member',
    `remarks`     varchar(20)   NOT NULL COMMENT 'Remarks',
    `extra`       varchar(1024) NOT NULL COMMENT 'Additional attributes',
    `status`      tinyint(255) NOT NULL COMMENT 'Status',
    `create_time` datetime      NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Creation time',
    `update_time` datetime      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Update time',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_group_id_user_id` (`group_id`, `user_id`) USING BTREE,
    KEY           `idx_user_id` (`user_id`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='Group member';

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`
(
    `id`           bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'Auto-increment primary key',
    `phone_number` varchar(20)   NOT NULL COMMENT 'Phone number',
    `nickname`     varchar(20)   NOT NULL COMMENT 'Nickname',
    `gender`       tinyint(4) NOT NULL COMMENT 'Gender, 0: Unknown; 1: Male; 2: Female',
    `avatar_url`   varchar(256)  NOT NULL COMMENT 'User avatar URL',
    `extra`        varchar(1024) NOT NULL COMMENT 'Additional attributes',
    `create_time`  datetime      NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Creation time',
    `update_time`  datetime      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Update time',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_phone_number` (`phone_number`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='User';

CREATE TABLE `seq`
(
    `id`          bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'Auto-increment primary key',
    `object_type` tinyint  NOT NULL COMMENT 'Object type, 1: User; 2: Group',
    `object_id`   bigint unsigned NOT NULL COMMENT 'Object ID',
    `seq`         bigint unsigned NOT NULL COMMENT 'Sequence number',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Creation time',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Update time',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_object` (`object_type`,`object_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='Sequence number';

-- ----------------------------
-- Table structure for message
-- ----------------------------
DROP TABLE IF EXISTS `message`;
CREATE TABLE `message`
(
    `id`          bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'Auto-increment primary key',
    `user_id`     bigint(20) unsigned NOT NULL COMMENT 'Associated type ID',
    `request_id`  bigint(20) NOT NULL COMMENT 'Request ID',
    `code`        tinyint(4) NOT NULL COMMENT 'Message type',
    `content`     blob     NOT NULL COMMENT 'Message content',
    `seq`         bigint(20) unsigned NOT NULL COMMENT 'Message sequence number',
    `send_time`   datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT 'Message send time',
    `status`      tinyint(255) NOT NULL DEFAULT '0' COMMENT 'Message status, 0: Unprocessed; 1: Recalled',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Creation time',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Update time',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_user_id_seq` (`user_id`, `seq`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='Message';