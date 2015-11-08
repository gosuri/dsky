package ui

import (
	"bytes"
	"fmt"
)

// DefaultTableColumnSeperator is the default seperater for tabluar columns
var DefaultTableColumnSeperator = "\t"

// Table is a UI component that renders the data in formatted in a table
type Table struct {
	// Rows is the collection of rows in the table
	Rows []*Row

	// MaxCellWidth is the maximum allowed with for cells in the table
	MaxCellWidth int

	// Seperater for tabluar columns
	Seperator string
}

// Row represents a row in a table
type Row struct {
	// Cells is the group of cell for the row
	Cells []*Cell
}

// Cell Represents a column in a row
type Cell struct {
	Witdh int
	// Data is the cell data
	Data      interface{}
	formatted string
}

// NewTable returns an instance of Table with the given headers
func NewTable(headers ...interface{}) *Table {
	t := &Table{Seperator: DefaultTableColumnSeperator}
	if len(headers) == 0 {
		return t
	}
	t.AddRow(headers...)
	return t
}

// Add row adds the data as a row to the table
func (t *Table) AddRow(data ...interface{}) *Table {
	cells := make([]*Cell, len(data))
	for i, col := range data {
		cells[i] = &Cell{Data: col, Witdh: t.MaxCellWidth}
	}
	t.Rows = append(t.Rows, &Row{Cells: cells})
	return t
}

// Format returns the formated table
func (t *Table) String() string {
	// determine no of columns
	var colLen int
	for _, row := range t.rows() {
		if rlen := len(row.cells()); rlen > colLen {
			colLen = rlen
		}
	}

	// determine width for each column
	colWidths := make([]int, colLen)
	for _, row := range t.rows() {
		for i, cell := range row.cells() {
			if cellWidth := cell.width(t.MaxCellWidth); cellWidth > colWidths[i] {
				colWidths[i] = cellWidth
			}
		}
	}

	var buf bytes.Buffer
	for _, row := range t.rows() {
		for i, cell := range row.cells() {
			// format the data by resizing each cell to match the width of the column
			dat := alignStringLeft(cell.Format(t.MaxCellWidth), colWidths[i])
			buf.WriteString(dat + t.Seperator)
		}
		buf.WriteString("\n")
	}
	return buf.String()
}

// Format converts the data to a string. It inserts an ellipsis when the width
// of the cell exceeds the maxwidth. It does not trim when the maxwidth is set to 0
func (c *Cell) Format(maxwidth int) string {
	if len(c.formatted) == 0 {
		s := fmt.Sprintf("%v", c.Data)
		if maxwidth != 0 && len(s) > maxwidth {
			var buf bytes.Buffer
			b := []byte(s)
			for i := 0; i < maxwidth-3; i++ {
				buf.WriteByte(b[i])
			}
			buf.WriteString("...")
			c.formatted = buf.String()
			return c.formatted
		}
		c.formatted = s
	}
	return c.formatted
}

func (c *Cell) width(maxwidth int) int {
	return len(c.Format(maxwidth))
}

func alignStringLeft(s string, width int) string {
	var buf bytes.Buffer
	b := []byte(s)
	for i := 0; i < width; i++ {
		if i >= len(b) {
			buf.WriteString(" ")
		} else {
			buf.WriteByte(b[i])
		}
	}
	return buf.String()
}

func (t *Table) rows() []*Row {
	if t.Rows == nil {
		t.Rows = make([]*Row, 0)
	}
	return t.Rows
}

func (r *Row) cells() []*Cell {
	if r.Cells == nil {
		r.Cells = make([]*Cell, 0)
	}
	return r.Cells
}
