package client

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
	"github.com/rickmoonex/gotigen/internal/models"
	"github.com/thingsdb/go-thingsdb"
)

// Client is a wrapper around the ThingsDB connection
type Client struct {
	*thingsdb.Conn
}

// NewClient creates a new client object
func NewClient(host string, port uint16, token string) (*Client, error) {
	conn := thingsdb.NewConn(host, port, nil)

	if err := conn.Connect(); err != nil {
		return nil, err
	}

	if err := conn.AuthToken(token); err != nil {
		return nil, err
	}

	return &Client{conn}, nil
}

// GetTypes returns all the type definitions in a given scope
func (c *Client) GetTypes(scope string) (*[]models.ThingsDBType, error) {
	var types []models.ThingsDBType

	res, err := c.Query(scope, "types_info();", nil)
	if err != nil {
		return nil, err
	}

	if err := mapstructure.Decode(res, &types); err != nil {
		return nil, fmt.Errorf("error decoding types info response: %s", err)
	}

	return &types, nil
}

// GetProcedures returns all the procedure definitions in a given scope
func (c *Client) GetProcedures(scope string) (*[]models.ThingsDBProcedure, error) {
	var procedures []models.ThingsDBProcedure

	res, err := c.Query(scope, "procedures_info();", nil)
	if err != nil {
		return nil, err
	}

	if err := mapstructure.Decode(res, &procedures); err != nil {
		return nil, fmt.Errorf("error decoding procedures info response: %s", err)
	}

	return &procedures, nil
}

// GetEnums returns all the enum definitions in a given scope
func (c *Client) GetEnums(scope string) (*[]models.ThingsDBEnum, error) {
	var enums []models.ThingsDBEnum

	res, err := c.Query(scope, "enums_info();", nil)
	if err != nil {
		return nil, err
	}

	if err := mapstructure.Decode(res, &enums); err != nil {
		return nil, fmt.Errorf("error decoding enums info response: %s", err)
	}

	return &enums, nil
}
