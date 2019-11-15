package utils

import (
	"fmt"
	"github.com/qingcc/goblog/databases"
	"io/ioutil"
)

func Create() {
	file_string := "config/data.sql"
	buf, err := ioutil.ReadFile(file_string)
	if err != nil {
		fmt.Println("读取db.sql文件失败")
	}

	_, err = databases.Orm.Exec(string(buf))

	if err != nil {
		fmt.Println("执行失败", err.Error())
		return
	}
	fmt.Println("执行成功")
	return
}

//只清除数据, 不修改序列
func CleanTableData(table string) {
	gsql := "delete from `" + table + "`" //删除数据
	exec(gsql)
	fmt.Println("表", table, "清空")
}

func TruncateTable(table string, id_max string) {
	table_seq := table + "_id_seq" //序列名

	if id_max == "" {
		id_max = "1" //重建序列时的初始值
	}

	gsql := "truncate table " + table //清空表
	exec(gsql)
	gsql = "ALTER TABLE " + table + " ALTER COLUMN id SET DEFAULT null" //解除绑定
	exec(gsql)
	gsql = "DROP SEQUENCE IF EXISTS " + table_seq //删除序列
	exec(gsql)
	gsql = "CREATE SEQUENCE " + table_seq + " START WITH " + id_max //重建序列
	exec(gsql)
	gsql = "ALTER TABLE " + table + " ALTER COLUMN id SET DEFAULT nextval('" + table_seq + "'::regclass)" //绑定序列
	exec(gsql)
	fmt.Println("表", table, "清空并初始化自增序列")
}

func exec(sql string) {
	_, err := databases.Orm.Exec(sql)
	if err != nil {
		fmt.Println("sql 执行失败::", sql)
	}
}

//endregion
