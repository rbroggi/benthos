---
title: xml
type: processor
status: beta
categories: ["Parsing"]
---

<!--
     THIS FILE IS AUTOGENERATED!

     To make changes please edit the contents of:
     lib/processor/xml.go
-->

import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';

:::caution BETA
This component is mostly stable but breaking changes could still be made outside of major version releases if a fundamental problem with the component is found.
:::

Parses messages as an XML document, performs a mutation on the data, and then
overwrites the previous contents with the new value.


<Tabs defaultValue="common" values={[
  { label: 'Common', value: 'common', },
  { label: 'Advanced', value: 'advanced', },
]}>

<TabItem value="common">

```yaml
# Common config fields, showing default values
label: ""
xml:
  operator: to_json
  cast: false
```

</TabItem>
<TabItem value="advanced">

```yaml
# All config fields, showing default values
label: ""
xml:
  operator: to_json
  cast: false
  parts: []
```

</TabItem>
</Tabs>

## Operators

### `to_json`

Converts an XML document into a JSON structure, where elements appear as keys of
an object according to the following rules:

- If an element contains attributes they are parsed by prefixing a hyphen,
  `-`, to the attribute label.
- If the element is a simple element and has attributes, the element value
  is given the key `#text`.
- XML comments, directives, and process instructions are ignored.
- When elements are repeated the resulting JSON value is an array.

For example, given the following XML:

```xml
<root>
  <title>This is a title</title>
  <description tone="boring">This is a description</description>
  <elements id="1">foo1</elements>
  <elements id="2">foo2</elements>
  <elements>foo3</elements>
</root>
```

The resulting JSON structure would look like this:

```json
{
  "root":{
    "title":"This is a title",
    "description":{
      "#text":"This is a description",
      "-tone":"boring"
    },
    "elements":[
      {"#text":"foo1","-id":"1"},
      {"#text":"foo2","-id":"2"},
      "foo3"
    ]
  }
}
```

With cast set to true, the resulting JSON structure would look like this:

```json
{
  "root":{
    "title":"This is a title",
    "description":{
      "#text":"This is a description",
      "-tone":"boring"
    },
    "elements":[
      {"#text":"foo1","-id":1},
      {"#text":"foo2","-id":2},
      "foo3"
    ]
  }
}
```

## Fields

### `operator`

An XML [operation](#operators) to apply to messages.


Type: `string`  
Default: `"to_json"`  
Options: `to_json`.

### `cast`

Whether to try to cast values that are numbers and booleans to the right type. Default: all values are strings.


Type: `bool`  
Default: `false`  

### `parts`

An optional array of message indexes of a batch that the processor should apply to.
If left empty all messages are processed. This field is only applicable when
batching messages [at the input level](/docs/configuration/batching).

Indexes can be negative, and if so the part will be selected from the end
counting backwards starting from -1.


Type: `array`  
Default: `[]`  


