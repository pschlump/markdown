/*
Package rtf implements RTF renderer of parsed markdown document.

# Configuring and customizing a renderer

A renderer can be configured with multiple options:

	import "github.com/pschlump/markdown/rtf"

	flags := rtf.CommonFlags | rtf.CompletePage | rtf.HrefTargetBlank
	opts := rtf.RendererOptions{
		Title: "A custom title",
		Flags: flags,
	}
	renderer := rtf.NewRenderer(opts)

You can also re-use most of the logic and customize rendering of selected nodes
by providing node render hook.
This is most useful for rendering nodes that allow for design choices, like
links or code blocks.

	import (
		"github.com/pschlump/markdown/rtf"
		"github.com/pschlump/markdown/ast"
	)

	// a very dummy render hook that will output "code_replacements" instead of
	// <code>${content}</code> emitted by rtf.Renderer
	func renderHookCodeBlock(w io.Writer, node ast.Node, depth int, entering bool) (ast.WalkStatus, bool) {
		_, ok := node.(*ast.CodeBlock)
		if !ok {
			return ast.GoToNext, false
		}
		io.WriteString(w, "code_replacement")
		return ast.GoToNext, true
	}

	opts := rtf.RendererOptions{
		RenderNodeHook: renderHookCodeBlock,
	}
	renderer := rtf.NewRenderer(opts)
*/
package rtf
