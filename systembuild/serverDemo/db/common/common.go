package common

import (
	"database/sql"
	"fmt"
	"reflect"
	"serverDemo/common/dstruct"
	"serverDemo/common/log"
	"serverDemo/common/util"
	"serverDemo/db"
	"sync"

	"xorm.io/xorm"
)

// 加载静态数据到缓存
func InitMap() {
	defer util.Profiling("db.InitMap")()

	loadTables := []interface{}{}
	w := sync.WaitGroup{}
	for i, v := range loadTables {
		w.Add(1)
		go func(tableStruct interface{}, i int) {
			defer w.Done()
			defer func() {
				if err := recover(); err != nil {
					log.Error("Recover in IninRedisNew:", i)
				}
			}()
			m := reflect.ValueOf(tableStruct).MethodByName("LoadDatas")
			if m.IsValid() && m.Kind() == reflect.Func {
				m.Call(nil)
			}
		}(v, i)
		w.Wait() // 并行运行
	}
	// w.Wait() // 并发运行
}

// 通过事务提交sql
// isCheck，更新有效记录是否必须不为0
func ExeSqlByTransaction(engine xorm.Engine, isCheck bool, sqlSlice ...*db.SqlData) error {
	var sqlResult sql.Result
	session := engine.NewSession()
	defer session.Close()
	err := session.Begin()
	for i := range sqlSlice {
		if sqlSlice[i].Sql == "" {
			continue
		}
		execSql := []interface{}{sqlSlice[i].Sql}
		execSql = append(execSql, sqlSlice[i].Param...)
		log.Info("execute sql:", execSql)
		if sqlResult, err = session.Exec(execSql...); err != nil {
			log.Error(err)
			break
		} else {
			num, _ := sqlResult.RowsAffected()
			log.Info("num:", num)
			if isCheck && num < 1 {
				err = fmt.Errorf("rowAffect num:%v", num)
				log.Warn(err)
				break
			}
		}
	}
	if err != nil {
		err = fmt.Errorf("[ExeSqlByTransaction.fail]:%v", err)
		log.Error(err)
		session.Rollback()
		return err
	}
	err = session.Commit()
	if err != nil {
		err = fmt.Errorf("[ExeSqlByTransaction.fail]:%v", err)
		log.Error(err)
		return err
	}
	return nil
}

// 通过反射调用结构对应的TableName函数,达到返回表名的目的
func GetTableName(tableStruct interface{}) string {
	m := reflect.ValueOf(tableStruct).MethodByName("TableName")
	if m.IsValid() && m.Kind() == reflect.Func {
		re := m.Call(nil)
		for _, v := range re {
			if v.IsValid() {
				return v.String()
			}
		}
	}
	return "unknown"
}

// 根据主键 id通用更新表数据
// tableStruct:为表映射后的结构指针
// updateMap:为更新表数据的 map 结构，期中必须包含主键 id
func UpdateTableById(engine *xorm.Engine, tableStruct interface{}, updateMap map[string]interface{}) error {
	if reflect.TypeOf(tableStruct).Kind() != reflect.Ptr {
		err := fmt.Errorf("tableStruct must ptr")
		log.Error(err)
		return err
	}
	id, ok := updateMap["id"]
	if !ok {
		err := fmt.Errorf("updateMap not found id")
		log.Warn(err)
		return err
	}
	num, err := engine.Table(tableStruct).ID(id).Update(updateMap)
	if err != nil {
		log.Warn(err)
	}
	log.Infof("[UpdateTableById.%v] id:%v effectNum:%v updateMap:%v\n", GetTableName(tableStruct), id, num, updateMap)
	return nil
}

func ViewDayData(rows []*dstruct.AddUserNum) []*dstruct.DayAddUserNum {
	ret := []*dstruct.DayAddUserNum{}
	item := &dstruct.DayAddUserNum{}
	for _, row := range rows {
		if len(item.Date) > 0 && item.Date != row.Date {
			ret = append(ret, item)
			item = &dstruct.DayAddUserNum{}
		}
		item.Date = row.Date
		item.All += row.Num
		switch row.UserType {
		case "99u":
			item.U99 = row.Num
		case "wx":
			item.Wx = row.Num
		}
	}
	return ret
}
