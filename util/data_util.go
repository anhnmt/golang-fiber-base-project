package util

import (
	"fmt"

	"github.com/bwmarrin/snowflake"
	"github.com/gofiber/fiber/v2/utils"
)

// GenerateSnowflakeID generate a new Snowflake UUID
func GenerateSnowflakeID(node ...int64) snowflake.ID {
	var nodeNum int64 = 0

	if node != nil && len(node) > 0 {
		nodeNum = node[0]
	}

	// Create a new Node with a Node number of 1
	newNode, err := snowflake.NewNode(nodeNum)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	// Generate a snowflake ID.
	return newNode.Generate()
}

// GenerateUUID generate a new UUIDv4
func GenerateUUID() string {
	// Generate a UUID.
	return utils.UUIDv4()
}

// GetFirst get the first index if is array interface
func GetFirst(args interface{}, num ...int) interface{} {
	var check = 0

	if num != nil && len(num) > 0 {
		check = num[0]
	}

	if args == nil || check > 3 {
		return nil
	}

	switch args.(type) {
	case []interface{}:
		first := args.([]interface{})

		if len(first) == 0 {
			return nil
		}

		if result, ok := first[0].([]interface{}); ok {
			return GetFirst(result, check+1)
		}

		return first[0]
	case interface{}:
		return args

	default:
		return nil
	}
}
