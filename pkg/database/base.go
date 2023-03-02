package database

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"guinevere/pkg/database/arangodb"
	"guinevere/pkg/utils"
)

var _ context.Context

func Create(colName string, doc interface{}) ([]map[string]interface{}, error) {

	h := arangodb.NewQueryExecutor()
	query := fmt.Sprintf(`INSERT @doc IN %s RETURN UNSET(NEW, "_id", "_key", "_rev")`, colName)
	meta, err := h.Execute(query, map[string]interface{}{
		"doc": doc,
	})

	if err != nil {
		logrus.Error("error executing query: ", err)
		return nil, err
	}
	return meta, nil
}

func Update(colName string, key string, doc interface{}) ([]map[string]interface{}, error) {
	h := arangodb.NewQueryExecutor()
	query := fmt.Sprintf(`UPDATE @id WITH @doc IN %s RETURN UNSET(NEW, "_id", "_key", "_rev")`, colName)
	meta, err := h.Execute(query, map[string]interface{}{
		"id":  key,
		"doc": doc,
	})
	if err != nil {
		logrus.Error("error executing query: ", err)
		return nil, err
	}

	return meta, nil
}

func Get(colName string, key string) (map[string]interface{}, error) {
	h := arangodb.NewQueryExecutor()
	query := fmt.Sprintf(`RETURN UNSET(DOCUMENT(%s, @id), '_id', '_rev', '_key', 'password')`, colName)
	meta, err := h.Execute(query, map[string]interface{}{
		"id": key,
	})
	if err != nil {
		logrus.Error("error executing query: ", err)
		return nil, err
	}
	if len(meta) < 1 || meta[0] == nil {
		logrus.Info("invalid id")
		return nil, utils.NewErrJson("invalid id")
	}
	logrus.Info("meta: ", meta[0])
	return meta[0], err
}

func Remove(key string, colName string) (map[string]interface{}, error) {
	h := arangodb.NewQueryExecutor()
	query := fmt.Sprintf(`REMOVE @id IN %s RETURN OLD`, colName)
	meta, err := h.Execute(query, map[string]interface{}{
		"id": key,
	})
	if err != nil {
		logrus.Error("error executing query: ", err)
		return nil, err
	}
	logrus.Info("meta: ", meta[0])
	return meta[0], err
}
