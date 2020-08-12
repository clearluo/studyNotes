package admin

import (
	"fmt"
	"serverDemo/common/consts"
	"serverDemo/common/log"
	"serverDemo/common/util"
	"serverDemo/db"
	"time"
)

type TAdmin struct {
	Id       int       `xorm:"not null pk autoincr INT(11)" json:"id"`
	Username string    `xorm:"default '' comment('用户名') unique VARCHAR(32)" json:"username"`
	Password string    `xorm:"default '' comment('密码') VARCHAR(32)" json:"-"`
	Role     string    `xorm:"default '' comment('用户角色【superadmin,admin,app,,guest】') VARCHAR(16)" json:"role"`
	Status   int       `xorm:"default 1 comment('1-生效；2-无效') TINYINT(1)" json:"status"`
	UpdateAt time.Time `xorm:"name updateAt default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP" json:"updateAt"`
	CreateAt time.Time `xorm:"name createAt default CURRENT_TIMESTAMP TIMESTAMP" json:"createAt"`
	Token    string    `xorm:"default '' VARCHAR(128)" json:"-"`
	FatherId int       `xorm:"name fatherId default 0 comment('创建者的id') INT(11)" json:"fatherId"`
}

func (t *TAdmin) TableName() string {
	return "t_admin"
}

// 查询单条数据
func 例子SelectAdminOne() (*TAdmin, error) {
	row := &TAdmin{}
	ok, err := db.GetAppDb().SQL("SELECT * FROM t_admin WHERE id = ?", 6).Get(row)
	log.Warn("ok,err:", ok, err)
	log.Info("row:", util.AssertMarshal(row))

	return row, nil
}

// 例子查询多条数据
func SelectAdminList() ([]*TAdmin, error) {
	rows := []*TAdmin{}
	db.GetAppDb().SQL("SELECT * FROM t_admin").Find(&rows)
	log.Info("rows:", util.AssertMarshal(rows))

	return rows, nil
}

// 例子执行sql
func UpdateAdmin() error {
	sql := `UPDATE t_admin SET password='123456' WHERE id = ?`
	res, err := db.GetAppDb().Exec(sql, 6)
	if err != nil {
		log.Warn(err)
		return err
	}
	num, _ := res.RowsAffected()
	log.Warn("更新:", num)
	return nil
}

func GetAdminById(id int) (*TAdmin, error) {
	row := &TAdmin{}
	sql := "SELECT * FROM t_admin WHERE id = ? "
	ok, err := db.GetAppDb().SQL(sql, id).Get(row)
	if err != nil {
		log.Warn(ok, err)
		return nil, err
	}
	if !ok {
		err := fmt.Errorf("admin not found:%v", id)
		log.Warn(err)
		return nil, err
	}
	return row, nil
}

func SelectAdminByName(name string) (*TAdmin, error) {
	row := &TAdmin{}
	sql := "SELECT * FROM t_admin WHERE username = ? AND status = ?"
	ok, err := db.GetAppDb().SQL(sql, name, consts.USER_STATUS_ON).Get(row)
	if err != nil {
		log.Warn(ok, err)
		return nil, err
	}
	if !ok {
		err := fmt.Errorf("admin not found:%v", name)
		log.Warn(err)
		return nil, err
	}
	return row, nil
}

func AddAdmin(name string, pwd string, role string) (err error) {
	sql := `INSERT INTO t_admin(username,password,role)VALUES(?,?,?)`
	_, err = db.GetAppDb().Exec(sql, name, pwd, role)
	return
}
