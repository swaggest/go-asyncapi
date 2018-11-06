# My Lovely Messaging API 0.0.0 documentation




## Table of Contents

* [Connection Details](#servers)
* [Topics](#topics)
* [Schemas](#schemas)


<a name="servers"></a>
## Connection details

<table class="table">
  <thead class="table__head">
    <tr class="table__head__row">
      <th class="table__head__cell">URL</th>
      <th class="table__head__cell">Scheme</th>
      <th class="table__head__cell">Description</th>
    </tr>
  </thead>
  <tbody class="table__body">
    <tr class="table__body__row">
      <td class="table__body__cell">api.lovely.com:{port}</td>
      <td class="table__body__cell">amqp</td>
      <td class="table__body__cell"></td>
    </tr>


  </tbody>
</table>


## Topics

<a name="topic-another.one"></a>
<h3><code>subscribe</code>
another.one
</h3>


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
        <td></td>
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
        <td></td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>item.values </td>
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
  "item": {
    "key": "string",
    "values": [
      0
    ]
  }
}
```

</div>
<a name="topic-one.{name}.two"></a>
<h3><code>publish</code>
one.{name}.two
</h3>

#### Topic Parameters

##### name


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

This is a sample schema


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
        <td></td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>items </td>
        <td>
          array(object)
        </td>
        <td></td>
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
  "createdAt": "2018-11-06T07:30:41Z",
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

</div>

## Messages


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
        <td></td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>item.values </td>
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
        <td></td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>items </td>
        <td>
          array(object)
        </td>
        <td></td>
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
  "createdAt": "2018-11-06T07:30:41Z",
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
        <td></td>
        <td><em>Any</em></td>
      </tr>
      <tr>
        <td>values </td>
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
  "key": "string",
  "values": [
    0
  ]
}
```
