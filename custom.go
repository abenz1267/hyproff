package main

type Custom struct {
	Label         string  `json:"label"`
	CustomEntries []Entry `json:"entries"`
}

func (c Custom) Identifier() string {
	return "custom"
}

func (c Custom) defaultConfig() Custom {
	return Custom{
		Label:         "Custom",
		CustomEntries: []Entry{},
	}
}

func (c Custom) IsAvailable(config Config) bool {
	return config.containsModule(c.Identifier())
}

func (c Custom) Entries() []Entry {
	return c.CustomEntries
}
