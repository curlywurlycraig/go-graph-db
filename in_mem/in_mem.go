package in_mem

type NodeType struct {
	Name                string
	AttributeDefinition map[string]string
}

type NodeInstance struct {
	TypeName   string
	ID         string
	Attributes map[string]any
}

type EdgeType struct {
	Name                string
	AttributeDefinition map[string]string
	FromNodeType        string
	ToNodeType          string
}

type EdgeInstance struct {
	TypeName   string
	ID         string
	Attributes map[string]any
	FromNodeID string
	ToNodeID   string
}

type DBState struct {
	NodeTypes     map[string]NodeType
	NodeInstances map[string]map[string]NodeInstance
	EdgeTypes     map[string]EdgeType
	EdgeInstances map[string]map[string]EdgeInstance
}

func NewDB() DBState {
	return DBState{
		NodeTypes:     make(map[string]NodeType, 0),
		NodeInstances: make(map[string]map[string]NodeInstance, 0),
		EdgeTypes:     make(map[string]EdgeType, 0),
		EdgeInstances: make(map[string]map[string]EdgeInstance, 0),
	}
}

func (db *DBState) NewNodeType(name string, attributeDefinition map[string]string) NodeType {
	result := NodeType{
		Name:                name,
		AttributeDefinition: attributeDefinition,
	}
	db.NodeTypes[name] = result
	db.NodeInstances[name] = make(map[string]NodeInstance)
	return result
}

func (db *DBState) InsertNode(
	typeName string,
	ID string,
	attributes map[string]any,
) NodeInstance {
	// TODO Rudimentary type checking
	result := NodeInstance{
		TypeName:   typeName,
		ID:         ID,
		Attributes: attributes,
	}
	db.NodeInstances[typeName][ID] = result
	return result
}

func (db *DBState) NewEdgeType(
	name string,
	fromType string,
	toType string,
	attributeDefinition map[string]string,
) EdgeType {
	result := EdgeType{
		Name:                name,
		AttributeDefinition: attributeDefinition,
		FromNodeType:        fromType,
		ToNodeType:          toType,
	}
	db.EdgeTypes[name] = result
	db.EdgeInstances[name] = make(map[string]EdgeInstance)
	return result
}

func (db *DBState) InsertEdge(
	typeName string,
	ID string,
	fromID string,
	toID string,
	attributes map[string]any,
) EdgeInstance {
	// TODO Rudimentary type checking
	result := EdgeInstance{
		TypeName:   typeName,
		ID:         ID,
		Attributes: attributes,
		FromNodeID: fromID,
		ToNodeID:   toID,
	}
	db.EdgeInstances[typeName][ID] = result
	return result
}

func (db *DBState) Traverse(
	nodeInstance NodeInstance,
	edgeType string,
) NodeInstance {
	// This is interesting. Obviously iterating over all edges isn't ideal. This is the first
	// clear "problem" we have encountered. A traversal should be a constant time lookup.
	// Perhaps when adding an edge, the edge ID should be denormalised onto the node.

	// Additionally, an actual traversal would want to do something with
	// attributes of the edge, rather than just return the node on the other
	// end. So if we had some function here, the found edge would be used in some way
	for _, edge := range db.EdgeInstances[edgeType] {
		if edge.FromNodeID == nodeInstance.ID {
			return db.NodeInstances[nodeInstance.TypeName][edge.ToNodeID]
		} else if edge.ToNodeID == nodeInstance.ID {
			return db.NodeInstances[nodeInstance.TypeName][edge.FromNodeID]
		}
	}

	return NodeInstance{}
}
