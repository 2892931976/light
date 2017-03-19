package mapper

import (
	"testing"
	"time"

	"github.com/arstd/light/example/enum"
	"github.com/arstd/light/example/model"
	"github.com/arstd/log"
)

func TestCreateTable(t *testing.T) {
	_, err := db.Exec("drop table if exists models")
	if err != nil {
		log.Error(err)
	}
	_, err = db.Exec(`
		create table models (
			id serial primary key,
			name text not null,
			flag bool not null default false,
			score decimal(3,1) not null default 0.0,

			map jsonb not null default '{}',
			time timestamptz not null default now(),
			xarray text[] not null,
			slice text[] not null,

			status smallint not null default 0,
			state text not null default '',

			pointer jsonb not null default '{}',
			struct_slice jsonb not null default '[]',
			uint32 bigint not null default 0
		)
	`)
	if err != nil {
		log.Error(err)
	}
}

var mapper ModelMapper = &ModelMapperImpl{}
var id int = 1

func TestModelMapperInsert(t *testing.T) {
	m := &model.Model{
		Name:  "name",
		Flag:  true,
		Score: 1.23,

		Map:   map[string]interface{}{"a": 1},
		Time:  time.Now(),
		Array: []int64{1, 2, 3},
		Slice: []string{"Slice Elem 1", "Slice Elem 2"},

		Status:  enum.StatusNormal,
		Pointer: &model.Model{Name: "Pointer"},
		StructSlice: []*model.Model{
			{Name: "StructSlice"},
		},

		Uint32: 32,
	}
	tx, err := BeginTx()
	if err != nil {
		t.Fatalf("insert error: %s", err)
	}
	defer RollbackTx(tx)
	err = mapper.Insert(tx, m)
	if err != nil {
		t.Fatalf("insert error: %s", err)
	}

	CommitTx(tx)
	id = m.Id
	log.Infof("id=%d", m.Id)
}

func TestModelMapperBatchInsert(t *testing.T) {
	m := &model.Model{
		Name:  "name",
		Flag:  true,
		Score: 1.23,

		Map:   map[string]interface{}{"a": 1},
		Time:  time.Now(),
		Array: []int64{1, 2, 3},
		Slice: []string{"Slice Elem 1", "Slice Elem 2"},

		Status:  enum.StatusNormal,
		Pointer: &model.Model{Name: "Pointer"},
		StructSlice: []*model.Model{
			{Name: "StructSlice"},
		},

		Uint32: 32,
	}
	tx, err := BeginTx()
	if err != nil {
		t.Fatalf("insert error: %s", err)
	}
	defer RollbackTx(tx)
	a, err := mapper.BatchInsert(tx, []*model.Model{m, m, m})
	if err != nil {
		t.Fatalf("insert error: %s", err)
	}

	CommitTx(tx)
	log.Infof("affect %d rows", a)
}

func TestModelMapperGet(t *testing.T) {
	tx, err := BeginTx()
	if err != nil {
		t.Fatalf("insert error: %s", err)
	}
	defer RollbackTx(tx)
	m, err := mapper.Get(tx, id)
	if err != nil {
		t.Fatalf("get error: %s", err)
	}

	CommitTx(tx)
	log.JSON(m)
}

func TestModelMapperUpdate(t *testing.T) {
	m := &model.Model{
		Id:    id,
		Name:  "name update",
		Flag:  true,
		Score: 1.23,

		Map:   map[string]interface{}{"a": "1  update"},
		Time:  time.Now().Add(-3 * time.Hour),
		Slice: []string{"Slice Elem 1 update", "Slice Elem 2 update"},

		Status:  enum.StatusNormal,
		Pointer: &model.Model{Name: "Pointer update"},
		StructSlice: []*model.Model{
			{Name: "StructSlice update"},
		},
		Uint32: 32,
	}
	tx, err := BeginTx()
	if err != nil {
		t.Fatalf("insert error: %s", err)
	}
	defer RollbackTx(tx)
	a, err := mapper.Update(tx, m)
	if err != nil {
		t.Fatalf("update error: %s", err)
	}

	CommitTx(tx)
	log.Infof("affected=%d", a)
}

func TestModelMapperCount(t *testing.T) {
	m := &model.Model{
		Name:   "name%", // like 'name%'
		Flag:   true,
		Status: enum.StatusNormal,
		Slice:  []string{"Slice Elem 3", "xSlice Elem 2"},
	}
	tx, err := BeginTx()
	if err != nil {
		t.Fatalf("insert error: %s", err)
	}
	defer RollbackTx(tx)
	count, err := mapper.Count(tx, m, []enum.Status{enum.StatusNormal, enum.StatusDeleted})
	if err != nil {
		t.Fatalf("count(%+v) error: %s", m, err)
	}

	CommitTx(tx)
	log.JSON(count)
}

func TestModelMapperList(t *testing.T) {
	m := &model.Model{
		Name:  "name%", // like 'name%'
		Flag:  true,
		Array: []int64{11, 22, 3},
		// Slice: []string{"SliceElem1", "SliceElem2"},
	}
	ss := []enum.Status{enum.StatusNormal, enum.StatusDeleted}
	tx, err := BeginTx()
	if err != nil {
		t.Fatalf("insert error: %s", err)
	}
	defer RollbackTx(tx)
	ms, err := mapper.List(tx, m, ss, 0, 20)
	if err != nil {
		t.Fatalf("list(%+v) error: %s", m, err)
	}

	CommitTx(tx)
	log.JSON(ms)
}

func TestModelMapperDelete(t *testing.T) {
	tx, err := BeginTx()
	if err != nil {
		t.Fatalf("insert error: %s", err)
	}
	defer RollbackTx(tx)
	a, err := mapper.Delete(tx, id)
	CommitTx(tx)

	if err != nil {
		t.Fatalf("delete error: %s", err)
	}

	log.JSON(a)
}
