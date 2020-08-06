# Device Registry Service (DRS)
Written in python 3

## Usage
All responses will be of the form
'''json
{
	"data": "Mixed type, holding the content of the response",
	"message": "Description of activity",
}
'''
Subsequent response definitions will only detail the expected value of the 'data field'

### List all devices
**Definition**
'GET /devices'
**Response**
- '200 OK' on success
'''json
[
	{
		"identifier": "living-room-floor-lamp",
		"name": "Floor Lamp",
		"device_type": "switch",
		"controller_gateway": "192.168.0.2"
	},
	{
		"identifier": "living-room-tv",
		"name": "Living Room TV",
		"device_type": "television",
		"controller_gateway": "192.168.0.3"
	}
]
'''

### Registering a new device
**Definition**
'POST /devices'
**Arguments**
- '"identifier": string' a globally unique identifier for this device
- '"name": string' a common/display name for the device
- '"device_type": string' the type of the device, as understood by the client
- '"controller_gateway": string' the IP address of the device's controller
If a device with the given identifier already exists, the existing device will be overwritten (naive solution).
**Response**
- '201 Created' on success
'''json
{
	"identifier": "living-room-floor-lamp",
	"name": "Floor Lamp",
	"device_type": "switch",
	"controller_gateway": "192.168.0.2"
}
'''

### Look up device details
'GET /device/<identifier>'
**Response**
- '404 Not Found' if the device does not exist
- '200 OK' on success
'''json
{
	"identifier": "living-room-floor-lamp",
	"name": "Floor Lamp",
	"device_type": "switch",
	"controller_gateway": "192.168.0.2"
}
'''

### Delete a device
**Definition**
'DELETE /devices/<identifier>'
**Response**
- '404 Not Found' if the device does not exist
- '204 No Content' if the device was removed, but no useful data to return
