# Gophercraft/protoss

[![Go Reference](https://pkg.go.dev/badge/github.com/Gophercraft/protoss.svg)](https://pkg.go.dev/github.com/Gophercraft/protoss)
[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)
[![Chat on discord](https://img.shields.io/discord/556039662997733391.svg)](https://discord.gg/xPtuEjt)

Protoss is a tool used to decompile Protobuf message and BGS service descriptors from an executable file.

Protoss is specifically created for the purpose of decompiling BGS service definitions for use in Gophercraft, and might not be useful for other applications.

# Workflow

Protoss uses its own output to assist in subsequent decompilation.

BGS services have additional metadata tucked away in their service descriptors, which need to be extracted.

## 1. Extract definitions

```bash
protoss game.exe -o proto
```

## 2. Compile extensions

```bash
./script/compile_extensions
```

## 3. Rebuild Protoss

```bash
go install -v github.com/Gophercraft/protoss/cmd/protoss
```

## 4. Re-extract definitions now with proper extensions available

```bash
protoss game.exe -o proto
```

# Thanks

- https://github.com/marin-m/pbtk
- https://github.com/Shauren/protobuf-decompiler/