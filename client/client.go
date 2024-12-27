package client

import (
	ti "github.com/thingsdb/go-thingsdb"
)

// ThingsDBClient is a wrapper around the ThingsDB Connection
type ThingsDBClient struct {
	conn *ti.Conn
}

// NewThingsDBClient creates a new instance of ThingsDBClient
func NewThingsDBClient(host, token string, port uint16) (*ThingsDBClient, error) {
	conn := ti.NewConn(host, port, nil)

	if err := conn.Connect(); err != nil {
		return nil, err
	}

	if err := conn.AuthToken(token); err != nil {
		return nil, err
	}

	return &ThingsDBClient{conn: conn}, nil
}

// GetTypeInfoByName retrieves a type definition in ThingsDB by its name
func (t *ThingsDBClient) GetTypeInfoByName(collectionName string, typeName string) (interface{}, error) {
	vars := map[string]interface{}{
		"name": typeName,
	}

	res, err := t.conn.Query("//"+collectionName, "type_info(`{name}`);", vars)
	if err != nil {
		return nil, err
	}

	return res, nil
}
