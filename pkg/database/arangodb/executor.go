package arangodb

import (
	"context"
	"github.com/arangodb/go-driver"
)

//ArangoDbQueryExecutor is the primary type for executing queries
type QueryExecutorArangoDB struct {
	DB driver.Database
	db driver.Database
	QueryExecutor
}

//Execute runs the specified query
func (q *QueryExecutorArangoDB) Execute(queryText string, bindVars map[string]interface{}) ([]map[string]interface{}, error) {
	if q.db == nil {
		db, err := InitializeArangoDb()
		if err != nil {
			return nil, err
		}
		q.db = db
		q.DB = db
	}

	ctx := driver.WithQueryCount(context.Background())
	cursor, err := q.db.Query(ctx, queryText, bindVars)
	if err != nil {
		return nil, err
	}
	defer cursor.Close()

	count := cursor.Count()
	data := make([]map[string]interface{}, count)

	idx := 0
	for {
		var doc map[string]interface{}
		_, err := cursor.ReadDocument(ctx, &doc)
		if driver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			// handle other errors
		}

		data[idx] = doc
		idx++
	}

	return data, nil
}

//Execute runs the specified query and gets the full count from cursor
func (q *QueryExecutorArangoDB) ExecuteWithCount(queryText string, bindVars map[string]interface{}) ([]map[string]interface{}, int64, error) {
	if q.db == nil {
		db, err := InitializeArangoDb()
		if err != nil {
			return nil, 0, err
		}
		q.db = db
		q.DB = db
	}

	ctx := driver.WithQueryFullCount(context.Background())
	cursor, err := q.db.Query(ctx, queryText, bindVars)
	if err != nil {
		return nil, 0, err
	}

	defer cursor.Close()

	count := cursor.Statistics().FullCount()
	var data []map[string]interface{}

	idx := 0
	for {
		var doc map[string]interface{}
		_, err := cursor.ReadDocument(ctx, &doc)
		if driver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			// handle other errors
		}

		data = append(data, doc)
		idx++
	}

	return data, count, nil
}

//NewQueryExecutor returns a new query executor
func NewQueryExecutor() QueryExecutorArangoDB {
	qe := QueryExecutorArangoDB{}
	return qe
}
