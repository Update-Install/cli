<div align="center">
<h1>Update-Install CLI</h1>
Easy way to install package for LInux Distributions<br>
<code>ui install</code>
</div>

# Commands
`ui --help` show command information  
`ui config` show the config file of ui  
`ui config -n {name} -u {URL}` set a source for a package  
`ui install` download the package file in config file's list then install it.

# Installation
## Build from source
### Requirements
- Go (1.22)

### Installation Steps
Download the source files
```bash
git clone https://github.com/Update-Install/CLI.git
```
Install packages and build
```
go install
go build . -o ui
```
Then just add the executable file to PATH

## Install Pre-build file with script
```bash
wget https://github.com/Update-Install/CLI/releases/download/v0.1.0/ui-cli_0.1.0_linux_amd64.tar.gz
tar -xvf ui-cli_0.1.0_linux_amd64.tar.gz
sudo cp ui /usr/local/bin
```
