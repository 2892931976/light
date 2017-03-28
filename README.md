light
================================================================================

Generate go database code by SQL statement, sprit from MyBatis/ibatis, but no Reflect.

7 kind of methods
--------------------------------------------------------------------------------

* add: insert into table(name) values('name') returning id
* modify: update table set name='name' where id=1
* remove: delete from table where id=1
* get: select id, name from table where id=1
* count: select count(id) from table where id < 1000
* list: select id, name from table where id < 1000 order by id offset 10 limit 5
* page: select count(id) | id, name from table where id < 1000 [ order by id offset 10 limit 5 ]


Usage
--------------------------------------------------------------------------------

1. Code interface, Comment methods with SQL statement

```go
package persist

//go:generate light

// ModelMapper example model
type ModelMapper interface {

	// select id, name, third_field, status, content
	// from demos
	// where name=${d.Name}
	//   [?{ d.ThirdField != false } and third_field=${d.ThirdField} ]
	//   [?{ d.Content != nil } and content=${d.Content} ]
	//   [?{ len(d.Tags) != 0 } and tag in (${d.Tags}) ]
	// order by id
	// offset ${(d.Page-1)*d.Size} limit ${d.Size}
	List(d *domain.Demo, tx *sql.Tx) ([]*domain.Demo, error)
}
```

2. Execute go generate tool

    `go generate ./...`


more example: [example/mapper/model.go](example/mapper/model.go)

generated impl code: [example/mapper/modelimpl.go](example/mapper/modelimpl.go)


More
--------------------------------------------------------------------------------

```
# light -h
usage: light [flags] [file.go]
	//go:generate light [flags] [file.go]
examples:
	light -force -dbvar=db.DB -dbpath=github.com/arstd/light/example/mapper
	light -force -dbvar=db2.DB -dbpath=github.com/arstd/light/example/mapper
  -dbpath string
    	path of db to open transaction and execute SQL statements
  -dbvar string
    	variable of db to open transaction and execute SQL statements (default "db")
  -force
    	force to regenerate, even sourceimpl.go file newer than source.go file
  -v	variable of db to open transaction and execute SQL statements
```
