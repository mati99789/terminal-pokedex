# 🧭 Terminal Pokédex (Go CLI)

A **terminal-based Pokédex** built in **Go** as a learning project.  
This application runs in an interactive REPL and allows you to explore locations, catch Pokémon, and inspect your collection — all from the command line.

Built as part of a **Boot.dev exercise** to practice:
- Go fundamentals
- Structs & maps
- Variadic functions
- REPL design
- API-driven CLI apps

---

## 🚀 Features

- Interactive command-line REPL
- Explore Pokémon world locations
- Catch Pokémon
- Inspect caught Pokémon
- Pagination for map navigation
- Extensible command system

---

## 🖥️ Commands

| Command | Description |
|------|------------|
| `help` | Display available commands |
| `exit` | Exit the Pokédex |
| `map` | Show the next list of locations |
| `mapb` | Go back to the previous list of locations |
| `explore <location>` | Explore a location and list available Pokémon |
| `catch <pokemon>` | Attempt to catch a Pokémon |
| `inspect <pokemon>` | View details of a caught Pokémon |

---

## 🧠 Example Usage

```text
Pokedex > help
Pokedex > map
Pokedex > explore viridian-forest
Pokedex > catch pikachu
Pokedex > inspect pikachu
Pokedex > exit
