from sensor import Sensor

broker = "broker.hivemq.com"
port = 1883

sensor = Sensor("SPS30", 40.7128, 74.0060, 100.0, 60)

sensor.connect(host=broker, port=port, keepalive=60)

sensor.publish("sensor/1", sensor.measurement)

sensor.loop_forever()