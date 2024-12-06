package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

func DBUpgrade(ctx context.Context) {
	addHasCommitToUserDetail(ctx)
	addTableFundingCard(ctx)
	addTableRechargeList(ctx)
	addBalanceToApplyList(ctx)
	addFileToCardLevel(ctx)
	addTransactionRecord(ctx)
	addNoteToApplyList(ctx)
	addKYCStatusToAppUser(ctx)
	addUserCardStatusToAppUser(ctx)
	addBankCardStatusToAppUser(ctx)
	// addApplyStatusToAppUser(ctx)
}

func addHasCommitToUserDetail(ctx context.Context) {
	needUpgrade := true
	result, err := g.DB().GetAll(ctx, "desc app_user_detail")
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	list := result.List()
	for _, m := range list {
		if m["Field"] == "has_commit" {
			needUpgrade = false
			break
		}
	}
	if needUpgrade {
		sql := "ALTER TABLE `app_user_detail` CHANGE `has_commit` `has_commit` INT  NOT NULL  DEFAULT 1  COMMENT '是否提交到万事达,1否2是';"
		_, err = g.DB().Exec(ctx, sql)
		if err != nil {
			g.Log().Error(ctx, err)
		}
	}
}

func addTableFundingCard(ctx context.Context) {
	needUpgrade := true
	result, err := g.DB().GetAll(ctx, "show tables")
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	list := result.List()
	for _, m := range list {
		if m["Tables_in_upay"] == "funding_card" {
			needUpgrade = false
			break
		}
	}
	if needUpgrade {
		sql := "CREATE TABLE `funding_card` (`id` int NOT NULL AUTO_INCREMENT,`funding_card_no` varchar(255) NOT NULL COMMENT '资金卡号',`bank_uid` varchar(255) NOT NULL COMMENT '银行uid',`balance` decimal(10,2) NOT NULL COMMENT '余额（HKD）',`created_at` timestamp NOT NULL COMMENT '创建时间',`updated_at` timestamp NOT NULL COMMENT '修改时间',`deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',PRIMARY KEY (`id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='资金卡表';"
		_, err = g.DB().Exec(ctx, sql)
		if err != nil {
			g.Log().Error(ctx, err)
		}
	}
}

func addTableRechargeList(ctx context.Context) {
	needUpgrade := true
	result, err := g.DB().GetAll(ctx, "show tables")
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	list := result.List()
	for _, m := range list {
		if m["Tables_in_upay"] == "recharge_list" {
			needUpgrade = false
			break
		}
	}
	if needUpgrade {
		sql := "CREATE TABLE `recharge_list` (\n`id` int NOT NULL AUTO_INCREMENT,\n`fund_card_no` varchar(255) NOT NULL DEFAULT '' COMMENT '资金卡号',\n`card_no` varchar(255) NOT NULL DEFAULT '' COMMENT '对方卡号',\n`amount` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '金额（hkd）',\n`name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '对方姓名',\n`apply_list_id` int NOT NULL DEFAULT '0' COMMENT '申请列表id',\n`admin_username` varchar(255) NOT NULL DEFAULT '' COMMENT '操作账号',\n`created_at` timestamp NULL DEFAULT NULL COMMENT '充值时间',\n`updated_at` timestamp NULL DEFAULT NULL,\n`deleted_at` timestamp NULL DEFAULT NULL,\nPRIMARY KEY (`id`)\n) ENGINE=InnoDB AUTO_INCREMENT=44 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='充值记录'"
		_, err = g.DB().Exec(ctx, sql)
		if err != nil {
			g.Log().Error(ctx, err)
		}
	}
}

func addBalanceToApplyList(ctx context.Context) {
	needUpgrade := true
	result, err := g.DB().GetAll(ctx, "desc apply_list")
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	list := result.List()
	for _, m := range list {
		if m["Field"] == "balance" {
			needUpgrade = false
			break
		}
	}
	if needUpgrade {
		sql := "ALTER TABLE `apply_list` \nADD COLUMN `balance` decimal(10, 2) NOT NULL COMMENT '余额的信息';"
		_, err = g.DB().Exec(ctx, sql)
		if err != nil {
			g.Log().Error(ctx, err)
		}
	}
}

func addFileToCardLevel(ctx context.Context) {
	needUpgrade := true
	result, err := g.DB().GetAll(ctx, "desc card_level")
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	list := result.List()
	for _, m := range list {
		if m["Field"] == "file" {
			needUpgrade = false
			break
		}
	}
	if needUpgrade {
		sql := "ALTER TABLE `card_level` \nADD COLUMN `file` longtext NULL COMMENT '文件base64值';"
		_, err = g.DB().Exec(ctx, sql)
		if err != nil {
			g.Log().Error(ctx, err)
		}
	}
}

func addTransactionRecord(ctx context.Context) {
	needUpgrade := true
	result, err := g.DB().GetAll(ctx, "show tables")
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	list := result.List()
	for _, m := range list {
		if m["Tables_in_upay"] == "transaction_record" {
			needUpgrade = false
			break
		}
	}
	if needUpgrade {
		sql := "CREATE TABLE `transaction_record` (\n  `id` varchar(127) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',\n  `xid` varchar(127) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '用户xid',\n  `ref_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',\n  `card_account_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',\n  `card_last_4` varchar(7) NOT NULL DEFAULT '',\n  `card_embossed_name` varchar(127) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',\n  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',\n  `status` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '状态 pending 待处理  posted 已完成 declined 已拒绝  vodi 已取消',\n  `posted_at` timestamp NULL DEFAULT NULL,\n  `intent` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '交易目的/方式，charge收费 refund退款 topup充值 repay偿还 cashback现金返回 interest利息 fee费用 other其他',\n  `amount` decimal(10,2) NOT NULL DEFAULT '0.00',\n  `entry_type` varchar(31) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',\n  `currency` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',\n  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '响应消息（例如：响应被接受）'\n) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='交易记录';"
		_, err = g.DB().Exec(ctx, sql)
		if err != nil {
			g.Log().Error(ctx, err)
		}
	}

}

func addNoteToApplyList(ctx context.Context) {
	needUpgrade := true
	result, err := g.DB().GetAll(ctx, "desc apply_list")
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	list := result.List()
	for _, m := range list {
		if m["Field"] == "note" {
			needUpgrade = false
			break
		}
	}
	if needUpgrade {
		sql := "ALTER TABLE `upay`.`apply_list` \nADD COLUMN `note` varchar(255) NULL COMMENT '备注信息';"
		_, err = g.DB().Exec(ctx, sql)
		if err != nil {
			g.Log().Error(ctx, err)
		}
	}
}

func addKYCStatusToAppUser(ctx context.Context) {
	needUpgrade := true
	result, err := g.DB().GetAll(ctx, "desc app_user")
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	list := result.List()
	for _, m := range list {
		if m["Field"] == "kyc_status" {
			needUpgrade = false
			break
		}
	}
	if needUpgrade {
		sql := "ALTER TABLE `upay`.`app_user` ADD COLUMN `kyc_status` tinyint(1) NOT NULL DEFAULT 0 COMMENT '0-待kyc 1-已kyc' AFTER `status`;"
		_, err = g.DB().Exec(ctx, sql)
		if err != nil {
			g.Log().Error(ctx, err)
			return
		}

		sql2 := "UPDATE `upay`.`app_user` SET `kyc_status` = 1 WHERE `status`>1;"
		_, err = g.DB().Exec(ctx, sql2)
		if err != nil {
			g.Log().Error(ctx, err)
		}
	}
}

func addUserCardStatusToAppUser(ctx context.Context) {
	needUpgrade := true
	result, err := g.DB().GetAll(ctx, "desc app_user")
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	list := result.List()
	for _, m := range list {
		if m["Field"] == "user_card_status" {
			needUpgrade = false
			break
		}
	}
	if needUpgrade {
		sql := "ALTER TABLE `upay`.`app_user` ADD COLUMN `user_card_status` tinyint(1) NOT NULL DEFAULT 0 COMMENT '0-未上传 1-已上传' AFTER `kyc_status`;"
		_, err = g.DB().Exec(ctx, sql)
		if err != nil {
			g.Log().Error(ctx, err)
			return
		}

		sql2 := "UPDATE `upay`.`app_user` SET `user_card_status` = 1 WHERE `status`>2;"
		_, err = g.DB().Exec(ctx, sql2)
		if err != nil {
			g.Log().Error(ctx, err)
		}
	}
}

func addBankCardStatusToAppUser(ctx context.Context) {
	needUpgrade := true
	result, err := g.DB().GetAll(ctx, "desc app_user")
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	list := result.List()
	for _, m := range list {
		if m["Field"] == "bank_card_status" {
			needUpgrade = false
			break
		}
	}
	if needUpgrade {
		sql := "ALTER TABLE `upay`.`app_user` ADD COLUMN `bank_card_status` tinyint(1) NOT NULL COMMENT '0-未配卡 1-已配卡' AFTER `user_card_status`;"
		_, err = g.DB().Exec(ctx, sql)
		if err != nil {
			g.Log().Error(ctx, err)
			return
		}

		sql2 := "UPDATE `upay`.`app_user` SET `bank_card_status` = 1 WHERE `status`>3;"
		_, err = g.DB().Exec(ctx, sql2)
		if err != nil {
			g.Log().Error(ctx, err)
		}
	}
}

// func addApplyStatusToAppUser(ctx context.Context) {
// 	needUpgrade := true
// 	result, err := g.DB().GetAll(ctx, "desc app_user")
// 	if err != nil {
// 		g.Log().Error(ctx, err)
// 		return
// 	}
// 	list := result.List()
// 	for _, m := range list {
// 		if m["Field"] == "apply_status" {
// 			needUpgrade = false
// 			break
// 		}
// 	}
// 	if needUpgrade {
// 		sql := "ALTER TABLE `upay`.`app_user` ADD COLUMN `apply_status` tinyint(1) NOT NULL COMMENT '0-待申请 1-已申请' AFTER `bank_card_status`;"
// 		_, err = g.DB().Exec(ctx, sql)
// 		if err != nil {
// 			g.Log().Error(ctx, err)
// 			return
// 		}

// 		sql2 := "UPDATE `upay`.`app_user` SET `apply_status` = 1 WHERE `status`=5;"
// 		_, err = g.DB().Exec(ctx, sql2)
// 		if err != nil {
// 			g.Log().Error(ctx, err)
// 		}
// 	}
// }
