package csvutil

// Config is a struct that holds the configuration of a CSV file.
type Config struct {
	cols          map[string]*CSVCol
	defaultValues map[string]any
	silent        bool
	hooks         []func(*CSVCol, any) any

	// calculate info.
}

// SetSilent sets the silent flag of the CSV file.
func (c *Config) SetSilent(silent bool) *Config {
	c.silent = silent
	return c
}

// SetCols sets the columns of the CSV file.
func (c *Config) SetCols(cols ...*CSVCol) *Config {
	c.cols = make(map[string]*CSVCol)
	for i, col := range cols {
		col.index = i
		c.cols[col.Key] = col
	}

	return c
}

// ApplyHook applies a hook to a column.
func (c *Config) ApplyHook(hooks ...func(*CSVCol, any) any) *Config {
	c.hooks = append(c.hooks, hooks...)

	return c
}
