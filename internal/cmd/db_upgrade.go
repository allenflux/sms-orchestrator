package cmd

import (
	"context"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/util/grand"

	"github.com/gogf/gf/v2/frame/g"
)

func DBUpgrade(ctx context.Context) {
	addTableDeviceList(ctx)
	addTableProjectList(ctx)
	addTableSubGroup(ctx)
	addTableSmsMissionReport(ctx)
	addTableSmsMissionRecord(ctx)
	addTableSmsChartLog(ctx)
	addTableUser(ctx)
	addDataToUser(ctx)
	addTableUserProject(ctx)
	addTableRole(ctx)
	addDataToRole(ctx)
	addTablePermission(ctx)
	addDataToPermission(ctx)
	addTableRolePermission(ctx)
	addDataToRolePermission(ctx)
	addTableLog(ctx)
}

func addTableDeviceList(ctx context.Context) {
	needUpgrade := true
	result, err := g.DB().GetAll(ctx, "show tables")
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	list := result.List()
	for _, m := range list {
		if m["Tables_in_sms"] == "device_list" {
			needUpgrade = false
			break
		}
	}
	if needUpgrade {
		sql := "CREATE TABLE `device_list` (`id` int NOT NULL AUTO_INCREMENT,`device_id` int NOT NULL COMMENT '设备id',`number` varchar(255) DEFAULT NULL COMMENT '设备号码',`active_time` timestamp NULL DEFAULT NULL COMMENT '激活时间',`owner_account` varchar(255) NOT NULL COMMENT '所属账号',`assigned_items` varchar(255) NOT NULL COMMENT '所属项目',`sent_status` tinyint NOT NULL COMMENT '发送状态，1-异常 2-占用 3-空闲',`quantity_sent` int NOT NULL COMMENT '发送数量',`quantity_all` int NOT NULL COMMENT '全部数量',`device_status` tinyint NOT NULL COMMENT '设备状态，1-异常 2-正常',`group_name` varchar(255) DEFAULT NULL COMMENT '分组名称',`group_id` int DEFAULT NULL COMMENT '分组id',`created_at` timestamp NOT NULL COMMENT '创建时间',`update_at` timestamp NOT NULL COMMENT '修改时间',`delete_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',PRIMARY KEY (`id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;"
		if _, err = g.DB().Exec(ctx, sql); err != nil {
			g.Log().Error(ctx, err)
		}
	}
}

func addTableProjectList(ctx context.Context) {
	needUpgrade := true
	result, err := g.DB().GetAll(ctx, "show tables")
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	list := result.List()
	for _, m := range list {
		if m["Tables_in_sms"] == "project_list" {
			needUpgrade = false
			break
		}
	}
	if needUpgrade {
		sql := "CREATE TABLE `project_list` (`id` int NOT NULL AUTO_INCREMENT,`project_name` varchar(255) DEFAULT NULL COMMENT '项目名称',`quantity_device` int NOT NULL COMMENT '设备数量',`associated_account` varchar(255) DEFAULT NULL COMMENT '关联账号',`note` varchar(255) DEFAULT NULL COMMENT '备注',`created_at` timestamp NOT NULL COMMENT '创建时间',`update_at` timestamp NOT NULL COMMENT '修改时间',`delete_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',PRIMARY KEY (`id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;"
		if _, err = g.DB().Exec(ctx, sql); err != nil {
			g.Log().Error(ctx, err)
		}
	}
}

func addTableSubGroup(ctx context.Context) {
	needUpgrade := true
	result, err := g.DB().GetAll(ctx, "show tables")
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	list := result.List()
	for _, m := range list {
		if m["Tables_in_sms"] == "sub_group" {
			needUpgrade = false
			break
		}
	}
	if needUpgrade {
		sql := "CREATE TABLE `sub_group` (`id` int NOT NULL AUTO_INCREMENT,`sub_user_id` int NOT NULL COMMENT '子账号id',`project_id` int NOT NULL COMMENT '项目id',`sub_group_name` varchar(255) NOT NULL COMMENT '分组名称',`created_at` timestamp NOT NULL COMMENT '创建时间',`update_at` timestamp NOT NULL COMMENT '修改时间',`delete_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',PRIMARY KEY (`id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;"
		if _, err = g.DB().Exec(ctx, sql); err != nil {
			g.Log().Error(ctx, err)
		}
	}
}

func addTableSmsMissionReport(ctx context.Context) {
	needUpgrade := true
	result, err := g.DB().GetAll(ctx, "show tables")
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	list := result.List()
	for _, m := range list {
		if m["Tables_in_sms"] == "sms_mission_report" {
			needUpgrade = false
			break
		}
	}
	if needUpgrade {
		sql := "CREATE TABLE `sms_mission_report` (`id` int NOT NULL AUTO_INCREMENT,`task_id` int NOT NULL COMMENT '任务id',`task_name` varchar(255) NOT NULL COMMENT '任务名称',`file_name` varchar(255) NOT NULL COMMENT '文件名',`device_quota` varchar(255) NOT NULL COMMENT '执行设备',`task_status` tinyint NOT NULL COMMENT '任务状态，1-待发送 2-发送中 3-已发送 4-已撤销',`sms_quantity` int NOT NULL COMMENT '短信总条数',`surplus_quantity` int NOT NULL COMMENT '剩余数量',`quantity_sent` int NOT NULL COMMENT '已发送数量',`associated_account` varchar(255) NOT NULL COMMENT '所属子账号',`interval_time` varchar(255) NOT NULL COMMENT '间隔时间(秒)',`start_time` timestamp NOT NULL COMMENT '任务开始时间',`created_at` timestamp NOT NULL COMMENT '创建时间',`update_at` timestamp NOT NULL COMMENT '修改时间',`delete_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',PRIMARY KEY (`id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;"
		if _, err = g.DB().Exec(ctx, sql); err != nil {
			g.Log().Error(ctx, err)
		}
	}
}

func addTableSmsMissionRecord(ctx context.Context) {
	needUpgrade := true
	result, err := g.DB().GetAll(ctx, "show tables")
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	list := result.List()
	for _, m := range list {
		if m["Tables_in_sms"] == "sms_mission_record" {
			needUpgrade = false
			break
		}
	}
	if needUpgrade {
		sql := "CREATE TABLE `sms_mission_record` (`id` int NOT NULL AUTO_INCREMENT,`task_name` varchar(255) NOT NULL COMMENT '任务名称',`sub_task_id` int NOT NULL COMMENT '子任务id',`target_phone_number` varchar(255) NOT NULL COMMENT '目标手机号',`device_number` varchar(255) NOT NULL COMMENT '执行设备号',`sms_topic` varchar(255) NOT NULL COMMENT '短信主题',`sms_content` varchar(255) NOT NULL COMMENT '短信内容',`sms_status` varchar(255) NOT NULL COMMENT '短信发送状态，1-失败 2-成功',`associated_account` varchar(255) NOT NULL COMMENT '所属子账号',`project_name` varchar(255) NOT NULL COMMENT '所属项目名称',`project_id` int NOT NULL COMMENT '所属项目id',`start_time` timestamp NOT NULL COMMENT '开始时间',`created_at` timestamp NOT NULL COMMENT '创建时间',`update_at` timestamp NOT NULL COMMENT '修改时间',`delete_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',PRIMARY KEY (`id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;"
		if _, err = g.DB().Exec(ctx, sql); err != nil {
			g.Log().Error(ctx, err)
		}
	}
}

func addTableSmsChartLog(ctx context.Context) {
	needUpgrade := true
	result, err := g.DB().GetAll(ctx, "show tables")
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	list := result.List()
	for _, m := range list {
		if m["Tables_in_sms"] == "sms_chart_log" {
			needUpgrade = false
			break
		}
	}
	if needUpgrade {
		sql := "CREATE TABLE `sms_chart_log` (`id` int NOT NULL AUTO_INCREMENT,`project_name` varchar(255) NOT NULL COMMENT '项目名称',`project_id` int NOT NULL COMMENT '项目id',`target_phone_number` varchar(255) NOT NULL COMMENT '目标手机号',`device_number` varchar(255) NOT NULL COMMENT '执行设备号',`sms_topic` varchar(255) NOT NULL COMMENT '短信主题',`sms_content` varchar(255) NOT NULL COMMENT '短信内容',`sms_status` tinyint NOT NULL COMMENT '短信发送状态，1-失败 2-成',`associated_account` varchar(255) NOT NULL COMMENT '所属子账号',`created_at` timestamp NOT NULL COMMENT '创建时间',`update_at` timestamp NOT NULL COMMENT '修改时间',`delete_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',PRIMARY KEY (`id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;"
		if _, err = g.DB().Exec(ctx, sql); err != nil {
			g.Log().Error(ctx, err)
		}
	}
}

func addTableUser(ctx context.Context) {
	needUpgrade := true
	result, err := g.DB().GetAll(ctx, "show tables")
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	list := result.List()
	for _, m := range list {
		if m["Tables_in_sms"] == "user" {
			needUpgrade = false
			break
		}
	}
	if needUpgrade {
		sql := "CREATE TABLE `user` (`id` int NOT NULL AUTO_INCREMENT,`name` varchar(255) NOT NULL COMMENT '用户名',`password` varchar(255) NOT NULL COMMENT '用户密码',`salt` varchar(255) NOT NULL COMMENT '密码盐',`status` tinyint NOT NULL COMMENT '账号状态，1-停用 2-启用',`role_id` int NOT NULL COMMENT '角色权限id',`system_id` int NOT NULL COMMENT '所属后台id',`note` varchar(255) DEFAULT NULL COMMENT '备注',`created_at` timestamp NOT NULL COMMENT '创建时间',`update_at` timestamp NOT NULL COMMENT '更新时间',`delete_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',PRIMARY KEY (`id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;"
		if _, err = g.DB().Exec(ctx, sql); err != nil {
			g.Log().Error(ctx, err)
		}
	}
}

func addDataToUser(ctx context.Context) {
	count, err := g.DB().Model("user").Where("name=?", "admin").Count()
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	if count == 0 {
		salt := grand.Letters(10)
		password := gmd5.MustEncryptString(gmd5.MustEncryptString("123456") + gmd5.MustEncryptString(salt))
		_, err = g.DB().Model("user").Ctx(ctx).Data(g.Map{
			"name":     "admin",
			"password": password,
			"salt":     salt,
			"status":   2,
			"role_id":  1,
		}).Save()
		if err != nil {
			g.Log().Error(ctx, err)
		}
	}
}

func addTableUserProject(ctx context.Context) {
	needUpgrade := true
	result, err := g.DB().GetAll(ctx, "show tables")
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	list := result.List()
	for _, m := range list {
		if m["Tables_in_sms"] == "user_project" {
			needUpgrade = false
			break
		}
	}
	if needUpgrade {
		sql := "CREATE TABLE `user_project` (`id` int NOT NULL AUTO_INCREMENT,`user_id` int NOT NULL COMMENT '用户id',`project_id` int NOT NULL COMMENT '项目id',`created_at` timestamp NOT NULL COMMENT '创建时间',`update_at` timestamp NOT NULL COMMENT '修改时间',`delete_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',PRIMARY KEY (`id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;"
		if _, err = g.DB().Exec(ctx, sql); err != nil {
			g.Log().Error(ctx, err)
		}
	}
}

func addTableRole(ctx context.Context) {
	needUpgrade := true
	result, err := g.DB().GetAll(ctx, "show tables")
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	list := result.List()
	for _, m := range list {
		if m["Tables_in_sms"] == "role" {
			needUpgrade = false
			break
		}
	}
	if needUpgrade {
		sql := "CREATE TABLE `role` (`id` int NOT NULL AUTO_INCREMENT,`name` varchar(255) NOT NULL COMMENT '角色名称',`created_at` timestamp NOT NULL COMMENT '创建时间',`update_at` timestamp NOT NULL COMMENT '修改时间',`delete_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',PRIMARY KEY (`id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;"
		if _, err = g.DB().Exec(ctx, sql); err != nil {
			g.Log().Error(ctx, err)
		}
	}
}

func addDataToRole(ctx context.Context) {
	count, err := g.DB().Model("role").Where("name=?", "超级管理员").Count()
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	if count == 0 {
		if _, err = g.DB().Model("role").Ctx(ctx).Data(g.Map{
			"id":   1,
			"name": "超级管理员",
		}).Save(); err != nil {
			g.Log().Error(ctx, err)
		}
	}
}

func addTablePermission(ctx context.Context) {
	needUpgrade := true
	result, err := g.DB().GetAll(ctx, "show tables")
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	list := result.List()
	for _, m := range list {
		if m["Tables_in_sms"] == "permission" {
			needUpgrade = false
			break
		}
	}
	if needUpgrade {
		sql := "CREATE TABLE `permission` (`id` int NOT NULL AUTO_INCREMENT,`name` varchar(255) NOT NULL COMMENT '功能名称',`created_ad` timestamp NOT NULL COMMENT '创建时间',`update_at` timestamp NOT NULL COMMENT '修改时间',`delete_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',PRIMARY KEY (`id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;"
		if _, err = g.DB().Exec(ctx, sql); err != nil {
			g.Log().Error(ctx, err)
		}

	}
}

func addDataToPermission(ctx context.Context) {
	count, err := g.DB().Model("permission").Count()
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	if count == 0 {
		if _, err = g.DB().Model("permission").Ctx(ctx).Data(g.List{
			{"id": 1, "name": "群发短信"},
			{"id": 2, "name": "群发记录"},
			{"id": 3, "name": "消息对话"},
			{"id": 4, "name": "设备列表"},
			{"id": 5, "name": "项目管理"},
			{"id": 6, "name": "子后台账号"},
		}).Save(); err != nil {
			g.Log().Error(ctx, err)
		}

	}
}

func addTableRolePermission(ctx context.Context) {
	needUpgrade := true
	result, err := g.DB().GetAll(ctx, "show tables")
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	list := result.List()
	for _, m := range list {
		if m["Tables_in_sms"] == "role_permission" {
			needUpgrade = false
			break
		}
	}
	if needUpgrade {
		sql := "CREATE TABLE `role_permission` (`id` int NOT NULL AUTO_INCREMENT,`role_id` int NOT NULL COMMENT '角色id',`permission_id` int NOT NULL COMMENT '权限id',`created_at` timestamp NOT NULL COMMENT '创建时间',`update_at` timestamp NOT NULL COMMENT '修改时间',`delete_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',PRIMARY KEY (`id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;"
		if _, err = g.DB().Exec(ctx, sql); err != nil {
			g.Log().Error(ctx, err)
		}
	}
}

func addDataToRolePermission(ctx context.Context) {
	count, err := g.DB().Model("role_permission").Where("role_id=?", 1).Count()
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	if count < 6 {
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
	needUpgrade := true
	result, err := g.DB().GetAll(ctx, "show tables")
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	list := result.List()
	for _, m := range list {
		if m["Tables_in_sms"] == "log" {
			needUpgrade = false
			break
		}
	}
	if needUpgrade {
		sql := "CREATE TABLE `log` (`id` int NOT NULL AUTO_INCREMENT,`user_id` int NOT NULL COMMENT '操作者的用户id',`user_name` varchar(255) NOT NULL COMMENT '操作者的用户名',`client_ip` varchar(255) NOT NULL COMMENT '操作ip',`function` varchar(255) NOT NULL COMMENT '操作功能',`note` varchar(255) NOT NULL COMMENT '操作内容',`created_at` timestamp NOT NULL COMMENT '创建时间',`update_at` timestamp NOT NULL COMMENT '修改时间',`delete_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',PRIMARY KEY (`id`),KEY `idx_created_at` (`created_at`) USING BTREE COMMENT '加速根据时间筛选',KEY `idx_user_id` (`user_id`) COMMENT '加速根据用户筛选') ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;"
		if _, err = g.DB().Exec(ctx, sql); err != nil {
			g.Log().Error(ctx, err)
		}
	}
}
