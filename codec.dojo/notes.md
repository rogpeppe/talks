techniques:

- straight values
- structs
- maps
- tags
	- omitempty
	- rename
	- omit
- UseNumber
- base64
- embedding
- interfaces
- custom encode/decode
- TextMarshaler/TextUnmarshaler
- Raw
- streaming


# Exercises:

1. Encode a Go string to JSON.
	- import encoding/json
	- import os
	- json.Marshal
	- os.Stdout.Write(data)

2. Try it with a different types
	- an integer?
	- a float?
	- a slice?
	- a function?
	- a map?
	- Always check your errors!

3. Decode a JSON string into a Go string.
	- use back-quoted string literal to put JSON in code.
	- json.Unmarshal
	- remember to unmarshal into a pointer to the object!

3. Encode a Go struct into JSON
	- type T struct { ... }
	- what's unusual about the JSON produced?
	- what happens when you try to fix it?

4. Rename fields with struct tags
	- A "struct tag" is a string literal following a struct field
	- See https://golang.org/pkg/reflect/#StructTag
	- Use backquotes
	- go vet will tell you about malformed values.

5.  Decode a JSON object into a Go struct

6. Encode a Go struct as JSON but omit empty strings when encoded.
	- use omitempty flag.

7. Decode the JSON values held in the file json1.json, json2.json and json3.json.
	- Try decoding into the empty interface (interface{})
	- Print some of the values in the data structure.
	- Experiment with http://json2struct.mervine.net/
	- When might that *not* be appropriate?

8. Decode the JSON value held in the file json4.json
	- See https://golang.org/pkg/encoding/json/#Decoder.UseNumber
	- https://golang.org/pkg/math/big/#Int
	- big.Int does not know about JSON? How does it manage to do it.
	- https://golang.org/pkg/encoding/#TextUnmarshaler

9. Given the following struct type, add methods to make it encode and decode as a
dot-separated version number. For example Version{1, 2, 4} should encode
as "1.2.4".

		type Version struct {
			Major, Minor, Patch int
		}

	- See https://golang.org/pkg/strings/#Split
	- See https://golang.org/pkg/strconv/#Atoi

10. Given the following struct type, define methods on it to decode the JSON in json5.json to it.

		type Shapes []Shape
		type Shape interface {
			Draw()
		}

	- You'll need to decode in two stages.
	- See https://golang.org/pkg/encoding/json/#Marshaler
	- See https://golang.org/pkg/encoding/json/#Unmarshaler
	- See https://golang.org/pkg/encoding/json/#RawMessage

12. Experiment with other codecs.
	- See gopkg.in/yaml.v2
	- See gopg.in/mgo.v2/bson
	- encoding/xml
	- How are they the same/different?

11. (advanced!) Define a codec for this type: https://godoc.org/github.com/mvdan/sh/syntax#Stmt

	- problem described by Daniel https://groups.google.com/d/msg/golang-nuts/iAJEj7kIbtc/DxNmSBkOFQAJ
	- See https://golang.org/pkg/reflect/#StructOf
	- or fork encoding/json
	- or... ?
