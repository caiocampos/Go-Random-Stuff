package tree

import (
	"errors"
	"sort"
)

// Node type defines a node of a tree
type Node struct {
	ID       int
	Children []*Node
}

// Record type defines a record to be restructured
type Record struct {
	ID     int
	Parent int
}

// Build function restructures the RecordList in a tree structure
func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}
	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})
	if records[0].ID != 0 || records[0].Parent > 0 {
		return nil, errors.New("Invalid root")
	}
	nodes := make(map[int]*Node)
	nodes[0] = &Node{ID: 0}
	for i, record := range records {
		if i == 0 {
			continue
		}
		if record.ID != i {
			return nil, errors.New("Non-continuous Tree")
		}
		if record.Parent >= i {
			return nil, errors.New("Invalid Child ID")
		}
		nodes[i] = &Node{ID: i}
		nodes[record.Parent].Children = append(nodes[record.Parent].Children, nodes[i])
	}
	return nodes[0], nil
}
