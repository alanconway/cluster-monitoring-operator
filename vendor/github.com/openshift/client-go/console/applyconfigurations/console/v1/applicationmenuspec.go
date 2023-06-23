// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// ApplicationMenuSpecApplyConfiguration represents an declarative configuration of the ApplicationMenuSpec type for use
// with apply.
type ApplicationMenuSpecApplyConfiguration struct {
	Section  *string `json:"section,omitempty"`
	ImageURL *string `json:"imageURL,omitempty"`
}

// ApplicationMenuSpecApplyConfiguration constructs an declarative configuration of the ApplicationMenuSpec type for use with
// apply.
func ApplicationMenuSpec() *ApplicationMenuSpecApplyConfiguration {
	return &ApplicationMenuSpecApplyConfiguration{}
}

// WithSection sets the Section field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Section field is set to the value of the last call.
func (b *ApplicationMenuSpecApplyConfiguration) WithSection(value string) *ApplicationMenuSpecApplyConfiguration {
	b.Section = &value
	return b
}

// WithImageURL sets the ImageURL field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ImageURL field is set to the value of the last call.
func (b *ApplicationMenuSpecApplyConfiguration) WithImageURL(value string) *ApplicationMenuSpecApplyConfiguration {
	b.ImageURL = &value
	return b
}