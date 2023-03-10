# PasswordsGenerator

Generate random passwords

## Installation

### From Github releases page

Go to [Release page](https://github.com/danny270793/passwordsgenerator/releases) then download the binary which fits your environment

### From terminal

Get the last versión available on github

```bash
LAST_VERSION=$(curl https://api.github.com/repos/danny270793/PasswordsGenerator/releases/latest | grep tag_name | cut -d '"' -f 4)
```

Download the last version directly to the binaries folder

For Linux (linux):

```bash
curl -L https://github.com/danny270793/PasswordsGenerator/releases/download/${LAST_VERSION}/PasswordsGenerator_${LAST_VERSION}_linux_amd64.tar.gz -o ./passwordssgenerator.tar.gz
```

For MacOS (darwin):

```bash
curl -L https://github.com/danny270793/PasswordsGenerator/releases/download/${LAST_VERSION}/PasswordsGenerator_${LAST_VERSION}_darwin_amd64.tar.gz -o ./passwordssgenerator.tar.gz
```

Untar the downloaded file

```bash
tar -xvf ./passwordssgenerator.tar.gz
```

Then copy the binary to the binaries folder

```bash
sudo cp ./PasswordsGenerator /usr/local/bin/passwords-generator
```

Make it executable the binary

```bash
sudo chmod +x /usr/local/bin/passwords-generator
```

```bash
passwords-generator --version
```

## Ussage

Run the binary and pass the path to the folders where you want to search for git projects

```bash
passwordsgenerator -length=50 -symbols=false -numbers=false
```

![command output](./images/output.png)

## Follow me

- [Youtube](https://www.youtube.com/channel/UC5MAQWU2s2VESTXaUo-ysgg)
- [Github](https://www.github.com/danny270793/)
- [LinkedIn](https://www.linkedin.com/in/danny270793)

## LICENSE

Licensed under the [MIT](license.md) License

## Version

PasswordsGenerator version 1.0.0

Last update 10/03/2023
