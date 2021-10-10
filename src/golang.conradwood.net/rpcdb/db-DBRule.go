package db

/*
 This file was created by mkdb-client.
 The intention is not to modify thils file, but you may extend the struct DBDBRule
 in a seperate file (so that you can regenerate this one from time to time)
*/

/*
 PRIMARY KEY: ID
*/

/*
 postgres:
 create sequence dbrule_seq;

Main Table:

 CREATE TABLE dbrule (id integer primary key default nextval('dbrule_seq'),serviceid bigint not null,resultonmatch boolean not null);

Archive Table: (structs can be moved from main to archive using Archive() function)

 CREATE TABLE dbrule_archive (id integer unique not null,serviceid bigint not null,resultonmatch boolean not null);
*/

import (
	gosql "database/sql"
	"fmt"
	savepb "golang.conradwood.net/apis/rpcaclapi"
	"golang.conradwood.net/go-easyops/sql"
	"golang.org/x/net/context"
)

type DBDBRule struct {
	DB *sql.DB
}

func NewDBDBRule(db *sql.DB) *DBDBRule {
	foo := DBDBRule{DB: db}
	return &foo
}

// archive. It is NOT transactionally save.
func (a *DBDBRule) Archive(ctx context.Context, id uint64) error {

	// load it
	p, err := a.ByID(ctx, id)
	if err != nil {
		return err
	}

	// now save it to archive:
	_, e := a.DB.ExecContext(ctx, "insert_DBDBRule", "insert into dbrule_archive (id,serviceid, resultonmatch) values ($1,$2, $3) ", p.ID, p.ServiceID, p.ResultOnMatch)
	if e != nil {
		return e
	}

	// now delete it.
	a.DeleteByID(ctx, id)
	return nil
}

// Save (and use database default ID generation)
func (a *DBDBRule) Save(ctx context.Context, p *savepb.DBRule) (uint64, error) {
	rows, e := a.DB.QueryContext(ctx, "DBDBRule_Save", "insert into dbrule (serviceid, resultonmatch) values ($1, $2) returning id", p.ServiceID, p.ResultOnMatch)
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
func (a *DBDBRule) SaveWithID(ctx context.Context, p *savepb.DBRule) error {
	_, e := a.DB.ExecContext(ctx, "insert_DBDBRule", "insert into dbrule (id,serviceid, resultonmatch) values ($1,$2, $3) ", p.ID, p.ServiceID, p.ResultOnMatch)
	return e
}

func (a *DBDBRule) Update(ctx context.Context, p *savepb.DBRule) error {
	_, e := a.DB.ExecContext(ctx, "DBDBRule_Update", "update dbrule set serviceid=$1, resultonmatch=$2 where id = $3", p.ServiceID, p.ResultOnMatch, p.ID)

	return e
}

// delete by id field
func (a *DBDBRule) DeleteByID(ctx context.Context, p uint64) error {
	_, e := a.DB.ExecContext(ctx, "deleteDBDBRule_ByID", "delete from dbrule where id = $1", p)
	return e
}

// get it by primary id
func (a *DBDBRule) ByID(ctx context.Context, p uint64) (*savepb.DBRule, error) {
	rows, e := a.DB.QueryContext(ctx, "DBDBRule_ByID", "select id,serviceid, resultonmatch from dbrule where id = $1", p)
	if e != nil {
		return nil, fmt.Errorf("ByID: error querying (%s)", e)
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, fmt.Errorf("ByID: error scanning (%s)", e)
	}
	if len(l) == 0 {
		return nil, fmt.Errorf("No DBRule with id %d", p)
	}
	if len(l) != 1 {
		return nil, fmt.Errorf("Multiple (%d) DBRule with id %d", len(l), p)
	}
	return l[0], nil
}

// get all rows
func (a *DBDBRule) All(ctx context.Context) ([]*savepb.DBRule, error) {
	rows, e := a.DB.QueryContext(ctx, "DBDBRule_all", "select id,serviceid, resultonmatch from dbrule order by id")
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

// get all "DBDBRule" rows with matching ServiceID
func (a *DBDBRule) ByServiceID(ctx context.Context, p uint64) ([]*savepb.DBRule, error) {
	rows, e := a.DB.QueryContext(ctx, "DBDBRule_ByServiceID", "select id,serviceid, resultonmatch from dbrule where serviceid = $1", p)
	if e != nil {
		return nil, fmt.Errorf("ByServiceID: error querying (%s)", e)
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, fmt.Errorf("ByServiceID: error scanning (%s)", e)
	}
	return l, nil
}

// get all "DBDBRule" rows with matching ResultOnMatch
func (a *DBDBRule) ByResultOnMatch(ctx context.Context, p bool) ([]*savepb.DBRule, error) {
	rows, e := a.DB.QueryContext(ctx, "DBDBRule_ByResultOnMatch", "select id,serviceid, resultonmatch from dbrule where resultonmatch = $1", p)
	if e != nil {
		return nil, fmt.Errorf("ByResultOnMatch: error querying (%s)", e)
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, fmt.Errorf("ByResultOnMatch: error scanning (%s)", e)
	}
	return l, nil
}

/**********************************************************************
* Helper to convert from an SQL Row to struct
**********************************************************************/
func (a *DBDBRule) Tablename() string {
	return "dbrule"
}

func (a *DBDBRule) SelectCols() string {
	return "id,serviceid, resultonmatch"
}

func (a *DBDBRule) FromRows(ctx context.Context, rows *gosql.Rows) ([]*savepb.DBRule, error) {
	var res []*savepb.DBRule
	for rows.Next() {
		foo := savepb.DBRule{}
		err := rows.Scan(&foo.ID, &foo.ServiceID, &foo.ResultOnMatch)
		if err != nil {
			return nil, err
		}
		res = append(res, &foo)
	}
	return res, nil
}
