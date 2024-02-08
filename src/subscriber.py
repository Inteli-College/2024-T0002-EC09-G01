import paho.mqtt.client as mqtt

client = mqtt.Client()

broker = "broker.hivemq.com"
port = 1883

def on_message(client, userdata, message):
    print(f"Received: {message.payload.decode()} on topic {message.topic}")

client.on_message = on_message

client.connect(host=broker, port=port, keepalive=60)

client.subscribe("sensor/1")

client.loop_forever()