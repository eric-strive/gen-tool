#### db_ad_system.ts_ad_cs_white 
白名单表

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| ---- | ---- | ---- | ---- | ---- | ---- | ---- | ---- |
| 1 | id | 主键 | int(11) | PRI | NO | auto_increment |  |
| 2 | name |  | varchar(128) |  | YES |  |  |
| 3 | os | 系统名称 | varchar(16) |  | NO |  |  |
| 4 | idfa | idfa | varchar(64) |  | NO |  |  |
| 5 | oaid |  | varchar(128) |  | YES |  |  |
| 6 | device_id | device_id | varchar(64) |  | NO |  |  |
| 7 | user_id | 用户ID | varchar(64) |  | NO |  |  |
| 8 | type |  | varchar(64) |  | YES |  |  |
| 9 | verify_sign | 对应校验的标识cs就是计划ID、api就是api标识、SDK就是SDK标识 | varchar(64) | MUL | NO |  |  |
| 10 | description |  | varchar(64) |  | YES |  |  |
| 11 | remark |  | varchar(64) |  | YES |  |  |
| 12 | is_deleted |  | tinyint(4) |  | YES |  |  |
| 13 | create_time | 创建时间 | bigint(20) |  | NO |  | 0 |
| 14 | create_by |  | varchar(64) |  | YES |  |  |
| 15 | update_by |  | varchar(64) |  | YES |  |  |
