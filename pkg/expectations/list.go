package expectations

type List struct {
	BodyExpectations BodyExpectations `yaml:"body,omitempty" json:"body,omitempty"`
}

type BodyExpectations struct {
	BodyEqual    *BodyEqual    `yaml:"eq,omitempty" json:"eq,omitempty"`
	BodyContains *BodyContains `yaml:"contains,omitempty" json:"contains,omitempty"`
}
