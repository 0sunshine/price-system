package db

import (
	"backend/protocol"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"strings"
)

var DB *sql.DB

func InitDB() error {

	//"用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	DB, _ = sql.Open("mysql", "root:xPza9LA3k6gykjzyvJS5D%#b@tcp(120.24.255.225:3306)/db_price_system?charset=utf8")

	DB.SetMaxIdleConns(2)

	if err := DB.Ping(); err != nil {
		fmt.Println("open database fail")
		return err
	}

	fmt.Println("connnect success")

	return nil
}

func appendAndCond(conds []string) string {

	first := true
	s := ""

	for _, v := range conds {
		if first {
			first = false
		} else {
			s += " and "
		}
		s += v
	}

	return s
}

func GetProjectKindList(parentProjectKindId string) ([]protocol.ProjectKindItem, error) {

	list := []protocol.ProjectKindItem{}

	strSql := "select * from table_project_kind"

	conds := []string{}
	condsValue := []any{}

	if len(parentProjectKindId) == 0 {
		parentProjectKindId = "000000000000000000000000000000000000"
	}

	conds = append(conds, "parent_project_kind_id=?")
	condsValue = append(condsValue, parentProjectKindId)

	strConds := appendAndCond(conds)
	if len(strConds) > 0 {
		strSql += " where "
		strSql += strConds
		strSql += ";"
	}

	//fmt.Println(strSql)
	rows, err := DB.Query(strSql, condsValue...)
	if err != nil {
		fmt.Println(err)
		return list, err
	}

	defer rows.Close()

	for rows.Next() {
		var item protocol.ProjectKindItem
		err := rows.Scan(&item.ProjectKindId, &item.ProjectName,
			&item.ParentProjectKindId, &item.Notes)
		if err != nil {
			fmt.Println(err)
			return []protocol.ProjectKindItem{}, err
		}

		subList, err := GetProjectKindList(item.ProjectKindId)
		if err != nil {
			fmt.Println(err)
			return []protocol.ProjectKindItem{}, err
		}
		item.Children = subList

		list = append(list, item)
	}

	return list, nil
}

func getMarkByLevel(parentProjectKindID string, level int) (int, error) {
	if level > 5 {
		err := errors.New("level > 5")
		fmt.Println(err)
		return 0, err
	}

	mark := parentProjectKindID[level*6 : (level+1)*6]
	n, err := strconv.Atoi(mark)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return n, err
}

func getNextProjectKindIDRange(parentProjectKindID string) (string, string, error) {
	level := 0
	n := 0

	var err error = nil

	if parentProjectKindID == "000000000000000000000000000000000000" {
		level = -1
	} else {
		for {
			n, err = getMarkByLevel(parentProjectKindID, level)
			if err != nil {
				fmt.Println(err)
				return "", "", err
			}

			if n == 0 {
				level--
				break
			}

			level++
		}
	}

	fmt.Println(parentProjectKindID, " level: ", level)

	var strMin string
	var strMax string

	if level >= 0 {
		strMin += parentProjectKindID[0 : (level+1)*6]
		strMax += parentProjectKindID[0 : (level+1)*6]

		//strMin += fmt.Sprintf("%06d", n)
		//strMax += fmt.Sprintf("%06d", n)
	}

	strMin += "000001"
	strMax += "999999"

	strMin += strings.Repeat("000000", 5-level-1)
	strMax += strings.Repeat("000000", 5-level-1)

	return strMin, strMax, nil
}

func getNextProjectKindID(parentProjectKindID string) (string, error) {
	min, max, err := getNextProjectKindIDRange(parentProjectKindID)
	if err != nil {
		return "", err
	}

	fmt.Println("min: ", min, "  max: ", max)

	row := DB.QueryRow("select project_kind_id from table_project_kind "+
		"where project_kind_id >= ? and project_kind_id <= ? order by project_kind_id desc limit 1", min, max)

	curr_max_project_kind_id := ""

	err = row.Scan(&curr_max_project_kind_id)
	if err != nil {
		if err == sql.ErrNoRows {
			curr_max_project_kind_id = min
		} else {
			return "", err
		}
	}

	fmt.Println("curr_max_project_kind_id ", curr_max_project_kind_id)

	level := 0
	new_max_project_kind_id := ""
	lastN := 0

	for {
		n, err := getMarkByLevel(curr_max_project_kind_id, level)
		if err != nil {
			fmt.Println(err)
			return "", err
		}

		if n == 0 {
			level--
			break
		}

		lastN = n
		level++
	}

	fmt.Println("level ", level)
	new_max_project_kind_id += parentProjectKindID[0 : level*6]

	lastN++

	if lastN > 999999 {
		return "", errors.New("lastN > 999999")
	}

	new_max_project_kind_id += fmt.Sprintf("%06d", lastN)

	new_max_project_kind_id += strings.Repeat("000000", 5-level)

	return new_max_project_kind_id, nil
}

func getProjectKindID(parentProjectKindID string) (string, error) {
	//"000000 000000 000000 000000 000000 000000"

	if len(parentProjectKindID) == 0 {
		parentProjectKindID = "000000000000000000000000000000000000"
	}

	level := 0

	for level < 6 {
		n, err := getMarkByLevel(parentProjectKindID, level)
		if err != nil {
			fmt.Println(err)
			return "", err
		}

		if n != 0 {
			break
		}

		level++
	}

	if level >= 5 {
		err := errors.New("level >= 5")
		fmt.Println(err)
		return "", err
	}

	return "", nil
}

func AddProjectKind(req protocol.AddProjectKindListReq) error {

	var err error
	if len(req.ParentProjectKindId) == 0 {
		req.ParentProjectKindId = "000000000000000000000000000000000000"
	}

	fmt.Println("ParentProjectKindId: ", req.ParentProjectKindId)

	req.ProjectKindId, err = getNextProjectKindID(req.ParentProjectKindId)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("ProjectKindId:", req.ProjectKindId)

	_, err = DB.Exec(
		"insert into table_project_kind(project_kind_id, project_name, parent_project_kind_id,notes) "+
			"values (?,?,?,?)", req.ProjectKindId, req.ProjectName, req.ParentProjectKindId, req.Notes)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func UpdatedProjectKind(req protocol.UpdateProjectKindReq) error {

	var err error

	_, err = DB.Exec(
		"update table_project_kind set project_name = ?, notes=? where project_kind_id=?",
		req.ProjectName, req.Notes, req.ProjectKindId)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func DeleteProjectKind(req protocol.DeleteProjectKindReq) error {

	var err error

	tx, err := DB.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	for _, id := range req.ProjectKindIds {
		list, err := GetProjectKindList(id)
		if err != nil {
			fmt.Println(err)
			return err
		}

		if len(list) > 0 {
			return errors.New(fmt.Sprintf("project_kind_id: %s ,has child", id))
		}

		_, err = DB.Exec(
			"delete from table_project_kind where project_kind_id=?", id)

		if err != nil {
			fmt.Println(err)
			return err
		}
	}

	return nil
}

func GetMaterialKindList(req protocol.GetMaterialKindReq) ([]protocol.MaterialKindItem, error) {
	list := []protocol.MaterialKindItem{}

	strSql := "select * from table_material_kind"

	conds := []string{}
	condsValue := []any{}

	//conds = append(conds, "parent_project_kind_id=?")
	//condsValue = append(condsValue, parentProjectKindId)

	strConds := appendAndCond(conds)
	if len(strConds) > 0 {
		strSql += " where "
		strSql += strConds
		strSql += ";"
	}

	//fmt.Println(strSql)
	rows, err := DB.Query(strSql, condsValue...)
	if err != nil {
		fmt.Println(err)
		return list, err
	}

	defer rows.Close()

	for rows.Next() {
		var item protocol.MaterialKindItem
		err := rows.Scan(&item.MaterialKindId, &item.ProjectKindId,
			&item.MaterialName, &item.Unit, &item.Notes)
		if err != nil {
			fmt.Println(err)
			return []protocol.MaterialKindItem{}, err
		}

		list = append(list, item)
	}

	return list, nil
}

func AddMaterialKind(req protocol.AddMaterialKindReq) error {

	var err error

	_, err = DB.Exec(
		"insert into table_material_kind(project_kind_id, material_name, unit, notes) "+
			"values (?,?,?,?)", req.ProjectKindId, req.MaterialName, req.Unit, req.Notes)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func UpdatedMaterialKind(req protocol.UpdateMaterialKindReq) error {

	var err error

	_, err = DB.Exec(
		"update table_material_kind set material_name = ?,unit=?,notes=? where material_kind_id=?",
		req.MaterialName, req.Unit, req.Notes, req.MaterialKindId)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func DeleteMaterialKind(req protocol.DeleteMaterialKindReq) error {

	var err error

	tx, err := DB.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	for _, id := range req.MaterialKindIds {

		_, err = DB.Exec(
			"delete from table_material_kind where material_kind_id=?", id)

		if err != nil {
			fmt.Println(err)
			return err
		}
	}

	return nil
}
