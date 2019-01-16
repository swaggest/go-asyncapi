# My Lovely Messaging API 1.2.3 documentation




## Table of Contents

* [Connection Details](#servers)
* [Topics](#topics)
* [Messages](#messages)
* [Schemas](#schemas)


<a name="servers"></a>
## Connection details

<table>
  <thead>
    <tr>
      <th>URL</th>
      <th>Scheme</th>
      <th>Description</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td>api.lovely.com:{port}</td>
      <td>amqp</td>
      <td></td>
    </tr>


  </tbody>
</table>


## Topics

<a name="topic-another.one"></a>

### `subscribe` another.one


#### Message

Sample consumer

This is another sample schema

##### Headers


<table>
  <thead>
    <tr>
      <th>Name</th>
      <th>Type</th>
      <th>Description</th>
      <th>Accepted values</th>
    </tr>
  </thead>
  <tbody>
      <tr>
        <td>X-Trace-ID <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td><p>Tracing header</p>
      </td>
        <td><em>Any</em></td>
      </tr>
  </tbody>
</table>


###### Example of headers _(generated)_

```json
{
  "X-Trace-ID": "string"
}
```

##### Payload


<table>
  <thead>
    <tr>
      <th>Name</th>
      <th>Type</th>
      <th>Description</th>
      <th>Accepted values</th>
    </tr>
  </thead>
  <tbody>
      <tr>
        <td>item </td>
        <td>
          object
        </td>
        <td></td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>item.key </td>
        <td>
          string
        </td>
        <td><p>Item key</p>
      </td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>item.values </td>
        <td>
          array(integer)
        </td>
        <td><p>List of item values</p>
      </td>
        <td><em>Any</em></td>
      </tr>
  </tbody>
</table>


###### Example of payload _(generated)_

```json
{
  "item": {
    "key": "string",
    "values": [
      0
    ]
  }
}
```

<a name="topic-one.{name}.two"></a>

### `publish` one.{name}.two

#### Topic Parameters

##### name

Name

<table>
  <thead>
    <tr>
      <th>Name</th>
      <th>Type</th>
      <th>Description</th>
      <th>Accepted values</th>
    </tr>
  </thead>
  <tbody>
      <tr>
        <td>name </td>
        <td>
          string
        </td>
        <td></td>
        <td><em>Any</em></td>
      </tr>
  </tbody>
</table>


#### Message

Sample publisher

This is a sample schema, AMQP VHost: some-vhost, AMQP Exchange: some-exchange, AMQP RoutingKey: some-key


##### Payload


<table>
  <thead>
    <tr>
      <th>Name</th>
      <th>Type</th>
      <th>Description</th>
      <th>Accepted values</th>
    </tr>
  </thead>
  <tbody>
      <tr>
        <td>createdAt </td>
        <td>
          string
        </td>
        <td><p>Creation time</p>
      </td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>items </td>
        <td>
          array(object)
        </td>
        <td><p>List of items</p>
      </td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>items.key </td>
        <td>
          string
        </td>
        <td></td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>items.values </td>
        <td>
          array(integer)
        </td>
        <td></td>
        <td><em>Any</em></td>
      </tr>
  </tbody>
</table>


###### Example of payload _(generated)_

```json
{
  "createdAt": "2019-01-16T13:20:16Z",
  "items": [
    {
      "key": "string",
      "values": [
        0
      ]
    }
  ]
}
```



## Messages

### MyAnotherMessage 
Sample consumer

This is another sample schema

#### Headers


<table>
  <thead>
    <tr>
      <th>Name</th>
      <th>Type</th>
      <th>Description</th>
      <th>Accepted values</th>
    </tr>
  </thead>
  <tbody>
      <tr>
        <td>X-Trace-ID <strong>(required)</strong></td>
        <td>
          string
        </td>
        <td><p>Tracing header</p>
      </td>
        <td><em>Any</em></td>
      </tr>
  </tbody>
</table>


##### Example of headers _(generated)_

```json
{
  "X-Trace-ID": "string"
}
```

#### Payload


<table>
  <thead>
    <tr>
      <th>Name</th>
      <th>Type</th>
      <th>Description</th>
      <th>Accepted values</th>
    </tr>
  </thead>
  <tbody>
      <tr>
        <td>item </td>
        <td>
          object
        </td>
        <td></td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>item.key </td>
        <td>
          string
        </td>
        <td><p>Item key</p>
      </td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>item.values </td>
        <td>
          array(integer)
        </td>
        <td><p>List of item values</p>
      </td>
        <td><em>Any</em></td>
      </tr>
  </tbody>
</table>


##### Example of payload _(generated)_

```json
{
  "item": {
    "key": "string",
    "values": [
      0
    ]
  }
}
```

### MyMessage 
Sample publisher

This is a sample schema, AMQP VHost: some-vhost, AMQP Exchange: some-exchange, AMQP RoutingKey: some-key



#### Payload


<table>
  <thead>
    <tr>
      <th>Name</th>
      <th>Type</th>
      <th>Description</th>
      <th>Accepted values</th>
    </tr>
  </thead>
  <tbody>
      <tr>
        <td>createdAt </td>
        <td>
          string
        </td>
        <td><p>Creation time</p>
      </td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>items </td>
        <td>
          array(object)
        </td>
        <td><p>List of items</p>
      </td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>items.key </td>
        <td>
          string
        </td>
        <td></td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>items.values </td>
        <td>
          array(integer)
        </td>
        <td></td>
        <td><em>Any</em></td>
      </tr>
  </tbody>
</table>


##### Example of payload _(generated)_

```json
{
  "createdAt": "2019-01-16T13:20:16Z",
  "items": [
    {
      "key": "string",
      "values": [
        0
      ]
    }
  ]
}
```


## Schemas

#### MyAnotherMessage

<table>
  <thead>
    <tr>
      <th>Name</th>
      <th>Type</th>
      <th>Description</th>
      <th>Accepted values</th>
    </tr>
  </thead>
  <tbody>
      <tr>
        <td>item </td>
        <td>
          object
        </td>
        <td></td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>item.key </td>
        <td>
          string
        </td>
        <td><p>Item key</p>
      </td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>item.values </td>
        <td>
          array(integer)
        </td>
        <td><p>List of item values</p>
      </td>
        <td><em>Any</em></td>
      </tr>
  </tbody>
</table>

##### Example _(generated)_

```json
{
  "item": {
    "key": "string",
    "values": [
      0
    ]
  }
}
```
#### MyMessage

<table>
  <thead>
    <tr>
      <th>Name</th>
      <th>Type</th>
      <th>Description</th>
      <th>Accepted values</th>
    </tr>
  </thead>
  <tbody>
      <tr>
        <td>createdAt </td>
        <td>
          string
        </td>
        <td><p>Creation time</p>
      </td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>items </td>
        <td>
          array(object)
        </td>
        <td><p>List of items</p>
      </td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>items.key </td>
        <td>
          string
        </td>
        <td></td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>items.values </td>
        <td>
          array(integer)
        </td>
        <td></td>
        <td><em>Any</em></td>
      </tr>
  </tbody>
</table>

##### Example _(generated)_

```json
{
  "createdAt": "2019-01-16T13:20:16Z",
  "items": [
    {
      "key": "string",
      "values": [
        0
      ]
    }
  ]
}
```
#### SubItem

<table>
  <thead>
    <tr>
      <th>Name</th>
      <th>Type</th>
      <th>Description</th>
      <th>Accepted values</th>
    </tr>
  </thead>
  <tbody>
      <tr>
        <td>key </td>
        <td>
          string
        </td>
        <td><p>Item key</p>
      </td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>values </td>
        <td>
          array(integer)
        </td>
        <td><p>List of item values</p>
      </td>
        <td><em>Any</em></td>
      </tr>
  </tbody>
</table>

##### Example _(generated)_

```json
{
  "key": "string",
  "values": [
    0
  ]
}
```
