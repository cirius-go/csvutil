package csvutil

import (
	"encoding/json"
	"fmt"
	"sync"
)

// CSVRow is a struct that holds the data of a row in a CSV file.
type CSVRow struct {
	c *Config

	mu   sync.Mutex
	data map[string]any
}

// Set sets the value of a column in a row.
func (c *CSVRow) Set(key string, value any) *CSVRow {
	c.mu.Lock()
	defer c.mu.Unlock()

	csvRow, ok := c.c.cols[key]
	if !ok {
		if !c.c.silent {
			fmt.Printf("Column %s not found\n", key)
		}
		return c
	}

	for _, hookFn := range c.c.hooks {
		value = hookFn(csvRow, value)
	}

	if c.data == nil {
		c.data = make(map[string]any)
	}

	c.data[key] = value
	return c
}

// EncodeJSON encodes a JSON object to a row.
func (c *CSVRow) EncodeJSON(v any) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	m := make(map[string]any)
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}

	for k, v := range m {
		c.Set(k, v)
	}

	return nil
}
