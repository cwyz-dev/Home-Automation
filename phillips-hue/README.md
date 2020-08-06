### Get device state
**Definition**
'GET service.controller.hue/device/<device-identifier>'
**Response**
- '200 OK' on success
- '404 Not Found' if the device does not exists
'''json
{
	"identifier": "table-lamp",
	"name": "Table Lamp",
	"type": "hueLight",
	"state":
	{
		"brightness":
		{
			"type": "int",
			"min": 0,
			"max": 100,
			"value": 73
		}
	}
}
'''
### Update a device
**Definition**
'PATCH service.controller.hue/device/<device-identifier>'
**Arguments**
The properties to change and their new values
**Response**
- '200 OK' on success
- '400 Bad Request' if request validation fails
- '404 Not Found' if the device does not exist
'''json
{
	"identifier": "table-lamp",
	"name": "Table Lamp",
	"type": "hueLight",
	"state":
	{
		"brightness":
		{
			"type": "int",
			"min": 0,
			"max": 100,
			"value": 73
		}
	}
}
