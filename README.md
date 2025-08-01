# LinkID [![Go Report Card](https://goreportcard.com/badge/github.com/shuban-789/LinkID)](https://goreportcard.com/report/github.com/shuban-789/LinkID)

The LinkID blockchain system is an AES encryption enhanced Blockchain system for MediLink to ensure secure containment and transfer of medical records.

## Contributions

The LinkID blockchain system was originally created and programmed by me as a subsystem of the MediLink project for the 2024 San Diego Big Data Hackathon. 

The original repository is licensed under MIT with my name (Shuban Pal) and acts as an archive. Although MediLink is now inactive, I wish to update LinkID further. Active updates and changes will now go to this repository which is still licensed under the MIT license with my name. As of 8/1/2025, I am archiving this repository. Although this hackathon project taught me a lot about blockchain, I am unsure whether this is valid to be called an actual blockchain system. For now, I do not see this project continuing any further.

All my preceding commits and contributions prior to this version can be found on the archive repository [HERE](https://github.com/TEAM-GOJO/LinkID)

> DISCLAIMER: This program was written for educational purposes. As stated by the MIT license, any liability for damages is waived. This program still has multiple steps to go before it achieves robust security (Right now, it uses AES which is a symmetric encryption algorithm)

## Usage

```
Usage: ./linkid [OPTION1] [ARGUMENT1] ... [OPTIONn] [ARGUMENTn]


Options:
  -c, Create a new blockchain with the provided JSON file.
                -E, Save the output as JSON
  -a, Access an existing blockchain with the provided ID and key.
                -E, Save the output as JSON
  -A, Add a new block to an existing blockchain with the provided ID and key.

Format:
  ./linkid -c <GENESIS.json>
  ./linkid -a <ID> <KEY>
  ./linkid -A <BLOCK.json> <ID> <KEY>

Examples:
  ./linkid -c genesis.json
  ./linkid -a 12345678 1234567890abcdef1234567890abcdef
  ./linkid -A block.json 12345678 1234567890abcdef1234567890abcdef
```

## Makefile Variables for `OS` (GOOS)

Common Operating System configurations for compiling the LinkID source code via Makefile. If you want to compile the code on an operating system not listed below, please check out this [list](https://pkg.go.dev/internal/platform) for a list of valid `GOOS` and `GOARCH` combinations.

### Linux (default) 🐧
```
OS=linux
```

### MacOS 🍎
```
OS=darwin
```

### Windows 🪟
```
OS=windows
```

## Makefile Variables for `ARCH` (GOARCH)

| GOARCH Variable       | Processor Name   | 32-bit    | 64-bit    |
| :-------------------- | :--------------: | :-------: | :-------: |
| `ARCH=386`            | Intel 386        | ✅        |           |
| `ARCH=amd64`          | AMD64            |           | ✅        |
| `ARCH=amd64p32`       | AMD64 (32-bit)   | ✅        |           |
| `ARCH=arm`            | ARM              | ✅        |           |
| `ARCH=arm64`          | ARM64            |           | ✅        |
| `ARCH=arm64be`        | ARM64 (big-endian)|          | ✅        |
| `ARCH=armbe`          | ARM (big-endian) | ✅        |           |
| `ARCH=loong64`        | Loongson 64-bit  |           | ✅        |
| `ARCH=mips`           | MIPS             | ✅        |           |
| `ARCH=mips64`         | MIPS64           |           | ✅        |
| `ARCH=mips64le`       | MIPS64 (little-endian) |    | ✅        |
| `ARCH=mips64p32`      | MIPS64 (32-bit)  | ✅        |           |
| `ARCH=mips64p32le`    | MIPS64 (32-bit little-endian)| ✅      |   |
| `ARCH=mipsle`         | MIPS (little-endian)| ✅      |          |
| `ARCH=ppc`            | PowerPC          | ✅        |           |
| `ARCH=ppc64`          | PowerPC 64       |           | ✅        |
| `ARCH=ppc64le`        | PowerPC 64 (little-endian) | | ✅        |
| `ARCH=riscv`          | RISC-V           | ✅        |           |
| `ARCH=riscv64`        | RISC-V 64        |           | ✅        |
| `ARCH=s390`           | IBM System/390   | ✅        |           |
| `ARCH=s390x`          | IBM System/390x  |           | ✅        |
| `ARCH=parc`           | SPARC            | ✅        |           |
| `ARCH=sparc64`        | SPARC64          |           | ✅        |
| `ARCH=wasm`           | WebAssembly      | ✅        |           |

## Notes

Earlier it was attempted to implement an asymmetric solution into this project (Hybrid RSA/AES encryption) which is why the code was played around with for a bit concerning the genration of key pairs through RSA, the format at which the pairs could be generated and given out, and even some RSA experimentation. As of now, it would seem a bit out of place to use this hybrid encryption for the purpose and use set which this program was developed for. However, more development in the future may be gaurunteed. Here is how this could work out:

- An RSA private key and public key is generated
- The chain json file is still enrypted via AES, but the AES key itself is encrypted via the public key.
- The program will require the private key and the encrypted AES key to unlock the chain

The purpose of this could possibly be that the AES key itself may stay constant or may change not so often, however the public key and private key pair which encrypts and decrypts the key itself could alternate.
