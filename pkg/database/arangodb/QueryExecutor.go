package arangodb

//QueryExecutor is an abstraction to simplify mocking queries for testing
type QueryExecutor interface {
	Execute(queryText string, bindVars map[string]interface{}) ([]map[string]interface{}, error)
}
