# Godocdown [![GoDoc](https://godoc.org/github.com/dotpy3/godocdown?status.svg)](https://godoc.org/github.com/dotpy3/godocdown)

Package `godocdown` generates Go documentation in a GitHub-friendly Markdown format.

This is a fork of [robertkrimen/godocdown](https://github.com/robertkrimen/godocdown). The original is only available as a command, whereas this one is available as a Go package.

## Usage

```
// Executed in the dotpy3/godocdown directory
doc, err := godocdown.GenerateDocumentation(".", "")
if err != nil {
    panic(err)
}

// Prints the documentation of this package
fmt.Println(doc)
```

## Templating

In addition to Markdown rendering, godocdown provides templating via text/template (http://golang.org/pkg/text/template/) for further customization. By putting a file named ".godocdown.template" (or one from the list below) in the same directory as your package/command, godocdown will know to use the file as a template.

```
# text/template
.godocdown.markdown
.godocdown.md
.godocdown.template
.godocdown.tmpl
```

A template file can also be specified with the `templatePath` parameter of `GenerateDocumentation()`.

Along with the standard template functionality, the starting data argument has the following interface:

```
{{ .Emit }}                                                                                       
// Emit the standard documentation (what godocdown would emit without a template)                 
                                                                                                    
{{ .EmitHeader }}                                                                                 
// Emit the package name and an import line (if one is present/needed)                            
                                                                                                    
{{ .EmitSynopsis }}                                                                               
// Emit the package declaration                                                                   
                                                                                                    
{{ .EmitUsage }}                                                                                  
// Emit package usage, which includes a constants section, a variables section,                   
// a functions section, and a types section. In addition, each type may have its own constant,    
// variable, and/or function/method listing.                                                      
                                                                                                    
{{ if .IsCommand  }} ... {{ end }}                                                                
// A boolean indicating whether the given package is a command or a plain package                 
                                                                                                    
{{ .Name }}                                                                                       
// The name of the package/command (string)                                                       
                                                                                                    
{{ .ImportPath }}                                                                                 
// The import path for the package (string)                                                       
// (This field will be the empty string if godocdown is unable to guess it)
```
