package db

/*
 This file was created by mkdb-client.
 The intention is not to modify thils file, but you may extend the struct DBDBRuleGroup
 in a seperate file (so that you can regenerate this one from time to time)
*/

/*
 PRIMARY KEY: ID
*/

/*
 postgres:
 create sequence dbrulegroup_seq;

Main Table:

 CREATE TABLE dbrulegroup (id integer primary key default nextval('dbrulegroup_seq'),ruleid bigint not null,groupid varchar(2000) not null);

Archive Table: (structs can be moved from main to archive using Archive() function)

 CREATE TABLE dbrulegroup_archive (id integer unique not null,ruleid bigint not null,groupid varchar(2000) not null);
*/

import (
	gosql "database/sql"
	"fmt"
	savepb "golang.conradwood.net/apis/rpcaclapi"
	"golang.conradwood.net/go-easyops/sql"
	"golang.org/x/net/context"
)

type DBDBRuleGroup struct {
	DB *sql.DB
}

func NewDBDBRuleGroup(db *sql.DB) *DBDBRuleGroup {
	foo := DBDBRuleGroup{DB: db}
	return &foo
}

// archive. It is NOT transactionally save.
func (a *DBDBRuleGroup) Archive(ctx context.Context, id uint64) error {

	// load it
	p, err := a.ByID(ctx, id)
	if err != nil {
		return err
	}

	// now save it to archive:
	_, e := a.DB.ExecContext(ctx, "insert_DBDBRuleGroup", "insert into dbrulegroup_archive (id,ruleid, groupid) values ($1,$2, $3) ", p.ID, p.RuleID, p.GroupID)
	if e != nil {
		return e
	}

	// now delete it.
	a.DeleteByID(ctx, id)
	return nil
}

// Save (and use database default ID generation)
func (a *DBDBRuleGroup) Save(ctx context.Context, p *savepb.DBRuleGroup) (uint64, error) {
	rows, e := a.DB.QueryContext(ctx, "DBDBRuleGroup_Save", "insert into dbrulegroup (ruleid, groupid) values ($1, $2) returning id", p.RuleID, p.GroupID)
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
func (a *DBDBRuleGroup) SaveWithID(ctx context.Context, p *savepb.DBRuleGroup) error {
	_, e := a.DB.ExecContext(ctx, "insert_DBDBRuleGroup", "insert into dbrulegroup (id,ruleid, groupid) values ($1,$2, $3) ", p.ID, p.RuleID, p.GroupID)
	return e
}

func (a *DBDBRuleGroup) Update(ctx context.Context, p *savepb.DBRuleGroup) error {
	_, e := a.DB.ExecContext(ctx, "DBDBRuleGroup_Update", "update dbrulegroup set ruleid=$1, groupid=$2 where id = $3", p.RuleID, p.GroupID, p.ID)

	return e
}

// delete by id field
func (a *DBDBRuleGroup) DeleteByID(ctx context.Context, p uint64) error {
	_, e := a.DB.ExecContext(ctx, "deleteDBDBRuleGroup_ByID", "delete from dbrulegroup where id = $1", p)
	return e
}

// get it by primary id
func (a *DBDBRuleGroup) ByID(ctx context.Context, p uint64) (*savepb.DBRuleGroup, error) {
	rows, e := a.DB.QueryContext(ctx, "DBDBRuleGroup_ByID", "select id,ruleid, groupid from dbrulegroup where id = $1", p)
	if e != nil {
		return nil, fmt.Errorf("ByID: error querying (%s)", e)
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, fmt.Errorf("ByID: error scanning (%s)", e)
	}
	if len(l) == 0 {
		return nil, fmt.Errorf("No DBRuleGroup with id %d", p)
	}
	if len(l) != 1 {
		return nil, fmt.Errorf("Multiple (%d) DBRuleGroup with id %d", len(l), p)
	}
	return l[0], nil
}

// get all rows
func (a *DBDBRuleGroup) All(ctx context.Context) ([]*savepb.DBRuleGroup, error) {
	rows, e := a.DB.QueryContext(ctx, "DBDBRuleGroup_all", "select id,ruleid, groupid from dbrulegroup order by id")
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

// get all "DBDBRuleGroup" rows with matching RuleID
func (a *DBDBRuleGroup) ByRuleID(ctx context.Context, p uint64) ([]*savepb.DBRuleGroup, error) {
	rows, e := a.DB.QueryContext(ctx, "DBDBRuleGroup_ByRuleID", "select id,ruleid, groupid from dbrulegroup where ruleid = $1", p)
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

// get all "DBDBRuleGroup" rows with matching GroupID
func (a *DBDBRuleGroup) ByGroupID(ctx context.Context, p string) ([]*savepb.DBRuleGroup, error) {
	rows, e := a.DB.QueryContext(ctx, "DBDBRuleGroup_ByGroupID", "select id,ruleid, groupid from dbrulegroup where groupid = $1", p)
	if e != nil {
		return nil, fmt.Errorf("ByGroupID: error querying (%s)", e)
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, fmt.Errorf("ByGroupID: error scanning (%s)", e)
	}
	return l, nil
}

/**********************************************************************
* Helper to convert from an SQL Row to struct
**********************************************************************/
func (a *DBDBRuleGroup) Tablename() string {
	return "dbrulegroup"
}

func (a *DBDBRuleGroup) SelectCols() string {
	return "id,ruleid, groupid"
}

func (a *DBDBRuleGroup) FromRows(ctx context.Context, rows *gosql.Rows) ([]*savepb.DBRuleGroup, error) {
	var res []*savepb.DBRuleGroup
	for rows.Next() {
		foo := savepb.DBRuleGroup{}
		err := rows.Scan(&foo.ID, &foo.RuleID, &foo.GroupID)
		if err != nil {
			return nil, err
		}
		res = append(res, &foo)
	}
	return res, nil
}
