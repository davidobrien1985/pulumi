package hcl2

import (
	"fmt"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/pulumi/pulumi/pkg/codegen/hcl2/model"
)

func errorf(subject hcl.Range, f string, args ...interface{}) *hcl.Diagnostic {
	return diagf(hcl.DiagError, subject, f, args...)
}

func diagf(severity hcl.DiagnosticSeverity, subject hcl.Range, f string, args ...interface{}) *hcl.Diagnostic {
	message := fmt.Sprintf(f, args...)
	return &hcl.Diagnostic{
		Severity: severity,
		Summary:  message,
		Detail:   message,
		Subject:  &subject,
	}
}

func labelsErrorf(block *hclsyntax.Block, f string, args ...interface{}) *hcl.Diagnostic {
	startRange := block.LabelRanges[0]

	diagRange := hcl.Range{
		Filename: startRange.Filename,
		Start:    startRange.Start,
		End:      block.LabelRanges[len(block.LabelRanges)-1].End,
	}
	return errorf(diagRange, f, args...)
}

func notYetImplemented(v interface{}) hcl.Diagnostics {
	var subject hcl.Range
	switch v := v.(type) {
	case model.Definition:
		subject = v.SyntaxNode().Range()
	case interface{ Range() hcl.Range }:
		subject = v.Range()
	}
	return hcl.Diagnostics{errorf(subject, "NYI: %v", v)}
}

func malformedToken(token string, sourceRange hcl.Range) *hcl.Diagnostic {
	return errorf(sourceRange, "malformed token '%v': expected 'pkg:module:member'", token)
}

func unknownPackage(pkg string, tokenRange hcl.Range) *hcl.Diagnostic {
	return errorf(tokenRange, "unknown package '%s'", pkg)
}

func unknownResourceType(token string, tokenRange hcl.Range) *hcl.Diagnostic {
	return errorf(tokenRange, "unknown resource type '%s'", token)
}

func circularReference(stack []hclsyntax.Node, referent hclsyntax.Node) *hcl.Diagnostic {
	// TODO(pdg): stack trace
	return errorf(referent.Range(), "circular reference to node")
}

func unsupportedBlock(blockType string, typeRange hcl.Range) *hcl.Diagnostic {
	return errorf(typeRange, "unsupported block of type '%v'", blockType)
}

func unsupportedAttribute(attrName string, nameRange hcl.Range) *hcl.Diagnostic {
	return errorf(nameRange, "unsupported attribute '%v'", attrName)
}

func missingRequiredAttribute(attrName string, missingRange hcl.Range) *hcl.Diagnostic {
	return errorf(missingRange, "missing required attribute '%v'", attrName)
}
