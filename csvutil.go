package csvutil

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"sync"
)

// CSVData is a struct that holds the data of a CSV file.
type CSVData struct {
	cfg *Config

	mu   sync.Mutex
	rows []*CSVRow
}

// New creates a new CSVData struct with a configuration.
func New(cfg *Config) *CSVData {
	return &CSVData{cfg: cfg}
}

// NewRow creates a new CSVRow struct.
func (c *CSVData) NewRow() *CSVRow {
	if c.mu.TryLock() {
		defer c.mu.Unlock()
	}
	r := &CSVRow{c: c.cfg}
	c.rows = append(c.rows, r)

	return r
}

// EncodeJSON encode a JSON to CSVRows.
func (c *CSVData) EncodeJSON(persist bool, v any) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if !persist {
		c.rows = make([]*CSVRow, 0)
	}

	if len(c.cfg.cols) == 0 {
		return nil
	}

	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	m := make([]any, 0)
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}

	for _, v := range m {
		r := c.NewRow()
		// TODO: should support []any as well
		d, ok := v.(map[string]any)
		if !ok {
			return fmt.Errorf("invalid JSON object")
		}

		for k, v := range d {
			r.Set(k, v)
		}
	}

	return nil
}

// Write writes the CSVData to a writer.
func (c *CSVData) Write(w io.Writer) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	wt := csv.NewWriter(w)
	defer wt.Flush()

	if c.cfg.writeHeader {
		headers := make([]string, len(c.cfg.cols))
		for _, c := range c.cfg.cols {
			headers[c.index] = c.Key
		}

		if err := wt.Write(headers); err != nil {
			return err
		}
	}

	for _, row := range c.rows {
		data := make([]string, len(c.cfg.cols))
		for _, c := range c.cfg.cols {
			if d, ok := row.data[c.Key]; ok {
				data[c.index] = fmt.Sprintf("%v", d)
			}
		}

		if err := wt.Write(data); err != nil {
			return err
		}
	}

	return nil
}
