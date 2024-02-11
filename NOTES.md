# Disclaimer:

The following notes are based on scenario that you already have connected yourself to the network of the client/target either compromised it during your pentest activity or the access is granted to the network by your client.

## Tools used:
- netdiscover
- nmap
- snmp-check
- plcscan
- modbus
- ICSSecurityScripts
- Shodan
- metasploit
- Google Dorking
  

# Let's begin:

### Network Discovery:

Assuming you are on a  `10.1.0.0/24` network.

Let's run `netdiscover` to identify all other devices up in this network range. For that we use the following command:

```sh
netdiscover -r 10.1.0.0/24
```
By default `netdiscover` runs in active mode, checking all hosts in a network. The output will provide you with list of IP addresses reachable in your network, MAC Address, and MAC Vendor or Hostname. MAC Vendor or Hostname determine what type of system is running on that IP and it is important information for your pentest activity.

Once we identify what IP addresses are reachable with your `netdiscover` network, we run `nmap` scan to discover which ports are open on that IP address. We run the following simple command to accomplish this task.

To scan all TCP ports:
```bash
nmap -Pn 10.1.0.11 -p 1-65535 
```

To scan UDP ports
```sh
nmap -Pn 10.1.0.11 -sU -p 1-65535
```
The results from `nmap` will provide you with all open ports on the specific IP Address.

Next, we run `snmp-check` tool to gain more information. You will run this when you have port `161/udp` open on the network. This tool will provide you with information such as vendor name, model series, fimware version, model code, serial number and other network data.

```sh
snmp-check 10.1.0.11
```

### Open Source Tools for Information Gathering:

We use tools such as `Shodan` and `Google Dorking` to identify hosts connected to the internet and accessible remotely. We later try to compromise them. An example of the queries is available below:

Shodan: `port: 102 Siemens SIMANTIC 6ES7`

Google Dork: `inurl:/portal/portal.mwsl`

You may try this and get a gist of what this dorks can provide you with.

Once you identify different portals accessible via the internet, you can try default passowrds to access them.




