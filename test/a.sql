CREATE TABLE `t_online_class_course` (
  `t_class_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '小班ID',
  `t_seqno` tinyint(4) DEFAULT '0' COMMENT '内部序号',
  `t_cid` int(11) DEFAULT NULL COMMENT '课程ID',
  `t_term_id` int(11) DEFAULT NULL COMMENT '班级ID',
  `t_tu_list` text COMMENT '助教列表',
  `t_cg_list` text COMMENT '付费服务群',
  `t_student_total` int(11) DEFAULT '0' COMMENT '限制人数',
  `t_real_apply_total` int(255) DEFAULT '0' COMMENT '真实报名人数',
  `t_apply_total` int(255) DEFAULT '0' COMMENT '总的报名人数',
  `t_create_time` bigint(20) DEFAULT NULL COMMENT '创建时间',
  `t_last_update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后更新时间',
  `t_del_flag` tinyint(4) NOT NULL DEFAULT '0',
  PRIMARY KEY (`t_class_id`)
) ENGINE=InnoDB AUTO_INCREMENT=34944 DEFAULT CHARSET=utf8 COMMENT '测试1';

CREATE TABLE `t_online_class_course_a` (
  `t_class_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '小班ID',
  `t_seqno` tinyint(4) DEFAULT '0' COMMENT '内部序号',
  `t_cid` int(11) DEFAULT NULL COMMENT '课程ID',
  `t_term_id` int(11) DEFAULT NULL COMMENT '班级ID',
  `t_tu_list` text COMMENT '助教列表',
  `t_cg_list` text COMMENT '付费服务群',
  `t_student_total` int(11) DEFAULT '0' COMMENT '限制人数',
  `t_real_apply_total` int(255) DEFAULT '0' COMMENT '真实报名人数',
  `t_apply_total` int(255) DEFAULT '0' COMMENT '总的报名人数',
  `t_create_time` bigint(20) DEFAULT NULL COMMENT '创建时间',
  `t_last_update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后更新时间',
  `t_del_flag` tinyint(4) NOT NULL DEFAULT '0',
  PRIMARY KEY (`t_class_id`)
) ENGINE=InnoDB AUTO_INCREMENT=34944 DEFAULT CHARSET=utf8 COMMENT '测试2';