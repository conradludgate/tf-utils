package tfutils

func (s SimpleSchema) Required(status bool) SimpleSchema {
	s.s.Required = status
	return s
}

func (s SimpleSchema) Optional(status bool) SimpleSchema {
	s.s.Optional = status
	return s
}

func (s SimpleSchema) Computed(status bool) SimpleSchema {
	s.s.Computed = status
	return s
}

func (s SimpleSchema) Default(d interface{}) SimpleSchema {
	s.s.Default = d
	return s.Optional(true)
}

func (s SimpleSchema) ConflictsWith(t ...string) SimpleSchema {
	s.s.ConflictsWith = t
	return s
}

func (s SimpleSchema) ExactlyOneOf(t ...string) SimpleSchema {
	s.s.ExactlyOneOf = t
	return s
}

func (s SimpleSchema) AtLeastOneOf(t ...string) SimpleSchema {
	s.s.AtLeastOneOf = t
	return s
}

func (s SimpleSchema) RequiredWith(t ...string) SimpleSchema {
	s.s.RequiredWith = t
	return s
}

func (s SimpleSchema) Sensitive(status bool) SimpleSchema {
	s.s.Sensitive = status
	return s
}

func (s ListSchema) Required(status bool) ListSchema {
	s.s.Required = status
	return s
}

func (s ListSchema) Optional(status bool) ListSchema {
	s.s.Optional = status
	return s
}

func (s ListSchema) Computed(status bool) ListSchema {
	s.s.Computed = status
	return s
}

func (s ListSchema) Default(d interface{}) ListSchema {
	s.s.Default = d
	return s.Optional(true)
}

func (s ListSchema) ConflictsWith(t ...string) ListSchema {
	s.s.ConflictsWith = t
	return s
}

func (s ListSchema) ExactlyOneOf(t ...string) ListSchema {
	s.s.ExactlyOneOf = t
	return s
}

func (s ListSchema) AtLeastOneOf(t ...string) ListSchema {
	s.s.AtLeastOneOf = t
	return s
}

func (s ListSchema) RequiredWith(t ...string) ListSchema {
	s.s.RequiredWith = t
	return s
}

func (s ListSchema) Sensitive(status bool) ListSchema {
	s.s.Sensitive = status
	return s
}

func (s SetSchema) Required(status bool) SetSchema {
	s.s.Required = status
	return s
}

func (s SetSchema) Optional(status bool) SetSchema {
	s.s.Optional = status
	return s
}

func (s SetSchema) Computed(status bool) SetSchema {
	s.s.Computed = status
	return s
}

func (s SetSchema) Default(d interface{}) SetSchema {
	s.s.Default = d
	return s.Optional(true)
}

func (s SetSchema) ConflictsWith(t ...string) SetSchema {
	s.s.ConflictsWith = t
	return s
}

func (s SetSchema) ExactlyOneOf(t ...string) SetSchema {
	s.s.ExactlyOneOf = t
	return s
}

func (s SetSchema) AtLeastOneOf(t ...string) SetSchema {
	s.s.AtLeastOneOf = t
	return s
}

func (s SetSchema) RequiredWith(t ...string) SetSchema {
	s.s.RequiredWith = t
	return s
}

func (s SetSchema) Sensitive(status bool) SetSchema {
	s.s.Sensitive = status
	return s
}

func (s MapSchema) Required(status bool) MapSchema {
	s.s.Required = status
	return s
}

func (s MapSchema) Optional(status bool) MapSchema {
	s.s.Optional = status
	return s
}

func (s MapSchema) Computed(status bool) MapSchema {
	s.s.Computed = status
	return s
}

func (s MapSchema) Default(d interface{}) MapSchema {
	s.s.Default = d
	return s.Optional(true)
}

func (s MapSchema) ConflictsWith(t ...string) MapSchema {
	s.s.ConflictsWith = t
	return s
}

func (s MapSchema) ExactlyOneOf(t ...string) MapSchema {
	s.s.ExactlyOneOf = t
	return s
}

func (s MapSchema) AtLeastOneOf(t ...string) MapSchema {
	s.s.AtLeastOneOf = t
	return s
}

func (s MapSchema) RequiredWith(t ...string) MapSchema {
	s.s.RequiredWith = t
	return s
}

func (s MapSchema) Sensitive(status bool) MapSchema {
	s.s.Sensitive = status
	return s
}
