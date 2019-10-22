# My Lovely Messaging API 1.2.3 documentation





## Table of Contents



* [Servers](#servers)


* [Channels](#channels)





<a name="servers"></a>
## Servers

<table>
  <thead>
    <tr>
      <th>URL</th>
      <th>Protocol</th>
      <th>Description</th>
    </tr>
  </thead>
  <tbody>
  <tr>
      <td>api.lovely.com:{port}</td>
      <td>amqp</td>
      <td></td>
    </tr>
    <tr>
      <td colspan="3">
        <details>
          <summary>URL Variables</summary>
          <table>
            <thead>
              <tr>
                <th>Name</th>
                <th>Default value</th>
                <th>Possible values</th>
                <th>Description</th>
              </tr>
            </thead>
            <tbody>
              </tbody>
          </table>
        </details>
      </td>
    </tr>
    </tbody>
</table>






## Channels



<a name="channel-another.one"></a>





#### Channel Parameters







###  `subscribe` another.one

#### Message



Sample consumer



This is another sample schema.



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
  <td>X-Trace-ID </td>
  <td>string</td>
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
  <td>object</td>
  <td></td>
  <td><em>Any</em></td>
</tr>





<tr>
  <td>item.key </td>
  <td>string</td>
  <td><p>Item key</p>
</td>
  <td><em>Any</em></td>
</tr>









<tr>
  <td>item.values </td>
  <td>array(integer)</td>
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








<a name="channel-one.{name}.two"></a>





#### Channel Parameters



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
  <td>string</td>
  <td><p>Name</p>
</td>
  <td><em>Any</em></td>
</tr>







    
  </tbody>
</table>





###  `publish` one.{name}.two

#### Message



Sample publisher



This is a sample schema.





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
  <td>string</td>
  <td><p>Creation time</p>
</td>
  <td><em>Any</em></td>
</tr>







    
      
<tr>
  <td>items </td>
  <td>array(object)</td>
  <td><p>List of items</p>
</td>
  <td><em>Any</em></td>
</tr>








<tr>
  <td>items.key </td>
  <td>string</td>
  <td><p>Item key</p>
</td>
  <td><em>Any</em></td>
</tr>









<tr>
  <td>items.values </td>
  <td>array(integer)</td>
  <td><p>List of item values</p>
</td>
  <td><em>Any</em></td>
</tr>












    
  </tbody>
</table>



###### Example of payload _(generated)_

```json
{
  "createdAt": "2019-10-22T12:15:29Z",
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










