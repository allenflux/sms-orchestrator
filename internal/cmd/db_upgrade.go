package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

func DBUpgrade(ctx context.Context) {
	addTableDeviceList(ctx)
	addTableProjectList(ctx)
	addTableSubGroup(ctx)
	addTableSmsMissionReport(ctx)
	addTableSmsMissionRecord(ctx)
	addTableSmsChartLog(ctx)
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
		_, err = g.DB().Exec(ctx, sql)
		if err != nil {
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
		_, err = g.DB().Exec(ctx, sql)
		if err != nil {
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
		_, err = g.DB().Exec(ctx, sql)
		if err != nil {
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
		_, err = g.DB().Exec(ctx, sql)
		if err != nil {
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
		_, err = g.DB().Exec(ctx, sql)
		if err != nil {
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
		_, err = g.DB().Exec(ctx, sql)
		if err != nil {
			g.Log().Error(ctx, err)
		}
	}
}
