package db

/*
 This file was created by mkdb-client.
 The intention is not to modify thils file, but you may extend the struct DBDBRuleMatch
 in a seperate file (so that you can regenerate this one from time to time)
*/

/*
 PRIMARY KEY: ID
*/

/*
 postgres:
 create sequence dbrulematch_seq;

Main Table:

 CREATE TABLE dbrulematch (id integer primary key default nextval('dbrulematch_seq'),ruleid bigint not null,rulematch integer not null);

Archive Table: (structs can be moved from main to archive using Archive() function)

 CREATE TABLE dbrulematch_archive (id integer unique not null,ruleid bigint not null,rulematch integer not null);
*/

import (
	gosql "database/sql"
	"fmt"
	savepb "golang.conradwood.net/apis/rpcaclapi"
	"golang.conradwood.net/go-easyops/sql"
	"golang.org/x/net/context"
)

type DBDBRuleMatch struct {
	DB *sql.DB
}

func NewDBDBRuleMatch(db *sql.DB) *DBDBRuleMatch {
	foo := DBDBRuleMatch{DB: db}
	return &foo
}

// archive. It is NOT transactionally save.
func (a *DBDBRuleMatch) Archive(ctx context.Context, id uint64) error {

	// load it
	p, err := a.ByID(ctx, id)
	if err != nil {
		return err
	}

	// now save it to archive:
	_, e := a.DB.ExecContext(ctx, "insert_DBDBRuleMatch", "insert into dbrulematch_archive (id,ruleid, rulematch) values ($1,$2, $3) ", p.ID, p.RuleID, p.RuleMatch)
	if e != nil {
		return e
	}

	// now delete it.
	a.DeleteByID(ctx, id)
	return nil
}

// Save (and use database default ID generation)
func (a *DBDBRuleMatch) Save(ctx context.Context, p *savepb.DBRuleMatch) (uint64, error) {
	rows, e := a.DB.QueryContext(ctx, "DBDBRuleMatch_Save", "insert into dbrulematch (ruleid, rulematch) values ($1, $2) returning id", p.RuleID, p.RuleMatch)
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
func (a *DBDBRuleMatch) SaveWithID(ctx context.Context, p *savepb.DBRuleMatch) error {
	_, e := a.DB.ExecContext(ctx, "insert_DBDBRuleMatch", "insert into dbrulematch (id,ruleid, rulematch) values ($1,$2, $3) ", p.ID, p.RuleID, p.RuleMatch)
	return e
}

func (a *DBDBRuleMatch) Update(ctx context.Context, p *savepb.DBRuleMatch) error {
	_, e := a.DB.ExecContext(ctx, "DBDBRuleMatch_Update", "update dbrulematch set ruleid=$1, rulematch=$2 where id = $3", p.RuleID, p.RuleMatch, p.ID)

	return e
}

// delete by id field
func (a *DBDBRuleMatch) DeleteByID(ctx context.Context, p uint64) error {
	_, e := a.DB.ExecContext(ctx, "deleteDBDBRuleMatch_ByID", "delete from dbrulematch where id = $1", p)
	return e
}

// get it by primary id
func (a *DBDBRuleMatch) ByID(ctx context.Context, p uint64) (*savepb.DBRuleMatch, error) {
	rows, e := a.DB.QueryContext(ctx, "DBDBRuleMatch_ByID", "select id,ruleid, rulematch from dbrulematch where id = $1", p)
	if e != nil {
		return nil, fmt.Errorf("ByID: error querying (%s)", e)
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, fmt.Errorf("ByID: error scanning (%s)", e)
	}
	if len(l) == 0 {
		return nil, fmt.Errorf("No DBRuleMatch with id %d", p)
	}
	if len(l) != 1 {
		return nil, fmt.Errorf("Multiple (%d) DBRuleMatch with id %d", len(l), p)
	}
	return l[0], nil
}

// get all rows
func (a *DBDBRuleMatch) All(ctx context.Context) ([]*savepb.DBRuleMatch, error) {
	rows, e := a.DB.QueryContext(ctx, "DBDBRuleMatch_all", "select id,ruleid, rulematch from dbrulematch order by id")
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

// get all "DBDBRuleMatch" rows with matching RuleID
func (a *DBDBRuleMatch) ByRuleID(ctx context.Context, p uint64) ([]*savepb.DBRuleMatch, error) {
	rows, e := a.DB.QueryContext(ctx, "DBDBRuleMatch_ByRuleID", "select id,ruleid, rulematch from dbrulematch where ruleid = $1", p)
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

// get all "DBDBRuleMatch" rows with matching RuleMatch
func (a *DBDBRuleMatch) ByRuleMatch(ctx context.Context, p uint32) ([]*savepb.DBRuleMatch, error) {
	rows, e := a.DB.QueryContext(ctx, "DBDBRuleMatch_ByRuleMatch", "select id,ruleid, rulematch from dbrulematch where rulematch = $1", p)
	if e != nil {
		return nil, fmt.Errorf("ByRuleMatch: error querying (%s)", e)
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, fmt.Errorf("ByRuleMatch: error scanning (%s)", e)
	}
	return l, nil
}

/**********************************************************************
* Helper to convert from an SQL Row to struct
**********************************************************************/
func (a *DBDBRuleMatch) Tablename() string {
	return "dbrulematch"
}

func (a *DBDBRuleMatch) SelectCols() string {
	return "id,ruleid, rulematch"
}

func (a *DBDBRuleMatch) FromRows(ctx context.Context, rows *gosql.Rows) ([]*savepb.DBRuleMatch, error) {
	var res []*savepb.DBRuleMatch
	for rows.Next() {
		foo := savepb.DBRuleMatch{}
		err := rows.Scan(&foo.ID, &foo.RuleID, &foo.RuleMatch)
		if err != nil {
			return nil, err
		}
		res = append(res, &foo)
	}
	return res, nil
}
