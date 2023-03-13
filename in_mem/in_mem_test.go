package in_mem

import "testing"

func TestInMem(t *testing.T) {
	db := NewDB()

	db.NewNodeType("Person", map[string]string{
		"age": "number",
	})
	db.NewEdgeType("friend", "Person", "Person", map[string]string{
		"since": "number",
	})

	nodeInstance := db.InsertNode("Person", "1", map[string]any{
		"age": 29,
	})
	db.InsertNode("Person", "2", map[string]any{
		"age": 29,
	})
	db.InsertEdge("friend", "1", "1", "2", map[string]any{
		"since": 2022,
	})

	traversalResult := db.Traverse(nodeInstance, "friend")
	if traversalResult.ID != "2" {
		t.Errorf("Expected traversal result to be \"2\" but was %s\n", traversalResult.ID)
	}
}
