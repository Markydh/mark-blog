package dao

import (
	"mark-blog-backend/entity"
	"database/sql"
	"fmt"
	"mark-blog-backend/config"	
	_ "github.com/go-sql-driver/mysql"
)

//连接池对象
var DB *sql.DB


//初始化连接池，连接Mysql
func Init(){
	DB,_ =sql.Open("mysql",config.Mysqldb)
	DB.SetConnMaxLifetime(100)
	DB.SetConnMaxIdleTime(10)
	//判断连接是否成功
	if err :=DB.Ping(); err!=nil{
		fmt.Println("open database fail")
	}
	fmt.Println("open database success")
}



//查询语句
func Query(){
	sqlStr := "select * from users"
	rows, err := DB.Query(sqlStr)
	if err != nil {
		fmt.Println()
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	// 非常重要：关闭rows释放持有的数据库链接
	defer rows.Close()

	for rows.Next() {
		user := &entity.User{}
		err := rows.Scan(&user.User_id, &user.UserName,&user.Password,&user.Email,&user.Password,&user.CreatedTime,&user.UpdatedTime)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("id:%d name:%s level:%d\n", user.User_id,user.UserName,user.Level)
	}
}


