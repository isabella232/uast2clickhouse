package main

var typesToGoDeeper = map[string]struct{}{
	"": {},
	// Bash
	"bash:arithmetic_command":                  {},
	"bash:arithmetic_equality":                 {},
	"bash:arithmetic_exponent":                 {},
	"bash:arithmetic_multiplication":           {},
	"bash:arithmetic_post_incr":                {},
	"bash:arithmetic_simple":                   {},
	"bash:arithmetic_sum":                      {},
	"bash:arith_compund_comparision":           {},
	"bash:array_assignment_list":               {},
	"bash:backquote":                           {},
	"bash:backquote_shellcommand":              {},
	"bash:composed_command":                    {},
	"bash:composed_variable":                   {},
	"bash:file_ref":                            {},
	"bash:heredoc_content_element":             {},
	"bash:heredoc_end_element":                 {},
	"bash:heredoc_end_element_(ignoring_tabs)": {},
	"bash:heredoc_end_marker":                  {},
	"bash:heredoc_end_marker_(ignoring_tabs)":  {},
	"bash:heredoc_marker_tag":                  {},
	"bash:heredoc_start_element":               {},
	"bash:heredoc_start_marker":                {},
	"bash:let_command":                         {},
	"bash:pipeline_command":                    {},
	"bash:process_substitution_element":        {},
	"bash:redirect_element":                    {},
	"bash:redirect_list":                       {},
	"bash:string":                              {},
	"bash:subshell_shellcommand":               {},
	"bash:var_substitution":                    {},
	// C++
	"cpp:CPPASTQualifiedName":    {},
	"cpp:CPPASTUsingDeclaration": {},
	// C#
	"csharp:CloseBraceToken":   {},
	"csharp:CloseBracketToken": {},
	"csharp:CloseParenToken":   {},
	"csharp:ColonColonToken":   {},
	"csharp:ColonToken":        {},
	"csharp:DotToken":          {},
	"csharp:OpenBraceToken":    {},
	"csharp:OpenBracketToken":  {},
	"csharp:OpenParenToken":    {},
	"csharp:SemicolonToken":    {},
}
