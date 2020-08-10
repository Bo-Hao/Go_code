package main

import (
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    fmt.Println("開啟資料")
    db, err := sql.Open("sqlite3", "./foo.db")
    checkErr(err)

    fmt.Println("生成資料表")
    sql_table := `
CREATE TABLE IF NOT EXISTS "userinfo" (
   "uid" INTEGER PRIMARY KEY AUTOINCREMENT,
   "username" VARCHAR(64) NULL,
   "departname" VARCHAR(64) NULL,
   "created" TIMESTAMP default (datetime('now', 'localtime'))  
);
CREATE TABLE IF NOT EXISTS "userdeatail" (
   "uid" INT(10) NULL,
   "intro" TEXT NULL,
   "profile" TEXT NULL,
   PRIMARY KEY (uid)
);
   `
    db.Exec(sql_table)

    //插入資料
    fmt.Print("插入資料, ID=")
    stmt, err := db.Prepare("INSERT INTO userinfo(username, departname)  values(?, ?)")
    checkErr(err)
    res, err := stmt.Exec("astaxie", "研發部門")
    checkErr(err)
    id, err := res.LastInsertId()
    checkErr(err)
    fmt.Println(id)

    //更新資料
    fmt.Print("更新資料 ")
    stmt, err = db.Prepare("update userinfo set username=? where uid=?")
    checkErr(err)
    res, err = stmt.Exec("astaxieupdate", id)
    checkErr(err)
    affect, err := res.RowsAffected()
    checkErr(err)
    fmt.Println(affect)

    //查詢資料
    fmt.Println("查詢資料")
    rows, err := db.Query("SELECT * FROM userinfo")
    checkErr(err)
    for rows.Next() {
        var uid int
        var username string
        var department string
        var created string
        err = rows.Scan(&uid, &username, &department, &created)
        checkErr(err)
        fmt.Println(uid, username, department, created)
    }

    /*
        //刪除資料
        fmt.Println("刪除資料")
        stmt, err = db.Prepare("delete from userinfo where uid=?")
        checkErr(err)
        res, err = stmt.Exec(id)
        checkErr(err)
        affect, err = res.RowsAffected()
        checkErr(err)
        fmt.Println(affect)
    */
    db.Close()
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}