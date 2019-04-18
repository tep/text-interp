

# interp [![Mit License][mit-img]][mit] [![GitHub Release][release-img]][release] [![GoDoc][godoc-img]][godoc] [![Go Report Card][reportcard-img]][reportcard] [![Build Status][travis-img]][travis]

`import "toolman.org/text/interp"`

* [Install](#pkg-install)
* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-install">Install</a>

```sh
    go get toolman.org/text/interp
```

## <a name="pkg-overview">Overview</a>

Package interp provides a simple string variable interpolator.

## <a name="pkg-index">Index</a>
* [type Interpolator](#Interpolator)
  * [func New(g Resolver) *Interpolator](#New)
  * [func NewWithFormat(g Resolver, b *VarFormat) *Interpolator](#NewWithFormat)
  * [func (i *Interpolator) Interpolate(s string) (string, error)](#Interpolator.Interpolate)
* [type Resolver](#Resolver)
* [type Value](#Value)
* [type VarFormat](#VarFormat)


#### <a name="pkg-files">Package files</a>
[cat.go](/src/toolman.org/text/interp/cat.go) [interp.go](/src/toolman.org/text/interp/interp.go) [token.go](/src/toolman.org/text/interp/token.go) [varfmt.go](/src/toolman.org/text/interp/varfmt.go) 


## <a name="Interpolator">type</a> [Interpolator](/src/target/interp.go?s=915:981#L25)
``` go
type Interpolator struct {
    // contains filtered or unexported fields
}

```
An Interpolator is used to interpolate variable expressions embedded within
a body of text. By default, variable expression are of the form `${varname}`
and use a backslash as an escape character. An alternate expression format
may also be specified by calling `NewWithFormat`.

For both Interpolator constructors, the caller must provide a Resolver for
resolving variable names to their associated values.







### <a name="New">func</a> [New](/src/target/interp.go?s=1089:1123#L32)
``` go
func New(g Resolver) *Interpolator
```
New returns a new Interpolator for the given Resolver using the default
variable expression format.


### <a name="NewWithFormat">func</a> [NewWithFormat](/src/target/interp.go?s=1260:1318#L38)
``` go
func NewWithFormat(g Resolver, b *VarFormat) *Interpolator
```
NewWithFormat returns a new Interpolator for the given Resolver and
specified VarFormat.





### <a name="Interpolator.Interpolate">func</a> (\*Interpolator) [Interpolate](/src/target/interp.go?s=1807:1867#L50)
``` go
func (i *Interpolator) Interpolate(s string) (string, error)
```
Interpolate repeatedly searchs the given string for the inner-most variable
expression replacing each one with its associated value acquired from the
Interpolator's Resolver. If the Resolver returns an error, or if its
returned Value cannot be stringified, an empty string and error are
returned.

Interpolate continues this process until it no longer finds a variable
expression in the provided string, or an error is encountered.




## <a name="Resolver">type</a> [Resolver](/src/target/interp.go?s=211:486#L11)
``` go
type Resolver interface {
    // Resolve returns the Value associated with the given variable. If no Value
    // is found, Resolve may return nil and an error. Any error returned here
    // will ultimately be returned by Replace.
    Resolve(variable string) (value Value, err error)
}
```
Resolver is used by Interpolator to lookup values for interpolated variables.










## <a name="Value">type</a> [Value](/src/target/interp.go?s=2639:2661#L85)
``` go
type Value interface{}
```
Value is an enpty interface type returned by a Resolvers' Resolve method.
Its underlying value should be something that is ultimately stringable with
the following guidelines:


	* string values are used directly
	* String() is called for fmt.Stringer values
	* If Value is an encoding.TextMarshaler, its MarshalText()
	  method is called and the results are converted to a string
	* Otherwise the value is passed through fmt.Sprintf("%v")










## <a name="VarFormat">type</a> [VarFormat](/src/target/varfmt.go?s=175:405#L12)
``` go
type VarFormat struct {
    Begin  string // The string token that preceeds a variable name
    End    string // The string token that follows a variable name
    Escape byte   // The escale character used to skip one of the above tokens
}

```
VarFormat defines what a variable expression should look like.

[mit-img]: http://img.shields.io/badge/License-MIT-c41e3a.svg
[mit]: https://github.com/toolmanorg/text-interp/blob/master/LICENSE

[release-img]: https://img.shields.io/github/release/toolmanorg/text-interp/all.svg
[release]: https://github.com/toolmanorg/text-interp/releases

[godoc-img]: https://godoc.org/toolman.org/text/interp?status.svg
[godoc]: https://godoc.org/toolman.org/text/interp

[reportcard-img]: https://goreportcard.com/badge/toolman.org/text/interp
[reportcard]: https://goreportcard.com/report/toolman.org/text/interp

[travis-img]: https://travis-ci.org/tep/text-interp.svg?branch=master
[travis]: https://travis-ci.org/tep/text-interp

