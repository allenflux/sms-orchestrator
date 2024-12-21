package cmd

import (
	"context"
	"github.com/gogf/gf/v2/util/grand"
	"sms_backend/library/libUtils"

	"github.com/gogf/gf/v2/frame/g"
)

func DBUpgrade(ctx context.Context) {
	//addTableDeviceList(ctx)
	//addTableProjectList(ctx)
	//addTableSubGroup(ctx)
	//addTableSmsMissionReport(ctx)
	//addTableSmsMissionRecord(ctx)
	//addTableSmsChartLog(ctx)
	//addTableUser(ctx)
	//addDataToUser(ctx)
	//addTableUserProject(ctx)
	//addTableRole(ctx)
	//addDataToRole(ctx)
	//addTablePermission(ctx)
	//addDataToPermission(ctx)
	//addTableRolePermission(ctx)
	//addDataToRolePermission(ctx)
	//addTableLog(ctx)
}

func addTableDeviceList(ctx context.Context) {
	result, err := g.DB().GetAll(ctx, "show tables like 'device_list';")
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	g.Log().Info(ctx, result)
	if len(result) == 0 {
		sql := "CREATE TABLE `device_list` (`id` int NOT NULL AUTO_INCREMENT,`device_id` int NOT NULL COMMENT '设备id',`number` varchar(255) DEFAULT NULL COMMENT '设备号码',`active_time` timestamp NULL DEFAULT NULL COMMENT '激活时间',`owner_account` varchar(255) NOT NULL COMMENT '所属账号',`assigned_items` varchar(255) NOT NULL COMMENT '所属项目',`sent_status` tinyint NOT NULL COMMENT '发送状态，1-异常 2-占用 3-空闲',`quantity_sent` int NOT NULL COMMENT '发送数量',`quantity_all` int NOT NULL COMMENT '全部数量',`device_status` tinyint NOT NULL COMMENT '设备状态，1-异常 2-正常',`group_name` varchar(255) DEFAULT NULL COMMENT '分组名称',`group_id` int DEFAULT NULL COMMENT '分组id',`created_at` timestamp NOT NULL COMMENT '创建时间',`update_at` timestamp NOT NULL COMMENT '修改时间',`delete_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',PRIMARY KEY (`id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;"
		if _, err = g.DB().Exec(ctx, sql); err != nil {
			g.Log().Error(ctx, err)
			return
		}
	}
}

func addTableProjectList(ctx context.Context) {
	result, err := g.DB().GetAll(ctx, "show tables like 'project_list'")
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	g.Log().Info(ctx, result)
	if len(result) == 0 {
		sql := "CREATE TABLE `project_list` (`id` int NOT NULL AUTO_INCREMENT,`project_name` varchar(255) DEFAULT NULL COMMENT '项目名称',`quantity_device` int NOT NULL COMMENT '设备数量',`associated_account` varchar(255) DEFAULT NULL COMMENT '关联账号',`note` varchar(255) DEFAULT NULL COMMENT '备注',`created_at` timestamp NOT NULL COMMENT '创建时间',`update_at` timestamp NOT NULL COMMENT '修改时间',`delete_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',PRIMARY KEY (`id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;"
		if _, err = g.DB().Exec(ctx, sql); err != nil {
			g.Log().Error(ctx, err)
			return
		}
	}
}

func addTableSubGroup(ctx context.Context) {
	result, err := g.DB().GetAll(ctx, "show tables like 'sub_group'")
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	g.Log().Info(ctx, result)
	if len(result) == 0 {
		sql := "CREATE TABLE `sub_group` (`id` int NOT NULL AUTO_INCREMENT,`sub_user_id` int NOT NULL COMMENT '子账号id',`project_id` int NOT NULL COMMENT '项目id',`sub_group_name` varchar(255) NOT NULL COMMENT '分组名称',`created_at` timestamp NOT NULL COMMENT '创建时间',`update_at` timestamp NOT NULL COMMENT '修改时间',`delete_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',PRIMARY KEY (`id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;"
		if _, err = g.DB().Exec(ctx, sql); err != nil {
			g.Log().Error(ctx, err)
			return
		}
	}
}

func addTableSmsMissionReport(ctx context.Context) {
	result, err := g.DB().GetAll(ctx, "show tables like 'sms_mission_report'")
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	g.Log().Info(ctx, result)
	if len(result) == 0 {
		sql := "CREATE TABLE `sms_mission_report` (`id` int NOT NULL AUTO_INCREMENT,`task_id` int NOT NULL COMMENT '任务id',`task_name` varchar(255) NOT NULL COMMENT '任务名称',`file_name` varchar(255) NOT NULL COMMENT '文件名',`device_quota` varchar(255) NOT NULL COMMENT '执行设备',`task_status` tinyint NOT NULL COMMENT '任务状态，1-待发送 2-发送中 3-已发送 4-已撤销',`sms_quantity` int NOT NULL COMMENT '短信总条数',`surplus_quantity` int NOT NULL COMMENT '剩余数量',`quantity_sent` int NOT NULL COMMENT '已发送数量',`associated_account` varchar(255) NOT NULL COMMENT '所属子账号',`interval_time` varchar(255) NOT NULL COMMENT '间隔时间(秒)',`start_time` timestamp NOT NULL COMMENT '任务开始时间',`created_at` timestamp NOT NULL COMMENT '创建时间',`update_at` timestamp NOT NULL COMMENT '修改时间',`delete_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',PRIMARY KEY (`id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;"
		if _, err = g.DB().Exec(ctx, sql); err != nil {
			g.Log().Error(ctx, err)
			return
		}
	}
}

func addTableSmsMissionRecord(ctx context.Context) {
	result, err := g.DB().GetAll(ctx, "show tables like 'sms_mission_record'")
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	g.Log().Info(ctx, result)
	if len(result) == 0 {
		sql := "CREATE TABLE `sms_mission_record` (`id` int NOT NULL AUTO_INCREMENT,`task_name` varchar(255) NOT NULL COMMENT '任务名称',`sub_task_id` int NOT NULL COMMENT '子任务id',`target_phone_number` varchar(255) NOT NULL COMMENT '目标手机号',`device_number` varchar(255) NOT NULL COMMENT '执行设备号',`sms_topic` varchar(255) NOT NULL COMMENT '短信主题',`sms_content` varchar(255) NOT NULL COMMENT '短信内容',`sms_status` varchar(255) NOT NULL COMMENT '短信发送状态，1-失败 2-成功',`associated_account` varchar(255) NOT NULL COMMENT '所属子账号',`project_name` varchar(255) NOT NULL COMMENT '所属项目名称',`project_id` int NOT NULL COMMENT '所属项目id',`start_time` timestamp NOT NULL COMMENT '开始时间',`created_at` timestamp NOT NULL COMMENT '创建时间',`update_at` timestamp NOT NULL COMMENT '修改时间',`delete_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',PRIMARY KEY (`id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;"
		if _, err = g.DB().Exec(ctx, sql); err != nil {
			g.Log().Error(ctx, err)
			return
		}
	}
}

func addTableSmsChartLog(ctx context.Context) {
	result, err := g.DB().GetAll(ctx, "show tables like 'sms_chart_log'")
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	g.Log().Info(ctx, result)
	if len(result) == 0 {
		sql := "CREATE TABLE `sms_chart_log` (`id` int NOT NULL AUTO_INCREMENT,`project_name` varchar(255) NOT NULL COMMENT '项目名称',`project_id` int NOT NULL COMMENT '项目id',`target_phone_number` varchar(255) NOT NULL COMMENT '目标手机号',`device_number` varchar(255) NOT NULL COMMENT '执行设备号',`sms_topic` varchar(255) NOT NULL COMMENT '短信主题',`sms_content` varchar(255) NOT NULL COMMENT '短信内容',`sms_status` tinyint NOT NULL COMMENT '短信发送状态，1-失败 2-成',`associated_account` varchar(255) NOT NULL COMMENT '所属子账号',`created_at` timestamp NOT NULL COMMENT '创建时间',`update_at` timestamp NOT NULL COMMENT '修改时间',`delete_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',PRIMARY KEY (`id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;"
		if _, err = g.DB().Exec(ctx, sql); err != nil {
			g.Log().Error(ctx, err)
			return
		}
	}
}

func addTableUser(ctx context.Context) {
	result, err := g.DB().GetAll(ctx, "show tables like 'user'")
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	g.Log().Info(ctx, result)
	if len(result) == 0 {
		sql := "CREATE TABLE `user` (`id` int NOT NULL AUTO_INCREMENT,`name` varchar(255) NOT NULL COMMENT '用户名',`password` varchar(255) NOT NULL COMMENT '用户密码',`salt` varchar(255) NOT NULL COMMENT '密码盐',`status` tinyint NOT NULL COMMENT '账号状态，1-停用 2-启用',`role_id` int NOT NULL COMMENT '角色权限id',`system_id` int NOT NULL COMMENT '所属后台id',`note` varchar(255) DEFAULT NULL COMMENT '备注',`created_at` timestamp NOT NULL COMMENT '创建时间',`update_at` timestamp NOT NULL COMMENT '更新时间',`delete_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',PRIMARY KEY (`id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;"
		if _, err = g.DB().Exec(ctx, sql); err != nil {
			g.Log().Error(ctx, err)
			return
		}
	}
}

func addDataToUser(ctx context.Context) {
	exists, err := g.DB().Model("user").Where("name=?", "admin1").Exist()
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	if !exists {
		salt := grand.Letters(10)
		password := libUtils.EncryptPassword("123456", salt)
		_, err = g.DB().Model("user").Ctx(ctx).Data(g.Map{
			"id":        1,
			"name":      "surper",
			"password":  password,
			"salt":      salt,
			"status":    2,
			"role_id":   1,
			"system_id": 1,
			"note":      "总后台管理员",
		}).Save()
		if err != nil {
			g.Log().Error(ctx, err)
		}
	}
}

func addTableUserProject(ctx context.Context) {
	result, err := g.DB().GetAll(ctx, "show tables like 'user_project'")
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	g.Log().Info(ctx, result)
	if len(result) == 0 {
		sql := "CREATE TABLE `user_project` (`id` int NOT NULL AUTO_INCREMENT,`user_id` int NOT NULL COMMENT '用户id',`project_id` int NOT NULL COMMENT '项目id',`created_at` timestamp NOT NULL COMMENT '创建时间',`update_at` timestamp NOT NULL COMMENT '修改时间',`delete_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',PRIMARY KEY (`id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;"
		if _, err = g.DB().Exec(ctx, sql); err != nil {
			g.Log().Error(ctx, err)
		}
	}
}

func addTableRole(ctx context.Context) {
	result, err := g.DB().GetAll(ctx, "show tables like 'role'")
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	g.Log().Info(ctx, result)
	if len(result) == 0 {
		sql := "CREATE TABLE `role` (`id` int NOT NULL AUTO_INCREMENT,`name` varchar(255) NOT NULL COMMENT '角色名称',`note` varchar(255) DEFAULT NULL COMMENT '备注',`created_at` timestamp NOT NULL COMMENT '创建时间',`update_at` timestamp NOT NULL COMMENT '修改时间',`delete_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',PRIMARY KEY (`id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;"
		if _, err = g.DB().Exec(ctx, sql); err != nil {
			g.Log().Error(ctx, err)
		}
	}
}

func addDataToRole(ctx context.Context) {
	exist, err := g.DB().Model("role").Where("name=?", "总后台管理员").Exist()
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	if !exist {
		if _, err = g.DB().Model("role").Ctx(ctx).Data(g.Map{
			"id":   1,
			"name": "总后台管理员",
			"note": "固定为admin",
		}).Save(); err != nil {
			g.Log().Error(ctx, err)
		}
	}
}

func addTablePermission(ctx context.Context) {
	result, err := g.DB().GetAll(ctx, "show tables like 'permission'")
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	g.Log().Info(ctx, result)
	if len(result) == 0 {
		sql := "CREATE TABLE `permission` (`id` int NOT NULL AUTO_INCREMENT,`name` varchar(255) NOT NULL COMMENT '功能名称',`redirect` varchar(255) DEFAULT NULL COMMENT '路由重定向地址',`created_at` timestamp NOT NULL COMMENT '创建时间',`update_at` timestamp NOT NULL COMMENT '修改时间',`delete_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',PRIMARY KEY (`id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;"
		if _, err = g.DB().Exec(ctx, sql); err != nil {
			g.Log().Error(ctx, err)
			return
		}

	}
}

func addDataToPermission(ctx context.Context) {
	exist, err := g.DB().Model("permission").Exist()
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	if !exist {
		if _, err = g.DB().Model("permission").Ctx(ctx).Data(g.List{
			{"id": 1, "name": "群发短信"},
			{"id": 2, "name": "群发记录"},
			{"id": 3, "name": "消息对话"},
			{"id": 4, "name": "设备列表"},
			{"id": 5, "name": "项目管理"},
			{"id": 6, "name": "子后台账号"},
			{"id": 7, "name": "账号管理"},
			{"id": 8, "name": "角色管理"},
			{"id": 9, "name": "操作日志"},
		}).Save(); err != nil {
			g.Log().Error(ctx, err)
		}
	}
}

func addTableRolePermission(ctx context.Context) {
	result, err := g.DB().GetAll(ctx, "show tables like 'role_permission'")
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	g.Log().Info(ctx, result)
	if len(result) == 0 {
		sql := "CREATE TABLE `role_permission` (`id` int NOT NULL AUTO_INCREMENT,`role_id` int NOT NULL COMMENT '角色id',`permission_id` int NOT NULL COMMENT '权限id',`created_at` timestamp NOT NULL COMMENT '创建时间',`update_at` timestamp NOT NULL COMMENT '修改时间',`delete_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',PRIMARY KEY (`id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;"
		if _, err = g.DB().Exec(ctx, sql); err != nil {
			g.Log().Error(ctx, err)
			return
		}
	}
}

func addDataToRolePermission(ctx context.Context) {
	exist, err := g.DB().Model("role_permission").Where("role_id=?", 1).Exist()
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	if !exist {
		if _, err = g.DB().Model("role_permission").Ctx(ctx).Data(g.List{
			{"role_id": 1, "permission_id": 1},
			{"role_id": 1, "permission_id": 2},
			{"role_id": 1, "permission_id": 3},
			{"role_id": 1, "permission_id": 4},
			{"role_id": 1, "permission_id": 5},
			{"role_id": 1, "permission_id": 6},
		}).Save(); err != nil {
			g.Log().Error(ctx, err)
		}
	}
}

func addTableLog(ctx context.Context) {
	result, err := g.DB().GetAll(ctx, "show tables like 'log'")
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	g.Log().Info(ctx, result)
	if len(result) == 0 {
		sql := "CREATE TABLE `log` (`id` int NOT NULL AUTO_INCREMENT,`user_id` int NOT NULL COMMENT '操作者的用户id',`user_name` varchar(255) NOT NULL COMMENT '操作者的用户名',`client_ip` varchar(255) NOT NULL COMMENT '操作ip',`function` varchar(255) NOT NULL COMMENT '操作功能',`note` varchar(255) NOT NULL COMMENT '操作内容',`created_at` timestamp NOT NULL COMMENT '创建时间',`update_at` timestamp NOT NULL COMMENT '修改时间',`delete_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',PRIMARY KEY (`id`),KEY `idx_created_at` (`created_at`) USING BTREE COMMENT '加速根据时间筛选',KEY `idx_user_id` (`user_id`) COMMENT '加速根据用户筛选') ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;"
		if _, err = g.DB().Exec(ctx, sql); err != nil {
			g.Log().Error(ctx, err)
			return
		}
	}
}
