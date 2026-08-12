package scalar

import (
	"fmt"
	"strconv"
)

type IntID int

func (IntID) ImplementsGraphQLType(name string) bool {
	return name == "ID"
}

func (id *IntID) UnmarshalGraphQL(input any) error {
	switch input := input.(type) {
	case string:
		i, err := strconv.Atoi(input)
		if err != nil {
			return err
		}
		*id = IntID(i)
	case int32:
		*id = IntID(input)
	default:
		return fmt.Errorf("wrong type for IntID: %T", input)
	}
	return nil
}

func (id IntID) MarshalJSON() ([]byte, error) {
	return strconv.AppendQuote(nil, strconv.Itoa(int(id))), nil
}

type NullIntID struct {
	Value *IntID
	Set   bool
}

func (NullIntID) ImplementsGraphQLType(name string) bool {
	return name == "ID"
}

func (s *NullIntID) UnmarshalGraphQL(input any) error {
	s.Set = true

	if input == nil {
		return nil
	}

	s.Value = new(IntID)
	return s.Value.UnmarshalGraphQL(input)
}

func (s *NullIntID) Nullable() {}
