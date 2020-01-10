package hcl2template

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/ext/typeexpr"
	"github.com/zclconf/go-cty/cty"
)

type InputVariable struct {
	Default cty.Value
	Type    cty.Type

	block *hcl.Block
}

func (v *InputVariable) Value() cty.Value {
	return v.Default
}

type InputVariables map[string]InputVariable

func (variables InputVariables) Values() map[string]cty.Value {
	res := map[string]cty.Value{}
	for k, v := range variables {
		res[k] = v.Value()
	}
	return res
}

// decodeConfig decodes a "variables" section the way packer 1 used to
func (variables *InputVariables) decodeConfigMap(block *hcl.Block) hcl.Diagnostics {
	if (*variables) == nil {
		(*variables) = InputVariables{}
	}
	attrs, diags := block.Body.JustAttributes()

	if diags.HasErrors() {
		return diags
	}

	for key, attr := range attrs {
		if _, found := (*variables)[key]; found {
			diags = append(diags, &hcl.Diagnostic{
				Severity: hcl.DiagError,
				Summary:  "Duplicate variable",
				Detail:   "Duplicate " + key + " variable found.",
				Subject:  attr.NameRange.Ptr(),
				Context:  block.DefRange.Ptr(),
			})
			continue
		}
		value, moreDiags := attr.Expr.Value(nil)
		diags = append(diags, moreDiags...)
		if moreDiags.HasErrors() {
			continue
		}
		(*variables)[key] = InputVariable{
			Default: value,
			Type:    value.Type(),
		}
	}

	return diags
}

// decodeConfig decodes a "variables" section the way packer 1 used to
func (variables *InputVariables) decodeConfig(block *hcl.Block) hcl.Diagnostics {
	if (*variables) == nil {
		(*variables) = InputVariables{}
	}

	attrs, diags := block.Body.JustAttributes()

	if diags.HasErrors() {
		return diags
	}

	res := InputVariable{
		block: block,
	}
	if def, ok := attrs["default"]; ok {
		defaultValue, moreDiags := def.Expr.Value(nil)
		diags = append(diags, moreDiags...)
		if moreDiags.HasErrors() {
			return diags
		}
		res.Default = defaultValue
		res.Type = defaultValue.Type()
	}
	if t, ok := attrs["type"]; ok {
		tp, moreDiags := typeexpr.Type(t.Expr)
		diags = append(diags, moreDiags...)
		if moreDiags.HasErrors() {
			return diags
		}

		res.Type = tp
	}

	(*variables)[block.Labels[0]] = res

	return diags
}
