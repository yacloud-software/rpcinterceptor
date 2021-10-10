package db

/*
 This file was created by mkdb-client.
 The intention is not to modify thils file, but you may extend the struct DBDBRuleService
 in a seperate file (so that you can regenerate this one from time to time)
*/

/*
 PRIMARY KEY: ID
*/

/*
 postgres:
 create sequence dbruleservice_seq;

Main Table:

 CREATE TABLE dbruleservice (id integer primary key default nextval('dbruleservice_seq'),ruleid bigint not null,serviceuserid varchar(2000) not null);

Archive Table: (structs can be moved from main to archive using Archive() function)

 CREATE TABLE dbruleservice_archive (id integer unique not null,ruleid bigint not null,serviceuserid varchar(2000) not null);
*/

import (
	gosql "database/sql"
	"fmt"
	savepb "golang.conradwood.net/apis/rpcaclapi"
	"golang.conradwood.net/go-easyops/sql"
	"golang.org/x/net/context"
)

type DBDBRuleService struct {
	DB *sql.DB
}

func NewDBDBRuleService(db *sql.DB) *DBDBRuleService {
	foo := DBDBRuleService{DB: db}
	return &foo
}

// archive. It is NOT transactionally save.
func (a *DBDBRuleService) Archive(ctx context.Context, id uint64) error {

	// load it
	p, err := a.ByID(ctx, id)
	if err != nil {
		return err
	}

	// now save it to archive:
	_, e := a.DB.ExecContext(ctx, "insert_DBDBRuleService", "insert into dbruleservice_archive (id,ruleid, serviceuserid) values ($1,$2, $3) ", p.ID, p.RuleID, p.ServiceUserID)
	if e != nil {
		return e
	}

	// now delete it.
	a.DeleteByID(ctx, id)
	return nil
}

// Save (and use database default ID generation)
func (a *DBDBRuleService) Save(ctx context.Context, p *savepb.DBRuleService) (uint64, error) {
	rows, e := a.DB.QueryContext(ctx, "DBDBRuleService_Save", "insert into dbruleservice (ruleid, serviceuserid) values ($1, $2) returning id", p.RuleID, p.ServiceUserID)
	if e != nil {
		return 0, e
	}
	defer rows.Close()
	if !rows.Next() {
		return 0, fmt.Errorf("No rows after insert")
	}
	var id uint64
	e = rows.Scan(&id)
	if e != nil {
		return 0, fmt.Errorf("failed to scan id after insert: %s", e)
	}
	p.ID = id
	return id, nil
}

// Save using the ID specified
func (a *DBDBRuleService) SaveWithID(ctx context.Context, p *savepb.DBRuleService) error {
	_, e := a.DB.ExecContext(ctx, "insert_DBDBRuleService", "insert into dbruleservice (id,ruleid, serviceuserid) values ($1,$2, $3) ", p.ID, p.RuleID, p.ServiceUserID)
	return e
}

func (a *DBDBRuleService) Update(ctx context.Context, p *savepb.DBRuleService) error {
	_, e := a.DB.ExecContext(ctx, "DBDBRuleService_Update", "update dbruleservice set ruleid=$1, serviceuserid=$2 where id = $3", p.RuleID, p.ServiceUserID, p.ID)

	return e
}

// delete by id field
func (a *DBDBRuleService) DeleteByID(ctx context.Context, p uint64) error {
	_, e := a.DB.ExecContext(ctx, "deleteDBDBRuleService_ByID", "delete from dbruleservice where id = $1", p)
	return e
}

// get it by primary id
func (a *DBDBRuleService) ByID(ctx context.Context, p uint64) (*savepb.DBRuleService, error) {
	rows, e := a.DB.QueryContext(ctx, "DBDBRuleService_ByID", "select id,ruleid, serviceuserid from dbruleservice where id = $1", p)
	if e != nil {
		return nil, fmt.Errorf("ByID: error querying (%s)", e)
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, fmt.Errorf("ByID: error scanning (%s)", e)
	}
	if len(l) == 0 {
		return nil, fmt.Errorf("No DBRuleService with id %d", p)
	}
	if len(l) != 1 {
		return nil, fmt.Errorf("Multiple (%d) DBRuleService with id %d", len(l), p)
	}
	return l[0], nil
}

// get all rows
func (a *DBDBRuleService) All(ctx context.Context) ([]*savepb.DBRuleService, error) {
	rows, e := a.DB.QueryContext(ctx, "DBDBRuleService_all", "select id,ruleid, serviceuserid from dbruleservice order by id")
	if e != nil {
		return nil, fmt.Errorf("All: error querying (%s)", e)
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, fmt.Errorf("All: error scanning (%s)", e)
	}
	return l, nil
}

/**********************************************************************
* GetBy[FIELD] functions
**********************************************************************/

// get all "DBDBRuleService" rows with matching RuleID
func (a *DBDBRuleService) ByRuleID(ctx context.Context, p uint64) ([]*savepb.DBRuleService, error) {
	rows, e := a.DB.QueryContext(ctx, "DBDBRuleService_ByRuleID", "select id,ruleid, serviceuserid from dbruleservice where ruleid = $1", p)
	if e != nil {
		return nil, fmt.Errorf("ByRuleID: error querying (%s)", e)
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, fmt.Errorf("ByRuleID: error scanning (%s)", e)
	}
	return l, nil
}

// get all "DBDBRuleService" rows with matching ServiceUserID
func (a *DBDBRuleService) ByServiceUserID(ctx context.Context, p string) ([]*savepb.DBRuleService, error) {
	rows, e := a.DB.QueryContext(ctx, "DBDBRuleService_ByServiceUserID", "select id,ruleid, serviceuserid from dbruleservice where serviceuserid = $1", p)
	if e != nil {
		return nil, fmt.Errorf("ByServiceUserID: error querying (%s)", e)
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, fmt.Errorf("ByServiceUserID: error scanning (%s)", e)
	}
	return l, nil
}

/**********************************************************************
* Helper to convert from an SQL Row to struct
**********************************************************************/
func (a *DBDBRuleService) Tablename() string {
	return "dbruleservice"
}

func (a *DBDBRuleService) SelectCols() string {
	return "id,ruleid, serviceuserid"
}

func (a *DBDBRuleService) FromRows(ctx context.Context, rows *gosql.Rows) ([]*savepb.DBRuleService, error) {
	var res []*savepb.DBRuleService
	for rows.Next() {
		foo := savepb.DBRuleService{}
		err := rows.Scan(&foo.ID, &foo.RuleID, &foo.ServiceUserID)
		if err != nil {
			return nil, err
		}
		res = append(res, &foo)
	}
	return res, nil
}
