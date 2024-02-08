import paho.mqtt.client as mqtt
from typing import Optional

class Sensor(mqtt.Client):

    def __init__(
            self,
            name: str,
            latitude: float,
            longitude: float,
            measurement: float,
            rate: int = 60
        ):
        
        super().__init__()
        self.set_attributes(name, latitude, longitude, measurement, rate)

    @staticmethod
    def on_message(client, userdata, message):
        print(f"Received: {message.payload.decode()} on topic {message.topic}")
    
    def set_attributes(
            self,
            name: Optional[str] = None,
            latitude: Optional[float]=None,
            longitude: Optional[float]=None,
            measurement: Optional[float]=None,
            rate: Optional[float]=None
        ):

        self.name = name or self.name
        self.latitude = latitude or self.latitude
        self.longitude = longitude or self.longitude
        self.measurement = measurement or self.measurement
        self.rate = rate or rate or self.rate

    def __str__(self) -> str:
        return f'{self.name} is at {self.latitude} and {self.longitude}'