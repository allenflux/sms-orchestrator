# 数据库表信息文档

## 表信息概览

| 表名                | 行数     | 数据大小       | 存储引擎 | 创建日期        |
|---------------------|----------|---------------|----------|-----------------|
| device_list         | 83       | 16 KB         | InnoDB   | 2024-12-12      |
| log                 | 162      | 16 KB         | InnoDB   | 2024-12-12      |
| permission          | 12       | 16 KB         | InnoDB   | 2024-12-12      |
| project_list        | 12       | 16 KB         | InnoDB   | 2024-12-12      |
| role                | 5        | 16 KB         | InnoDB   | 2024-12-12      |
| role_permission     | 34       | 16 KB         | InnoDB   | 2024-12-12      |
| sms_chart_log       | 1,030,000+| 472.9 MB      | InnoDB   | 2024-12-12      |
| sms_mission_record  | 60,000+  | 21.5 MB       | InnoDB   | 2024-12-12      |
| sms_mission_report  | 107      | 48 KB         | InnoDB   | 2024-12-12      |
| sms_trace_task      | 38,000+  | 3.5 MB        | InnoDB   | 2024-12-12      |
| sub_group           | 20       | 16 KB         | InnoDB   | 2024-12-12      |
| user                | 5        | 16 KB         | InnoDB   | 2024-12-12      |

---

**说明：**
- **行数**：每个表中的数据行数（部分为估算值，如 `sms_chart_log` 和 `sms_mission_record`）。
- **数据大小**：表的存储大小。
- **存储引擎**：所有表均使用 InnoDB。
- **创建日期**：所有表的创建时间相同，为 2024 年 12 月 12 日。

# 数据库表结构文档



## device_list 表
| 字段名            | 类型              | 说明           |
|-------------------|-------------------|----------------|
| id                | int              |                |
| device_id         | varchar(255)     | 设备id         |
| number            | varchar(255)     | 设备号码       |
| active_time       | timestamp        | 激活时间       |
| owner_account     | varchar(255)     | 子账号         |
| assigned_items    | varchar(255)     | 所属项目       |
| sent_status       | tinyint          | 发送状态       |
| quantity_sent     | int              | 发送数量       |
| quantity_all      | int              | 全部数量       |
| device_status     | tinyint          | 设备状态       |
| group_name        | varchar(255)     | 分组名称       |
| group_id          | int              | 分组id         |
| created_at        | timestamp        | 创建时间       |
| update_at         | timestamp        | 修改时间       |
| delete_at         | timestamp        | 删除时间       |
| device_number     | varchar(255)     | 设备序列号     |
| start_at          | timestamp        | 发送时间       |
| assigned_items_id | int              | 项目id         |
| owner_account_id  | int              | 子账号id       |

---

## log 表
| 字段名      | 类型          | 说明             |
|-------------|---------------|------------------|
| id          | int           |                  |
| user_id     | int           | 操作者用户id     |
| user_name   | varchar(255)  | 操作者用户名     |
| client_ip   | varchar(255)  | 操作ip           |
| function    | varchar(255)  | 操作功能         |
| note        | varchar(255)  | 操作内容         |
| system_id   | int           | 后台id           |
| created_at  | timestamp     | 创建时间         |
| update_at   | timestamp     | 修改时间         |
| delete_at   | timestamp     | 删除时间         |

---

## permission 表
| 字段名     | 类型              | 说明         |
|------------|-------------------|--------------|
| id         | int              |              |
| pid        | int              | 功能父id     |
| name       | varchar(255)     | 规则名称     |
| title      | varchar(255)     | 功能名称     |
| icon       | varchar(255)     | icon图标     |
| menu_type  | tinyint          | 菜单类型     |
| is_hide    | tinyint          | 显示状态     |
| path       | varchar(255)     | 路由地址     |
| component  | varchar(255)     | 组件路径     |
| is_cached  | tinyint          | 是否缓存     |
| created_at | timestamp        | 创建时间     |
| update_at  | timestamp        | 修改时间     |
| delete_at  | timestamp        | 删除时间     |

---

## project_list 表
| 字段名                | 类型          | 说明           |
|-----------------------|---------------|----------------|
| id                   | int           |                |
| project_name         | varchar(255)  | 项目名称       |
| quantity_device      | int           | 设备数量       |
| associated_account   | varchar(255)  | 关联账号       |
| note                 | varchar(255)  | 备注           |
| created_at           | timestamp     | 创建时间       |
| update_at            | timestamp     | 修改时间       |
| delete_at            | timestamp     | 删除时间       |
| associated_account_id| int           |                |

---

## role 表
| 字段名     | 类型          | 说明       |
|------------|---------------|------------|
| id         | int           |            |
| name       | varchar(255)  | 角色名称   |
| note       | varchar(255)  | 备注       |
| created_at | timestamp     | 创建时间   |
| update_at  | timestamp     | 修改时间   |
| delete_at  | timestamp     | 删除时间   |

---

## role_permission 表
| 字段名        | 类型          | 说明       |
|---------------|---------------|------------|
| id            | int           |            |
| role_id       | int           | 角色id     |
| permission_id | int           | 权限id     |
| created_at    | timestamp     | 创建时间   |
| update_at     | timestamp     | 修改时间   |
| delete_at     | timestamp     | 删除时间   |

---

## sms_chart_log 表
| 字段名              | 类型          | 说明           |
|---------------------|---------------|----------------|
| id                 | int           |                |
| project_name       | varchar(255)  | 项目名称       |
| project_id         | int           | 项目id         |
| target_phone_number| varchar(255)  | 目标手机号     |
| device_number      | varchar(255)  | 执行设备号     |
| sms_topic          | varchar(255)  | 短信主题       |
| sms_content        | varchar(255)  | 短信内容       |
| sms_status         | tinyint       | 短信发送状态   |
| associated_account | varchar(255)  | 所属子账号     |
| created_at         | timestamp     | 创建时间       |
| update_at          | timestamp     | 修改时间       |
| delete_at          | timestamp     | 删除时间       |
| sent_or_receive    | int           | 短信方向       |
| associated_account_id| int         | 子账号id       |
| row_hash           | varchar(255)  | log哈希值      |
| task_id            | int           | 任务id         |

---

## sms_mission_record 表
| 字段名              | 类型          | 说明           |
|---------------------|---------------|----------------|
| id                 | int           |                |
| task_name          | varchar(255)  | 任务名称       |
| sub_task_id        | int           | 子任务id       |
| target_phone_number| varchar(255)  | 目标手机号     |
| device_number      | varchar(255)  | 执行设备号     |
| sms_topic          | varchar(255)  | 短信主题       |
| sms_content        | varchar(255)  | 短信内容       |
| sms_status         | varchar(255)  | 短信发送状态   |
| associated_account | varchar(255)  | 所属子账号     |
| project_name       | varchar(255)  | 所属项目名称   |
| project_id         | int           | 所属项目id     |
| start_time         | timestamp     | 开始时间       |
| created_at         | timestamp     | 创建时间       |
| update_at          | timestamp     | 修改时间       |
| delete_at          | timestamp     | 删除时间       |
| associated_account_id| int         | 子账号id       |
| row_hash           | varchar(255)  | hash 串         |
| reason             | varchar(255)  | 失败原因       |

---

## sms_mission_report 表
| 字段名              | 类型          | 说明           |
|---------------------|---------------|----------------|
| id                 | int           |                |
| project_id         | int           | 项目id         |
| task_name          | varchar(255)  | 任务名称       |
| file_name          | varchar(255)  | 文件名         |
| device_quota       | varchar(255)  | 执行设备       |
| task_status        | tinyint       | 任务状态       |
| sms_quantity       | int           | 短信总条数     |
| surplus_quantity   | int           | 剩余数量       |
| quantity_sent      | int           | 已发送数量     |
| associated_account | varchar(255)  | 所属子账号     |
| interval_time      | varchar(255)  | 间隔时间       |
| start_time         | timestamp     | 任务开始时间   |
| created_at         | timestamp     | 创建时间       |
| update_at          | timestamp     | 修改时间       |
| delete_at          | timestamp     | 删除时间       |
| project_name       | varchar(255)  | 项目名称       |
| associated_account_id| int         | 子账号id       |
| group_id           | int           | 分组id         |
| sent_success_quantity| int         | 发送成功数量   |
| sent_fail_quantity | int           | 发送失败数量   |

---

## sms_trace_task 表
| 字段名        | 类型          | 说明           |
|---------------|---------------|----------------|
| id            | int           |                |
| target_number | varchar(255)  | 目标号码       |
| device_number | varchar(255)  | 设备号         |
| task_id       | int           | 任务id         |
| create_at     | timestamp     | 创建时间       |
| update_at     | timestamp     | 修改时间       |
| delete_at     | timestamp     | 删除时间       |

---

## sub_group 表
| 字段名         | 类型          | 说明         |
|----------------|---------------|--------------|
| id            | int           |              |
| sub_user_id   | int           | 子账号id     |
| project_id    | int           | 项目id       |
| sub_group_name| varchar(255)  | 分组名称     |
| created_at    | timestamp     | 创建时间     |
| update_at     | timestamp     | 修改时间     |
| delete_at     | timestamp     | 删除时间     |

---

## user 表
| 字段名       | 类型          | 说明         |
|--------------|---------------|--------------|
| id          | int           |              |
| name        | varchar(255)  | 用户名       |
| password    | varchar(255)  | 用户密码     |
| salt        | varchar(255)  | 密码盐       |
| status      | tinyint       | 账号状态     |
| role_id     | int           | 角色权限id   |
| system_id   | int           | 所属后台id   |
| note        | varchar(255)  | 备注         |
| created_at  | timestamp     | 创建时间     |
| update_at   | timestamp     | 修改时间     |
| delete_at   | timestamp     | 删除时间     |