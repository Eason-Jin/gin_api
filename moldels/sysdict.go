package moldels

import db "gin_api/database"

type SysDict struct {
	id    int    `json:"id" form:"id"`
	name  string `json:"name" form:"name"`
	value string `json:"value" form:"value"`
}

func (p *SysDict) GetSysDict()(sysDicts []SysDict, err error)  {
	sysDicts = make([]SysDict, 0)
	rows, err := db.SqlDB.Query("SELECT id, name, value FROM sys_dict")
	defer rows.Close() //查询关闭连接
	if err!=nil{
		return
	}
	//拿到结果
	for rows.Next(){
		var sysDict SysDict
		rows.Scan(&sysDict.id,&sysDict.name,&sysDict.value)
		sysDicts=append(sysDicts,sysDict)
		if err = rows.Err() ;err!=nil{
			return
		}
	}
	return sysDicts,nil
}


