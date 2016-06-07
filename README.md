gobatis
================================================================================

根据接口和 SQL 生成数据库 CRUD 实现方法


支持 6 种操作
--------------------------------------------------------------------------------

* add: insert into table(name) values('name') returning id
* modify: update table set name='name' where id=1
* remove: delete from table where id=1
* get: select id,name from table where id=1
* list: select id,name from table where id < 1000 offset 10 limit 5
* count/sum: select count(id) from table where id < 1000


Usage
--------------------------------------------------------------------------------

1. 编写接口

    ```go
    package persist

    //go:generate gobatis

    // DemoPersister 示例接口
    type DemoPersister interface {

    	// select id, name, third_field, status, content
    	// from demos
    	// where name=${d.Name}
    	//   [?{d.ThirdField != false} and third_field=${d.ThirdField} ]
    	//   [?{d.Content != nil} and content=${d.Content} ]
    	//   [?{len(statuses) != 0} and status in (${statuses}) ]
    	List(tx *sql.Tx, d *domain.Demo, statuses []enums.Status, page, size int) ([]*domain.Demo, error)
    }
    ```

2. 生成代码

    `go generate ./...`

3. Over。

更多示例见： [examples/persist/demo.go](examples/persist/demo.go)
生成的文件： [examples/persist/demoimpl.go](examples/persist/demoimpl.go)


gobatis 更多参数
--------------------------------------------------------------------------------


```sh
gobatis -h
Usage of gobatis:
  -db string
    	db variable to Query/QueryRow/Exec (default "db")
  -path string
    	db variable path
```
