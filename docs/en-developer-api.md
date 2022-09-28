# Developer API Manual
  * [Cluster Template API](#----api)
    + [1. Create cluster](#1-----)
    + [2. Get the list of clusters](#2-------)
    + [3. Create VPC](#3---vpc)
    + [4. View VPC](#4---vpc)
    + [5. Create subnets](#5-----)
    + [6. View subnet](#6-----)
    + [7. Create security groups](#7------)
    + [8. View Security Group](#8------)
    + [9. Create network configuration](#9-------)
    + [10. View the list of region](#10---region--)
    + [11. View zone list](#11---zone--)
    + [12. View the list of models](#12-------)
    + [13. Get the list of mirrors](#13-------)
  * [Scaling Up And Scaling Down Task API](#-----api)
    + [1. Create scale-up task](#1-------)
    + [2. Create scale-down task](#2-------)
    + [3. View task list](#3-------)
  * [Machine API](#--api)
    + [1. Machine list](#1-----)
    + [2. Machine details](#2-----)
    + [3. Get the number of machines](#3-------)
  * [Fees API](#--api)
    + [1. Total machine hours used for the day](#1----------)
    + [2. Breakdown of machine hours used for the day](#2-----------)

## Cluster Template API
### 1. Create cluster
Create cluster template according to the user's needs.<br>

**Request Address**
<table>
  <tr>
    <td>POST Method</td>
  </tr>
  <tr>
    <td>POST /api/v1/cluster/create </td>
  </tr>
</table>

**Request Parameters**
<table>
  <tr>
    <td>Name</td>
    <td>Type</td>
    <td>Required Field</td>
    <td>Description</td>
    <td>Sample Value</td>
  </tr>
  <tr>
    <td>name</td>
    <td>string</td>
    <td>Yes</td>
    <td>Cluster name</td>
    <td>test_cluster</td>
  </tr>
  <tr>
    <td>provider</td>
    <td>string</td>
    <td>Yes</td>
    <td>Cloud vendor</td>
    <td>AlibabaCloud | HuaweiCloud | TencentCloud | BaiduCloud | AWSCloud</td>
  </tr>
  <tr>
    <td>region_id</td>
    <td>string</td>
    <td>Yes</td>
    <td>Region</td>
    <td>cn-beijing</td>
  </tr>
  <tr>
    <td>zone_id</td>
    <td>string</td>
    <td>Yes</td>
    <td>Zone</td>
    <td>cn-beijing-h</td>
  </tr>
  <tr>
    <td>network_config</td>
    <td>object{}</td>
    <td>Yes</td>
    <td>Network configuration info</td>
    <td>{}</td>
  </tr>
  <tr>
    <td>instance_type</td>
    <td>string</td>
    <td>Yes</td>
    <td>Sample specificaitions</td>
    <td>ecs.s6-c1m1.small</td>
  </tr>
  <tr>
    <td>charge_config</td>
    <td>object{}</td>
    <td>Yes</td>
    <td>Payment info configuration</td>
    <td>{}</td>
  </tr>
  <tr>
    <td>image</td>
    <td>string</td>
    <td>Yes</td>
    <td>System image id</td>
    <td>m-2ze14bof6m3aadve22aq</td>
  </tr>
  <tr>
    <td>disks</td>
    <td>object{}</td>
    <td>Yes</td>
    <td>Storage configuration info </td>
    <td>{}</td>
  </tr>
  <tr>
    <td>password</td>
    <td>string</td>
    <td>Yes</td>
    <td>Default instance password</td>
    <td>ASDqwe123</td>
  </tr>
  <tr>
    <td>account_key</td>
    <td>string</td>
    <td>Yes</td>
    <td>Cloud account ak</td>
    <td>LTAI5tAwAMpXAQ78pePcRb6t</td>
  </tr>
  <tr>
    <td>desc</td>
    <td>string</td>
    <td>No</td>
    <td>Cluster description</td>
    <td>For responding to unexpected operations</td>
  </tr>
    <tr>
    <td>tags</td>
    <td>string</td>
    <td>No</td>
    <td>Cluster tags</td>
    <td>null</td>
</table>

**Content in "network_config"**
<table>
  <tr>
    <td>Name</td>
    <td>Type</td>
    <td>Required Field</td>
    <td>Description</td>
    <td>Sample Value</td>
  </tr>
  <tr>
    <td>vpc</td>
    <td>string</td>
    <td>Yes</td>
    <td>Virtual private network vpc's ID</td>
    <td>vpc-2zelmmlfd5c5duibc2xb2</td>
  </tr>
  <tr>
    <td>subnet_id</td>
    <td>string</td>
    <td>Yes</td>
    <td>Subnet ID</td>
    <td>vsw-2zennaxawzq6sa2fdj8l5</td>
  </tr>
  <tr>
    <td>security_group</td>
    <td>string</td>
    <td>Yes</td>
    <td>Security group ID</td>
    <td>sg-2zefbt9tw0yo1r7vc3ac</td>
  </tr>
  <tr>
    <td>internet_charge_type</td>
    <td>string</td>
    <td>Yes(mandatory when public bandwidth is required)</td>
    <td>Newwork billing type.Range of values：<br>
PayByBandwidth：fixed bandwidth billing<br>
PayByTraffic（default）：usage-based billing.</td>
    <td>PayByTraffic</td>
  </tr>
  <tr>
    <td>internet_charge_type</td>
    <td>string</td>
    <td>Yes(mandatory when public bandwidth is required)</td>
    <td>Maximum network bandwidth(M)</td>
    <td>10</td>
  </tr>
  
</table>

**Content in "charge_config"**
<table>
  <tr>
    <td>Name</td>
    <td>Type</td>
    <td>Required Field</td>
    <td>Description</td>
    <td>Sample Value</td>
  </tr>
  <tr>
    <td>charge_type</td>
    <td>string</td>
    <td>Yes</td>
    <td>Payment type：<br>
PostPaid: pay-as-you-go<br>
PrePaid: annual or monthly package</td>
    <td>PostPaid</td>
  </tr>
<tr>
    <td>period_unit</td>
    <td>string</td>
    <td>Required when charge_type is PrePaid</td>
    <td>The unit of length of purchased resources. the range of values:<br>Week: week<br> Month: month</td>
    <td>Month</td>
  </tr>
  <tr>
    <td>period</td>
    <td>int</td>
    <td>Required when charge_type is PrePaid</td>
    <td>The length of the purchased resource,the range of values:<br>when period_unit is Week: [1, 2, 3, 4]<br>when period_unit is Month: [1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36, 48, 60]</td>
    <td>1</td>
  </tr>
</table>

**Content in "disks"**
<table>
  <tr>
    <td>Name</td>
    <td>Type</td>
    <td>Required Field</td>
    <td>Description</td>
    <td>Sample Value</td>
  </tr>
  <tr>
    <td>system_disk</td>
    <td>object</td>
    <td>Yes</td>
    <td>System disk configuration</td>
    <td></td>
  </tr>
  <tr>
    <td>data_disk</td>
    <td>object</td>
    <td>Yes</td>
    <td>Data disk configuration</td>
    <td></td>
  </tr>
</table>

**Conetent in "system_disk" and "data_disk"**
<table>
  <tr>
    <td>Name</td>
    <td>Type</td>
    <td>Required Field</td>
    <td>Description</td>
    <td>Sample Value</td>
  </tr>
  <tr>
    <td>category</td>
    <td>string</td>
    <td>Yes</td>
    <td>System disk(data disk)</td>
    <td>cloud_efficiency</td>
  </tr>
  <tr>
    <td>size</td>
    <td>int</td>
    <td>Yes</td>
    <td>System disk（data disk）size (G)</td>
    <td>40</td>
  </tr>
</table>


**Return parameters**
<table>
  <tr>
    <td>Name</td>
    <td>Type</td>
    <td>Required Field</td>
    <td>Description</td>
    <td>Sample Value</td>
  </tr>
  <tr>
    <td>code</td>
    <td>int</td>
    <td>Yes</td>
    <td>Return code</td>
    <td>0</td>
  </tr>
  <tr>
    <td>msg</td>
    <td>string</td>
    <td>Yes</td>
    <td>Error message</td>
    <td>null</td>
  </tr>
  <tr>
    <td>data</td>
    <td>object</td>
    <td>Yes</td>
    <td>Normal message</td>
    <td> {}</td>
  </tr>
</table>

**Request Example**
```JSON
{
    "name":"test cluster",
    "provider":"AlibabaCloud",
    "account_key":"LTAI5t7qCv**Fh3hzSYpSv",
    "charge_type":"PostPaid",
    "region_id":"cn-qingdao",
    "zone_id":"cn-qingdao-b",
    "instance_type":"ecs.n1.tiny",
    "image":"centos_7_6_x64_20G_alibase_20211030.vhd",
    "password":"********",
    "desc":"cluster for testing",
    "network_config":{
        "vpc":"vpc-m5***6tgd",
        "subnet_id":"vsw-m5***4y3xs6ivwc",
        "security_group":"sg-m5***2cu2f5x8"
    },
    "disks":{
        "system_disk":{
            "category":"cloud_efficiency",
            "size":50,
        },
        "data_disk":[
            {
                "category":"cloud_efficiency",
                "size":50,
            }
        ]
    }
}
```

**Example response**

Normal return result：
```JSON
{
    "code":200,
    "msg":"success",
    "data":"test cluster"
}
```
Exception return result：
```JSON
{
    "code":400,
    "msg":"param_invalid",
    "data":null
}
```
**Return code explanation**
<table>
  <tr>
    <td>Return code</td>
    <td>Status</td>
    <td>Explanation</td>
  </tr>
  <tr>
    <td>200</td>
    <td>success</td>
    <td>Successful implementation</td>
  </tr>
  <tr>
    <td>400</td>
    <td>param_invalid</td>
    <td>Wrong parameters</td>
  </tr>
</table>



### 2. Get the list of clusters
Get information about all clusters under this account.<br>
**Request Address**
<table>
  <tr>
    <td>GET method</td>
  </tr>
  <tr>
    <td>GET /api/v1/cluster/describe_all </td>
  </tr>
</table>

**Request Parameters**
<table>
  <tr>
    <td>Name</td>
    <td>Type</td>
    <td>Required Field</td>
    <td>Description</td>
    <td>Sample Value</td>
  </tr>
  <tr>
    <td>account</td>
    <td>String</td>
    <td>No</td>
    <td>Cloud account(if not passed,all clusters under the organization will be queried by default)</td>
    <td>LTAI5tAWAM</td>
  </tr>
  <tr>
    <td>page_number</td>
    <td>int32</td>
    <td>No</td>
    <td>Default start page</td>
    <td>1</td>
  </tr>
  <tr>
    <td>page_size</td>
    <td>int32</td>
    <td>No</td>
    <td>Default 10 Maximum 50</td>
    <td>15</td>
  </tr>
</table>

**Return parameters**
<table>
  <tr>
    <td>Name</td>
    <td>Type</td>
    <td>Required Field</td>
    <td>Description</td>
    <td>Sample Value</td>

  </tr>
  <tr>
    <td>code</td>
    <td>int</td>
    <td>Yes</td>
    <td>Return code</td>
    <td>0</td>
  </tr>
  <tr>
    <td>msg</td>
    <td>string</td>
    <td>Yes</td>
    <td>Error message</td>
    <td>null</td>
  </tr>
  <tr>
    <td>data</td>
    <td>object</td>
    <td>Yes</td>
    <td>Normal message</td>
    <td> {}</td>
  </tr>
</table>

**Important parameter in "data"**
<table>
  <tr>
    <td>Name</td>
    <td>Sub-attributes</td>
    <td>Type</td>
    <td>Required field</td>
    <td>Description</td>
    <td>Sample Value</td>
  </tr>
  <tr>
    <td>cluster_list[]</td>
    <td></td>
    <td>Array</td>
    <td>Yes</td>
    <td>Cluster List</td>
    <td></td>
  </tr>
  <tr>
    <td></td>
    <td>cluster_name</td>
    <td>String</td>
    <td>Yes</td>
    <td>Cluster Name</td>
    <td></td>
  </tr>
  <tr>
    <td></td>
    <td>provider</td>
    <td>String</td>
    <td>Yes</td>
    <td>AlibabaCloud</td>
    <td></td>
  </tr>
  <tr>
    <td></td>
    <td>account</td>
    <td>String</td>
    <td>Yes</td>
    <td>Cloud Account</td>
    <td>aagjege</td>
  </tr>
  <tr>
    <td></td>
    <td>total_remainder</td>
    <td>String</td>
    <td>Yes</td>
    <td>Quota usage/residual</td>
    <td>8/200</td>
  </tr>
  <tr>
    <td></td>
    <td>tcreate_at</td>
    <td>String</td>
    <td>Yes</td>
    <td>Creation time</td>
    <td>2021-11-03 17:01:44</td>
  </tr>
  <tr>
    <td>pager</td>
    <td>Pager</td>
    <td>String</td>
    <td>Yes</td>
    <td>Paging parameters</td>
    <td></td>
  </tr>
</table>


**Request Example**
<table>
  <tr>
    <td>Name</td>
    <td>Sample Value</td>
  </tr>
  <tr>
    <td>account</td>
    <td>LTAI5tAWAM</td>
  </tr>
  <tr>
    <td>page_number</td>
    <td>1</td>
  </tr>
  <tr>
    <td>page_size</td>
    <td>15</td>
  </tr>
</table>


**Example response**

Normal return result：<br>
```JSON
{
  "code": 200,
  "data": {
    "cluster_list": [
      {
        "cluster_id": "1319",
        "cluster_name": "gf.bridgx.online",
        "provider": "AlibabaCloud",
        "account": "LTAI5tAwAMpXAQ78pePcRb6t",
        "create_at": "2021-11-02 06:09:38 +0800 CST",
        "create_by": ""
      },
      {
        "cluster_id": "1332",
        "cluster_name": "test6",
        "provider": "AlibabaCloud",
        "account": "LTAI5tAwAMpXAQ78pePcRb6t",
        "create_at": "2021-11-04 14:22:15 +0800 CST",
        "create_by": ""
      }
    ],
    "pager": {
      "page_number": 1,
      "page_size": 10,
      "total": 5
    }
  },
  "msg": "success"
}
```

Exception return result：

```JSON
{
    "code":400,
    "msg":"param_invalid",
    "data":null
}
```
**Return code explanation**

<table>
  <tr>
    <td>Return code</td>
    <td>Status</td>
    <td>Explanation</td>
  </tr>
  <tr>
    <td>200</td>
    <td>success</td>
    <td>Successful implementation</td>
  </tr>
</table>


### 3. Create VPC
A VPC can only specify one network segment, the range of the segment includes 10.0.0.0/8, 172.16.0.0/12 and 192.168.0.0/16 and their subnets, the mask of the segment is 8~24 bits, and the default is 172.16.0.0/12. <br>
The network segment cannot be modified after the VPC is created. <br>
The number of private network addresses per VPC to support cloud resource usage is 60,000, and the quota cannot be upgraded. >br?
Once a VPC is created, a router and a routing table will be automatically created. <br>
Each VPC supports three user-side network segments. If there is an inclusion relationship between multiple user-side segments, the segment with the shorter mask will actually take effect. For example, for 10.0.0.0/8 and 10.1.0.0/16, 10.0.0.0/8 will actually take effect.<br>

**Request Address**
<table>
  <tr>
    <td>POST Method</td>
  </tr>
  <tr>
    <td>POST /api/v1/vpc/create</td>
  </tr>
</table>

**Request Parameters**
<table>
  <tr>
    <td>Name</td>
    <td>Type</td>
    <td>Required Field</td>
    <td>Description</td>
    <td>Sample Value</td>
  </tr>
  <tr>
    <td>provider</td>
    <td>String</td>
    <td>Yes</td>
    <td>Cloud vendor</td>
    <td>AlibabaCloud</td>
  </tr>
  <tr>
    <td>region_id</td>
    <td>String</td>
    <td>Yes</td>
    <td>Target Region ID</td>
    <td>cn-hangzhou</td>
  </tr>
  <tr>
    <td>cidr_block</td>
    <td>String</td>
    <td>否</td>
    <td>VPC's network segment.You can use the following network segment or its subnet for the VPC：<br>
172.16.0.0/12（default）<br>
10.0.0.0/8<br>
192.168.0.0/16。</td>
    <td>172.16.0.0/12</td>
  </tr>  
  <tr>
    <td>vpc_name</td>
    <td>String</td>
    <td>Yes</td>
    <td>VPC name: length 2 to 64 characters,must start with a letter or a Chinese character,<br>
and can contain numbers,semi-colon periods（.）、underscore（_）and dash（-），<br>
but cannot start with http:// 或https://.</td>
    <td>abc</td>
  </tr>
  <tr>
    <td>ak</td>
    <td>String</td>
    <td>Yes</td>
    <td>AK under cloud vendor account</td>
    <td>dasdadasdasd</td>
  </tr>
</table>

**Return parameters**
<table>
  <tr>
    <td>Name</td>
    <td>Type</td>
    <td>Required Field</td>
    <td>Description</td>
    <td>Sample Value</td>
  </tr>
  <tr>
    <td>code</td>
    <td>int</td>
    <td>Yes</td>
    <td>Return code</td>
    <td>0</td>
  </tr>
  <tr>
    <td>msg</td>
    <td>string</td>
    <td>Yes</td>
    <td>Error message</td>
    <td>null</td>
  </tr>
  <tr>
    <td>data</td>
    <td>string</td>
    <td>Yes</td>
    <td>The successfully created "vpcid"</td>
    <td>vpc-asd**asdasda</td>
  </tr>
</table>

**Request Example**
<table>
  <tr>
    <td>Name</td>
    <td>Sample Value</td>
  </tr>
  <tr>
    <td>provider</td>
    <td>AlibabaCloud</td>
  </tr>
  <tr>
    <td>region_id</td>
    <td>cn-hangzhou</td>
  </tr>
  <tr>
    <td>cidr_block</td>
    <td>172.16.0.0/12</td>
  </tr>
  <tr>
    <td>vpc_name</td>
    <td>abc</td>
  </tr>
  <tr>
    <td>ak</td>
    <td>dasdadasdasd</td>
  </tr>
</table>

**Example response**

Normal return result：
```JSON
{
    "code":200,
    "msg":"success",
    "data":"vsw-m5evh*****s6ivwc"
}
```

Exception return result：
```JSON
{
    "code":400,
    "msg":"param_invalid",
    "data":null
}
```
**Return code explanation**

<table>
  <tr>
    <td>Return code</td>
    <td>Status</td>
    <td>Explanation</td>
  </tr>
  <tr>
    <td>200</td>
    <td>success</td>
    <td>Successful implementation</td>
  </tr>
  <tr>
    <td>400</td>
    <td>param_invalid</td>
    <td>Wrong parameter</td>
  </tr>
</table>


### 4. View VPC
Find information on the VPCs that the user has created based on the name of the VPC.

**Request Address**
<table>
  <tr>
    <td>GET method</td>
  </tr>
  <tr>
    <td>GET /api/v1/vpc/describe</td>
  </tr>
</table>

**Request Parameters**
<table>
  <tr>
    <td>Name</td>
    <td>Type</td>
    <td>Required Field</td>
    <td>Description</td>
    <td>Sample Value</td>
  </tr>
  <tr>
    <td>provider</td>
    <td>String</td>
    <td>No</td>
    <td>Cloud vendor</td>
    <td>AlibabaCloud</td>
  </tr>
  <tr>
    <td>region_id</td>
    <td>String</td>
    <td>No</td>
    <td>Region ID where the VPC is located</td>
    <td>cn-hangzhou</td>
  </tr>
  <tr>
    <td>vpc_name</td>
    <td>String</td>
    <td>No</td>
    <td>Name of VPC</td>
    <td>vpc-1</td>
  </tr>
  <tr>
    <td>page_number</td>
    <td>int32</td>
    <td>No</td>
    <td>Default start page</td>
    <td>1</td>
  </tr>
  <tr>
    <td>page_size</td>
    <td>int32</td>
    <td>No</td>
    <td>Default 10 Maximun 50</td>
    <td>15</td>
  </tr>
</table>


**Return parameters**
<table>
  <tr>
    <td>Name</td>
    <td>Type</td>
    <td>Required Field</td>
    <td>Description</td>
    <td>Sample Value</td>
  </tr>
  <tr>
    <td>code</td>
    <td>int</td>
    <td>Yes</td>
    <td>Return code</td>
    <td>0</td>
  </tr>
  <tr>
    <td>msg</td>
    <td>string</td>
    <td>Yes</td>
    <td>Error message</td>
    <td>null</td>
  </tr>
  <tr>
    <td>data</td>
    <td>object</td>
    <td>Yes</td>
    <td>Normal message</td>
    <td> {}</td>
  </tr>
</table>

**Important parameter in "data"**
<table>
  <tr>
    <td>Name</td>
    <td>Type</td>
    <td>Required Field</td>
    <td>Description</td>
    <td>Sample Value</td>
  </tr>
  <tr>
    <td>create_at</td>
    <td>String</td>
    <td>Yes</td>
    <td>VPC creation time</td>
    <td>2018-04-18T15:02:37Z</td>
  </tr>
  <tr>
    <td>status</td>
    <td>String</td>
    <td>Yes</td>
    <td>The status of the VPC，range of values：<br>
Pending：within the configuration。<br>
Available：available。</td>
    <td>Available</td>
  </tr>
  <tr>
    <td>vpc_id</td>
    <td>String</td>
    <td>Yes</td>
    <td>VPC's ID</td>
    <td>vpc-bp1qpo0kug3a20qqe****</td>
  </tr>
  <tr>
    <td>cidr_block</td>
    <td>String</td>
    <td>Yes</td>
    <td>VPC's IPv4 network segment</td>
    <td>192.168.0.0/16</td>
  </tr>
  <tr>
    <td>switch_ids</td>
    <td>String []</td>
    <td>Yes</td>
    <td>List of switches in VPC</td>
    <td>vsw-bp1nhbnpv2blyz8dl****</td>
  </tr>
  <tr>
    <td>provider</td>
    <td>String []</td>
    <td>Yes</td>
    <td>Affiliated Cloud Vendor</td>
    <td>AlibabaCloud</td>
  </tr>
  <tr>
    <td>vpc_name</td>
    <td>String []</td>
    <td>No</td>
    <td>VPC name</td>
    <td>vpc-1</td>
  </tr>
</table>

**Request Example**
<table>
  <tr>
    <td>Name</td>
    <td>Sample Value</td>
  </tr>
  <tr>
    <td>provider</td>
    <td>AlibabaCloud</td>
  </tr>
  <tr>
    <td>region_id</td>
    <td>cn-qingdao</td>
  </tr>
  <tr>
    <td>page_number</td>
    <td>1</td>
  </tr>
  <tr>
    <td>page_size</td>
    <td>50</td>
  </tr>
</table>


**Example response**

Normal return result：
```JSON
{
    "code":200,
    "data":{
        "Vpcs":[
            {
                "VpcId":"vpc-m5**swmv796tgd",
                "VpcName":"vpctest1",
                "CidrBlock":"",
                "SwitchIds":"",
                "Provider":"AlibabaCloud",
                "Status":"",
                "CreateAt":"2021-11-11 11:13:34 +0800 CST"
            },
            {
                "VpcId":"vpc-m5e50cwcefjgxvbjs1ud5",
                "VpcName":"vpctest2",
                "CidrBlock":"",
                "SwitchIds":"",
                "Provider":"AlibabaCloud",
                "Status":"",
                "CreateAt":"2021-11-05 11:40:16 +0800 CST"
            }
        ],
        "Pager":{
            "PageNumber":1,
            "PageSize":50,
            "Total":2
        }
    },
    "msg":"success"
}
```

Exception return result：
```JSON
{
    "code":400,
    "msg":"param_invalid",
    "data":null
}
```
**Return code explanation**
<table>
  <tr>
    <td>Return code </td>
    <td>Status</td>
    <td>Explanation</td>
  </tr>
  <tr>
    <td>200</td>
    <td>success</td>
    <td>Successful implementation</td>
  </tr>
  <tr>
    <td>400</td>
    <td>param_invalid</td>
    <td>Wrong parameters</td>
  </tr>
</table>



### 5. Create subnets
- The number of switches in each VPC cannot exceed 150.
- The 1st and the last 3 IP addresses of each switch segment are reserved for the system. For example, the system reserved addresses for 192.168.1.0/24 are 192.168.1.0, 192.168.1.253, 192.168.1.254 and 192.168.1.255.
- The number of cloud product instances under the switch is not allowed to exceed the number of remaining available cloud product instances for the VPC (15000 minus the current number of cloud product instances).
- A cloud product instance can only belong to one switch.
- The switch does not support multicast and broadcast.
- After the switch is created successfully, the network segment cannot be modified.


**Request Address**
<table>
  <tr>
    <td>GET method</td>
  </tr>
  <tr>
    <td>GET /api/v1/subnet/create</td>
  </tr>
</table>

**Request Address**
<table>
  <tr>
    <td>Name</td>
    <td>Type</td>
    <td>Required Field</td>
    <td>Description</td>
    <td>Sample Value</td>
  </tr>
  <tr>
    <td>zone_id</td>
    <td>String</td>
    <td>Yes</td>
    <td>Available Zone ID</td>
    <td>cn-hangzhou-g</td>
  </tr>
  <tr>
    <td>cidr_block</td>
    <td>String</td>
    <td>Yes</td>
    <td>The network segment of the swith.The switch segment requirements are as follows：<br>
The mask length range of the switch's segment is 16～29 bits。<br>
The network segment of the switch must belong to the network segment of the VPC.<br>
The network segment of the switch cannot be the same as the target network segment of <br>
the routing entry in the VPC where it is located.but it can be a subnet of the target network segment.<br>
    <td>172.16.0.0/24</td>
  </tr>
  <tr>
    <td>vpc_id</td>
    <td>String</td>
    <td>Yes</td>
    <td>The VPC ID of the switch to be created</td>
    <td>vpc-257gqcdfvx6n****</td>
  </tr>
  <tr>
    <td>switch_name</td>
    <td>String</td>
    <td>Yes</td>
    <td>THe name of the switch,with a length of 2 to 128 characters, <br>
      must start with a letter or a Chinese character,but not with http:// or https://.<br>
    <td>VSwitch-1</td>
  </tr>
</table>


**Return parameters**
<table>
  <tr>
    <td>Name</td>
    <td>Type</td>
    <td>Required Field</td>
    <td>Description</td>
    <td>Sample Value</td>
  </tr>
  <tr>
    <td>code</td>
    <td>int</td>
    <td>Yes</td>
    <td>Return code</td>
    <td>0</td>
  </tr>
  <tr>
    <td>msg</td>
    <td>string</td>
    <td>Yes</td>
    <td>Error message</td>
    <td>null</td>
  </tr>
  <tr>
    <td>data</td>
    <td>string</td>
    <td>Yes</td>
    <td>Switch ID created successfully </td>
    <td> "asdasdasdas"</td>
  </tr>
</table>


**Request Example**
<table>
  <tr>
    <td>Name</td>
    <td>Sample Value</td>
  </tr>
  <tr>
    <td>zone_id</td>
    <td>cn-qingdao-b</td>
  </tr>
  <tr>
    <td>cidr_block</td>
    <td>172.16.0.0/24</td>
  </tr>
  <tr>
    <td>vpc_id</td>
    <td>vpc-257gqcdfvx6n****</td>
  </tr>
  <tr>
    <td>switch_name</td>
    <td>VSwitch-1</td>
  </tr>
</table>

**Example response**

Normal return result：
```JSON
{
    "code":200,
    "msg":"success",
    "data":"vsw-m5evh***4y3xs6ivwc"
}
```

Exception return result    
```JSON
{
    "code":400,
    "msg":"param_invalid",
    "data":null
}
```
**Return code explanation**

<table>
  <tr>
    <td>Return code</td>
    <td>Status</td>
    <td>Explanation</td>
  </tr>
  <tr>
    <td>200</td>
    <td>success</td>
    <td>Successful implementation</td>
  </tr>
  <tr>
    <td>400</td>
    <td>param_invalid</td>
    <td>Wrong parameters</td>
  </tr>
</table>


### 6. View subnet
Get information on the subnet ID under this account<br>
**Request Address**
<table>
  <tr>
    <td>GET method</td>
  </tr>
  <tr>
    <td>GET /api/v1/subnet/describe</td>
  </tr>
</table>

**Request Parameters**
<table>
  <tr>
    <td>Name</td>
    <td>Type</td>
    <td>Required Field</td>
    <td>Description</td>
    <td>Sample Value</td>
  </tr>
  <tr>
    <td>vpc_id</td>
    <td>String</td>
    <td>Yes</td>
    <td>ID of the VPC to which the switch to be queried belongs</td>
    <td>vpc-257gqcdfvx6n****</td>
  </tr>
  <tr>
    <td>switch_name</td>
    <td>String</td>
    <td>No</td>
    <td>The name of the switch to query.(If not passed, all switcher under the VPC will be queried by default)</td>
    <td>VSwitch-1</td>
  </tr>
  <tr>
    <td>page_number</td>
    <td>int32</td>
    <td>No</td>
    <td>Default start page</td>
    <td>1</td>
  </tr>
  <tr>
    <td>page_size</td>
    <td>int32</td>
    <td>No</td>
    <td>Default 10 Maximum 50</td>
    <td>15</td>
  </tr>
</table>

**Return parameters**
<table>
  <tr>
    <td>Name</td>
    <td>Type</td>
    <td>Required Field</td>
    <td>Description</td>
    <td>Sample Value</td>
  </tr>
  <tr>
    <td>code</td>
    <td>int</td>
    <td>Yes</td>
    <td>Return code</td>
    <td>0</td>
  </tr>
  <tr>
    <td>msg</td>
    <td>string</td>
    <td>Yes</td>
    <td>Error message</td>
    <td>null</td>
  </tr>
  <tr>
    <td>data</td>
    <td>object</td>
    <td>Yes</td>
    <td>Normal message</td>
    <td> {}</td>
  </tr>
</table>

**Important parameter in "data"**
<table>
  <tr>
    <td>Name</td>
    <td>Type</td>
    <td>Required Field</td>
    <td>Description</td>
    <td>Sample Value</td>
  </tr>
  <tr>
    <td>vpc_id</td>
    <td>String</td>
    <td>Yes</td>
    <td>ID of the VPC to which the switch belongs</td>
    <td>vpc-257gqcdfvx6n****</td>
  </tr>
  <tr>
    <td>switch_name</td>
    <td>String</td>
    <td>Yes</td>
    <td>Name of the switch to be queried</td>
    <td>VSwitch-1</td>
  </tr>
  <tr>
    <td>status</td>
    <td>String</td>
    <td>Yes</td>
    <td>The status of the switch,range of values:<br>
Pending：in configuration<br>
Available：available。</td>
    <td>Available</td>
  </tr>
  <tr>
    <td>create_at</td>
    <td>String</td>
    <td>Yes</td>
    <td>Switch creation time</td>
    <td>2018-01-18T12:43:57Z</td>
  </tr>
  <tr>
    <td>is_default</td>
    <td>Boolean</td>
    <td>Yes</td>
    <td>Whether it is the default switch.<br>
true：is the default switch.<br>
false：Non-default switch.</td>
    <td>true</td>
  </tr>
  <tr>
    <td>available_ip_address_count</td>
    <td>Long</td>
    <td>Yes</td>
    <td>ID of the switch </td>
    <td>1</td>
  </tr>
  <tr>
    <td>switch_id</td>
    <td>String</td>
    <td>Yes</td>
    <td>ID of the switch</td>
    <td>vsw-25bcdxs7pv1****</td>
  </tr>
  <tr>
    <td>cidr_block</td>
    <td>String</td>
    <td>Yes</td>
    <td>IPv4 network segment of the switch</td>
    <td>172.16.0.0/24</td>
  </tr>
  <tr>
    <td>zone_id</td>
    <td>String</td>
    <td>Yes</td>
    <td>Available ID</td>
    <td>cn-hangzhou-g</td>
  </tr>
</table>

**Request Example**
<table>
  <tr>
    <td>Name</td>
    <td>Sample Value</td>
  </tr>
  <tr>
    <td>vpc_id</td>
    <td>vpc-257gqcdfvx6n****</td>
  </tr>
  <tr>
    <td>switch_name</td>
    <td>"The first switch"</td>
  </tr>
  <tr>
    <td>page_number</td>
    <td>1</td>
  </tr>
  <tr>
    <td>page_size</td>
    <td>15</td>
  </tr>
</table>

**Example response**

Normal return result：
```JSON
{
    "code":200,
    "data":{
        "Switches":[
            {
                "VpcId":"vpc-257gqcdfvx6n****",
                "SwitchId":"vsw-m5evh***4y3xs6ivwc",
                "ZoneId":"cn-qingdao-b",
                "SwitchName":"The first switch",
                "CidrBlock":"172.16.0.0/24",
                "VStatus":"",
                "CreateAt":"2021-11-02 17:50:37 +0800 CST",
                "IsDefault":"N",
                "AvailableIpAddressCount":0
            }
        ],
        "Pager":{
            "PageNumber":1,
            "PageSize":50,
            "Total":1
        }
    },
    "msg":"success"
}
```

Exception return result：
```JSON
{
    "code":400,
    "msg":"param_invalid",
    "data":null
}
```
**Return code explanation**

<table>
  <tr>
    <td>Return code</td>
    <td>Status</td>
    <td>Explanation</td>
  </tr>
  <tr>
    <td>200</td>
    <td>success</td>
    <td>Successful implementation</td>
  </tr>
  <tr>
    <td>400</td>
    <td>param_invalid</td>
    <td>Wrong parameters</td>
  </tr>
</table>


### 7. Create security groups
- In the API documentation of the security group, the originating end of the traffic is the Source and the receiving end of the data transmission is the Dest.
- The total number of safety group rules in the outbound and inbound directions cannot exceed 200.
- You can select a value from 1 to 100 for the security group rule priority. The smaller the number, the higher the priority.
- Security group rules with the same priority take precedence over rules that deny access (drop).
- The source device can be a specified IP address range (SourceCidrIp, Ipv6SourceCidrIp, SourcePrefixListId) or an ECS instance in another security group (SourceGroupId).
- If the matching security group rule already exists, this AuthorizeSecurityGroup call succeeds, but does not increase the number of rules.


**Request Address**
<table>
  <tr>
    <td>POST Method</td>
  </tr>
  <tr>
    <td>POST /api/v1/security_group/create_with_rule </td>
  </tr>
</table>

**Request Parameters**
<table>
  <tr>
    <td>Name</td>
    <td>Type</td>
    <td>Required Field</td>
    <td>Description</td>
    <td>Sample Value</td>
  </tr>
  <tr>
    <td>vpc_id</td>
    <td>String</td>
    <td>Yes</td>
    <td>VPC ID of the security group</td>
    <td>vpc-bp1opxu1zkhn00gzv****</td>
  </tr>
  <tr>
    <td>security_group_name</td>
    <td>String</td>
    <td>Yes</td>
    <td>The name of the security group.The length should be 2 to 128 English or Chinese characters.<br>
      Must start with a upper or lower case letter or a Chinese character,but not with http:// and https://.<br>
      It can contain numbers,semi-colons(:),underscores(_),or hypthens(-).Default value:null.  </td>
    <td>testSecurityGroupName</td>
  </tr>
  <tr>
    <td>rules</td>
    <td>[]Rule</td>
    <td>Yes</td>
    <td>Rules</td>
    <td></td>
  </tr>
  <tr>
    <td>rules</td>
    <td>[]Rule</td>
    <td>Yes</td>
    <td>Rules</td>
    <td></td>
  </tr>
</table>



**Rule**
<table>
  <tr>
    <td>Name</td>
    <td>Type</td>
    <td>Required Field</td>
    <td>Description</td>
    <td>Sample Value</td>
  </tr>
  <tr>
    <td>protocol</td>
    <td>String</td>
    <td>Yes</td>
    <td>Transport layer protocols.The value is case-sensitive.Range of values：<br>
    tcp<br>
    udp<br>
    icmp<br>
    gre<br>
    all：support</td>
    <td>tcp</td>
  </tr>
  <tr>
    <td>port_range</td>
    <td>String</td>
    <td>Yes</td>
    <td>The range of ports associated with the transport layer protocol opened by the security group at the destination.Value range：<br>
TCP/UDP：the range of values is 1~65535.Use a slash(/) to separate the start port from the end port. For example:1/200.<br>
ICMP：-1/-1<br>
GRE：-1/-1<br>
"IpProtocol" is "all"：-1/-1</td>
    <td>22/22</td>
  </tr>
  <tr>
    <td>direction</td>
    <td>String</td>
    <td>Yes</td>
    <td>The direction of the security group rule:the range of values:<br>
egress：outgoing direction<br>
ingress：incoming direction</td>
    <td>ingress</td>
  </tr>
  <tr>
    <td>group_id</td>
    <td>String</td>
    <td>Choose one of two</td>
    <td>The (incoming or outgoing) security group ID for which access rights need to be set.</td>
    <td>sg-bp67acfmxazb4p****</td>
  </tr>
  <tr>
    <td>cidr_ip</td>
    <td>String</td>
    <td>Choose one of two</td>
    <td>Block of (incoming or outgoing) IPv4 CIDR address requiring access privileges.</td>
    <td>10.0.0.0/8</td>
  </tr>
</table>

**Return parameters**
<table>
  <tr>
    <td>Name</td>
    <td>Type</td>
    <td>Required Field</td>
    <td>Description</td>
    <td>Sample Value</td>
  </tr>
  <tr>
    <td>code</td>
    <td>int</td>
    <td>Yes</td>
    <td>Return code</td>
    <td>0</td>
  </tr>
  <tr>
    <td>msg</td>
    <td>string</td>
    <td>Yes</td>
    <td>Error message</td>
    <td>null</td>
  </tr>
  <tr>
    <td>data</td>
    <td>string</td>
    <td>Yes</td>
    <td>Create successful security group ID</td>
    <td> "asdadasdasdad"</td>
  </tr>
</table>

**Request Example**
<table>
  <tr>
    <td>Name</td>
    <td>Sample Value</td>
  </tr>
  <tr>
    <td>vpc_id</td>
    <td>vpc-bp1opxu1zkhn00gzv****</td>
  </tr>
  <tr>
    <td>security_group_name</td>
    <td>testSecurityGroupName</td>
  </tr>
  <tr>
    <td>page_size</td>
    <td>15</td>
  </tr>
  <tr>
    <td>rules</td>
    <td>
      [{
"protocol":"tcp",<br>
"port_range":"22/22",<br>
"direction":"ingress",<br>
"group_id":"sg-bp67acfmxazb4p****",<br>
"cidr_ip":"10.0.0.0/8"
    }]
    </td>
  </tr>
</table>



**Example response**

Normal return result：
```JSON
{
    "code":200,
    "msg":"success",
    "data":"vsw-m5evh***4y3xs6ivwc"
}
```
Exception return result：
```JSON
{
    "code":400,
    "msg":"param_invalid",
    "data":null
}
```
**Return code explanation**

<table>
  <tr>
    <td>Return code</td>
    <td>Status</td>
    <td>Explanation</td>
  </tr>
  <tr>
    <td>200</td>
    <td>success</td>
    <td>Suceessful implemention</td>
  </tr>
  <tr>
    <td>400</td>
    <td>param_invalid</td>
    <td>Wrong parameters</td>
  </tr>
</table>



### 8. View Security Group
View the security groups that have been created.<br>

**Request Address**
<table>
  <tr>
    <td>GET method </td>
  </tr>
  <tr>
    <td>GET /api/v1/security_group/describe </td>
  </tr>
</table>

**Request Parameters**
<table>
  <tr>
    <td>Name</td>
    <td>Type</td>
    <td>Required Field</td>
    <td>Description</td>
    <td>Sample Value</td>
  </tr>
  <tr>
    <td>region_id</td>
    <td>String</td>
    <td>No</td>
    <td>Region ID of the security group </td>
    <td>cn-hangzhou</td>
  </tr>
  <tr>
    <td>vpc_id</td>
    <td>String</td>
    <td>Yes</td>
    <td>VPC ID of the security group</td>
    <td>vpc-bp1opxu1zkhn00gzv****</td>
  </tr>
  <tr>
    <td>security_group_name</td>
    <td>String</td>
    <td>No</td>
    <td>The name of the security group.(If not passed,all under the VPC will be queried by default)</td>
    <td>testSecurityGroupName</td>
  </tr>
  <tr>
    <td>page_number</td>
    <td>int32</td>
    <td>No</td>
    <td>Default start page</td>
    <td>t1</td>
  </tr>
  <tr>
    <td>page_size</td>
    <td>int32</td>
    <td>No</td>
    <td>Default 10 Maximum 50</td>
    <td>15</td>
  </tr>
</table>

**Return parameters**
<table>
  <tr>
    <td>Name</td>
    <td>Type</td>
    <td>Required Field</td>
    <td>Description</td>
    <td>Sample Value</td>
  </tr>
  <tr>
    <td>code</td>
    <td>int</td>
    <td>Yes</td>
    <td>Return code</td>
    <td>0</td>
  </tr>
  <tr>
    <td>msg</td>
    <td>string</td>
    <td>Yes</td>
    <td>Error message</td>
    <td>null</td>
  </tr>
  <tr>
    <td>data</td>
    <td>object</td>
    <td>Yes</td>
    <td>Normal message</td>
    <td> {}</td>
  </tr>
</table>

**Important parameter in "data"**
<table>
  <tr>
    <td>Name</td>
    <td>Type</td>
    <td>Required Field</td>
    <td>Description</td>
    <td>Sample Value</td>
  </tr>
  <tr>
    <td>create_at</td>
    <td>String</td>
    <td>Yes</td>
    <td>Creation time.Expressed according to the ISO 8601 standard and requires the use of UTC time.<br>
      The format is ：yyyy-MM-ddThh:mmZ。</td>
    <td>2021-08-31T03:12:29Z</td>
  </tr>
  <tr>
    <td>vpc_id</td>
    <td>String</td>
    <td>Yes</td>
    <td>VPC to which the security group </td>
    <td>vpc-bp67acfmxazb4p****</td>
  </tr>
  <tr>
    <td>security_group_id</td>
    <td>String</td>
    <td>Yes</td>
    <td>Security group ID</td>
    <td>sg-bp67acfmxazb4p****</td>
  </tr>
  <tr>
    <td>security_group_name</td>
    <td>String</td>
    <td>Yes</td>
    <td>Security group ID</td>
    <td>SGTestName</td>
  </tr>
  <tr>
    <td>security_group_type</td>
    <td>String</td>
    <td>Yes</td>
    <td>Security Group Type.Possible values:<br>
normal：normal security group enterprise:<br>
enterprise：enterprise security group</td>
    <td>normal</td>
  </tr>
</table>

**Request Example**
<table>
  <tr>
    <td>Name</td>
    <td>Sample Value</td>
  </tr>
  <tr>
    <td>region_id</td>
    <td>cn-hangzhou</td>
  </tr>
  <tr>
    <td>vpc_id</td>
    <td>vpc-bp1opxu1zkhn00gzv****</td>
  </tr>
  <tr>
    <td>security_group_name</td>
    <td></td>
  </tr>
  <tr>
    <td>spage_number</td>
    <td>1</td>
  </tr>
  <tr>
    <td>page_size</td>
    <td>15</td>
  </tr>
</table>


**Example response**

Normal return result：
```JSON
{
    "code":200,
    "data":{
        "Groups":[
            {
                "VpcId":"vpc-bp1opxu1zkhn00gzv****",
                "SecurityGroupId":"sg-m5ebc***v2cu2f5x8",
                "SecurityGroupName":"The first security group tested",
                "SecurityGroupType":"normal",
                "CreateAt":"2021-11-02 18:03:01 +0800 CST"
            }
        ],
        "Pager":{
            "PageNumber":1,
            "PageSize":15,
            "Total":1
        }
    },
    "msg":"success"
}
```

Exception return result：
```JSON
{
    "code":400,
    "msg":"param_invalid",
    "data":null
}
```

**Return code explanation**

<table>
  <tr>
    <td>Return code</td>
    <td>Status</td>
    <td>Explanation</td>
  </tr>
  <tr>
    <td>200</td>
    <td>success</td>
    <td>Successful implementation</td>
  </tr>
  <tr>
    <td>400</td>
    <td>param_invalid</td>
    <td>Wrong parameters</td>
  </tr>
</table>




### 9. Create network configuration
Create VPC,subnet and security groups with a single API.<br>
**Request Address**
<table>
  <tr>
    <td>POST Method</td>
  </tr>
  <tr>
    <td>POST /api/v1/network_config/create</td>
  </tr>
</table>

**Request Parameters**
<table>
  <tr>
    <td>Name</td>
    <td>Type</td>
    <td>Required Field</td>
    <td>Description</td>
    <td>Sample Value</td>
  </tr>
  <tr>
    <td>provider</td>
    <td>String</td>
    <td>Yes</td>
    <td>Cloud vendor</td>
    <td>AlibabaCloud</td>
  </tr>
  <tr>
    <td>region_id</td>
    <td>String</td>
    <td>Yes</td>
    <td>Target Region ID</td>
    <td>cn-qingdao</td>
  </tr>
  <tr>
    <td>cidr_block</td>
    <td>String</td>
    <td>Yes</td>
    <td>VPC's network segment.You can use the following network segment or subset of the network segment of the VPC：<br>
172.16.0.0/12（default）。<br>
10.0.0.0/8。<br>
192.168.0.0/16。</td>
    <td>172.16.0.0/12</td>
  </tr>
  <tr>
    <td>vpc_name</td>
    <td>String</td>
    <td>Yes</td>
    <td>VPC name:length is 2 to 128 characters，must start with a letter or a Chinese character,<br>
      and can contain numbers, semicolon period(.),underscore(_) and dash(-),
      but cannot start with http:// or https://<br>
    </td>
    <td>abc</td>
  </tr>
  <tr>
    <td>zone_id</td>
    <td>String</td>
    <td>Yes</td>
    <td>Available Zone ID </td>
    <td>cn-qigndao-b</td>
  </tr>
  <tr>
    <td>switch_cidr_block</td>
    <td>String</td>
    <td>Yes</td>
    <td>The network segment of the switch.The switch segment requirements are as follows:：<br>
       The mask length range of the switch's segment is 16 to 29 bits. The network segment of <br>
      the switch must belong to the network segment of the VPC.The network segment of the switch 
      cannot be the same as the target segment of the routing enty in the VPC where it is located,<br>
      but it can be a subset of the target segment.
    </td>
    <td>172.16.0.0/24</td>
  </tr>
  <tr>
    <td>switch_name</td>
    <td>String</td>
    <td>Yes</td>
    <td>The name of the switch:The length is 2 to 128 characters and must start <br>
      with a letter or a Chinese character,but not with http:// or https://. <br>
    </td>
    <td>VSwitch-1</td>
  </tr>
  <tr>
    <td>security_group_name</td>
    <td>String</td>
    <td>Yes</td>
    <td>
      The name of the security group. The length should be 2 to 128 English or Chinese characters.<br>
       Must start with a upper or lower case letter or a Chinese character,but not with http:// and https://.<br>
       It can contain numbers ,semi-colons(:),underscors(_),or hyphens(-).Default value:null.
    </td>
    <td>testSecurityGroupName</td>
  </tr>
  <tr>
    <td>security_group_type</td>
    <td>String</td>
    <td>Yes</td>
    <td>Security group type.Possible values：<br>
normal：normal<br>
enterprise：enterprise security group</td>
    <td>normal</td>
  </tr>
  <tr>
    <td>ak</td>
    <td>String</td>
    <td>Yes</td>
    <td></td>
    <td></td>
  </tr>
</table>


**Return parameters**
<table>
  <tr>
    <td>Name</td>
    <td>Type</td>
    <td>Required Field</td>
    <td>Description</td>
    <td>Sample Value</td>
  </tr>
  <tr>
    <td>code</td>
    <td>int</td>
    <td>Yes</td>
    <td>Return code</td>
    <td>0</td>
  </tr>
  <tr>
    <td>msg</td>
    <td>string</td>
    <td>Yes</td>
    <td>Error message</td>
    <td>null</td>
  </tr>
  <tr>
    <td>data</td>
    <td>string</td>
    <td>Yes</td>
    <td>Create a successfull vpc_id</td>
    <td>sdasd1qwdasdasd</td>
  </tr>
</table>

**Request Example**
<table>
  <tr>
    <td>Name</td>
    <td>Sample Value</td>
  </tr>
  <tr>
    <td>provider</td>
    <td>AlibabaCloud</td>
  </tr>
  <tr>
    <td>region_id</td>
    <td>cn-qingdao</td>
  </tr>
  <tr>
    <td>cidr_block</td>
    <td>172.16.0.0/12</td>
  </tr>
  <tr>
    <td>vpc_name</td>
    <td>abc</td>
  </tr>
  <tr>
    <td>zone_id</td>
    <td>cn-qigndao-b</td>
  </tr>
  <tr>
    <td>switch_cidr_block</td>
    <td>172.16.0.0/24</td>
  </tr>
  <tr>
    <td>switch_name</td>
    <td>VSwitch-1</td>
  </tr>
  <tr>
    <td>security_group_name</td>
    <td>testSecurityGroupName</td>
  </tr>
  <tr>
    <td>security_group_type</td>
    <td>normal</td>
  </tr>
  <tr>
    <td>ak</td>
    <td>asdadasdsadas</td>
  </tr>
</table>

**Example response**

Normal return result：
```JSON
{
    "code":200,
    "msg":"success",
    "data":"vsw-m5evh*****s6ivwc"
}
```

Exception return result：
```JSON
{
    "code":400,
    "msg":"param_invalid",
    "data":null
}
```

**Return code explanation**

<table>
  <tr>
    <td>Return code</td>
    <td>Status</td>
    <td>Explanation</td>
  </tr>
  <tr>
    <td>200</td>
    <td>success</td>
    <td>Successful implementation</td>
  </tr>
  <tr>
    <td>400</td>
    <td>param_invalid</td>
    <td>Wrong parameters</td>
  </tr>
</table>


### 10. View the list of regions
  View a list of the cloud vendor's region.<br>
**Request Address**
<table>
  <tr>
    <td>GET method</td>
  </tr>
  <tr>
    <td>GET /api/v1/region/list </td>
  </tr>
</table>

**Request Parameters**
<table>
  <tr>
    <td>Name</td>
    <td>Type</td>
    <td>Required Field</td>
    <td>Description</td>
    <td>Sample Value</td>
  </tr>
  <tr>
    <td>provider</td>
    <td>String</td>
    <td>Yes</td>
    <td>Cloud vendor</td>
    <td>AlibabaCloud</td>
  </tr>
</table>
 
**Return parameters**
<table>
  <tr>
    <td>Name</td>
    <td>Type</td>
    <td>Required Field</td>
    <td>Description</td>
    <td>Sample Value</td>
  </tr>
  <tr>
    <td>code</td>
    <td>int</td>
    <td>Yes</td>
    <td>Return code</td>
    <td>0</td>
  </tr>
  <tr>
    <td>msg</td>
    <td>string</td>
    <td>Yes</td>
    <td>Error message</td>
    <td>null</td>
  </tr>
  <tr>
    <td>data</td>
    <td>object</td>
    <td>Yes</td>
    <td>Normal message</td>
    <td> {}</td>
  </tr>
</table>

**Request Example**
<table>
  <tr>
    <td>Name</td>
    <td>Sample Value</td>
  </tr>
  <tr>
    <td>provider</td>
    <td>AlibabaCloud</td>
  </tr>
</table>

**Example response**

Normal return result：
```JSON
{
    "code":200,
    "data":[
        {
            "RegionId":"cn-qingdao",
            "LocalName":"North China 1"
        },
        {
            "RegionId":"cn-beijing",
            "LocalName":"Norh China 2"
        }
    ],
    "msg":"success"
}
```

Exception return result：
```JSON
{
    "code":400,
    "msg":"param_invalid",
    "data": null
}
```

**Return code explanation**

<table>
  <tr>
    <td>Return code</td>
    <td>Status</td>
    <td>Explanation</td>
  </tr>
  <tr>
    <td>200</td>
    <td>success</td>
    <td>Successful implementation</td>
  </tr>
  <tr>
    <td>400</td>
    <td>param_invalid</td>
    <td>Wrong parameters</td>
  </tr>
</table>


### 11. View zone list
View a list of available zones under the cloud vendor's region.<br>
**Request Address**
<table>
  <tr>
    <td>GET method</td>
  </tr>
  <tr>
    <td>GET /api/v1/zone/list</td>
  </tr>
</table>

**Request Address**
<table>
  <tr>
    <td>Name</td>
    <td>Type</td>
    <td>Required Field</td>
    <td>Description</td>
    <td>Sample Value</td>
  </tr>
  <tr>
    <td>provider</td>
    <td>String</td>
    <td>Yes</td>
    <td>Cloud vendor</td>
    <td>AlibabaCloud</td>
  </tr>
  <tr>
    <td>region_id</td>
    <td>String</td>
    <td>Yes</td>
    <td>Region ID</td>
    <td> cn-qingdao</td>
  </tr>
</table>

**Return parameters**
<table>
  <tr>
    <td>Name</td>
    <td>Type</td>
    <td>Required Field</td>
    <td>Description</td>
    <td>Sample Value</td>
  </tr>
  <tr>
    <td>code</td>
    <td>int</td>
    <td>Yes</td>
    <td>Return code</td>
    <td>0</td>
  </tr>
  <tr>
    <td>msg</td>
    <td>string</td>
    <td>Yes</td>
    <td>Error message</td>
    <td>null</td>
  </tr>
  <tr>
    <td>data</td>
    <td>object</td>
    <td>Yes</td>
    <td>Normal message</td>
    <td> {}</td>
  </tr>
</table>

**Request Example**
<table>
  <tr>
    <td>Name</td>
    <td>Sample Value</td>
  </tr>
  <tr>
    <td>provider</td>
    <td>AlibabaCloud</td>
  </tr>
  <tr>
    <td>region_id</td>
    <td>cn-qingdao</td>
  </tr>
</table>


**Request Example**

Example response：
```JSON
{
    "code":200,
    "data":[
        {
            "ZoneId":"cn-qingdao-b",
            "LocalName":"North China 1 Available Zone B"
        },
        {
            "ZoneId":"cn-qingdao-c",
            "LocalName":"North China 1 Available Zone C"
        }
    ],
    "msg":"success"
}
```

Exception return result：
```JSON
{
    "code":400,
    "msg":"param_invalid",
    "data":null
}
```

**Return code explanation**

<table>
  <tr>
    <td>Return</td>
    <td>Status</td>
    <td>Explanation</td>
  </tr>
  <tr>
    <td>200</td>
    <td>success</td>
    <td>Successful implementation</td>
  </tr>
  <tr>
    <td>400</td>
    <td>param_invalid</td>
    <td>Wrong parameters</td>
  </tr>
</table>


### 12. View the list of models
View the list of models under that region and the available zones, through the cloud vendor's region and available zones.<br>

**Request Address**
<table>
  <tr>
    <td>GET method</td>
  </tr>
  <tr>
    <td>GET /api/v1/instance_type/list</td>
  </tr>
</table>

**Request Parameters**
<table>
  <tr>
    <td>Name</td>
    <td>Type</td>
    <td>Required Field</td>
    <td>Description</td>
    <td>Sample Value</td>
  </tr>
  <tr>
    <td>provider</td>
    <td>String</td>
    <td>Yes</td>
    <td>Cloud vendor</td>
    <td>AlibabaCloud</td>
  </tr>
  <tr>
    <td>region_id</td>
    <td>String</td>
    <td>Yes</td>
    <td>Region ID</td>
    <td>cn-qingdao</td>
  </tr>
  <tr>
    <td>zone_id</td>
    <td>String</td>
    <td>Yes</td>
    <td>Available zone</td>
    <td>cn-qingdao-b</td>
  </tr>
</table>

**Return parameters**
<table>
  <tr>
    <td>Name</td>
    <td>Type</td>
    <td>Required Field</td>
    <td>Description</td>
    <td>Sample Value</td>
  </tr>
  <tr>
    <td>code</td>
    <td>int</td>
    <td>Yes</td>
    <td>Return code</td>
    <td>0</td>
  </tr>
  <tr>
    <td>msg</td>
    <td>string</td>
    <td>Yes</td>
    <td>Error message</td>
    <td>null</td>
  </tr>
  <tr>
    <td>data</td>
    <td>object</td>
    <td>Yes</td>
    <td>Normal message</td>
    <td> {}</td>
  </tr>
</table>

**Important parameter in "data"**
<table>
  <tr>
    <td>Name</td>
    <td>Type</td>
    <td>Required Field</td>
    <td>Description</td>
    <td>Sample Value</td>
  </tr>
  <tr>
    <td>instance_type_family</td>
    <td>String</td>
    <td>Yes</td>
    <td>Model Specification Family</td>
    <td>ecs.g6</td>
  </tr>
  <tr>
    <td>instance_type</td>
    <td>String</td>
    <td>Yes</td>
    <td>Model Name</td>
    <td>ecs.g6.large</td>
  </tr>
  <tr>
    <td>core</td>
    <td>int</td>
    <td>Yes</td>
    <td>cpu cores</td>
    <td>4</td>
  </tr>
  <tr>
    <td>memory</td>
    <td>int</td>
    <td>Yes</td>
    <td>Memery size in G</td>
    <td>8</td>
  </tr>
</table>

**Request Example**
<table>
  <tr>
    <td>Name</td>
    <td>Sample Value</td>
  </tr>
  <tr>
    <td>provider</td>
    <td>AlibabaCloud</td>
  </tr>
  <tr>
    <td>region_id</td>
    <td>cn-qingdao</td>
  </tr>
  <tr>
    <td>zone_id</td>
    <td>cn-qingdao-b</td>
  </tr>
</table>


**Example response**

Normal return result：
```JSON
{
    "code":200,
    "data":[
        {
            "instance_type_family":"ecs.e3",
            "instance_type":"ecs.e3.large",
            "core":4,
            "memory":32
        },
        {
            "instance_type_family":"ecs.r6",
            "instance_type":"ecs.r6.13xlarge",
            "core":52,
            "memory":384
        }
    ],
    "msg":"success"
}
```

Exception return result：
```JSON
{
    "code":400,
    "msg":"param_invalid",
    "data":null
}
```

**Return code explanation**

<table>
  <tr>
    <td>Return code</td>
    <td>Status</td>
    <td>Explanation</td>
  </tr>
  <tr>
    <td>200</td>
    <td>success</td>
    <td>Successful implementation</td>
  </tr>
  <tr>
    <td>400</td>
    <td>param_invalid</td>
    <td>Wrong parameters</td>
  </tr>
</table>



### 13. Get the list of mirrors
View the list of models under that region and the available zones, through the cloud vendor's region and available zones.<br>
**Request Address**
<table>
  <tr>
    <td>GET method</td>
  </tr>
  <tr>
    <td>GET /api/v1/image/list</td>
  </tr>
</table>

**Request Parameters**
<table>
  <tr>
    <td>Name</td>
    <td>Type</td>
    <td>Required Field</td>
    <td>Description</td>
    <td>Sample Value</td>
  </tr>
  <tr>
    <td>provider</td>
    <td>String</td>
    <td>Yes</td>
    <td>Cloud vendor</td>
    <td>AlibabaCloud</td>
  </tr>
  <tr>
    <td>region_id</td>
    <td>String</td>
    <td>Yes</td>
    <td>Region</td>
    <td>cn-qingdao</td>
  </tr>
</table>

**Return parameters**
<table>
  <tr>
    <td>Name</td>
    <td>Type</td>
    <td>Required Field</td>
    <td>Description</td>
    <td>Sample Value</td>
  </tr>
  <tr>
    <td>code</td>
    <td>int</td>
    <td>Yes</td>
    <td>Return</td>
    <td>0</td>
  </tr>
  <tr>
    <td>msg</td>
    <td>string</td>
    <td>Yes</td>
    <td>Error message</td>
    <td>null</td>
  </tr>
  <tr>
    <td>data</td>
    <td>object</td>
    <td>Yes</td>
    <td>Normal</td>
    <td> {}</td>
  </tr>
</table>


**Import parameter in "data"**
<table>
  <tr>
    <td>Name</td>
    <td>Type</td>
    <td>Required Field</td>
    <td>Description</td>
    <td>Sample Value</td>
  </tr>
  <tr>
    <td>os_type</td>
    <td>String</td>
    <td>Yes</td>
    <td>Operating system</td>
    <td>linux</td>
  </tr>
  <tr>
    <td>os_name</td>
    <td>String</td>
    <td>Yes</td>
    <td>Operating system name</td>
    <td>Windows Server 2016 Data Center 64-bit Chinese Version</td>
  </tr>
  <tr>
    <td>image_id</td>
    <td>String</td>
    <td>Yes</td>
    <td>Mirror ID</td>
    <td>m-bp1g7004ksh0oeuc****</td>
  </tr>
</table>


**Request Example**
<table>
  <tr>
    <td>Name</td>
    <td>Sample Value</td>
  </tr>
  <tr>
    <td>provider</td>
    <td>AlibabaCloud</td>
  </tr>
  <tr>
    <td>region_id</td>
    <td>cn-qingdao</td>
  </tr>
</table>

**Example response**

Normal return result：
```JSON
{
    "code":200,
    "data":[
        {
            "OsType":"linux",
            "OsName":"CentOS  7.6 64位",
            "ImageId":"centos_7_6_x64_20G_alibase_20211030.vhd"
        },
        {
            "OsType":"linux",
            "OsName":"Gentoo  13  64bit",
            "ImageId":"gentoo13_64_40G_aliaegis_20160222.vhd"
        }
    ],
    "msg":"success"
}
```

Exception return result：
```JSON
{
    "code":400,
    "msg":"param_invalid",
    "data":null
}
```

**Return code explanation**

<table>
  <tr>
    <td>Return code</td>
    <td>Status</td>
    <td>Explanation</td>
  </tr>
  <tr>
    <td>200</td>
    <td>success</td>
    <td>Successful implementation</td>
  </tr>
  <tr>
    <td>400</td>
    <td>param_invalid</td>
    <td>Wrong parameters</td>
  </tr>
</table>



## Scaling Up And Scaling Down Task API
### 1. Create scale-up task 
Scale up the number of machines in a cluster.<br>

**Request Address**
<table>
  <tr>
    <td>POST Method </td>
  </tr>
  <tr>
    <td>POST /api/v1/cluster/expand </td>
  </tr>
</table>

**Request Parameters**
<table>
  <tr>
    <td>Name</td>
    <td>Type</td>
    <td>Required Field</td>
    <td>Description</td>
    <td>Sample Value</td>
  </tr>
  <tr>
    <td>task_name</td>
    <td>String</td>
    <td>Yes</td>
    <td>Task Name </td>
    <td>expand_task</td>
  </tr>
  <tr>
    <td>cluster_name</td>
    <td>String</td>
    <td>Yes</td>
    <td>Name of the cluster</td>
    <td>gf.metrics.test</td>
  </tr>
  <tr>
    <td>count</td>
    <td>Int</td>
    <td>Yes</td>
    <td>Number of machines for scaling up</td>
    <td>10</td>
  </tr>
</table>

**Return parameters**
<table>
  <tr>
    <td>Name</td>
    <td>Type</td>
    <td>Required Field</td>
    <td>Description</td>
    <td>Sample Value</td>
  </tr>
  <tr>
    <td>code</td>
    <td>int</td>
    <td>Yes</td>
    <td>Return code</td>
    <td>0</td>
  </tr>
  <tr>
    <td>msg</td>
    <td>string</td>
    <td>Yes</td>
    <td>Error message</td>
    <td>null</td>
  </tr>
  <tr>
    <td>data</td>
    <td>object</td>
    <td>Yes</td>
    <td>Normal message</td>
    <td> {}</td>
  </tr>
</table>


**Request Example**
```JSON
{
    "cluster_name":"gf.bridgx.online",
    "task_name":"aaa",
    "count":1
}
```
**Example response**

Normal return result：
```JSON
{
  "code": 200,
  "data": "69762**4493870",
  "msg": "success"
}
```

Exception return result：
```JSON
{
    "code":400,
    "msg":"param_invalid",
    "data":null
}
```

**Return code explanation**

<table>
  <tr>
    <td>Return code</td>
    <td>Status</td>
    <td>Explanation</td>
  </tr>
  <tr>
    <td>200</td>
    <td>success</td>
    <td>Successful implementation</td>
  </tr>
  <tr>
    <td>400</td>
    <td>param_invalid</td>
    <td>Wrong parameters</td>
  </tr>
</table>




### 2. Create scale-down task
Scale-down the number of machines in a cluster. If IP is specified, it will be shrunk according to the specified IP, and if IP is not specified, "count" machines will be randomly selected for scaling down.<br>
**Request Address**
<table>
  <tr>
    <td>POST method</td>
  </tr>
  <tr>
    <td>POST /api/v1/cluster/expand </td>
  </tr>
</table>

**Request Address**
<table>
  <tr>
    <td>Name</td>
    <td>Type</td>
    <td>Required Field</td>
    <td>Description</td>
    <td>Sample Value</td>
  </tr>
  <tr>
    <td>task_name</td>
    <td>String</td>
    <td>Yes</td>
    <td>Task Name</td>
    <td>expand_task</td>
  </tr>
  <tr>
    <td>cluster_name</td>
    <td>String</td>
    <td>Yes</td>
    <td>Name of the cluster</td>
    <td>gf.metrics.test</td>
  </tr>
  <tr>
    <td>ips</td>
    <td>String</td>
    <td>No</td>
    <td>IP address for the scale-down</td>
    <td>["10.192.220.195", "10.192.220.196", "10.192.220.197"]</td>
  </tr>
  <tr>
    <td>count</td>
    <td>Int</td>
    <td>Yes</td>
    <td>IP address for the scale-down </td>
    <td>10</td>
  </tr>
</table>

**Return parameters**
<table>
  <tr>
    <td>Name</td>
    <td>Type</td>
    <td>Required Field</td>
    <td>Description</td>
    <td>Sample Value</td>
  </tr>
  <tr>
    <td>code</td>
    <td>int</td>
    <td>Yes</td>
    <td>return code</td>
    <td>0</td>
  </tr>
  <tr>
    <td>msg</td>
    <td>string</td>
    <td>Yes</td>
    <td>Error message</td>
    <td>null</td>
  </tr>
  <tr>
    <td>data</td>
    <td>object</td>
    <td>Yes</td>
    <td>Normal message</td>
    <td> {}</td>
  </tr>
</table>

**Request Example**
```JSON
{
    "cluster_name":"gf.bridgx.online",
    "task_name":"asas",
    "count":1
}
```
**Example response**

Normal return result：
```JSON
{
  "code": 200,
  "data": "69762**4493871",
  "msg": "success"
}
```
Exception return result：
```JSON
{
    "code":400,
    "msg":"param_invalid",
    "data": null
}
```
**Return code explanation**

<table>
  <tr>
    <td>Return code</td>
    <td>Status</td>
    <td>Explanation</td>
  </tr>
  <tr>
    <td>200</td>
    <td>success</td>
    <td>Successful implementation</td>
  </tr>
  <tr>
    <td>400</td>
    <td>param_invalid</td>
    <td>Wrong parameters</td>
  </tr>
</table>



### 3. View task list
Query the list of tasks under the account. If no "account" is passed, all will be queried by default. If "account" is passed, those under the specified account will be queried.<br>
**Request Address**
<table>
  <tr>
    <td>GET method</td>
  </tr>
  <tr>
    <td>GET /api/v1/task/list</td>
  </tr>
</table>

**Request Parameters**
<table>
  <tr>
    <td>Name</td>
    <td>Type</td>
    <td>Required Field</td>
    <td>Description</td>
    <td>Sample Value</td>
  </tr>
  <tr>
    <td>account</td>
    <td>String</td>
    <td>No</td>
    <td>Cloud Account</td>
    <td>LTAI5tAWAM</td>
  </tr>
  <tr>
    <td>page_number</td>
    <td>int32</td>
    <td>No</td>
    <td>Default start page</td>
    <td>1</td>
  </tr>
  <tr>
    <td>page_size</td>
    <td>int32</td>
    <td>No</td>
    <td>Defult 10 Maximum 50</td>
    <td>15</td>
  </tr>
</table>

**Request Parameters**
<table>
  <tr>
    <td>Name</td>
    <td>Type</td>
    <td>Required Field</td>
    <td>Description</td>
    <td>Sample Value</td>
  </tr>
  <tr>
    <td>code</td>
    <td>int</td>
    <td>Yes</td>
    <td>Return code</td>
    <td>0</td>
  </tr>
  <tr>
    <td>msg</td>
    <td>string</td>
    <td>Yes</td>
    <td>Error message</td>
    <td>null</td>
  </tr>
  <tr>
    <td>data</td>
    <td>object</td>
    <td>Yes</td>
    <td>Normal message</td>
    <td> {}</td>
  </tr>
</table>

**Important content in "data"**
<table>
  <tr>
    <td>Name</td>
    <td>Type</td>
    <td>Required Field</td>
    <td>Description</td>
    <td>Sample Value</td>
  </tr>
  <tr>
    <td>task_id</td>
    <td>String</td>
    <td>Yes</td>
    <td>task ID</td>
    <td>4803680234646135</td>
  </tr>
  <tr>
    <td>task_name</td>
    <td>String</td>
    <td>Yes</td>
    <td>Task Name</td>
    <td>xx emergency scale-up xx</td>
  </tr>
  <tr>
    <td>task_action</td>
    <td>String</td>
    <td>Yes</td>
    <td>Task Action Type</td>
    <td>Scale-up,scale-down</td>
  </tr>
  <tr>
    <td>status</td>
    <td>String</td>
    <td>Yes</td>
    <td>Task Status</td>
    <td>Pending,in progress, implementation completed </td>
  </tr>
  <tr>
    <td>cluster_name</td>
    <td>String</td>
    <td>Yes</td>
    <td>Cluster name</td>
    <td>gf.xx.aa</td>
  </tr>
  <tr>
    <td>create_at</td>
    <td>String</td>
    <td>Yes</td>
    <td>Creation time</td>
    <td>2021-11-03 16:54:46</td>
  </tr>
  <tr>
    <td>execute_time</td>
    <td>int</td>
    <td>Yes</td>
    <td>Execution time(unit:seconds)</td>
    <td>18</td>
  </tr>
  <tr>
    <td>finish_at</td>
    <td>string</td>
    <td>Yes</td>
    <td>Completion time</td>
    <td>2021-11-03 16:55:36</td>
  </tr>
</table>

**Request Example**
<table>
  <tr>
    <td>Name</td>
    <td>Sample Value</td>
  </tr>
  <tr>
    <td>account</td>
    <td>LTAI5tAWAM</td>
  </tr>
  <tr>
    <td>page_number</td>
    <td>1</td>
  </tr>
  <tr>
    <td>page_size</td>
    <td>15</td>
  </tr>
</table>

**Example response**

Normal return result：
```JSON
{
  "code": 200,
  "data": {
    "task_list": [
      {
        "task_id": "6978292825513784",
        "task_name": "",
        "task_action": "EXPAND",
        "status": "SUCCESS",
        "cluster_name": "gf.scheduler.test",
        "create_at": "2021-11-18 11:23:07 +0800 CST",
        "execute_time": 11,
        "finish_at": "2021-11-18 11:23:18 +0800 CST"
      },
      {
        "task_id": "6692774019653071",
        "task_name": "test scale-down again",
        "task_action": "SHRINK",
        "status": "SUCCESS",
        "cluster_name": "gf.bridgxine",
        "create_at": "2021-11-16 12:06:44 +0800 CST",
        "execute_time": 4,
        "finish_at": "2021-11-16 12:06:48 +0800 CST"
      },
      {
        "task_id": "6692475385208271",
        "task_name": "test",
        "task_action": "SHRINK",
        "status": "SUCCESS",
        "cluster_name": "gf.bridgx.online",
        "create_at": "2021-11-16 12:03:46 +0800 CST",
        "execute_time": 5,
        "finish_at": "2021-11-16 12:03:51 +0800 CST"
      }
    ],
    "pager": {
      "page_number": 1,
      "page_size": 50,
      "total": 52
    }
  },
  "msg": "success"
}
```
Exception return result:
```JSON
{
    "code":400,
    "msg":"param_invalid",
    "data":null
}
```
**Return code explanation**

<table>
  <tr>
    <td>Return code</td>
    <td>Status</td>
    <td>Explanation</td>
  </tr>
  <tr>
    <td>200</td>
    <td>success</td>
    <td>Successful implementation</td>
  </tr>
  <tr>
    <td>400</td>
    <td>param_invalid</td>
    <td>Wrong parameters</td>
  </tr>
</table>


## Machine API
### 1. Machine list
Get information on all machines under this account.<br>
**Request Parameters**
<table>
  <tr>
    <td>GET method</td>
  </tr>
  <tr>
    <td>GET /api/v1/instance/describe_all</td>
  </tr>
</table>

**Request Parameters**
<table>
  <tr>
    <td>Name</td>
    <td>Type</td>
    <td>Required Field</td>
    <td>Description</td>
    <td>Sample Value</td>
  </tr>
  <tr>
    <td>account</td>
    <td>String</td>
    <td>No</td>
    <td>Cloud Account</td>
    <td>LTAI5tAWAM</td>
  </tr>
  <tr>
    <td>instance_id</td>
    <td>String</td>
    <td>No</td>
    <td>Instance ID, exact match</td>
    <td>i-xaf23fasdc1edg</td>
  </tr>
  <tr>
    <td>ip</td>
    <td>String</td>
    <td>No</td>
    <td>Intranet or extranet IP, exact match</td>
    <td>10.12.13.1</td>
  </tr>
  <tr>
    <td>provider</td>
    <td>String</td>
    <td>No</td>
    <td>Specific cloud vendor, exact match</td>
    <td>AlibabaCloud</td>
  </tr>
  <tr>
    <td>cluster_name</td>
    <td>String</td>
    <td>No</td>
    <td>Cluster name</td>
    <td>Cluster1</td>
  </tr>
  <tr>
    <td>status</td>
    <td>String</td>
    <td>No</td>
    <td>Status</td>
    <td>running,deleted</td>
  </tr>
  <tr>
    <td>page_number</td>
    <td>int32</td>
    <td>No</td>
    <td>Default start page</td>
    <td>1</td>
  </tr>
  <tr>
    <td>page_size</td>
    <td>int32</td>
    <td>No</td>
    <td>Default 10 Maximum 50</td>
    <td>15</td>
  </tr>
</table>

**Return parameters**
<table>
  <tr>
    <td>Name</td>
    <td>Type</td>
    <td>Required Field</td>
    <td>Description</td>
    <td>Sample Value</td>
  </tr>
  <tr>
    <td>code</td>
    <td>int</td>
    <td>Yes</td>
    <td>Return code</td>
    <td>0</td>
  </tr>
  <tr>
    <td>msg</td>
    <td>string</td>
    <td>Yes</td>
    <td>Error message</td>
    <td>null</td>
  </tr>
  <tr>
    <td>data</td>
    <td>object</td>
    <td>Yes</td>
    <td>Normal message</td>
    <td> {}</td>
  </tr>
</table>


**Important content in "data"**
<table>
  <tr>
    <td>Name</td>
    <td>Type</td>
    <td>Required Field</td>
    <td>Description</td>
    <td>Sample Value</td>
  </tr>
  <tr>
    <td>instance_list</td>
    <td>0</td>
    <td>[]</td>
    <td>Yes</td>
    <td>Machine List</td>
    <td></td>
  </tr>
  <tr>
    <td></td>
    <td>instance_id</td>
    <td>String</td>
    <td>Yes</td>
    <td>Instance ID</td>
    <td></td>
  </tr>
  <tr>
    <td></td>
    <td>ip_inner</td>
    <td>String</td>
    <td>Yes</td>
    <td>Intranet IP</td>
    <td>10.208.28.126[Intranet]</td>
  </tr>
  <tr>
    <td></td>
    <td>ip_outer</td>
    <td>String</td>
    <td>No</td>
    <td>Extranet IP</td>
    <td>106.14.169.121[Public network]</td>
  </tr>
  <tr>
    <td></td>
    <td>provider</td>
    <td>String</td>
    <td>Yes</td>
    <td>Cloud vendor</td>
    <td>AlibabaCloud</td>
  </tr>
  <tr>
    <td></td>
    <td>cluster_name</td>
    <td>String</td>
    <td>Yes</td>
    <td>Affiliated Clusters</td>
    <td>gf.bridgx.online</td>
  </tr>
  <tr>
    <td></td>
    <td>instance_type</td>
    <td>String</td>
    <td>Yes</td>
    <td>Model</td>
    <td>ecs.7c.large</td>
  </tr>
  <tr>
    <td></td>
    <td>create_at</td>
    <td>String</td>
    <td>Yes</td>
    <td>Creation time</td>
    <td>2021.10.29  18：25：13</td>
  </tr>
  <tr>
    <td></td>
    <td>status</td>
    <td>String</td>
    <td>Yes</td>
    <td>Machine Status</td>
    <td>Yes</td>
  </tr>
  <tr>
    <td>pager</td>
    <td>Pager</td>
    <td>String</td>
    <td>Yes</td>
    <td>Yes</td>
    <td>Paging parameters</td>
  </tr>
</table>

**Request Example**
<table>
  <tr>
    <td>Name</td>
    <td>Sample Value</td>
  </tr>
  <tr>
    <td>account</td>
    <td>LTAI5tAWAM</td>
  </tr>
  <tr>
    <td>instance_id</td>
    <td>i-xaf23fasdc1edg</td>
  </tr>
  <tr>
    <td>ip</td>
    <td>10.12.13.1</td>
  </tr>
  <tr>
    <td>provider</td>
    <td>AlibabaCloud</td>
  </tr>
  <tr>
    <td>cluster_name</td>
    <td>Cluster1</td>
  </tr>
  <tr>
    <td>status</td>
    <td>running,deleted</td>
  </tr>
  <tr>
    <td>page_number</td>
    <td>1</td>
  </tr>
  <tr>
    <td>page_size</td>
    <td>15</td>
  </tr>
</table>

**Example response**

Normal return result：
```JSON
{
    "code":200,
    "data":{
        "instance_list":[
            {
                "instance_id":"i-2ze40**rrjk7mi6",
                "ip_inner":"10.192.221.25",
                "ip_outer":"",
                "provider":"AlibabaCloud",
                "create_at":"2021-11-12 09:38:31 +0800 CST",
                "status":"Deleted",
                "startup_time":0,
                "cluster_name":"gf.bridgx.online",
                "instance_type":"ecs.s6-c1m1.small"
            },
            {
                "instance_id":"i-2ze25xv**vu06m0p2",
                "ip_inner":"10.192.221.123",
                "ip_outer":"",
                "provider":"AlibabaCloud",
                "create_at":"2021-11-12 19:58:21 +0800 CST",
                "status":"Deleted",
                "startup_time":5,
                "cluster_name":"gf.bridgx.online",
                "instance_type":"ecs.s6-c1m1.small"
            }
        ],
        "pager":{
            "page_number":1,
            "page_size":10,
            "total":1747
        }
    },
    "msg":"success"
}
```

Exception return result：
```JSON
{
    "code":400,
    "msg":"param_invalid",
    "data":null
}
```

**Return code explanation**

<table>
  <tr>
    <td>Return code</td>
    <td>Status</td>
    <td>Explanation</td>
  </tr>
  <tr>
    <td>200</td>
    <td>success</td>
    <td>Successful implementation</td>
  </tr>
  <tr>
    <td>400</td>
    <td>param_invalid</td>
    <td>Wrong parameters</td>
  </tr>
</table>


### 2. Machine details
Get the details of a certain machine.<br>
**Request Address**
<table>
  <tr>
    <td>GET method</td>
  </tr>
  <tr>
    <td>GET /api/v1/instance/id/describe</td>
  </tr>
</table>

**Request Parameters**
<table>
  <tr>
    <td>Name</td>
    <td>Type</td>
    <td>Required Field</td>
    <td>Description</td>
    <td>Sample Value</td>
  </tr>
  <tr>
    <td>instance_id</td>
    <td>String</td>
    <td>Yes</td>
    <td>Machine ID</td>
    <td>i-2ze40hb376ihrrjk7mi6</td>
  </tr>
</table>

**Return parameters**
<table>
  <tr>
    <td>Name</td>
    <td>Type</td>
    <td>Required Field</td>
    <td>Description</td>
    <td>Sample Value</td>
  </tr>
  <tr>
    <td>code</td>
    <td>int</td>
    <td>Yes</td>
    <td>Return code</td>
    <td>0</td>
  </tr>
  <tr>
    <td>msg</td>
    <td>string</td>
    <td>Yes</td>
    <td>Error message</td>
    <td>null</td>
  </tr>
  <tr>
    <td>data</td>
    <td>object</td>
    <td>Yes</td>
    <td>Normal message</td>
    <td>{}</td>
  </tr>
</table>


**Important content in "data"**

<table>
  <tr>
    <td>Name</td>
    <td>Sub-attributes</td>
    <td>Sub-sub-attributes</td>
    <td>Type</td>
    <td>Required field</td>
    <td>Description</td>
    <td>Sample Value</td>
  </tr>
  <tr>
    <td>instance_id</td>
    <td></td>
    <td></td>
    <td>String</td>
    <td>Yes</td>
    <td>Machine ID</td>
    <td></td>
  </tr>
  <tr>
    <td>provider</td>
    <td></td>
    <td></td>
    <td>String</td>
    <td>Yes</td>
    <td>Cloud vendor</td>
    <td>AlibabaCloud</td>
  </tr>
  <tr>
    <td>region_id</td>
    <td></td>
    <td></td>
    <td>String</td>
    <td>Yes</td>
    <td>Available zone</td>
    <td>cn-beijing-h</td>
  </tr>
  <tr>
    <td>create_at</td>
    <td></td>
    <td></td>
    <td>String</td>
    <td>Yes</td>
    <td>Creation time</td>
    <td>2021-10-29 16：23：24</td>
  </tr>
  <tr>
    <td>image_id</td>
    <td></td>
    <td></td>
    <td>String</td>
    <td>Yes</td>
    <td>Mirror ID</td>
    <td></td>
  </tr>
  <tr>
    <td>instance_type</td>
    <td></td>
    <td></td>
    <td>String</td>
    <td>Yes</td>
    <td>Sample specifications</td>
    <td>4-core 16G</td>
  </tr>
  <tr>
    <td>storage_config</td>
    <td></td>
    <td></td>
    <td>{object}</td>
    <td>Yes</td>
    <td>Storage Configuration</td>
    <td></td>
  </tr>
  <tr>
    <td></td>
    <td>system_disk_type</td>
    <td>String</td>
    <td>Yes</td>
    <td>System disk type</td>
    <td>cloud_efficiency</td>
  </tr>
  <tr>
    <td></td>
    <td>system_disk_size</td>
    <td></td>
    <td>String</td>
    <td>Yes</td>
    <td>System disk size</td>
    <td>40G</td>
  </tr>
  <tr>
    <td></td>
    <td>data_disks</td>
    <td></td>
    <td>[]</td>
    <td>No</td>
    <td></td>
    <td></td>
  </tr>
  <tr>
    <td></td>
    <td></td>
    <td>data_disk_type</td>
    <td>String </td>
    <td>No</td>
    <td>Data disk type</td>
    <td>cloud_efficiency</td>
  </tr>
  <tr>
    <td></td>
    <td></td>
    <td>data_disk_size</td>
    <td>String </td>
    <td>No</td>
    <td>Data disk size</td>
    <td>40G</td>
  </tr>
  <tr>
    <td></td>
    <td>data_disk_num</td>
    <td></td>
    <td>int</td>
    <td>Yes</td>
    <td>Number of data disks</td>
    <td>4</td>
  </tr>
  <tr>
    <td>network_config</td>
    <td></td>
    <td></td>
    <td>{object}</td>
    <td>Yes</td>
    <td>Network Configuration</td>
    <td></td>
  </tr>
  <tr>
    <td></td>
    <td>vpc_name</td>
    <td></td>
    <td>String</td>
    <td>Yes</td>
    <td>Name of VPC</td>
    <td>testvpc</td>
  </tr>
  <tr>
    <td></td>
    <td>subnet_id_name</td>
    <td></td>
    <td>String</td>
    <td>Yes</td>
    <td>Sebnet Name</td>
    <td></td>
  </tr>
  <tr>
    <td></td>
    <td>security_group_name</td>
    <td></td>
    <td>String</td>
    <td>Yes</td>
    <td>Security Group Name</td>
    <td></td>
  </tr>
  <tr>
    <td>ip_outer</td>
    <td></td>
    <td></td>
    <td>String</td>
    <td>Yes</td>
    <td>Public Network IP</td>
    <td></td>
  </tr>
  <tr>
    <td>ip_inner</td>
    <td></td>
    <td></td>
    <td>String</td>
    <td>Yes</td>
    <td>Intranet IP</td>
    <td></td>
  </tr>
</table>

**Request Example**
<table>
  <tr>
    <td>Name</td>
    <td>Sample Value</td>
  </tr>
  <tr>
    <td>instance_id</td>
    <td>i-2ze40hb**hrrjk7mi6</td>
  </tr>
</table>


**Example response**
Normal return result:
```JSON
{
    "code":200,
    "data":{
        "instance_id":"i-2ze40hb**hrrjk7mi6",
        "provider":"AlibabaCloud",
        "region_id":"cn-beijing",
        "image_id":"m-2ze**m3aadve22aq",
        "instance_type":"ecs.s6-c1m1.small",
        "ip_inner":"10.192.221.25",
        "ip_outer":"",
        "create_at":"2021-11-12 09:38:31 +0800 CST",
        "storage_config":{
            "system_disk_type":"cloud_efficiency",
            "system_disk_size":40,
            "data_disks":[
                {
                    "data_disk_type":"cloud_efficiency",
                    "data_disk_size":100
                }
            ],
            "data_disk_num":1
        },
        "network_config":{
            "vpc_name":"vpc-2zelmmlf**c2xb2",
            "subnet_id_name":"vsw-2ze**q6sa2fdj8l5",
            "security_group_name":"sg-2zefbt9tw0y***7vc3ac"
        }
    },
    "msg":"success"
}
```
Exception return result:
```JSON
{
    "code":400,
    "msg":"param_invalid",
    "data":null
}
```
<table>
  <tr>
    <td>Return code</td>
    <td>Status</td>
    <td>Explanation</td>
  </tr>
  <tr>
    <td>200</td>
    <td>success</td>
    <td>Successful implementation</td>
  </tr>
  <tr>
    <td>400</td>
    <td>param_invalid</td>
    <td>Wrong parameters</td>
  </tr>
</table>


### 3. Get the number of machines
Get the number of machines running under this account.<br>
**Request Address**
<table>
  <tr>
    <td>GET method</td>
  </tr>
  <tr>
    <td>GET /api/v1/instance/num</td>
  </tr>
</table>

**Request Parameters**
<table>
  <tr>
    <td>Name</td>
    <td>Type</td>
    <td>Required Field</td>
    <td>Description</td>
    <td>Sample Value</td>
  </tr>
  <tr>
    <td>account</td>
    <td>String</td>
    <td>No</td>
    <td>Cloud Account</td>
    <td>LTAI5tAWAM</td>
  </tr>
  <tr>
    <td>cluster_name</td>
    <td>String</td>
    <td>No</td>
    <td>Cluster name</td>
    <td>Test Cluster</td>
  </tr>
</table>

**Return parameters**
<table>
  <tr>
    <td>Name</td>
    <td>Type</td>
    <td>Required Field</td>
    <td>Description</td>
    <td>Sample Value</td>
  </tr>
  <tr>
    <td>code</td>
    <td>int</td>
    <td>Yes</td>
    <td>Return code</td>
    <td>0</td>
  </tr>
  <tr>
    <td>msg</td>
    <td>string</td>
    <td>Yes</td>
    <td>Error message</td>
    <td>null</td>
  </tr>
  <tr>
    <td>data</td>
    <td>object</td>
    <td>Yes</td>
    <td>Normal message</td>
    <td> {}</td>
  </tr>
</table>

**Important parameter in "data"**
<table>
  <tr>
    <td>Name</td>
    <td>Type</td>
    <td>Required Field</td>
    <td>Description</td>
    <td>Sample Value</td>
  </tr>
  <tr>
    <td>instance_num</td>
    <td>int64</td>
    <td>Yes</td>
    <td>Number of instances</td>
    <td>2</td>
  </tr>
</table>

**Request Example**
<table>
  <tr>
    <td>Name</td>
    <td>Sample Value</td>
  </tr>
  <tr>
    <td>account</td>
    <td>LTAI5tAWAM</td>
  </tr>
  <tr>
    <td>cluster_name</td>
    <td>Test Cluster</td>
  </tr>
</table>

**Example response**

Normal return result：
```JSON
{
    "code":200,
    "data":{
        "instance_num":1
    },
    "msg":"success"
}
```

Exception return result:
```JSON
{
    "code":400,
    "msg":"param_invalid",
    "data":null
}
```

**Return code explanation**

<table>
  <tr>
    <td>Return code</td>
    <td>Status</td>
    <td>Explanation</td>
  </tr>
  <tr>
    <td>200</td>
    <td>success</td>
    <td>Successful implementation</td>
  </tr>
  <tr>
    <td>400</td>
    <td>param_invalid</td>
    <td>Wrong parameters</td>
  </tr>
</table>




## Fees API
### 1. Total machine hours used for the day
If a cluster is specified, this returns the usage time of the specific cluster, otherwise it returns the total time of all clusters associated under the current account. <br>
**Request Address**
<table>
  <tr>
    <td>GET method</td>
  </tr>
  <tr>
    <td>GET /api/v1/instance/usage_total</td>
  </tr>
</table>

**Request Parameters**
<table>
  <tr>
    <td>Name</td>
    <td>Type</td>
    <td>Required Field</td>
    <td>Description</td>
    <td>Sample Value</td>
  </tr>
  <tr>
    <td>cluster_name</td>
    <td>String</td>
    <td>No</td>
    <td>Cluster Name</td>
    <td>gf.bridgx.online</td>
  </tr>
  <tr>
    <td>date</td>
    <td>String</td>
    <td>No</td>
    <td>yyyy-dd-mm</td>
    <td>2021-10-11</td>
  </tr>
</table>

**Return parameters**
<table>
  <tr>
    <td>Name</td>
    <td>Type</td>
    <td>Required Field</td>
    <td>Description</td>
    <td>Sample Value</td>
  </tr>
  <tr>
    <td>code</td>
    <td>int</td>
    <td>Yes</td>
    <td>Return code</td>
    <td>0</td>
  </tr>
  <tr>
    <td>msg</td>
    <td>string</td>
    <td>Yes</td>
    <td>Error message</td>
    <td>null</td>
  </tr>
  <tr>
    <td>data</td>
    <td>object</td>
    <td>Yes</td>
    <td>Normal message</td>
    <td> {}</td>
  </tr>
</table>


**Important parameter in "data"**
<table>
  <tr>
    <td>Name</td>
    <td>Type</td>
    <td>Required Field</td>
    <td>Description</td>
    <td>Sample Value</td>
  </tr>
  <tr>
    <td>usage_total</td>
    <td>int</td>
    <td>Yes</td>
    <td>Duration of use in seconds</td>
    <td>1800</td>
  </tr>
</table>

**Request Example**
<table>
  <tr>
    <td>Name</td>
    <td>Sample Value</td>
  </tr>
  <tr>
    <td>cluster_name</td>
    <td>gf.bridgx.online</td>
  </tr>
  <tr>
    <td>date</td>
    <td>2021-10-11</td>
  </tr>
  <tr>
    <td>page_size</td>
    <td>15</td>
  </tr>
</table>

**Example response**

Normal return result:
```JSON
{
    "code": 200,
    "msg": "success",
    "data": {
        "usage_total": 1800,
    }
}
```
Exception return result:
```JSON
{
    "code":400,
    "msg":"param_invalid",
    "data":null
}
```

**Return code explanation**

<table>
  <tr>
    <td>Return code</td>
    <td>Status</td>
    <td>Explanation</td>
  </tr>
  <tr>
    <td>200</td>
    <td>success</td>
    <td>Successful implementation</td>
  </tr>
  <tr>
    <td>400</td>
    <td>param_invalid</td>
    <td>Wrong parameters</td>
  </tr>
</table>


### 2. Breakdown of machine hours used for the day

**Request Address**
<table>
  <tr>
    <td>GET method</td>
  </tr>
  <tr>
    <td>GET /api/v1/instance/usage_statistics</td>
  </tr>
</table>

**Request Parameters**

<table>
  <tr>
    <td>Name</td>
    <td>Type</td>
    <td>Required Field</td>
    <td>Description</td>
    <td>Sample Value</td>
  </tr>
  <tr>
    <td>cluster_name</td>
    <td>String</td>
    <td>No</td>
    <td>Cluster Name</td>
    <td>gf.bridgx.online</td>
  </tr>
  <tr>
    <td>date</td>
    <td>String</td>
    <td>Yes</td>
    <td>yyyy-dd-mm</td>
    <td>2021-10-11</td>
  </tr>
  <tr>
    <td>page_number</td>
    <td>int32</td>
    <td>No</td>
    <td>Default start page</td>
    <td>1</td>
  </tr>
  <tr>
    <td>page_size</td>
    <td>int32</td>
    <td>No</td>
    <td>Default 10 Maximum 50</td>
    <td>15</td>
  </tr>
</table>

**Return parameters**
<table>
  <tr>
    <td>Name</td>
    <td>Type</td>
    <td>Required Field</td>
    <td>Description</td>
    <td>Sample Value</td>
  </tr>
  <tr>
    <td>code</td>
    <td>int</td>
    <td>Yes</td>
    <td>return code</td>
    <td>0</td>
  </tr>
  <tr>
    <td>msg</td>
    <td>string</td>
    <td>Yes</td>
    <td>Error message</td>
    <td>null</td>
  </tr>
  <tr>
    <td>data</td>
    <td>object</td>
    <td>Yes</td>
    <td>Normal message</td>
    <td> {}</td>
  </tr>
</table>

**Parameter in “data”**
<table>
  <tr>
    <td>Type</td>
    <td>Required Field</td>
    <td>Description</td>
    <td>Sample Value</td>
  </tr>
  <tr>
    <td>instance_list</td>
    <td></td>
    <td>[]</td>
    <td></td>
    <td></td>
  </tr>
  <tr>
    <td></td>
    <td>id</td>
    <td>string</td>
    <td>S/No.</td>
    <td>1</td>
  </tr>
    <td></td>
    <td>cluster_name</td>
    <td>string</td>
    <td>Cluster Name</td>
    <td>gf.bridgx.online</td>
  </tr>
  </tr>
    <td></td>
    <td>instance_id</td>
    <td>string</td>
    <td>yyyy-dd-mm</td>
    <td>2021-10-11</td>
  </tr>
  </tr>
    <td></td>
    <td>startup_at</td>
    <td>string</td>
    <td>Startup time</td>
    <td>2021-11-11 15:15:20</td>
  </tr>
  </tr>
    <td></td>
    <td>shutdown_at</td>
    <td>string</td>
    <td>Startup time</td>
    <td>2021-11-11 15:45:20</td>
  </tr>
  </tr>
    <td></td>
    <td>startup_time</td>
    <td>int</td>
    <td>Machine service time in seconds</td>
    <td>1800</td>
  </tr>
  </tr>
    <td></td>
    <td>instance_type</td>
    <td>string</td>
    <td>Machine Type</td>
    <td>esc.c6.large</td>
  </tr>
  </tr>
    <td>pager</td>
    <td></td>
    <td>Pagination Information</td>
    <td></td>
    <td></td>
  </tr>
</table>

**Request Example**
<table>
  <tr>
    <td>Name</td>
    <td>Sample Value</td>
  </tr>
  <tr>
    <td>cluster_name</td>
    <td>gf.bridgx.online</td>
  </tr>
  <tr>
    <td>date</td>
    <td>2021-10-11</td>
  </tr>
  <tr>
    <td>page_number</td>
    <td>1</td>
  </tr>
  <tr>
    <td>page_size</td>
    <td>15</td>
  </tr>
</table>

**Example response**

Normal return result:
```JSON
{
    "code":200,
    "data":{
        "instance_list":[
            {
                "id":"1754",
                "cluster_name":"gf.scheduler.test",
                "instance_id":"i-2zeaavw***it89yvqm",
                "startup_at":"2021-11-18 10:16:53 +0800 CST",
                "shutdown_at":"2021-11-18 11:21:27 +0800 CST",
                "startup_time":3874,
                "instance_type":"ecs.s6-c1m1.small"
            },
            {
                "id":"1755",
                "cluster_name":"gf.scheduler.test",
                "instance_id":"i-2zeaavwojb***t89yvqn",
                "startup_at":"2021-11-18 10:16:53 +0800 CST",
                "shutdown_at":"2021-11-18 11:21:27 +0800 CST",
                "startup_time":3874,
                "instance_type":"ecs.s6-c1m1.small"
            },
            {
                "id":"1756",
                "cluster_name":"gf.bridgx.online",
                "instance_id":"i-2ze3ccvjn****1d9tzd0",
                "startup_at":"2021-11-18 10:31:53 +0800 CST",
                "shutdown_at":"2021-11-18 11:05:57 +0800 CST",
                "startup_time":2044,
                "instance_type":"ecs.s6-c1m1.small"
            }
        ],
        "pager":{
            "page_number":1,
            "page_size":10,
            "total":8
        }
    },
    "msg":"success"
}
```

Exception return result:
```JSON
{
    "code":400,
    "msg":"param_invalid",
    "data":null
}
```
**Return code explanation**

<table>
  <tr>
    <td>Return code</td>
    <td>Status</td>
    <td>Explanation</td>
  </tr>
  <tr>
    <td>200</td>
    <td>success</td>
    <td>Successful implementation</td>
  </tr>
  <tr>
    <td>400</td>
    <td>param_invalid</td>
    <td>Wrong parameters</td>
  </tr>
</table> 



