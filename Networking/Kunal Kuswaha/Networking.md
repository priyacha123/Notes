# Network 
- When computers are connected together

# Internet 
- collection of these computer networks 

# Arpanet(US) 
- To communicate -> MIT - Stanford - UCLA - University of Utah

# Tim burners 
- Formed "www"

# Protocols
-  Protocols are just rules that are defined by the internet Society.
- these are the rules that tells how data is transferred and everything 
-  basic protocols
      1. TCP (transmission control protocol)
       - Okay so the idea is that it will ensure that the data will reach its
         destination and it will not be like uh corrupted on the way okay so there's something you definitely want the person to get completely you should use it will be using TCP in your applications  
      2. UDP (user datagram protocol)
       - when you do not care about if 100% of the data is reaching your friend or not whomever you want to send the data.
      3. HTTP (hyper text transfer protocol) 
       - This is being used by web browsers the worldwide web okay so it basically defines the format of the um you know uh the data that is being transferred between your like uh when we talking about www so web clients basically and web servers okay clients and servers Okay cool so when a client over here in this you can see sends a request to a server so these particular things like Client Center request to a server server will send something back all of these things are uh given inside HTTP how the server will send back the data that is also a rule that is given in HTTP.

# IP Address
- computers and servers they are all identified via something known as an IP address
- example : X-X-X-X -> this is a format and every single X can have the numbers from 0 - 255 
- IP address there are like classes of IP address that are reserved 
- ISP (Internet service provider) -> modem/router -> Global IP address
- Modem/router -> Multiple devices are connected -> Local IP address
- Modem assigns these IP addresses using DHCP protocol (Dynamic host configuration protocol)
- Rrequest send -> now when response comes to modem , it'll know which device requested through IP Address -> and which application requested it is known through Port Numbers

# Port Numbers
- Ports are basically 16 bit numbers 
- 16 bit basically means you can have 16 cells um and each every cell can contain what zero or one so how many total numbers can you create '2^16' around 65,000 port numbers are available.
- HTTP -> 80
- MongoDB -> 27017
- Ports from 0 - 1023 -> Reserved Port means that if you have your own applications created and you're  like hey I will host my application on Port 80 you can't do that it's already reserved for HTTP
- Ports 1024 - 49152 -> hey are registered for applications some specific applications like MongoDB my SQL okay so SQL has every SQL
- Every SQL Server -> 1433 Port number
- Remaining ones u can use.

# Internet Speed
-1 MBPS(Mega-Bites Per Second) -> 10^6 bits/s
-1 GBPS (Giga-Bites Per Second) -> 10^9 bits/s
-1 KBPS (Kilo     s Per Second) -> 10^3 bits/s
-Download speed -> when someone sends data to you and you're downloading that called Download speed.
-Upload Speed -> you are sending data from one computer to another computer that is known as upload speed.

# How does the communication between two computers and things like this happen?
- There are two ways via which it happens:
    - Guided way -> there's a set of path already defined. For example two computers are connected with a wire that's the guided way
    - Unguided way -> communication is happening but there's no one single path for example Wi-Fi and Bluetooth.

# Submarine Cables Map (Optical Fibre Cables)
- website -> submarinecable.com
- wires running across the ocean from one country to another wires are running down the ocean.
-  in India from Chennai coachin and somewhere from the south we are connected to Sri Lanka and from Mumbai we are connected to like Dubai Oman UAE and we're also connected to somewhere in Singapore so we're connected to Malaysia.
- he bigger entity they give it to like smaller entities and then they give it to internet service providers and internet service providers give the control to us that's how it works on a larger scale 
- Physically connection -> Optical fibre cables , co-axial cables
- Wireless -> Bluetooth , wi-fi , for longer langes (3G , 4G , LTE , 5G)
- why can't we just use satellites why do we have to put cables under the ocean so many long cables across the world because it's faster than satellite

# LAN , MAN , WAN
- LAN (local area network) -> small house , office
                           -> connected by ethernet , wifi
- MAN (Metropolitan area network) -> across a city
- WAN (Wide area network) -> across a country
                         -> connected by optical fibre cables
                         -> SONET (synchronous Optical networking) ->  it basically carries the data uh using optical fiber cables hence it can cover larger distances.
                         -> frame relay ->  it's basically a way for you to connect your local area network to The Wider area like the internet.

- Internet is actually a collection of all these -> lot of local area networks that are connected to each other using Metropolitan networks that are connected to each other using wide area network. 