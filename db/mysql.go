package main

import (
    "crypto/md5"  
    "database/sql"  
    "encoding/hex"  
    "fmt"  
    // 新手都会被这个_所迷惑，其实这个就是Go设计的巧妙之处，我们在变量赋值的时候经常看到这个符号，它是用来忽略变量赋值的占位符，那么包引入用到这个符号也是相似的作用，这儿使用_的意思是引入后面的包名而不直接使用这个包中定义的函数，变量等资源。  
    _ "github.com/go-sql-driver/mysql"  
    "time"  
)

func checkErr(err error) {  
    if err != nil {  
        panic(err)  
    }  
}

func main() {
    //dsn 数据源名称  
    //user:password@tcp([de:ad:be:ef::ca:fe]:80)/dbname?charset=utf8mb4,utf8  
  
    db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/laravel5?charset=utf8")
    checkErr(err)

    /**
    * CREATE TABLE `test_go` (
          `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
        name varchar(30) not null comment 'username',
        password char(32) not null comment 'password',
        other  varchar(30) not null comment 'username',
        `ctime` timestamp default current_timestamp comment '创建时间',
        `utime` timestamp default '00-00-00 00:00:00' comment '修改时间',
        PRIMARY KEY (`id`)
       ) ENGINE=InnoDB DEFAULT CHARSET=utf8 comment 'test';
    *
    */

    //插入数据
    stmt, err := db.Prepare("insert test_go set name = ?,password = ?,other = ?,ctime = ?,utime = ?")
    checkErr(err)

    //标准格式化时间
    ctime := time.Now().Format("2006-01-02 15:04:05")
    pwd := []byte("test")
    //密码md5加密
    p := md5.Sum(pwd)

    //执行插入
    res, err := stmt.Exec("test1", hex.EncodeToString(p[:]), "test", ctime, ctime)
    checkErr(err)

    //获取最新插入的insertid
    id, err := res.LastInsertId()
    checkErr(err)

    fmt.Println(id)

    //查询数据
    rows, err := db.Query("SELECT * FROM test_go")
    checkErr(err)

    for rows.Next() {
        var id int
        var name string
        var password string
        var other string
        var ctime string
        var utime string
        err = rows.Scan(&id, &name, &password, &other, &ctime, &utime)
        checkErr(err)
        //查询打印结果集
        fmt.Println(id)
        fmt.Println(name)
        fmt.Println(password)
        fmt.Println(other)
        fmt.Println(ctime)
        fmt.Println(utime)
    }

}