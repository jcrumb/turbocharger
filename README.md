# turbocharger
HTTP Proxy for full armour, tune, brakes, and turbo regardless of level/unlocks

# Usage 

* Requires the iFruit app to use
* Once app is installed, set your proxy on the device to turbo.jaycrumb.me, port 8228
* Log in to the app and head to LS customs
* Select your desired online car, and change the plate
* Press order, and you should recieve an in game notification shortly to go pick up the order at LS customs
* Fufill the order and you will have full performance upgrades
* Remove the proxy from your device

# Troubleshooting
If you don't receive a notification after a couple minutes, there may have been a problem processing your order.
this happens most often when using the app for the first time. Turn off the proxy, change a plate on one of your cars first and after that
order goes through, turn it back on and change again.

# How it works
The iFruit app communicates with a web api hosted by rockstar. When an order is placed, the POSTs to the web api, and the body
of this request contains a json object detailing the parts for the order. However, all validation for this request is done
on the client, meaning that if you can modify the request before it hits the server, you car order whatever you like. 

Details on the fields in this object and what the values represent can be found [here](https://docs.google.com/spreadsheet/ccc?key=0AixUkyNxN55gdF83LWI1MVFaeE9CY0ptdFEyYVFPV3c&usp=sharing#gid=0)

The aim of this proxy is to simplify adding the performance parts to your car. If that's all you need to do, there's no need to proxy and modify
the request yourself via charles or some other locally install proxy software. 
