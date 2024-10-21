package close

type FieldCondition struct {
	typeOfFieldCondition string  "json:`type`"
	objectType           *string "json:`object_type`"
	fieldName            *string "json:`field_name`"
	customFieldID        *string "json:`custom_field_id`"
}

type Condition struct {
	typeOfCondition string   "json:`type`"
	mode            *string  "json:`mode`"
	value           *string  "json:`value`"
	values          []string "json:`values` json:`omitempty`"
}

type Constraint struct {
	typeOfConstraint  string          "json:`type`"
	objectType        *string         "json:`object_type`"
	mode              *string         "json:`mode`"
	value             *string         "json:`value`"
	thisObjectType    *string         "json:`this_object_type`"
	relatedObjectType *string         "json:`related_object_type`"
	relatedQuery      *Query          "json:`related_query`"
	field             *FieldCondition "json:`field`"
	condition         *Condition      "json:`condition`"
}

type Query interface {
	//TODO: Implement
}

type Filter struct {
	TypeOfQuery string  "json:`type`"
	Queries     []Query "json:`queries`"
}
